

# Task Management API Doc

Base URLs:

- <a href="http://localhost:8000">Develop Env: http://localhost:8000</a>

# Authentication

- HTTP Authentication, scheme: bearer

# Tasks

## POST Create Task

POST /tasks

> Body Parameters

```json
{
  "title": "Some New Task",
  "description": "description of that Task",
  "status": "Urgent",
  "due_date": "2006-01-02"
}
```

### Params

| Name          | Location | Type   | Required | Description |
| ------------- | -------- | ------ | -------- | ----------- |
| body          | body     | object | no       | none        |
| » title       | body     | string | yes      | none        |
| » description | body     | string | yes      | none        |
| » status      | body     | string | yes      | none        |
| » due_date    | body     | string | yes      | none        |

> Response Examples

> Success

```json
{
  "id": "66b3625d168fde4756c42e43",
  "title": "Dummy Task 1",
  "description": "description of that Dummy Task 1",
  "due_date": "2006-01-02T15:04:05Z",
  "status": "Urgent",
  "user_id": "66b2ef7e61865834622b58ec"
}
```

> Unauthorized

```json
{
  "error": "unauthorized"
}
```

> User Not Found

```json
{
  "error": "user not found"
}
```

### Responses

| HTTP Status Code | Meaning                                                         | Description    | Data schema |
| ---------------- | --------------------------------------------------------------- | -------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)         | Success        | Inline      |
| 401              | [Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1) | Unauthorized   | Inline      |
| 404              | [Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)  | User Not Found | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name          | Type   | Required | Restrictions | Title | description |
| ------------- | ------ | -------- | ------------ | ----- | ----------- |
| » id          | string | true     | none         |       | none        |
| » title       | string | true     | none         |       | none        |
| » description | string | true     | none         |       | none        |
| » due_date    | string | true     | none         |       | none        |
| » status      | string | true     | none         |       | none        |
| » user_id     | string | true     | none         |       | none        |

HTTP Status Code **401**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

HTTP Status Code **404**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

## GET Get All Tasks

GET /tasks

> Response Examples

> Success

```json
[
  {
    "id": "66b33133cec8e35ebca9f5d6",
    "title": "Some New Task",
    "description": "description of that Task",
    "due_date": "2006-01-02T15:04:05Z",
    "status": "Urgent",
    "user_id": "66b2ef7e61865834622b58ec"
  }
]
```

> Unauthorized

```json
{
  "error": "unauthorized"
}
```

### Responses

| HTTP Status Code | Meaning                                                         | Description  | Data schema |
| ---------------- | --------------------------------------------------------------- | ------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)         | Success      | Inline      |
| 401              | [Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1) | Unauthorized | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name          | Type   | Required | Restrictions | Title | description |
| ------------- | ------ | -------- | ------------ | ----- | ----------- |
| » id          | string | false    | none         |       | none        |
| » title       | string | false    | none         |       | none        |
| » description | string | false    | none         |       | none        |
| » due_date    | string | false    | none         |       | none        |
| » status      | string | false    | none         |       | none        |
| » user_id     | string | false    | none         |       | none        |

HTTP Status Code **401**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

## GET Get Task

GET /tasks/{id}

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| id   | path     | string | yes      | none        |

> Response Examples

> Success

```json
{
  "id": "66b33133cec8e35ebca9f5d6",
  "title": "Some New Task",
  "description": "description of that Task",
  "due_date": "2006-01-02T15:04:05Z",
  "status": "Urgent",
  "user_id": "66b2ef7e61865834622b58ec"
}
```

> Unauthorized

```json
{
  "error": "unauthorized"
}
```

> Not Found

```json
{
  "error": "Task not found"
}
```

### Responses

| HTTP Status Code | Meaning                                                         | Description  | Data schema |
| ---------------- | --------------------------------------------------------------- | ------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)         | Success      | Inline      |
| 401              | [Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1) | Unauthorized | Inline      |
| 404              | [Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)  | Not Found    | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name          | Type   | Required | Restrictions | Title | description |
| ------------- | ------ | -------- | ------------ | ----- | ----------- |
| » id          | string | true     | none         |       | none        |
| » title       | string | true     | none         |       | none        |
| » description | string | true     | none         |       | none        |
| » due_date    | string | true     | none         |       | none        |
| » status      | string | true     | none         |       | none        |
| » user_id     | string | true     | none         |       | none        |

HTTP Status Code **401**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

HTTP Status Code **404**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

## DELETE Delete Task

DELETE /tasks/{id}

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| id   | path     | string | yes      | none        |

> Response Examples

> Success

```json
{
  "message": "Task deleted"
}
```

> Unauthorized

```json
{
  "error": "unauthorized"
}
```

> Not Found

```json
{
  "error": "task not found"
}
```

### Responses

| HTTP Status Code | Meaning                                                         | Description  | Data schema |
| ---------------- | --------------------------------------------------------------- | ------------ | ----------- |
| 202              | [Accepted](https://tools.ietf.org/html/rfc7231#section-6.3.3)   | Success      | Inline      |
| 401              | [Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1) | Unauthorized | Inline      |
| 404              | [Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)  | Not Found    | Inline      |

### Responses Data Schema

HTTP Status Code **202**

| Name      | Type   | Required | Restrictions | Title | description |
| --------- | ------ | -------- | ------------ | ----- | ----------- |
| » message | string | true     | none         |       | none        |

HTTP Status Code **401**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

HTTP Status Code **404**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

## PUT Update Task

PUT /tasks/{id}

> Body Parameters

```json
{
  "title": "Some New Task",
  "description": "description of that Task",
  "status": "Urgent",
  "due_date": "2006-01-02T15:04:05Z",
  "user_id": "66b2ef7e61865834622b58ec"
}
```

### Params

| Name          | Location | Type   | Required | Description |
| ------------- | -------- | ------ | -------- | ----------- |
| id            | path     | string | yes      | none        |
| body          | body     | object | no       | none        |
| » title       | body     | string | yes      | none        |
| » description | body     | string | yes      | none        |
| » status      | body     | string | yes      | none        |
| » due_date    | body     | string | yes      | none        |
| » user_id     | body     | string | yes      | none        |

> Response Examples

> Success

```json
{
  "id": "66b3638e168fde4756c42e44",
  "title": "Some New Task",
  "description": "description of that Task",
  "due_date": "2006-01-02T15:04:05Z",
  "status": "Urgent",
  "user_id": "66b2ef7e61865834622b58ec"
}
```

> Unauthorized

```json
{
  "error": "unauthorized"
}
```

> Not Found

```json
{
  "error": "Task not found"
}
```

### Responses

| HTTP Status Code | Meaning                                                         | Description  | Data schema |
| ---------------- | --------------------------------------------------------------- | ------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)         | Success      | Inline      |
| 401              | [Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1) | Unauthorized | Inline      |
| 404              | [Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4)  | Not Found    | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name          | Type   | Required | Restrictions | Title | description |
| ------------- | ------ | -------- | ------------ | ----- | ----------- |
| » id          | string | true     | none         |       | none        |
| » title       | string | true     | none         |       | none        |
| » description | string | true     | none         |       | none        |
| » due_date    | string | true     | none         |       | none        |
| » status      | string | true     | none         |       | none        |
| » user_id     | string | true     | none         |       | none        |

HTTP Status Code **401**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

HTTP Status Code **404**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

# Users

## GET Get Users

GET /users

> Response Examples

> Success

```json
[
  {
    "id": "66b4a223137908816d0322b2",
    "username": "nahomderese",
    "role": "admin"
  },
  {
    "id": "66b4a256137908816d0322b3",
    "username": "nahomd",
    "role": "regular"
  },
  {
    "id": "66b4b13c3ec210450b9428b0",
    "username": "abelw",
    "role": "admin"
  },
  {
    "id": "66b4b17aac199c6aed508dd0",
    "username": "ephyg",
    "role": "admin"
  },
  {
    "id": "66b4b2d48dc37790656dd224",
    "username": "moali",
    "role": "regular"
  },
  {
    "id": "66b4b2ed97b550152156f0f1",
    "username": "chalaolani",
    "role": "regular"
  },
  {
    "id": "66b4b318972fe33d7ad635c2",
    "username": "check",
    "role": "regular"
  },
  {
    "id": "66b4b4ab681587ee0fcfe486",
    "username": "checked",
    "role": "regular"
  },
  {
    "id": "66b4b5d3d0a4f6bf177cdd60",
    "username": "checking",
    "role": "regular"
  }
]
```

### Responses

| HTTP Status Code | Meaning                                                 | Description | Data schema |
| ---------------- | ------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success     | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type   | Required | Restrictions | Title | description |
| ---------- | ------ | -------- | ------------ | ----- | ----------- |
| » id       | string | true     | none         |       | none        |
| » username | string | true     | none         |       | none        |
| » role     | string | true     | none         |       | none        |

## GET Promote

GET /

> Response Examples

> Success

```json
{
  "id": "66b4b318972fe33d7ad635c2",
  "username": "check",
  "role": "admin"
}
```

> Not Found

```json
{
  "error": "user not found"
}
```

### Responses

| HTTP Status Code | Meaning                                                        | Description | Data schema |
| ---------------- | -------------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)        | Success     | Inline      |
| 404              | [Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4) | Not Found   | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type   | Required | Restrictions | Title | description |
| ---------- | ------ | -------- | ------------ | ----- | ----------- |
| » id       | string | true     | none         |       | none        |
| » username | string | true     | none         |       | none        |
| » role     | string | true     | none         |       | none        |

HTTP Status Code **404**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

## GET Get User

GET /users/{id}

### Params

| Name | Location | Type   | Required | Description |
| ---- | -------- | ------ | -------- | ----------- |
| id   | path     | string | yes      | none        |

> Response Examples

> Success

```json
{
  "id": "66b4b2ed97b550152156f0f1",
  "username": "chalaolani",
  "role": "regular"
}
```

> Not Found

```json
{
  "error": "user not found"
}
```

### Responses

| HTTP Status Code | Meaning                                                        | Description | Data schema |
| ---------------- | -------------------------------------------------------------- | ----------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)        | Success     | Inline      |
| 404              | [Not Found](https://tools.ietf.org/html/rfc7231#section-6.5.4) | Not Found   | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type   | Required | Restrictions | Title | description |
| ---------- | ------ | -------- | ------------ | ----- | ----------- |
| » id       | string | true     | none         |       | none        |
| » username | string | true     | none         |       | none        |
| » role     | string | true     | none         |       | none        |

HTTP Status Code **404**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

# Auth

## POST Login

POST /auth/login

> Body Parameters

```json
{
  "username": "nahomderese",
  "password": "nahomderese"
}
```

### Params

| Name       | Location | Type   | Required | Description |
| ---------- | -------- | ------ | -------- | ----------- |
| body       | body     | object | no       | none        |
| » username | body     | string | yes      | none        |
| » password | body     | string | yes      | none        |

> Response Examples

> Success Login

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjMxMTgzMDEsImlkIjoiNjZiMmVlNWQ2MTg2NTgzNDYyMmI1OGViIiwidXNlcm5hbWUiOiJuYWhvbWRlcmVzZSJ9.YFVTN9QRb-t8j1fE8nkOHxp2xw6BgVWlisrQex8Seto"
}
```

### Responses

| HTTP Status Code | Meaning                                                 | Description   | Data schema |
| ---------------- | ------------------------------------------------------- | ------------- | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | Success Login | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » token | string | true     | none         |       | none        |

## POST SignUp

POST /auth/register

> Body Parameters

```json
{
  "username": "nahom_de",
  "password": "check"
}
```

### Params

| Name       | Location | Type   | Required | Description |
| ---------- | -------- | ------ | -------- | ----------- |
| body       | body     | object | no       | none        |
| » username | body     | string | yes      | none        |
| » password | body     | string | yes      | none        |

> Response Examples

> Success

```json
{
  "id": "66b4c8c7c5dad55df0394e59",
  "username": "new_user",
  "role": "regular"
}
```

> Duplicate Username

```json
{
  "error": "user with this username already exists"
}
```

### Responses

| HTTP Status Code | Meaning                                                       | Description        | Data schema |
| ---------------- | ------------------------------------------------------------- | ------------------ | ----------- |
| 200              | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)       | Success            | Inline      |
| 409              | [Conflict](https://tools.ietf.org/html/rfc7231#section-6.5.8) | Duplicate Username | Inline      |

### Responses Data Schema

HTTP Status Code **200**

| Name       | Type   | Required | Restrictions | Title | description |
| ---------- | ------ | -------- | ------------ | ----- | ----------- |
| » id       | string | true     | none         |       | none        |
| » username | string | true     | none         |       | none        |
| » role     | string | true     | none         |       | none        |

HTTP Status Code **409**

| Name    | Type   | Required | Restrictions | Title | description |
| ------- | ------ | -------- | ------------ | ----- | ----------- |
| » error | string | true     | none         |       | none        |

<br>
<br>
<br>

## Read More

[Click Here](https://www.apidog.com/apidoc/shared-46a03815-ef9e-41d9-8d49-6205356e65ac)
