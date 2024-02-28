from flask import Flask
from package import responseWrapper

app = Flask(__name__)

@app.route("/")

def main():
    return responseWrapper.wrap()
