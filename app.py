import json
from flask import Flask, request


app = Flask(__name__)


@app.route('/nodes', methods=['GET', 'POST'])
def operations():
    if request.method == 'POST':
        node = request.get_json(silent=True)
        print(node)
        return json.dumps({
            'result': True,
        })
    else:
        return json.dumps({
            'result': True,
            'nodes': [],
        })


if __name__ == '__main__':
    app.run(host='0.0.0.0')
