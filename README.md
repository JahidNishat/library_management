# Library Management Backend

This project is a library management backend project written in Golang. It uses PostgreSQL as the database, Viper for configuration, Cobra for command line interface, Argon2id for password hashing, Gin for web framework, and JWT for authentication and authorization.

## Features

- **User Authentication:** The project provides a robust user authentication system utilizing the Argon2id hashing algorithm and JWT tokens for enhanced security.
- **Authorization:** Access to specific routes and resources is restricted based on user roles and permissions, ensuring proper authorization.
- **User Management:** CRUD operations for managing users, including signup, login, user details retrieval, logout, and token refresh.
- **Book Management:** Comprehensive CRUD operations for managing books, allowing for seamless addition, retrieval, updating, and deletion of book records.

## Technologies Used

- **Golang:** The backend is developed using the Go programming language, known for its efficiency and performance.
- **PostgreSQL:** A powerful and reliable relational database used for storing user and book data securely.
- **Viper:** A versatile configuration management library that simplifies the handling of application configuration.
- **Cobra:** A command-line interface (CLI) library used for building a user-friendly and interactive CLI for the project.
- **Argon2id:** A highly secure and efficient password hashing algorithm used for user password hashing.
- **Gin:** A lightweight and blazing-fast HTTP framework used for routing and handling HTTP requests with ease.
- **JWT Token:** JSON Web Tokens are utilized for user authentication and authorization, ensuring secure communication.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/library-management-backend.git
2.  Install the dependencies:
    ```bash
    go mod download
3.  Set up the PostgreSQL database:
    - Create a new database for the project.
    - Update the database configuration in `config/config.yaml`.
4. Build and run the project:
    ```bash
    go build -o library-management-backend cmd/main.go ./library-management-backend

## Available API Endpoints

| Method | Endpoint               | Description                           |
| ------ | ----------------------| ------------------------------------- |
| POST   | `/signup`              | Register a new user                   |
| POST   | `/users/login`         | Authenticate and generate a JWT token |
| GET    | `/users`               | Get all users                         |
| GET    | `/users/{user_id}`     | Get user by ID                        |
| GET    | `/users/logout`        | Invalidate user's JWT token           |
| GET    | `/users/refresh`       | Generate new JWT token using the refresh token |
| DELETE | `/users/{user_id}`     | Delete user by ID                     |
| GET    | `/books`               | Get all books                         |
| GET    | `/books/{bookId}`      | Get book by ID                        |
| POST   | `/books`               | Create a new book                     |
| PUT    | `/books/{bookId}`      | Update book by ID                     |
| DELETE | `/books/{bookId}`      | Delete book by ID                     |

