import base64

def getImage():
    base64Image = None
    try:
        imagePath = "./venv/Image/sampleImage.jpg"
        with open(imagePath, "rb") as imgFile:
            img = imgFile.read()
            base64Image = "data:image/png;base64, " + base64.b64encode(img).decode('utf-8')
        print("Image loaded and converted to base64 succesfully.")
    except FileNotFoundError:
        print("Image file not found at path: ", imagePath)

    return base64Image
