# PAYMENT MICROSERVICE

## Project Structure (API GATEWAY PATTERN)

This project includes authetication service and payment service accompanied with api gateway.

### `To run`

```bash
docker-compose up --build -d
```

# `API-GATEWAY`

The API gateway in my microservice project acts as a centralized entry point for accessing and managing the order service and auth service. It provides a unified interface for clients to interact with these services without needing to know their specific locations.

### `To run`

```bash
# Navigate into the project
cd ./RampNow-API-Gateway

# Please make sure to add the envs
PORT=:3005
AUTH_SVC_URL=localhost:50056
ORDER_SVC_URL=localhost:50057

# Alternative
make deps # Install the dependencies
make server  # Run the project
```

# `RampNow-Auth_Service`

Auth service is responsible for handling user authentication and authorization. Its primary purpose is to ensure secure access to protected resources within the system. The service offers various authentication mechanisms, such as username/password, JWT tokens, or OAuth, etc... It validates user credentials, generates and verifies tokens, and enforces access control policies. By integrating the Authentication Service into the project, we can ensure that only authorized users can interact with the system's resources.
We have used Postgresql database for auth service.
(link to auth service) - Authentication Service (grpc)

### `To run`

```bash
# Navigate into the project
cd ./RampNow-Auth-Service

make deps # Install the dependencies
make run  # Run the project
```

# `RampNow-Payment_Service`

Payment Service is dedicated to managing and processing transaction within the system. It provides functionalities related to sending money, getting details of user wallets etc. The service allows users or other microservices to create new transactions.
I have used Postgres database for auth service.
(Link to payment service) - Order Service (grpc)

### `To run`

```bash
# Navigate into the project
cd ./RampNow-Payment-Service

make deps # Install the dependencies
make run  # Run the project
```

API DOCUMENTATION-SWAGGER

### `http://localhost:3005/swagger/index.html`
