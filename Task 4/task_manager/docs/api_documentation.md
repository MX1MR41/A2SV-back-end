# Task Management REST API Documentation

## Overview

This Task Management REST API provides endpoints to create, read, update, and delete tasks. The API is built using Go and the Gin framework and uses an in-memory database to store task data.

## Endpoints

### GET /tasks

Retrieve a list of all tasks.

- **URL**: `/tasks`
- **Method**: `GET`
- **Response**: 
  - `200 OK`: Returns a list of tasks.
  ```json
  [
    {
      "id": 1,
      "title": "Task 1",
      "description": "Description for task 1",
      "due_date": "2024-08-15",
      "status": "Pending"
    },
    {
      "id": 2,
      "title": "Task 2",
      "description": "Description for task 2",
      "due_date": "2024-08-16",
      "status": "Completed"
    }
  ]
  ```

### GET /tasks/:id

Retrieve details of a specific task.

- **URL**: `/tasks/:id`
- **Method**: `GET`
- **URL Params**: 
  - `id` (integer): The ID of the task to retrieve.
- **Response**: 
  - `200 OK`: Returns the task details.
  - `400 Bad Request`: Invalid task ID.
  - `404 Not Found`: Task not found.
  ```json
  {
    "id": 1,
    "title": "Task 1",
    "description": "Description for task 1",
    "due_date": "2024-08-15",
    "status": "Pending"
  }
  ```

### POST /tasks

Create a new task.

- **URL**: `/tasks`
- **Method**: `POST`
- **Request Body**:
  - `title` (string): The title of the task.
  - `description` (string): The description of the task.
  - `due_date` (string): The due date of the task.
  - `status` (string): The status of the task.
  ```json
  {
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-20",
    "status": "Pending"
  }
  ```
- **Response**: 
  - `201 Created`: Returns the created task.
  - `400 Bad Request`: Invalid request body.
  ```json
  {
    "id": 4,
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-20",
    "status": "Pending"
  }
  ```

### PUT /tasks/:id

Update a specific task.

- **URL**: `/tasks/:id`
- **Method**: `PUT`
- **URL Params**: 
  - `id` (integer): The ID of the task to update.
- **Request Body**:
  - `title` (string, optional): The title of the task.
  - `description` (string, optional): The description of the task.
  - `due_date` (string, optional): The due date of the task.
  - `status` (string, optional): The status of the task.
  ```json
  {
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-21",
    "status": "Completed"
  }
  ```
- **Response**: 
  - `200 OK`: Returns the updated task.
  - `400 Bad Request`: Invalid task ID or request body.
  - `404 Not Found`: Task not found.
  ```json
  {
    "id": 1,
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-21",
    "status": "Completed"
  }
  ```

### DELETE /tasks/:id

Delete a specific task.

- **URL**: `/tasks/:id`
- **Method**: `DELETE`
- **URL Params**: 
  - `id` (integer): The ID of the task to delete.
- **Response**: 
  - `204 No Content`: Task successfully deleted.
  - `400 Bad Request`: Invalid task ID.
  - `404 Not Found`: Task not found.

## Error Handling

The API uses standard HTTP status codes to indicate the success or failure of an API request. Error responses include a message explaining the reason for the error.

- `400 Bad Request`: Returned when the client provides invalid data.
- `404 Not Found`: Returned when the requested resource is not found.
- `500 Internal Server Error`: Returned when an unexpected error occurs on the server.

## Testing

You can use Postman or curl to test the API endpoints. Below are example curl commands for testing each endpoint:

### GET /tasks
```bash
curl -X GET http://localhost:8080/tasks
```

### GET /tasks/:id
```bash
curl -X GET http://localhost:8080/tasks/1
```

### POST /tasks
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-20",
    "status": "Pending"
  }'
```

### PUT /tasks/:id
```bash
curl -X PUT http://localhost:8080/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-21",
    "status": "Completed"
  }'
```

### DELETE /tasks/:id
```bash
curl -X DELETE http://localhost:8080/tasks/1
```

## Running the Application

To run the application, use the following commands:

1. Navigate to the project directory:
   ```bash
   cd task_manager
   ```

2. Build and run the application:
   ```bash
   go run main.go
   ```

The API will be accessible at `http://localhost:8080`.