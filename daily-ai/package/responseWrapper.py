from . import emotion
from . import summary
from . import image

def wrap(daily):
    print("daily: ", daily)
    json = {
        "emotions": emotion.getEmotions(),
        "summary": summary.getSummary(),
        "image": image.getImage()
    }

    return json
