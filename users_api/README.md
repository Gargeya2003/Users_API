Users API (Go + Fiber + MySQL)

A RESTful API built with Go to manage users with their name and date of birth (DOB).
The API dynamically calculates and returns the age of users when fetching details.

This project demonstrates clean backend architecture, SQLC-based database access, input validation, structured logging, pagination, and unit testing.

Features

CRUD operations for users

Dynamic age calculation (no age stored in DB)

Pagination support for listing users

Input validation using go-playground/validator

Structured logging using Uber Zap

Request ID injection and request duration logging

Clean layered architecture (Handler → Service → Repository)

Unit tests for business logic

🧱 Tech Stack

Go

GoFiber (HTTP framework)

MySQL

SQLC (type-safe SQL access)

Uber Zap (logging)

go-playground/validator (input validation)

📁 Project Structure
users_api/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── config.go
├── db/
│   ├── migrations/
│   └── sqlc/
│       ├── queries.sql
│       └── generated/
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── routes/
│   ├── middleware/
│   ├── models/
│   └── logger/
├── go.mod
├── go.sum
└── README.md

🗄️ Database Schema

Table: users

Field	Type	Constraints
id	SERIAL	PRIMARY KEY
name	TEXT	NOT NULL
dob	DATE	NOT NULL
⚙️ Prerequisites

Make sure you have the following installed:

Go (1.20+ recommended)

MySQL (8+ recommended)

SQLC

Git

🧩 Step-by-Step Setup
1️⃣ Clone the Repository
git clone <your-repo-url>
cd users_api

2️⃣ Create the Database

Log into MySQL:

CREATE DATABASE usersdb;

3️⃣ Run Database Migrations

Create the users table:

CREATE TABLE users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);

4️⃣ Configure Environment Variables

Set the database connection string:

Windows (PowerShell)
$env:DB_DSN="root:yourpassword@tcp(localhost:3306)/usersdb?parseTime=true"
$env:PORT="8080"
$env:ENV="development"

macOS / Linux
export DB_DSN="root:yourpassword@tcp(localhost:3306)/usersdb?parseTime=true"
export PORT=8080
export ENV=development

5️⃣ Generate SQLC Code

From the project root:

sqlc generate


This generates type-safe Go code for all SQL queries.

6️⃣ Run the Application
go run ./cmd/server


You should see logs similar to:

starting users API
database connected
server listening port=8080

📡 API Endpoints
➕ Create User

POST /users

{
  "name": "Alice",
  "dob": "1990-05-10"
}

📄 Get User by ID

GET /users/:id

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

📋 List Users (Paginated)

GET /users?limit=10&offset=0

[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
  }
]

✏️ Update User

PUT /users/:id

{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

❌ Delete User

DELETE /users/:id

Response: 204 No Content

🧪 Testing
Unit Test (Age Calculation)

Run:

go test ./internal/service


Expected output:

ok   users_api/internal/service


This verifies the correctness of the age calculation logic.

🪵 Logging & Middleware

The API includes middleware that:

Injects a unique X-Request-ID header for each request

Logs request method, path, status, duration, and request ID using Zap

Example log:

{
  "level": "info",
  "msg": "http request",
  "request_id": "uuid",
  "method": "GET",
  "path": "/users",
  "status": 200,
  "duration": "120ms"
}

🧠 Architecture Overview
Client
  ↓
Routes
  ↓
Handlers (HTTP + validation)
  ↓
Services (business logic)
  ↓
Repositories (DB access)
  ↓
SQLC (type-safe SQL)
  ↓
MySQL


This separation ensures:

Maintainability

Testability

Scalability

Clean codebase

✅ Summary

This project demonstrates how to build a production-ready REST API in Go with:

Clean architecture

Strong typing with SQLC

Proper validation and logging

Pagination and unit testing

Industry-standard backend practices

📌 Notes

Age is not stored in the database — it is calculated dynamically.

Auto-increment IDs are not reused after deletion (expected DB behavior).

Docker is intentionally not included in this setup.