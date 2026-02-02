# MyStore API

MyStore is a RESTful API service for managing a store's inventory, products, customers, and orders. It is built using Go, Gin, and GORM.

## Tech Stack

*   **Language:** Go (Golang)
*   **Framework:** Gin Web Framework
*   **Database:** MySQL
*   **ORM:** GORM
*   **Documentation:** Swagger (Swaggo)
*   **Configuration:** Viper

## Prerequisites

*   Go 1.16+
*   MySQL Database
*   Make (optional, for running Makefile commands)

## Setup & Configuration

1.  **Clone the repository**

    ```bash
    git clone https://github.com/oTeeLeko/mystore.git
    cd mystore
    ```

2.  **Environment Variables**

    Create a `app.env` file in the root directory with the following variables:

    ```env
    DB_DRIVER=mysql
    DB_SOURCE=username:password@tcp(localhost:3306)/mystore?charset=utf8mb4&parseTime=True&loc=Local
    HTTP_SERVER_ADDRESS=0.0.0.0:8080
    ```

    Replace `username`, `password`, `localhost:3306`, and `mystore` with your actual database credentials and configuration.

3.  **Install Dependencies**

    ```bash
    go mod download
    ```

## Running the Application

You can run the application using `make` or directly with `go run`.

**Using Make:**

```bash
make server
```

**Using Go:**

```bash
go run main.go
```

The server will start at the address specified in `HTTP_SERVER_ADDRESS` (default: `0.0.0.0:8080`).

## API Documentation

Swagger API documentation is automatically generated. Once the server is running, you can access the documentation at:

http://localhost:8080/swagger/index.html

### updating Swagger

If you modify the API controllers, you need to execute the following command to update the Swagger documentation:

```bash
swag init
```

## Project Structure

*   `api/`: Contains API related code (Controllers, Routes, Server setup).
*   `configs/`: Configuration and Database connection setup.
*   `core/`: Core business logic (Contracts, Models, Repositories).
*   `docs/`: Generated Swagger documentation files.
*   `models/`: GORM Database Models.
*   `util/`: Utility functions and configuration loading.
