import re
from flask import Flask, request, render_template
import boto3
# import botocore.session

app = Flask(__name__)

@app.route('/', methods=['GET'])
def index():
  s = boto3.Session()
  c = s.client("sts")
  
  # session = botocore.session.get_session()
  # resp = {"profile": session.profile}
  asd = c.get_caller_identity()
  return render_template('index.html')

@app.route('/upload', methods=['GET', 'POST'])
def upload():
  result = request.form

  if 'file' not in request.files:
    print("asd")
  else:
    print("ss")
  return render_template('index.html')

if __name__ == '__main__':
	# app.run(host="0.0.0.0",port=8080)
  app.run(debug=True, port=5001)