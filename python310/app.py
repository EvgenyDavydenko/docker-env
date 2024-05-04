from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello():
    return "<p>Hello, World!</p>"

# main driver function
if __name__ == "__main__":
    app.run(debug=True, host='python310', port=5000)
