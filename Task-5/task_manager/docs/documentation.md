# Task API Documentation

## Description

The Task API allows users to manage tasks efficiently. It provides endpoints to create, retrieve, update, and delete tasks. Users can fetch all tasks or get details of a specific task by its ID. This makes it easy to integrate task management functionality into your applications.

## BaseUrl

```
http://localhost:8000
```

## Endpoints

### Get All Tasks

#### Request

`GET /tasks`

#### Response

**Status Code:** 200 OK

**Response Body:**

```json
[
  {
    "id": 1,
    "title": "Some New Task",
    "description": "Description of the task",
    "due_date": "2006-01-02T15:04:05Z",
    "status": "Urgent"
  },
  {
    "id": 2,
    "title": "Another Task",
    "description": "Description of another task",
    "due_date": "2006-01-02T15:04:05Z",
    "status": "Urgent"
  },
  {
    "id": 3,
    "title": "Third Task",
    "description": "Description of the third task",
    "due_date": "2006-01-02T15:04:05Z",
    "status": "Urgent"
  }
]
```

### Get Task by ID

#### Request

`GET /tasks/{id}`

**Path Parameters:** - `id` (integer, required): The ID of the task to retrieve.

#### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "id": 1,
  "title": "Complete Go project",
  "description": "Finish the distributed system project in Go",
  "due_date": "2023-08-07T12:34:56Z",
  "status": "pending"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "Invalid ID format"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "Task not found"
}
```

### Update Task

#### Request

`PUT /tasks/{id}`

**Path Parameters:** - `id` (integer, required): The ID of the task to retrieve.

**Request Body:**

```json
{
  "title": "New Task Title",
  "description": "New description",
  "due_date": "2020-02-21T02:04:50Z",
  "status": "completed"
}
```

#### Response

**Status Code:** 200 OK

**Response Body:**

```json
{
  "id": 1,
  "title": "New Task Title",
  "description": "New description",
  "due_date": "2020-02-21T02:04:50Z",
  "status": "completed"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "Invalid ID or request data"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "Task not found"
}
```

### Delete Task

#### Request

`DELETE /tasks/:id`

**Path Parameters:** - `id` (integer, required): The ID of the task to retrieve.

#### Response

**Status Code:** 204 No Content

**Response Body:**

( No Content )

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "Invalid ID format"
}
```

**Status Code:** 404 Not Found

**Response Body:**

```json
{
  "error": "Task not found"
}
```

### Create Task

#### Request

`POST /tasks`

**Request Body:**

```json
{
  "title": "Task Title",
  "description": "Description for the task",
  "due_date": "2024-01-20T01:38:56Z",
  "status": "done"
}
```

#### Response

**Status Code:** 201 Created

**Response Body:**

```json
{
  "id": 2,
  "title": "Task Title",
  "description": "Description for the task",
  "due_date": "2024-01-20T01:38:56Z",
  "status": "done"
}
```

**Status Code:** 400 Bad Request

**Response Body:**

```json
{
  "error": "Invalid request data"
}
```
