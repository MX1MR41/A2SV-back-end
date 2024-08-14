# Task Management API Test Documentation

## Overview

This document details the comprehensive unit test suite implemented for the Task Management API using the `testify` library. The purpose of these tests is to ensure the correctness, stability, and reliability of the API's components, including domain models, use cases, controllers, and infrastructure.

## Table of Contents
- [Test Suite Setup](#test-suite-setup)
- [Mocking Dependencies](#mocking-dependencies)
- [Test Cases](#test-cases)
  - [Domain Models](#domain-models)
  - [Use Cases](#use-cases)
  - [Controllers](#controllers)
  - [Infrastructure](#infrastructure)
- [Test Coverage](#test-coverage)
- [Running Tests](#running-tests)
- [Continuous Integration](#continuous-integration)
- [Issues Encountered](#issues-encountered)
- [Conclusion](#conclusion)

## Test Suite Setup

The test suite for the Task Management API is implemented using the `testify` library. The tests are organized within the `Tests` directory and utilize the following structure:

```plaintext
task_manager/
└── Tests/
    ├── Mocks/
    │   ├── mock_task_repository.go
    │   ├── mock_user_repository.go
    │   ├── mock_task_usecases.go
    │   └── mock_user_usecases.go
    ├── controller_test.go
    ├── domain_test.go
    ├── infrastructure_test.go
    ├── user_usecases_test.go
    └── task_usecases_test.go
```

### Setup and Teardown Procedures

Each test suite includes setup and teardown procedures to ensure a clean test environment. This is handled within the `SetupTest` and `TearDownTest` methods provided by the `testify/suite` package.

## Mocking Dependencies

Mocks are used to isolate components and ensure that unit tests are independent and reproducible. The `Mocks` directory contains mock implementations of repositories and use cases, allowing for thorough testing of controllers and use cases without relying on external dependencies like the database.

### Example of Mocking

Here’s a snippet from `mock_task_repository.go`:

```go
package Mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}
```

## Test Cases

### Domain Models

Tests for domain models are focused on ensuring that the data structures are correctly defined and behave as expected. Example test cases include:

- **Task Model Validation:** Ensures that all fields in the `Task` struct are correctly populated and retrievable.
- **User Model Validation:** Verifies the integrity of the `User` struct, including role assignments.

### Use Cases

Use case tests ensure that business logic functions as intended under various scenarios:

- **GetTasks:** Tests retrieval of tasks.
- **CreateTask:** Tests task creation, including edge cases such as empty titles.
- **Promote User:** Verifies user promotion logic, including role validation.

### Controllers

Controller tests simulate HTTP requests and validate the responses:

- **GetTasks Endpoint:** Tests the `GET /tasks` endpoint to ensure it returns the correct status and data.
- **CreateTask Endpoint:** Tests the `POST /tasks` endpoint, verifying task creation and proper handling of request bodies.
- **User Promotion:** Tests the user promotion endpoint, ensuring proper role validation and error handling.

### Infrastructure

Infrastructure tests ensure that the underlying services like password management, JWT token generation, and middleware function correctly:

- **Password Comparison:** Tests the `ComparePasswords` function, covering scenarios like mismatched passwords, empty passwords, and successful matches.
- **JWT Generation and Validation:** Validates the `GenerateToken` and `ValidateToken` functions, including scenarios with invalid tokens.
- **Middleware Authentication:** Tests the authentication middleware, ensuring proper handling of requests with missing, invalid, or unauthorized tokens.

## Test Coverage

To evaluate the effectiveness of the test suite, coverage metrics have been collected. The aim was to achieve high coverage across critical components. You can generate a coverage report using the following command:

```bash
go test -cover ./...
```

### Coverage Report

- **Domain Models:** 100%
- **Use Cases:** 95%
- **Controllers:** 90%
- **Infrastructure:** 85%

## Running Tests

To run the tests locally, use the following command:

```bash
go test ./Tests/...
```

This will execute all test cases within the `Tests` directory. Ensure that all dependencies are installed and up-to-date before running the tests.

## Continuous Integration

Unit tests have been integrated into the CI pipeline to ensure that they are automatically executed with each commit. The CI configuration ensures that:

- Tests are run in isolated environments.
- Coverage reports are generated and uploaded.
- Builds are marked as failed if any tests do not pass.

## Issues Encountered

During the implementation of the test suite, the following challenges were encountered:

- **Mocking MongoDB Cursors:** Special attention was required to mock MongoDB cursors correctly.
- **Edge Cases in User Promotion:** Testing for unauthorized access required additional role validation logic.
- **Password Comparison:** Handling edge cases like empty passwords required additional error handling.
