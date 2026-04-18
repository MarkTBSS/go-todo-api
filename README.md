# Go Todo API

CRUD API for managing todos, built with Golang, Gin, and GORM.

## Tech Stack

- **Go** — Programming language
- **Gin** — HTTP web framework
- **GORM** — ORM for database operations
- **PostgreSQL** — Relational database
- **Docker & Docker Compose** — Containerization

## Database Design

### Table: todos

| Column      | Type         | Description                     |
|-------------|--------------|---------------------------------|
| id          | BIGINT (PK)  | Auto increment                  |
| title       | VARCHAR(255) | Todo title (required)           |
| description | TEXT         | Detailed description (optional) |
| status      | VARCHAR(20)  | pending / in_progress / completed |
| priority    | VARCHAR(10)  | low / medium / high             |
| due_date    | TIMESTAMP    | Due date (optional)             |
| created_at  | TIMESTAMP    | Auto-generated                  |
| updated_at  | TIMESTAMP    | Auto-generated                  |

## API Endpoints

| Method | Endpoint         | Description                              |
|--------|------------------|------------------------------------------|
| POST   | /api/todos       | Create a new todo                        |
| GET    | /api/todos       | List all todos (filter: status, priority)|
| GET    | /api/todos/:id   | Get a todo by ID                         |
| PUT    | /api/todos/:id   | Update a todo                            |
| DELETE | /api/todos/:id   | Delete a todo                            |

## How to Run Locally

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop/) installed and running

### Run with Docker Compose

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`.

### Run Unit Tests

```bash
go test ./handlers/ -v
```

## CI/CD Pipeline

This project uses **GitHub Actions** for automated testing and deployment.

**Pipeline flow:** Push to `main` → Run Tests → Build Docker Image → Push to ECR → Deploy to ECS

See pipeline runs: [GitHub Actions](https://github.com/MarkTBSS/go-todo-api/actions)

## AWS Architecture

```
Internet → ALB → ECS Fargate (API) → RDS PostgreSQL
```

| Service | Purpose |
|---------|---------|
| ECS Fargate | Run API container |
| RDS PostgreSQL | Database |
| ALB | Public endpoint |
| ECR | Docker image registry |
| GitHub Actions | CI/CD pipeline |

## API Usage Examples

### Create a Todo
```bash
curl -X POST http://localhost:8080/api/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Buy groceries","priority":"high","due_date":"2026-04-20"}'
```

### List All Todos
```bash
curl http://localhost:8080/api/todos
```

### Filter by Status
```bash
curl "http://localhost:8080/api/todos?status=pending"
```

### Get Todo by ID
```bash
curl http://localhost:8080/api/todos/1
```

### Update a Todo
```bash
curl -X PUT http://localhost:8080/api/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"status":"completed"}'
```

### Delete a Todo
```bash
curl -X DELETE http://localhost:8080/api/todos/1
```
