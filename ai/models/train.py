import torch
import os
import numpy as np
import pandas as pd
import base64
from datasets import load_dataset, load_metric
from transformers import AutoTokenizer, RobertaForSequenceClassification, RobertaConfig, Trainer, TrainingArguments, DataCollatorWithPadding
from datetime import datetime
from io import BytesIO
from PIL import Image
from keybert import KeyBERT
from diffusers import DiffusionPipeline
from bertopic import BERTopic
from collections import Counter


class EmotionClassifier:
    def __init__(self, model_path=None):
        self.emotion_labels = ["sadness", "joy",
                               "love", "anger", "fear", "surprise"]
        self.tokenizer = AutoTokenizer.from_pretrained("roberta-base")
        config = RobertaConfig.from_pretrained("roberta-base", num_labels=6)
        if model_path:
            self.model = RobertaForSequenceClassification.from_pretrained(
                model_path, config=config
            )
        else:
            self.model = RobertaForSequenceClassification(config=config)
        self.data_collator = DataCollatorWithPadding(tokenizer=self.tokenizer)
        self.kw_model = KeyBERT()
        self.diffusion_model = DiffusionPipeline.from_pretrained(
            "stabilityai/stable-diffusion-xl-base-1.0",
            torch_dtype=torch.float16,
            use_safetensors=True,
            variant="fp16",
        )
        self.diffusion_model.to("cuda")

    def save_model(self, save_path="model/emotion_classifier"):
        self.model.save_pretrained(save_path)
        print(f"Model saved to {save_path}")

    def load_model(self, model_path="model/emotion_classifier"):
        if os.path.exists(model_path):
            self.model = RobertaForSequenceClassification.from_pretrained(
                model_path)
            print(f"Model loaded from {model_path}")
        else:
            print(f"Error: Model directory not found at {model_path}")

    def load_data(self, train_file, test_file):
        data_files = {"train": train_file, "test": test_file}
        dataset = load_dataset("csv", data_files=data_files)
        self.train_dataset = dataset["train"].map(
            self.tokenize_function, batched=True)
        self.test_dataset = dataset["test"].map(
            self.tokenize_function, batched=True)

    def tokenize_function(self, examples):
        return self.tokenizer(examples["text"], truncation=True)

    def train_model(self, output_dir="./model_save"):
        training_args = TrainingArguments(
            output_dir=output_dir,
            learning_rate=2e-5,
            per_device_train_batch_size=16,
            per_device_eval_batch_size=16,
            num_train_epochs=2,
            weight_decay=0.01,
            save_strategy="epoch",
            push_to_hub=False,
        )
        trainer = Trainer(
            model=self.model,
            args=training_args,
            train_dataset=self.train_dataset,
            eval_dataset=self.test_dataset,
            tokenizer=self.tokenizer,
            data_collator=self.data_collator,
            compute_metrics=self.compute_metrics,
        )
        trainer.train()
        trainer.evaluate()

    def compute_metrics(self, eval_pred):
        logits, labels = eval_pred
        predictions = np.argmax(logits, axis=-1)
        accuracy = load_metric("accuracy").compute(
            predictions=predictions, references=labels
        )
        f1 = load_metric("f1").compute(
            predictions=predictions, references=labels, average="micro"
        )
        return {"accuracy": accuracy["accuracy"], "f1": f1["f1"]}

    def generate_image(self, text):
        keywords = self.kw_model.extract_keywords(
            text,
            keyphrase_ngram_range=(1, 10),
            stop_words="english",
            use_mmr=True,
            diversity=0.7,
        )
        prompt = " ".join([keyword[0] for keyword in keywords])
        image = self.diffusion_model(prompt).images[0]
        return image

    def encode_image_to_base64(self, image):
        buffered = BytesIO()
        image.save(buffered, format="PNG")
        return base64.b64encode(buffered.getvalue()).decode()

    def analyze_text(self, text):
        encoding = self.tokenizer(
            text, truncation=True, padding=True, return_tensors="pt"
        ).to("cuda")
        with torch.no_grad():
            output = self.model(**encoding)
        probs = torch.softmax(output.logits, dim=1)
        probabilities = probs.cpu().numpy()[0]

        emotion_probabilities = {
            label: prob for label, prob in zip(self.emotion_labels, probabilities)
        }

        return emotion_probabilities

    def create_journal_entry(self, author, text, favourites, is_shared, viewers):
        emotions = self.analyze_text(text)

        image = self.generate_image(text)
        image_base64 = self.encode_image_to_base64(image)

        # Extract keywords using KeyBERT
        keywords = self.kw_model.extract_keywords(
            text,
            keyphrase_ngram_range=(1, 2),
            stop_words="english",
            use_mmr=True,
            diversity=0.7
        )
        keywords = [keyword[0] for keyword in keywords]
        # Create and return the journal entry
        journal_entry = {
            "author": author,
            "text": text,
            "keywords": keywords,
            "emotions": emotions,
            "favourites": favourites,
            "createdAt": datetime.now().isoformat(),
            "isShared": is_shared,
            "viewers": viewers,
            "image": image_base64,
        }
        return journal_entry


class TopicExtractor:
    def __init__(self, model_path=None):
        if model_path:
            self.topic_model = BERTopic.load(model_path)
        else:
            self.topic_model = BERTopic(
                embedding_model="all-MiniLM-L6-v2", language="english")

    def extract_topics(self, text):
        topics, _ = self.topic_model.transform([text])
        return topics[0]

    def extract_topics_from_journal_entries(self, journal_entries):
        all_topics = []
        for entry in journal_entries:
            topics = self.extract_topics(extry['text'])
            all_topics.extend(topics)
        return all_topics

    def frequent_topics(self, journal_entries, num_topics=5):
        topics = self.extract_topics_from_journal_entries(journal_entries)
        topic_counts = Counter(topics)
        return topic_counts.most_common(num_topics)

    def update_model(self, texts):
        self.topic_model.fit(texts)

    def save_model(self, save_path=None):
        if save_path:
            self.topic_model.save(save_path)
        else:
            self.topic_model.save("./model/topic_model")


if __name__ == "__main__":
    # Load data
    classifier = EmotionClassifier()
    classifier.load_data("train.csv", "test.csv")

    # Train model
    classifier.train_model()

    # Save the model
    classifier.save_model()

    author = "65ddd1faa21e8586ed7b3a32"
    text = "This is a test journal entry."
    favourites = 0
    is_shared = True
    viewers = []
    journal_entry = classifier.create_journal_entry(
        author, text, favourites, is_shared, viewers)
    print(journal_entry)

