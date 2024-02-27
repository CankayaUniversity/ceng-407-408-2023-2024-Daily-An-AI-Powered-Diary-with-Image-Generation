from flask import Flask
from PIL import Image
import os
import base64
import json

app = Flask(__name__)

"""To run this app 
1) python -m venv venv
2) .\venv\Scripts\activate
3) python -m pip install --upgrade pip
4) python -m pip install flask
5) python -m flask --app .\app.py run"""
@app.route("/")

def aiServerResponse():
    sampleJson = {}
    emotions = ["curiosity", "excitement", "fear"]
    summary = "I am excited to implement the 'Flask' micro framework in my project. I'm not sure if it will work, but I'm looking forward to it."
    image_base64 = None

    try:
        imagePath = "./venv/Image/sampleImage.jpg"
        with open(imagePath, "rb") as img_file:
            img_data = img_file.read()
            image_base64 = base64.b64encode(img_data).decode('utf-8')
        print("Image loaded and converted to base64 successfully.")
    except FileNotFoundError:
        print("Image file not found at path: ", imagePath)

    sampleJson = {
        "emotions": emotions,
        "summary": summary,
        "imageToBase64": "base64img"
    }

    return sampleJson
