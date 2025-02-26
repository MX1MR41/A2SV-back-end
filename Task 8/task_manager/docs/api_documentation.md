# Task Management API Documentation

## Overview
The Task Management API facilitates the management of tasks and users within a system. It incorporates JWT-based authentication and authorization, supporting role-based access control. This document provides detailed information on the available endpoints, including user registration, login, task management, and user management.

## Table of Contents
- [Authentication and Authorization](#authentication-and-authorization)
  - [User Registration](#post-register)
  - [User Login](#post-login)
  - [Promote User](#post-userspromoteid)
  - [Usage of Protected Endpoints](#usage-of-protected-endpoints)
- [Task Management](#task-management)
  - [Get All Tasks](#get-tasks)
  - [Get Task by ID](#get-tasksid)
  - [Create Task](#post-tasks)
  - [Update Task](#put-tasksid)
  - [Delete Task](#delete-tasksid)
- [User Management](#user-management)
  - [Get All Users](#get-users)
- [Folder Structure](#folder-structure)
- [Security Considerations](#security-considerations)
- [Testing](#testing)

## Authentication and Authorization

### Overview
This API uses JWT (JSON Web Tokens) for authentication and authorization. Users must authenticate to access most endpoints. Admin users have elevated permissions, allowing them to perform additional operations like creating, updating, and deleting tasks, as well as promoting other users.

### Instructions for User Registration, Login, and Usage of Protected Endpoints

#### 1. User Registration
- **Endpoint:** `POST /register`
- **Description:** Registers a new user account with a unique username and password.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  - **201 Created:** User created successfully.
  - **400 Bad Request:** Invalid payload.

#### 2. User Login
- **Endpoint:** `POST /login`
- **Description:** Authenticates a user and generates a JWT token upon successful login.
- **Request Body:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  - **200 OK:** Returns a JWT token on successful login.
  - **400 Bad Request:** Invalid credentials or payload.

  **Example Response:**
  ```json
  {
    "message": "Successfully logged in",
    "token": "string"
  }
  ```

#### 3. Promote User
- **Endpoint:** `POST /users/promote/:id`
- **Description:** Promotes a user to the admin role. This endpoint is only accessible to admin users.
- **URL Parameter:**
  - **id:** The ID of the user to be promoted.
- **Response:**
  - **200 OK:** User promoted successfully.
  - **400 Bad Request:** Invalid user ID.
  - **401 Unauthorized:** Unauthorized action.

### Usage of Protected Endpoints
- **Authentication Header:**
  - All requests to protected endpoints must include an `Authorization` header with the format:
    ```
    Authorization: Bearer <JWT_TOKEN>
    ```
- **Role-Based Access:**
  - **Admin Role:** Admins can create, update, and delete tasks, as well as promote users.
  - **User Role:** Regular users can retrieve tasks and retrieve task details by ID.

### JWT Token Claims
- The JWT token includes the following claims:
  - **username:** The username of the authenticated user.
  - **role:** The role of the user (e.g., `admin`, `user`).

## Task Management

### GET /tasks
- **Description:** Retrieves all tasks. Accessible by both admins and regular users.
- **Response:**
  - **200 OK:** Returns an array of tasks.

### GET /tasks/:id
- **Description:** Retrieves a task by its ID. Accessible by both admins and regular users.
- **URL Parameter:**
  - **id:** The ID of the task.
- **Response:**
  - **200 OK:** Returns the task.
  - **400 Bad Request:** Invalid task ID.
  - **404 Not Found:** Task not found.

### POST /tasks
- **Description:** Creates a new task. Only accessible by admin users.
- **Request Body:**
  ```json
  {
    "title": "string",
    "description": "string",
	"due_date": "string",
	"status": "string"
  }
  ```
- **Response:**
  - **201 Created:** Task created successfully.
  - **400 Bad Request:** Invalid payload.
  - **401 Unauthorized:** Unauthorized access.

### PUT /tasks/:id
- **Description:** Updates an existing task by its ID. Only accessible by admin users.
- **URL Parameter:**
  - **id:** The ID of the task to be updated.
- **Request Body:**
  ```json
  {
    "title": "string",
    "description": "string",
	"due_date": "string",
	"status": "string"
  }
  ```
- **Response:**
  - **200 OK:** Task updated successfully.
  - **400 Bad Request:** Invalid task ID or payload.
  - **404 Not Found:** Task not found.
  - **401 Unauthorized:** Unauthorized access.

### DELETE /tasks/:id
- **Description:** Deletes a task by its ID. Only accessible by admin users.
- **URL Parameter:**
  - **id:** The ID of the task to be deleted.
- **Response:**
  - **204 No Content:** Task deleted successfully.
  - **400 Bad Request:** Invalid task ID.
  - **404 Not Found:** Task not found.
  - **401 Unauthorized:** Unauthorized access.

## User Management

### GET /users
- **Description:** Retrieves all users. Only accessible by admin users.
- **Response:**
  - **200 OK:** Returns an array of users.

## Folder Structure

```plaintext
task_manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── database.go
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go
    └── user_usecases.go

```

**Key Files:**
- **main.go:** Entry point of the application.
- **controllers/task_controller.go & controllers/user_controller.go:** Manage the flow between the client requests and the data handling logic.
- **models/task.go & models/user.go:** Define the schema for tasks and users respectively.
- **data/task_repository.go & data/user_repository.go:** Contain data access logic for tasks and users.
- **data/database.go:** Manages the connection to MongoDB.
- **middleware/auth_middleware.go:** Provides middleware for handling authentication and authorization using JWT.
- **router/router.go:** Sets up the routing configuration for the API.
- **docs/api_documentation.md:** This document, detailing all available endpoints and how to interact with them.

### Security Considerations
- User passwords are hashed using a secure hashing algorithm before storage.
- JWT tokens are signed using a secure secret key to prevent tampering.
- Ensure that the secret key used for signing JWTs is kept secure and is not hard-coded in the source code.

## Testing
Use Postman or similar tools to test the API endpoints. Verify that:
- Users can register and log in successfully.
- JWT tokens are generated and validated correctly.
- Only authenticated users can access protected routes.
- Access control rules are enforced based on user roles.
- Tasks can be managed according to the permissions of the authenticated user.