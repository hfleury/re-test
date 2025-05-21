# re-test
Henry code test for RE Partners.

## Folder Structure

This project folder structure follows a **Layered Structure**.

### Layered Structure
In this approach, the application is organized into distinct layers: `handlers` (or controllers), `services`, `repositories`, and `models`. Each layer is responsible for a specific concern, promoting separation of concerns and maintainability.

📚 Inspired by this [Medium article on Go project folder structures](https://medium.com/@smart_byte_labs/organize-like-a-pro-a-simple-guide-to-go-project-folder-structures-e85e9c1769c2).

---

## Testing the backend

### ✅ Prerequisites

- [Go installed](https://go.dev/doc/install) (version 1.23 or later recommended)

### 🚀 Run the Server

```bash
go mod tidy
go run cmd/main.go
```

### Testing
The server will start on http://localhost:8081
Now you need to send a request

```bash
curl -X GET "http://localhost:8081/calculate?order_amount=12001" -H "Accept: application/json"
```

## Testing using the UI

### ✅ Prerequisites

- [Node.js and npm installed](https://nodejs.org/)

### 🚀 Run the Client

```bash
cd web
npm install
npm start
```

### Testing
The frontend will be available on (http://localhost:3000)[http://localhost:3000] and will communicate with the backend API running on port 8081.


## Testing using Docker

### ✅ Prerequisites

- [Docker and Docker Compse installed](https://docs.docker.com/compose/install/)

### 🚀 Start the containers

```bash
docker compose up
```