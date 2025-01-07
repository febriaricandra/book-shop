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

## Books
- GET `api/books` - Get all books
- GET `api/books/{id}` - Get a book by id
- POST `api/books` - Create a new book
- PUT `api/books/{id}` - Update a book by id
- DELETE `api/books/{id}` - Delete a book by id

## Orders
- GET `api/orders` - Get all orders
- GET `api/orders/{id}` - Get an order by id
- POST `api/orders` - Create a new order

## Users
- POST `api/login` - Login a user
- POST `api/register` - Register a new user
- GET `api/profile` - Get the user profile

## Rajaongkir API INTEGRATION
- GET `api/provinces` - Get all provinces
- GET `api/cities` - Get all cities
- GET `api/cost` - Get the shipping cost
