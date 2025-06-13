
````markdown
# 🎓 Student API

A lightweight RESTful API for managing students.

## 🚀 Endpoints

### ➕ Create Student
- **POST** `/api/students`
- **Body**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "age": 20
}
````

* **Response**:
  `201 Created`

```json
{
  "id": 1
}
```

---

### 🔍 Get Student by ID

* **GET** `/api/students/{id}`
* **Response**:
  `200 OK`

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 20
}
```

---

### 📋 Get All Students

* **GET** `/api/students`
* **Response**:
  `200 OK`

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 20
  },
  ...
]
```

---

## ❗ Validation Rules

All fields are required:

* `name` – string
* `email` – string
* `age` – integer

Invalid input returns `400 Bad Request` with validation errors.

---
