# omakase-rooms-go-backend

A GoLang + Docker backend service with clean architecture for the Omakase Rooms project.

## How to run locally

### Prerequisites

1. Clone the repository
2. Install Go
3. Add the required environment variables in a `.env` file in the root directory. The required variables can be found in the `.env.example` file.

### Without Docker

Run the following command to start the server

```bash
go run cmd/main.go
```

### With Docker

Install docker and run the following commands

```
docker build -t omakase-rooms-go-backend -f Dockerfile .
```

```
docker run --rm -p 8080:8080 omakase-rooms-go-backend
```

### Running the multistage version with Docker Compose

```
docker-compose down
```

```
docker-compose up -d
```

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

## Deployment

```
docker build -t omakase-rooms-go-backend .
```

or run the lightweight multistage build

```
docker build -t omakase-rooms-go-backend:multistage -f Dockerfile.multistage .
```
