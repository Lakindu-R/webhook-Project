# 🚀 Simple Webhook System (Go Lang)

A lightweight **Webhook Receiver System** built using Go (Golang).
This project demonstrates how real-world webhook systems work using a simple backend API.

---

## 📌 What is a Webhook?

A webhook is a method where one system sends real-time data to another system when an event occurs.

Instead of constantly checking for updates (polling), the server automatically sends data when something happens.

---

## 🔔 Example Use Cases

* 💳 Payment successful → notify application
* 🛒 Order placed → notify warehouse system
* 🧑‍💻 GitHub push → trigger CI/CD pipeline

---

## ⚙️ Project Features

* Accepts incoming webhook requests (POST API)
* Parses JSON payload
* Stores events in memory
* Retrieves all events via GET API
* Simple and lightweight Go server

---

## 🧱 Project Structure

```
webhook-project/
│
├── main.go
├── go.mod
│
├── models/
│   └── event.go
│
├── handlers/
│   └── webhook.go
│
└── storage/
    └── memory.go
```

---

## 🚀 How to Run the Project

### 1. Clone the repository

```
git clone https://github.com/your-username/webhook-project.git
cd webhook-project
```

### 2. Initialize Go modules

```
go mod init webhook-project
```

### 3. Run the server

```
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### 🔹 Receive Webhook (POST)

```
POST /webhook
```

**Request Body:**

```json
{
  "id": 1,
  "message": "Payment Success"
}
```

**Response:**

```json
{
  "status": "success",
  "message": "Webhook received"
}
```

---

### 🔹 Get All Events (GET)

```
GET /events
```

**Response:**

```json
[
  {
    "id": 1,
    "message": "Payment Success"
  }
]
```

---

## 🧠 Technologies Used

* Go (Golang)
* net/http package
* JSON encoding/decoding
* In-memory storage (slice)

---

## 🌍 Real-World Applications

This project simulates real webhook systems used in:

* Stripe / PayPal payment notifications
* GitHub webhook events
* E-commerce order systems
* CI/CD automation pipelines

---

## 📈 Future Improvements

* Add database (MySQL / PostgreSQL)
* Add webhook security (secret verification)
* Add logging system
* Use Gin framework
* Dockerize the application

---

## 👨‍💻 Author

Built for learning purposes to understand Webhooks and backend systems using Go.
