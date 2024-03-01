from flask import Flask, request, jsonify
from package import responseWrapper
import json

app = Flask(__name__)

@app.route('/', methods=['POST'])
def receiveData():
    try:
        data = request.get_json()
        return responseWrapper.wrap(data.get('daily')), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 400


if __name__ == "__main__":
    app.run(debug=True)
