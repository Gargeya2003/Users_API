Users API (Go + Fiber + MySQL)

A RESTful API built with Go to manage users with their name and date of birth (DOB).
The API dynamically calculates and returns the age of users when fetching details.

This project demonstrates clean backend architecture, SQLC-based database access, input validation, structured logging, pagination, and unit testing.

🚀 Features

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
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
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

cd db\sqlc
sqlc generate


This generates type-safe Go code for all SQL queries.

6️⃣ Run the Application
go run ./cmd/server


You should see logs similar to:

starting users API
database connected
server listening port=8080

API Examples (curl + PowerShell)
## 🧪 API Usage Examples

➕ Create User

curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","dob":"1990-05-10"}'

Windows PowerShell command

Invoke-RestMethod `
  -Uri http://localhost:8080/users `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Alice","dob":"1990-05-10"}'

📋 List Users (Pagination)

curl "http://localhost:8080/users?limit=5&offset=0"

Windows PowerShell command

Invoke-RestMethod "http://localhost:8080/users?limit=5&offset=0"

📄 Get User by ID

curl http://localhost:8080/users/1

Windows PowerShell command

Invoke-RestMethod http://localhost:8080/users/1

✏️ Update User

curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","dob":"1991-03-15"}'

  Windows PowerShell command

  Invoke-RestMethod `
  -Uri http://localhost:8080/users/1 `
  -Method PUT `
  -ContentType "application/json" `
  -Body '{"name":"Alice Updated","dob":"1991-03-15"}'


❌ Delete User
curl -X DELETE http://localhost:8080/users/1

Windows PowerShell command

Invoke-RestMethod `
  -Uri http://localhost:8080/users/1 `
  -Method DELETE

🧪 Testing
Unit Test (Age Calculation)

Run:

go test ./internal/service


Expected output:

ok   users_api/internal/service

This verifies the correctness of the age calculation logic.

📌 Notes

Age is not stored in the database — it is calculated dynamically.

Auto-increment IDs are not reused after deletion (expected DB behavior).
