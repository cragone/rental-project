from flask import Flask, jsonify, request, send_file, g, render_template, make_response, current_app, send_from_directory
from flask_cors import CORS

app = Flask(__name__)

CORS(app)
# Define the payment data as a list of dictionaries
amounts_due = [
    { 'id': 1, 'due': 100, 'due_date': '2023-12-01', 'type': 'Utility' },
    { 'id': 2, 'due': 200, 'due_date': '2023-12-05', 'type': 'Rent' },
    { 'id': 3, 'due': 150, 'due_date': '2023-12-10', 'type': 'Late Fee' }
    # Add more data as needed
]

# Endpoint to fetch payments
@app.route('/api/payments', methods=['GET'])
def get_payments():
    user = request.args.get("user",None,str)
    return jsonify(amounts_due)

if __name__ == '__main__':
    app.run(debug=True)
