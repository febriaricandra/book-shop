# how to run this go project
- Clone the project
- Run `go run cmd/app/main.go` in the project root directory
- Open your browser and go to `http://localhost:8080/`
- You should see the message `Hello, World!` in your browser

# how to run the tests
- Run `go test -v ./tests/features` in the project root directory
- You should see the test results in the terminal

# Running on Docker
- Run `docker build -t book-shop .` in the project root directory
- Run `docker run -p 8080:8080 book-shop` in the project root directory
- Run `docker run --name book-shop-container -p 8080:8080 book-shop` in the project root directory

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
    "cover_image": "IMGES.JPEG",
    "old_price": 10.99,
    "new_price": 9.99
}
```
- Response:
```json
{
    "data": {
        "id": 12,
        "created_at": "2024-12-02T16:34:02.877+07:00",
        "updated_at": "2024-12-02T16:34:02.877+07:00",
        "deleted_at": null,
        "title": "JAMET SEDOLOR",
        "description": "lorem ipsum dolor amet",
        "category": "Popular",
        "trending": true,
        "cover_image": "https://cloudflare.r2.dev/bookshop/2aaf0b77-fd69-48bf-a5fa-b7b363843c36.JPEG",
        "old_price": 12.02,
        "new_price": 11.04
    },
    "message": "Book created successfully",
    "status": true
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

- Order a book
- Request: POST
- URL: `/api/orders`
- Authorization: Bearer Token
- Request Body:
```json
{
  "name": "FebriMiauw",
  "email": "Febrimiauw@gmail.com",
  "address": {
    "city": "Anytown",
    "country": "United State",
    "state": "CA",
    "zipcode": "12345"
  },
  "phone": "123-456-7890",
  "total_price": 102.20,
  "book_ids": [7, 8, 9]
}
```
- Response:
```json
{
    "id": 9,
    "created_at": "2024-11-19T21:23:52.078+07:00",
    "updated_at": "2024-11-19T21:23:52.078+07:00",
    "deleted_at": null,
    "name": "FebriMiauw",
    "email": "Febrimiauw@gmail.com",
    "address": {
        "city": "Anytown",
        "country": "United State",
        "state": "CA",
        "zipcode": "12345"
    },
    "phone": "123-456-7890",
    "total_price": 102.2,
    "user_id": 2,
    "books": null,
    "user": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "email": "",
        "name": "",
        "is_admin": false
    }
}
```