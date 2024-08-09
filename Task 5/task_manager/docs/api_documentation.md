# Task Management API Documentation

## Overview
This document provides an overview and instructions for the Task Management REST API developed using Go and the Gin framework. The API supports basic CRUD operations for managing tasks and now includes persistent data storage using MongoDB.

## Objective
The objective of this task is to extend the existing Task Management API with persistent data storage using MongoDB and the Mongo Go Driver. This enhancement replaces the in-memory database with MongoDB to provide data persistence across API restarts.

## Requirements
- Integrate MongoDB as the persistent data storage solution for the Task Management API.
- Use the Mongo Go Driver to interact with the MongoDB database from the Go application.
- Update the existing API endpoints to perform CRUD operations using MongoDB for data storage.
- Ensure proper error handling and validation of MongoDB operations.
- Validate the correctness of data stored in MongoDB by retrieving and verifying task information.
- Update the documentation to include instructions for configuring and connecting to MongoDB.
- Test the API endpoints to verify that data is stored and retrieved correctly from MongoDB.
- Ensure that the API remains backward compatible with the previous version, maintaining the same endpoint structure and behavior.

## Instructions

### Setting Up MongoDB
1. **Install MongoDB:**
   - Follow the instructions for your operating system on the [MongoDB installation guide](https://docs.mongodb.com/manual/installation/).

2. **Run MongoDB:**
   - Start the MongoDB server using the command: `mongod`.

3. **Install Mongo Go Driver:**
   - If not already installed, add the Mongo Go Driver package to your project:
     ```sh
     go get go.mongodb.org/mongo-driver
     ```

### Folder Structure
```
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```
- **main.go:** Entry point of the application.
- **controllers/task_controller.go:** Handles incoming HTTP requests and invokes the appropriate service methods.
- **models/task.go:** Defines the data structures used in the application.
- **data/task_service.go:** Contains business logic and data manipulation functions. Implement the ORM/ODM code here.
- **router/router.go:** Sets up the routes and initializes the Gin router and defines the routing configuration for the API.
- **docs/api_documentation.md:** Contains API documentation and other related documentation.
- **go.mod:** Defines the module and its dependencies.

## API Endpoints
The API provides the following endpoints:

### Get All Tasks
- **URL:** `/tasks`
- **Method:** `GET`
- **Description:** Retrieves a list of all tasks.
- **Response:**
  ```json
  [
    {
      "id": 1,
      "title": "Task 1",
      "description": "Description for task 1",
      "due_date": "2024-08-15",
      "status": "Pending"
    },
    ...
  ]
  ```

### Get Task by ID
- **URL:** `/tasks/:id`
- **Method:** `GET`
- **Description:** Retrieves the details of a specific task.
- **Response:**
  ```json
  {
    "id": 1,
    "title": "Task 1",
    "description": "Description for task 1",
    "due_date": "2024-08-15",
    "status": "Pending"
  }
  ```

### Create Task
- **URL:** `/tasks`
- **Method:** `POST`
- **Description:** Creates a new task.
- **Request Body:**
  ```json
  {
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-20",
    "status": "Pending"
  }
  ```
- **Response:**
  ```json
  {
    "id": 4,
    "title": "New Task",
    "description": "Description of the new task",
    "due_date": "2024-08-20",
    "status": "Pending"
  }
  ```

### Update Task
- **URL:** `/tasks/:id`
- **Method:** `PUT`
- **Description:** Updates an existing task.
- **Request Body:**
  ```json
  {
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-25",
    "status": "Completed"
  }
  ```
- **Response:**
  ```json
  {
    "id": 1,
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-25",
    "status": "Completed"
  }
  ```

### Delete Task
- **URL:** `/tasks/:id`
- **Method:** `DELETE`
- **Description:** Deletes a specific task.
- **Response:** `204 No Content`

## MongoDB Integration

### Configuration
The MongoDB connection is configured in the `data/task_service.go` file. The connection is initialized during the package initialization:

```go
var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("task_manager").Collection("tasks")
}
```

### CRUD Operations with MongoDB
The following functions interact with MongoDB to perform CRUD operations:

#### Get All Tasks
```go
func GetTasks() []models.Task {
	var tasks []models.Task
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &tasks); err != nil {
		log.Fatal(err)
	}
	return tasks
}
```

#### Get Task by ID
```go
func GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	filter := bson.M{"id": id}
	err := collection.FindOne(ctx, filter).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}
```

#### Create Task
```go
func CreateTask(task models.Task) (models.Task, error) {
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	return task, nil
}
```

#### Update Task
```go
func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"dueDate":     updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}
	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if result.Err() != nil {
		return nil, result.Err()
	}
	var task models.Task
	if err := result.Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}
```

#### Delete Task
```go
func DeleteTask(id int) error {
	filter := bson.M{"id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
```

## Testing
Use Postman or similar tools to test the API endpoints. Verify that tasks can be created, retrieved, updated, and deleted successfully. Additionally, use MongoDB Compass or direct MongoDB queries to verify the correctness of data stored in MongoDB.
