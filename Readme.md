# how to run this go project
- Clone the project
- Run `go run cmd/app/main.go` in the project root directory
- Open your browser and go to `http://localhost:8080/`
- You should see the message `Hello, World!` in your browser

# how to run the tests
- Run `go test -v ./tests/features` in the project root directory
- You should see the test results in the terminal


# API Endpoints Documentation

## Get Books
- Request: GET
- URL: `/api/v1/books`
- Response: 
```json
{
    "data": [
        {
            "id": 7,
            "created_at": "2024-11-07T15:42:36.87+07:00",
            "updated_at": "2024-11-07T15:42:36.87+07:00",
            "deleted_at": null,
            "title": "The Great Gatsby",
            "description": "The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald.",
            "category": "Novel",
            "trending": true,
            "cover_image": "https://images-na.ssl-images-amazon.com/images/I/51Zymoq7UnL._AC_SY400_.jpg",
            "old_price": 10.99,
            "new_price": 9.99
        },
        {
            "id": 8,
            "created_at": "2024-11-07T15:42:36.871+07:00",
            "updated_at": "2024-11-07T15:42:36.871+07:00",
            "deleted_at": null,
            "title": "To Kill a Mockingbird",
            "description": "To Kill a Mockingbird is a novel by Harper Lee published in 1960.",
            "category": "Novel",
            "trending": true,
            "cover_image": "https://images-na.ssl-images-amazon.com/images/I/51Zymoq7UnL._AC_SY400_.jpg",
            "old_price": 10.99,
            "new_price": 9.99
        },
    ],
    "status": true
}
```

- Get a single book
- Request: GET
- URL: `/api/v1/books/{id}`
- Response:
```json
{
    "data": {
        "id": 7,
        "created_at": "2024-11-07T15:42:36.87+07:00",
        "updated_at": "2024-11-07T15:42:36.87+07:00",
        "deleted_at": null,
        "title": "The Great Gatsby",
        "description": "The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald.",
        "category": "Novel",
        "trending": true,
        "cover_image": "https://images-na.ssl-images-amazon.com/images/I/51Zymoq7UnL._AC_SY400_.jpg",
        "old_price": 10.99,
        "new_price": 9.99
    },
    "status": true
}
```

- Create a book
- Request: POST
- URL: `/api/v1/books`
- Authorization: Bearer Token (isAdmin = true Required)
- Request Body:
```json
{
    "title": "The Great Gatsby",
    "description": "The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald.",
    "category": "Novel",
    "trending": true,
    "cover_image": "https://images-na.ssl-images-amazon.com/images/I/51Zymoq7UnL._AC_SY400_.jpg",
    "old_price": 10.99,
    "new_price": 9.99
}
```
- Response:
```json
{
    "data": {
        "id": 10,
        "created_at": "2024-11-19T15:14:03.906+07:00",
        "updated_at": "2024-11-19T15:14:03.906+07:00",
        "deleted_at": null,
        "title": "Hello World: new world!",
        "description": "lorem ipsum dolor amet",
        "category": "Popular",
        "trending": true,
        "cover_image": "uploads/1da8d4b3-523b-408e-b634-cba959f28cd9.JPEG",
        "old_price": 12.02,
        "new_price": 11.04
    },
    "message": "Book created successfully",
    "status": true
}
```

- Register a user
- Request: POST
- URL: `/api/register`
- Request Body:
```json
{
    "name": "John Doe",
    "email": "john.doe@gmail.com",
    "password": "password12345"
}
```
- Response:
```json
{
    "message": "User registered successfully"
}
```

- Login a user
- Request: POST
- URL: `/api/login`
- Request Body:
```json
{
    "email": "john.doe@gmail.com",
    "password": "password12345"
}
```
- Response:
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzMyMDkwNjE5LCJpYXQiOjE3MzIwMDQyMTksImp0aSI6IjI4NDM4Yjg3LTQxNGQtNGI3OC1hYjMzLTFkYTNjM2Q4MWViZCIsImlzX2FkbWluIjp0cnVlLCJuYW1lIjoiZmVicmlnYXVsIiwiZW1haWwiOiJmZWJyaWFyaWdhdWxAZ21haWwuY29tIiwidXNlcl9pZCI6Mn0.SOvaZvEzNl_PwpYcPnDcArqEw98VzItVtN0Ag6v3PcY",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzMyMDkwNjE5LCJpYXQiOjE3MzIwMDQyMTksImp0aSI6IjU2MjU0M2RkLWUzODEtNDU1Zi1iNDdmLTk0NTg3NjM0YmJkNCIsImlzX2FkbWluIjp0cnVlLCJuYW1lIjoiZmVicmlnYXVsIiwiZW1haWwiOiJmZWJyaWFyaWdhdWxAZ21haWwuY29tIiwidXNlcl9pZCI6Mn0.20t3pC348mrIe07PWGisGV-3AM1N8AQmbJ9Vz3a0ugo"
}
```

- Refresh Token
- Request: POST
- URL: `/api/refresh`
- Request Body:
```json
{
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzMyMDkwNjE5LCJpYXQiOjE3MzIwMDQyMTksImp0aSI6IjU2MjU0M2RkLWUzODEtNDU1Zi1iNDdmLTk0NTg3NjM0YmJkNCIsImlzX2FkbWluIjp0cnVlLCJuYW1lIjoiZmVicmlnYXVsIiwiZW1haWwiOiJmZWJyaWFyaWdhdWxAZ21haWwuY29tIiwidXNlcl9pZCI6Mn0.20t3pC348mrIe07PWGisGV-3AM1N8AQmbJ9Vz3a0ugo"
}
```
- Response:
```json
{
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIiwiZXhwIjoxNzMyMDA1NTEyLCJpYXQiOjE3MzIwMDQ2MTIsImp0aSI6IjJjMjJmNzk3LWZjOWYtNGFlZC1iNjIzLTIwN2Q3MTM3MzFhMSIsImlzX2FkbWluIjp0cnVlLCJuYW1lIjoiZmVicmlnYXVsIiwiZW1haWwiOiJmZWJyaWFyaWdhdWxAZ21haWwuY29tIiwidXNlcl9pZCI6Mn0.i7V__R7NpLggL9yaUOIpb9UOnrKdKGItIgoOMA0jswU"
}
```
