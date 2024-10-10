# Inventory Management System

This is an inventory management system designed to manage products, categories, suppliers, and other functionalities related to inventory administration.

## Description

The **Inventory Management System** allows users to manage products and track stock levels efficiently. It offers a user-friendly interface to handle inventory operations, such as adding new products, editing existing items, categorizing products, and monitoring stock levels in real-time.

## Backend Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/amorindev/backend-inventory.git
    ```

2. Download dependencies:
    ```bash
    go mod tidy
    ```

3. Set environment variables, add a `.env` file based on `env.example`:
    ```bash
    DB_PASSWORD=your_password
    DB_PORT=5432
    DB_HOST=your_host
    DB_USER=your_user
    DB_NAME=your_db_name

    # Development environment dev-prod
    GIN_MODE=debug   # or release
    BACK_ENV=dev
    PORT=7000

    CORS=http://localhost:5173
    ```
4. Run the project:
    ```bash
    go run main.go
    ```
