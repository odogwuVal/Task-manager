# Simple Task Manager API

This is a simple Task Manager API built using the Gin framework in Golang. It allows you to manage tasks, including creating, retrieving, updating, and deleting tasks.

## Features

- Create a new task
- Retrieve a list of tasks
- Retrieve a single task by ID
- Update a task by ID
- Delete a task by ID

## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.
- [Git](https://git-scm.com/downloads) installed on your machine.

### Installing

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/task-manager-api.git
   cd task-manager-api

2. Install Dependencies:

    ```go mod download```

3. Run The application:

    ```go run main.go```

The API will be available at http://localhost:8080

### Example Usage

```
curl -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{
        "title": "New Task",
        "description": "Description of the new task",
        "due_date": "2024-07-01T12:00:00Z",
        "status": "Pending"
      }'
