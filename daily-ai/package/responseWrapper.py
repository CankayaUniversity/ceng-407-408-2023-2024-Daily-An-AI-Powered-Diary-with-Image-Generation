from . import emotion
from . import summary
from . import image

def wrap():
    json = {
        "emotions": emotion.getEmotions(),
        "summary": summary.getSummary(),
        "image": image.getImage()
    }

    return json
