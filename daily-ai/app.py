from flask import Flask, request, jsonify
from package import responseWrapper
import time

app = Flask(__name__)

@app.route('/', methods=['POST'])
def receiveData():
    try:
        data = request.get_json()
        time.sleep(5)
        return responseWrapper.wrap(data.get('daily')), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 400


if __name__ == "__main__":
    app.run(debug=True)
