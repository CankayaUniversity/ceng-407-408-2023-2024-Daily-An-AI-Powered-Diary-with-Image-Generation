from loguru import logger
from transformers import AutoTokenizer, AutoModelForSequenceClassification
from pydantic import BaseModel
import torch
from fastapi import FastAPI, Request, Response
from typing import List
from keybert import KeyBERT

app = FastAPI()


# Initialize logging
logger.add("file.log", rotation="500 MB")


# Load model directly
tokenizer = AutoTokenizer.from_pretrained(
    "daily-ai-emotion-classifier", local_files_only=True)
model = AutoModelForSequenceClassification.from_pretrained(
    "daily-ai-emotion-classifier", local_files_only=True)


# Initialize FastAPI app
app = FastAPI()


class Prediction(BaseModel):
    emotions: dict
    keywords: list
    topics: list


class Predictions(BaseModel):
    predictions: List[Prediction]


@app.post("/predict", response_model=Predictions, response_model_exclude_unset=True)
async def predict(request: Request):
    body = await request.json()

    logger.info(f"Received request: {body}")

    instances = body["instances"]
    instances = [x['text'] for x in instances]

    tf_batch = tokenizer(instances, padding=True,
                         truncation=True, return_tensors='pt')
    kw_model = KeyBERT()

    with torch.no_grad():
        outputs = model(**tf_batch)
        logits = outputs.logits
        probabilities = torch.nn.functional.softmax(logits, dim=1)

    emotion_labels = ["sadness", "joy", "love", "anger", "fear", "surprise"]
    scores = probabilities.tolist()  # Convert tensor to list

    outputs = []
    keyword_arr = []
    topic_arr = []

    for text in instances:
        keywords = kw_model.extract_keywords(
            text,
            keyphrase_ngram_range=(1, 1),
            stop_words="english",
            use_mmr=True,
            diversity=0.7
        )
        topic = [keyword[0] for keyword in keywords]

        topic_arr.append(topic)

    for text in instances:
        keywords = kw_model.extract_keywords(
            text,
            keyphrase_ngram_range=(1, 10),
            stop_words="english",
            use_mmr=True,
            diversity=0.7
        )
        keywords = [keyword[0] for keyword in keywords]
        keyword_arr.append(keywords)

    for i in range(len(scores)):
        emotion_probabilities = {label: score for label,
                                 score in zip(emotion_labels, scores[i])}

        outputs.append(Prediction(
            emotions=emotion_probabilities, keywords=keyword_arr[i], topics=topic_arr[i]))

    print(outputs)
    return Predictions(predictions=outputs)


@ app.get("/health", status_code=200)
def health():
    return {}


@ app.get("/")
async def read_root():
    return {"message": "Emotion classification API"}

if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=8080)
