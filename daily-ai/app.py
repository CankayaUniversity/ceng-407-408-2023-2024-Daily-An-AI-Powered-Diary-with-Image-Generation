from flask import Flask
from package import responseWrapper

app = Flask(__name__)

"""To run this app 
1) python -m venv venv
2) .\venv\Scripts\activate
3) python -m pip install --upgrade pip
4) python -m pip install flask
5) python -m flask --app .\app.py run"""
@app.route("/")

def main():
    return responseWrapper.wrap()
