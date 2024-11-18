# Golang APIs Kickstart

Golang APIs Kickstart is a starter project for building APIs using Go. It includes database initialization, configuration management, and templating for web pages.

## Features

- **GORM**: For ORM and database management.
- **SQLite**: As the default database.
- **HTML Templates**: For rendering web pages.
- **Configuration Management**: Using a custom configuration loader.

## Project Structure
```
.
├── README.md
├── env.sample
├── go.mod
├── go.sum
├── database.db
├── cmd
│   ├── api
│   │   └── main.go               # Main entry point for the API server
│   └── migrate
│       └── migrate.go            # Database migration script
├── internal
│   ├── config
│   │   └── config.go             # Configuration management
│   ├── controllers
│   │   └── users.go              # User-related API controllers
│   ├── database
│   │   └── db.go                 # Database initialization and connection
│   ├── dto
│   │   └── authInput.go          # Data Transfer Objects (DTOs)
│   ├── middleware
│   │   └── checkAuth.go          # Authentication middleware
│   ├── models
│   │   └── users.go              # Database models
│   ├── routes
│   │   └── main.go               # API route definitions
│   └── templates
│       ├── layouts
│       │   └── base.html         # Base layout for HTML templates
│       ├── pages
│       │   └── index.html        # Main page template
│       └── partials
│           ├── footer.html       # Footer partial template
│           └── header.html       # Header partial template
├── scripts
│   └── generate_html.sh          # Script for HTML generation
└── tmp
    ├── build-errors.log          # Log file for build errors
    └── main                      # Temporary main build file
```


18 directories, 21 files

## Getting Started

### Prerequisites

- Go 1.19 or later
- SQLite

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/golang-apis-kickstart.git
    cd golang-apis-kickstart
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Usage

1. Initialize the database:
    ```sh
    go run cmd/migrate/migrate.go
    ```

2. Run the application:
    ```sh
    go run main.go
    ```

### Configuration

Configuration is managed in the `internal/config` package. Ensure to load the configuration at the start of your application.

### Database

The database is initialized in the `internal/database` package. The `Init` function sets up the SQLite database.

### Templates

HTML templates are located in the `internal/templates` directory. The base layout is defined in `layouts/base.html`, and individual pages are in the `pages` directory.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. 
