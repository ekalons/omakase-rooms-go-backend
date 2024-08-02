# omakase-rooms-go-backend

A GoLang backend service with clean architecture for the Omakase Rooms project.

## How to run locally

### Prerequisites

- GoLang

### Steps

1. Clone the repository
2. Install Go
3. Add the required environment variables in a `.env` file in the root directory. The required variables can be found in the `.env.example` file.
4. Run the following command to start the server

```bash
go run cmd/main.go
```

5. The server will start on port 8080

## API Documentation

### Get all rooms

```http
GET /rooms
```

### Get room by ID

```http
GET /room/{id}
```

### Create a room (Auth)

```http
POST /room
```

### Get token (Auth)

```http
POST /token
```
