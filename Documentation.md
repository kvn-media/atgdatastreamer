# ATG Data Streamer Documentation

## Overview
The ATGDataStreamer application facilitates the interaction between data tanks and a central server through a serial port. It allows for the creation, retrieval, updating, and deletion of data tanks, as well as reading and writing data to a serial port. The application is built in Go and follows a modular structure to enhance maintainability and extensibility.

## Features

- CRUD operations for data tanks
- Read data from the serial port and store it in the database
- Send data to a specified HTTPS endpoint for external integration
- Logging middleware for HTTP requests
- Graceful shutdown handling for server
- Database initialization and migration
- Caching mechanism for optimized data retrieval
- Use of Gorilla Mux for routing

## Table of Contents
- [ATG Data Streamer Documentation](#atg-data-streamer-documentation)
  - [Overview](#overview)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation Package](#installation-package)
  - [Dependencies](#dependencies)
  - [Configuration](#configuration)
  - [Usage](#usage)
    - [How to Run the Application](#how-to-run-the-application)
    - [Build Application](#build-application)
    - [Access Endpoints](#access-endpoints)
    - [API Endpoints](#api-endpoints)
    - [Postman Usage](#postman-usage)
  - [Database Migration](#database-migration)
  - [Code Structure](#code-structure)
  - [Middleware](#middleware)
  - [Caching](#caching)
  - [Testing Scenarios](#testing-scenarios)
    - [Scenario 1: Create Tank Data](#scenario-1-create-tank-data)
    - [Scenario 2: Read Tank Data](#scenario-2-read-tank-data)
    - [Scenario 3: Update Tank Data](#scenario-3-update-tank-data)
    - [Scenario 4: Delete Tank Data](#scenario-4-delete-tank-data)
    - [Scenario 5: Read from Serial](#scenario-5-read-from-serial)
  - [Mock-up Testing](#mock-up-testing)
  - [To-Do List](#to-do-list)
  - [Future Improvements](#future-improvements)
  - [Program Flow Explanation](#program-flow-explanation)

## Prerequisites
- Go programming language installed
- Database (SQLite) for storing tank data
- Serial port connected to the ATG system
- HTTPS endpoint for data delivery

## Installation Package
Clone the repository and navigate to the project directory.

```bash
git clone <repository-url>
cd atgdatastreamer
```

## Dependencies
The project relies on the following external packages:
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP router and dispatcher
- [gorm.io/gorm](https://gorm.io/) - Object-relational mapping library for Golang
- [patrickmn/go-cache](https://github.com/patrickmn/go-cache) - In-memory key/value cache

Install dependencies using the following command:

```bash
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u github.com/patrickmn/go-cache
```

## Configuration
The configuration file `config.json` contains the settings for the application, including database path, serial port details, HTTPS endpoint, and graceful shutdown timeout.

```json
{
    "dbPath": "internal/database/schema/atg_data_stream.db",
    "serial_port_name": "COM1",
    "serial_port_baud": 9600,
    "https_endpoint": "https://localhost:3000/receive-data",
    "graceful_shutdown_timeout": 10
}
```

## Usage

### How to Run the Application
Run the following command to start the application:

```bash
go run main.go
```

### Build Application
To build the application, use the following command:

```bash
go build -o atgdatastreamer
```

### Access Endpoints
The application exposes endpoints for managing tank data. By default, it runs on port 8080.

Example:
- `http://localhost:8080/data-tank` (POST, GET)
- `http://localhost:8080/data-tank/{id}` (PUT, DELETE)
- `http://localhost:8080/read-serial` (GET)

### API Endpoints
- **POST /data-tank**: Create new tank data
- **GET /data-tank**: Retrieve all tank data
- **PUT /data-tank/{id}**: Update tank data by ID
- **DELETE /data-tank/{id}**: Delete tank data by ID
- **GET /read-serial**: Read data from the serial port

### Postman Usage
Use Postman or any API testing tool to interact with the provided API endpoints.

## Database Migration
The application automatically performs database migration using GORM. The SQLite database is created and the `DataTank` table is set up.

## Code Structure
- **/internal/application**: Main application logic
- **/internal/configs**: Configuration handling
- **/internal/controllers**: HTTP request handlers
- **/internal/database**: Database initialization and migration
- **/internal/delivery**: Data delivery to an HTTPS endpoint
- **/internal/models**: Data models
- **/internal/repository**: Database interaction
- **/internal/serial**: Serial port communication
- **/internal/usecase**: Business logic and use cases
- **main.go**: Entry point for the application

## Middleware
The application uses a logging middleware (`MyLoggingMiddleware`) to log information about incoming HTTP requests.

## Caching
Caching is implemented using an in-memory cache with a default expiration time of 5 minutes. The cache is initialized in `InitializeCache()`.

## Testing Scenarios

### Scenario 1: Create Tank Data

1. Send a POST request to `/data-tank` with valid tank data in the request body.

### Scenario 2: Read Tank Data

1. Send a GET request to `/data-tank` to retrieve all tank data.

### Scenario 3: Update Tank Data

1. Send a PUT request to `/data-tank/{id}` with the ID of an existing tank and updated data.

### Scenario 4: Delete Tank Data

1. Send a DELETE request to `/data-tank/{id}` with the ID of an existing tank to delete it.

### Scenario 5: Read from Serial

1. Send a GET request to `/read-serial` to read data from the serial port.

## Mock-up Testing

Mock-up testing scenarios should be implemented to simulate various conditions such as network failures, serial communication errors, and server unavailability.

## To-Do List
- [ ] Enhance error handling for HTTP requests.
- [ ] Implement validation for tank data.
- [ ] Enhance serial port error handling.
- [ ] Implement comprehensive unit tests.

## Future Improvements
- [ ] Implementing additional security measures
- [ ] Enhancing data validation and error handling
- [ ] Supporting multiple database backends

## Program Flow Explanation

1. **Initialization**: The application initializes the database, repository, serial port, and HTTPS delivery.
2. **Routing**: HTTP routes are defined using Gorilla mux in `internal/application/app.go`.
3. **Middleware**: Requests pass through a logging middleware defined in `internal/application/middleware.go`.
4. **Endpoint Handling**: Each endpoint is handled by the respective controller methods in `internal/controllers/dataTank_controller.go`.
5. **Database Interaction**: The controller methods interact with the database using the repository in `internal/repository/dataTank_repository.go`.
6. **Serial Communication**: The application reads data from the serial port and parses it using the SS160PLUS Protocol in `internal/serial/serialPort.go` and `internal/serial/ss160plus_parser.go`.
7. **HTTPS Delivery**: Tank data is sent to the configured HTTPS endpoint using the `HttpsDelivery` in `internal/delivery/https_delivery.go`.
8. **Caching**: A simple cache is initialized in `internal/application/cache.go` to store tank data temporarily.
9. **Shutdown Signal Handling**: The application gracefully handles shutdown signals, allowing ongoing tasks to finish.
10. **Run Application**: The application is run, and the server starts listening for incoming requests.

This flow represents the key components and their interactions within the ATGDataStreamer application.