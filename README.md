# Backend Golang Coding Test

Build a simple RESTful API in Golang that manages a list of users. Use MongoDB for persistence, JWT for authentication, and follow clean code practices.


## Project Setup and Run Instructions

This section guides you through setting up and running the Backend Golang Coding Test on your local machine.

### Prerequisites

* Go: `Version 1.20 or higher`

### Installation

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/ThanakitK/7solutions-backend.git
    cd <your-project-directory>
    ```

2.  **Install dependencies:**
    ```
    go mod tidy
    ```

3.  **Environment Variables (if applicable):**
    Create a `.env` file in the root directory of the project based 
    ```
    # .env
    DB_URI = your_database_connection_string
    DB_NAME = your_database_name
    SIGNATURE_KEY = your_signature_key
    SIGNATURE_EXP = your_expire_time
    ```

### Running the Application

```
go run main.go
```
* **Specify the URL where the application will be accessible (e.g., `http://localhost:3000`).**

---

## JWT Token Usage Guide

This project uses JSON Web Tokens (JWT) for authentication and authorization.

### What is a JWT?

A JWT is a compact, URL-safe means of representing claims to be transferred between two parties. The claims in a JWT are encoded as a JSON object that is digitally signed using a secret (or a public/private key pair).

### Obtaining a JWT

To get a JWT, you'll first need to register a user, then authenticate them.


**Endpoint:** `POST /api/create-user` 

**Request Body Example:**

```json
{
    "name": "user",
    "email": "user@example.com",
    "password": "your_password"
}
```
**Reponse Body Example:**
``` json
{
    "status": true,
    "message": "create user success",
    "code": 201,
    "data": {
        "id": "your_id",
        "name": "user",
        "email": "user@example.com",
        "password": "your_hashpassword",
        "createAt": "your_local_time"
    }
}
```

**Endpoint:** `POST /api/signin` 

**Request Body Example:**

```json
{
    "email": "user@example.com",
    "password": "your_password"
}
```
**Reponse Body Example:**
``` json
{
    "status": true,
    "message": "sign in success",
    "code": 200,
    "data": {
        "type": "Bearer",
        "accessToken": "your_accesstoken"
    }
}
```

**Endpoint:** `GET /api/user/:id` 

**Authorization:** Bearer <your_jwt_token>

**Reponse Body Example:**
``` json
{
    "status": true,
    "message": "get user success",
    "code": 200,
    "data": {
        "id": "your_id",
        "name": "user",
        "email": "user@example.com",
        "password": "your_hashpassword",
        "createAt": "your_local_time"
    }
}
```

**Endpoint:** `GET /api/users` 

**Authorization:** Bearer <your_jwt_token>

**Reponse Body Example:**
``` json
{
    "status": true,
    "message": "get users success",
    "code": 200,
    "data": [
        {
             "id": "your_id",
            "name": "user",
            "email": "user@example.com",
            "password": "your_hashpassword",
            "createAt": "your_local_time"
        },
    ]
}
```

**Endpoint:** `PUT /api/user/:id` 

**Authorization:** Bearer <your_jwt_token>

**Request Body Example:**

```json
{
    "name": "your_modified_name",
    "email": "your_modified_email"
}
```
**Reponse Body Example:**
``` json
{
    "status": true,
    "message": "update user success",
    "code": 200,
    "data": {
        "id": "your_id",
        "name": "your_modified_name",
        "email": "your_modified_email",
        "password": "your_hashpassword",
        "createAt": "your_local_time"
    }
}
```
**Endpoint:** `DELETE /api/user/:id` 

**Authorization:** Bearer <your_jwt_token>

**Reponse Body Example:**
``` json
{
    "status": true,
    "message": "delete user success",
    "code": 200,
    "data": null
}
```

## Assumptions or Decisions Made


* **Technology Stack**: Go was chosen for its performance, concurrency features (goroutines), and strong type system, making it suitable for building efficient APIs. **MongoDB** was selected as the database for its flexibility with schema-less data and good integration with Go's official driver.
* **Authentication Strategy**: **JWT (HMAC HS256)** was implemented for stateless authentication, allowing for scalability and easy integration with client-side applications. The token contains the user's `id` and has a fixed expiration time.
* **Password Hashing**: User passwords are **hashed using bcrypt** before being stored in the database. This is a crucial security measure to protect user credentials.
* **Error Handling**: The API provides consistent **JSON error responses** with a clear `error` message for various failure scenarios (e.g., unauthorized, is required).
* **Middleware**:
    * **Authentication Middleware**: A dedicated middleware is used to validate JWTs for all protected routes, ensuring only authenticated requests can access sensitive endpoints.
* **Concurrency Task**: A background **goroutine** runs every 10 seconds to log the current number of users in the database. This demonstrates Go's concurrency capabilities and provides basic insights into data growth.
* **Database Interactions**: The official `go.mongodb.org/mongo-driver` is used for all MongoDB operations, ensuring robust and idiomatic interaction with the database.
* **User Model**: The `CreatedAt` field for the user model is automatically populated upon user creation.
* **Input Validation**: Basic input validation is performed for user registration and update requests to ensure necessary fields are present and in a valid format.