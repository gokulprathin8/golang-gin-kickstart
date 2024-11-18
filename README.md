# Golang APIs Kickstart

Golang APIs Kickstart is a starter project for building APIs using Go. It includes database initialization, configuration management, and templating for web pages.

## Features

- **GORM**: For ORM and database management.
- **SQLite**: As the default database.
- **HTML Templates**: For rendering web pages.
- **Configuration Management**: Using a custom configuration loader.

## Project Structure
.
├── README.md
├── cmd
│   ├── api
│   │   └── main.go
│   └── migrate
│       └── migrate.go
├── database.db
├── env.sample
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── controllers
│   │   └── users.go
│   ├── database
│   │   └── db.go
│   ├── dto
│   │   └── authInput.go
│   ├── middleware
│   │   └── checkAuth.go
│   ├── models
│   │   └── users.go
│   ├── routes
│   │   └── main.go
│   └── templates
│       ├── layouts
│       │   └── base.html
│       ├── pages
│       │   └── index.html
│       └── partials
│           ├── footer.html
│           └── header.html
├── scripts
│   └── generate_html.sh
└── tmp
    ├── build-errors.log
    └── main

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

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
