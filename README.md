# ORDER_MICROSERVICE


## Project Structure (API GATEWAY PATTERN)
This project includes authenction service and order service accompanied with api gateway.

# `API-GATEWAY`
The API gateway in our microservice project acts as a centralized entry point for accessing and managing the order service and auth service. It provides a unified interface for clients to interact with these services without needing to know their specific locations.
(https://github.com/SethukumarJ/sellerapp_order_svc/tree/main/SellerApp-API-Gateway) - API Gateway (HTTP)
### `To run`
```bash
# Navigate into the project
cd ./SellerApp-API-Gateway

# Please make sure to add the envs
PORT=:3005
AUTH_SVC_URL=localhost:50056
ORDER_SVC_URL=localhost:50057

# Run the project in Development Mode
 docker build -t my-api-gateway .
 docker run -it --rm -p 3005:3005 my-api-gateway

# Alternative
make deps # Install the dependencies 
make server  # Run the project
```
# `SellerApp-Auth_Service`
Auth service is responsible for handling user authentication and authorization. Its primary purpose is to ensure secure access to protected resources within the system. The service offers various authentication mechanisms, such as username/password, JWT tokens, or OAuth, etc... It validates user credentials, generates and verifies tokens, and enforces access control policies. By integrating the Authentication Service into the project, we can ensure that only authorized users can interact with the system's resources.
We have used Postgresql database for auth service.
(https://github.com/SethukumarJ/sellerapp_order_svc/tree/main/SellerApp-Auth-Service) - Authentication Service (grpc)
### `To run`
```bash
# Navigate into the project
cd ./SellerApp-Auth-Service

# Please make sure to add the envs

# Run the project in Development Mode
 docker-compose build .
 docker-compose up

# Alternative
make deps # Install the dependencies 
make run  # Run the project
```
# `SellerApp-Order_Service`
Order Service is dedicated to managing and processing orders within the system. It provides functionalities related to creating, updating, and retrieving order information. The service allows users or other microservices to place new orders, modify existing orders, track order statuses, and retrieve order history. 
We have used Mysql database for auth service.
(https://github.com/SethukumarJ/sellerapp_order_svc/tree/main/SellerApp-Auth-Service) - Order Service (grpc)
### `To run`
```bash
# Navigate into the project
cd ./SellerApp-Order-Service

# Please make sure to add the envs

# Run the project in Development Mode
 docker-compose build .
 docker-compose up

# Alternative
make deps # Install the dependencies 
make run  # Run the project
```

## Template Structure (CLEAN ARCHITECTURE)

- [Gin](github.com/gin-gonic/gin) is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.
- [JWT](github.com/golang-jwt/jwt) A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.
- [GORM](https://gorm.io/index.html) with [PostgresSQL & MySQL](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)The fantastic ORM library for Golang aims to be developer friendly.
- [Wire](https://github.com/google/wire) is a code generation tool that automates connecting components using dependency injection.
- [Viper](https://github.com/spf13/viper) is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats.
- [swag](https://github.com/swaggo/swag) converts Go annotations to Swagger Documentation 2.0 with [gin-swagger](https://github.com/swaggo/gin-swagger) and [swaggo files](github.com/swaggo/files)
Additional commands:
### `Other cammands`
```bash
➔ make help
build                          Compile the code, build Executable File
run                            Start application
test                           Run tests
test-coverage                  Run tests and generate coverage file
deps                           Install dependencies
deps-cleancache                Clear cache in Go module
wire                           Generate wire_gen.go
swag                           Generate swagger docs
help                           Display this help screen
```

API DOCUMENTATION-SWAGGER
### `http://localhost:3005/swagger/index.html`
![Screenshot from 2023-06-03 07-10-26](https://github.com/SethukumarJ/sellerapp_order_svc/assets/114211073/4bac5c76-bb13-4f1f-896b-fe75deb65a93)
