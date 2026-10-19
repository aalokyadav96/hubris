from flask import Flask, jsonify, request
from flask_cors import CORS
from flask_sqlalchemy import SQLAlchemy
from flask_jwt_extended import (
    JWTManager, create_access_token, jwt_required, get_jwt_identity
)
from werkzeug.security import generate_password_hash, check_password_hash

app = Flask(__name__)

# -----------------------------------------------------------------------------
# Configuration
# -----------------------------------------------------------------------------
# PostgreSQL Connection URL format: postgresql://username:password@localhost:5432/database_name
app.config['SQLALCHEMY_DATABASE_URI'] = 'postgresql://postgres:password@localhost:5432/mydatabase'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

# Change this secret key in production
app.config['JWT_SECRET_KEY'] = 'super-secret-key-change-me' 

db = SQLAlchemy(app)
jwt = JWTManager(app)

# CORS updated to allow Authorization headers required for JWT
CORS(app, resources={r"/api/*": {"origins": "*"}}, allow_headers=["Content-Type", "Authorization"])

# -----------------------------------------------------------------------------
# Database Models
# -----------------------------------------------------------------------------
class User(db.Model):
    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), unique=True, nullable=False)
    password_hash = db.Column(db.String(255), nullable=False)

class Item(db.Model):
    __tablename__ = 'items'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(100), nullable=False)
    description = db.Column(db.Text, nullable=True)
    user_id = db.Column(db.Integer, db.ForeignKey('users.id'), nullable=False) # Owner of the item

# Create tables automatically (for dev testing)
with app.app_context():
    db.create_all()

# -----------------------------------------------------------------------------
# Authentication Endpoints
# -----------------------------------------------------------------------------

@app.route('/api/auth/register', methods=['POST'])
def register():
    data = request.get_json()
    if not data or 'username' not in data or 'password' not in data:
        return jsonify({"status": "error", "message": "Missing username or password"}), 400
        
    if User.query.filter_by(username=data['username']).first():
        return jsonify({"status": "error", "message": "Username already exists"}), 400

    hashed_password = generate_password_hash(data['password'])
    new_user = User(username=data['username'], password_hash=hashed_password)
    
    db.session.add(new_user)
    db.session.commit()
    
    return jsonify({"status": "success", "message": "User registered successfully"}), 201

@app.route('/api/auth/login', methods=['POST'])
def login():
    data = request.get_json()
    user = User.query.filter_by(username=data.get('username')).first()
    
    if not user or not check_password_hash(user.password_hash, data.get('password')):
        return jsonify({"status": "error", "message": "Invalid credentials"}), 401
        
    # Generate JWT token with the user ID as the identity
    access_token = create_access_token(identity=str(user.id))
    return jsonify({"status": "success", "access_token": access_token}), 200

# -----------------------------------------------------------------------------
# Protected Item Endpoints (Requires valid JWT)
# -----------------------------------------------------------------------------

# GET: Fetch all items belonging to the authenticated user
@app.route('/api/items', methods=['GET'])
@jwt_required()
def get_items():
    current_user_id = int(get_jwt_identity())
    user_items = Item.query.filter_by(user_id=current_user_id).all()
    
    output = [{"id": item.id, "name": item.name, "description": item.description} for item in user_items]
    return jsonify({"status": "success", "data": output}), 200

# POST: Create an item for the authenticated user
@app.route('/api/items', methods=['POST'])
@jwt_required()
def create_item():
    current_user_id = int(get_jwt_identity())
    data = request.get_json()
    
    if not data or 'name' not in data:
        return jsonify({"status": "error", "message": "Missing required field: name"}), 400
        
    new_item = Item(name=data['name'], description=data.get('description', ''), user_id=current_user_id)
    db.session.add(new_item)
    db.session.commit()
    
    return jsonify({
        "status": "success", 
        "data": {"id": new_item.id, "name": new_item.name, "description": new_item.description}
    }), 201

# DELETE: Remove an item (Only if it belongs to the logged-in user)
@app.route('/api/items/<int:item_id>', methods=['DELETE'])
@jwt_required()
def delete_item(item_id):
    current_user_id = int(get_jwt_identity())
    item = Item.query.filter_by(id=item_id, user_id=current_user_id).first()
    
    if not item:
        return jsonify({"status": "error", "message": "Item not found or unauthorized"}), 404
        
    db.session.delete(item)
    db.session.commit()
    return jsonify({"status": "success", "message": f"Item {item_id} successfully deleted"}), 200

if __name__ == '__main__':
    app.run(debug=True, port=5000)
