# ATG Data Streamer Documentation

## Overview
The ATGDataStreamer application facilitates the interaction between data tanks and a central server through a serial port. It allows for the creation, retrieval, updating, and deletion of data tanks, as well as reading and writing data to a serial port. The application is built in Go and follows a modular structure to enhance maintainability and extensibility.

## Features
- **DataTank Management:** CRUD operations for managing data tanks.
- **Serial Communication:** Read and write data to a serial port.
- **HTTPS Delivery:** Send data to an HTTPS endpoint.
- **Database Integration:** SQLite database for persistent storage.
- **Graceful Shutdown:** Gracefully shutdown the server on interrupt signals.

## Table of Contents
- [ATG Data Streamer Documentation](#atg-data-streamer-documentation)
  - [Overview](#overview)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Dependencies](#dependencies)
  - [Configuration](#configuration)
  - [Usage](#usage)
    - [How to Run](#how-to-run)
    - [Access Endpoints](#access-endpoints)
  - [API Endpoints](#api-endpoints)
  - [Postman Usage](#postman-usage)
  - [Database Migration](#database-migration)
  - [Code Structure](#code-structure)
  - [Testing Scenarios](#testing-scenarios)
  - [To-Do List](#to-do-list)
  - [Future Improvements](#future-improvements)
  - [Code Flow](#code-flow)

## Prerequisites
- Go programming language installed.
- SQLite database.

## Installation
1. Clone the repository: `git clone <repository-url>`
2. Navigate to the project directory: `cd atgdatastreamer`
3. Run the application: `go run main.go`

## Dependencies
- [github.com/gorilla/mux](https://github.com/gorilla/mux) - HTTP router
- [github.com/patrickmn/go-cache](https://github.com/patrickmn/go-cache) - In-memory cache
- [github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver
- [github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate) - Database migrations

## Configuration
The application reads its configuration from an external JSON file. An example configuration file (`config.json`) is provided in the repository. Adjust the configuration values as needed.

```json
{
    "db_path": "synchub.db",
    "serial_port_name": "COM1",
    "serial_port_baud": 9600,
    "https_endpoint": "https://localhost:3000/receive-data"
}
```

## Usage

### How to Run
```bash
go run main.go
```

### Access Endpoints
- Local server: [http://localhost:8080](http://localhost:8080)

## API Endpoints
- **POST /data-tank:** Create a new data tank.
- **GET /data-tank:** Retrieve all data tanks.
- **PUT /data-tank/{id}:** Update a data tank by ID.
- **DELETE /data-tank/{id}:** Delete a data tank by ID.
- **GET /read-serial:** Read data from the serial port.

## Postman Usage
Use Postman or a similar tool to interact with the API endpoints.

## Database Migration
The application performs automatic database migration on startup. Ensure the SQLite database is created and accessible.

## Code Structure
The project follows a modular structure:
- **application:** Main application logic and server setup.
- **configs:** Configuration loading.
- **controllers:** HTTP request handlers.
- **database:** Database initialization and migration.
- **delivery:** Data delivery mechanisms (HTTPS, cache).
- **models:** Data structures used in the application.
- **repository:** Database interactions.
- **serial:** Serial port communication.
- **usecase:** Business logic and use case implementations.

## Testing Scenarios
- Ensure data tank creation, retrieval, update, and deletion work as expected.
- Verify serial port communication by reading data from the serial port.
- Check the integration with external systems through HTTPS delivery.

## To-Do List
- [ ] Implement caching for improved performance.
- [ ] Enhance error handling and logging.
- [ ] Add validation for API inputs.
- [ ] Implement unit tests for critical components.

## Future Improvements
- [ ] Support multiple serial ports.
- [ ] Implement secure communication protocols.
- [ ] Integrate with additional data sources and sinks.

## Code Flow
The following is a simplified flow of the application code:

1. **Main Execution (main.go

):**
   - Load configuration.
   - Initialize the database.
   - Create repository, serial port, HTTPS delivery, use case, and controller instances.
   - Initialize and run the application.

2. **Application Initialization (application/app.go):**
   - Set up routes and handlers.
   - Perform database migration.
   - Start the HTTP server.

3. **Controller Handling (controllers/data_tank_controller.go):**
   - Receive HTTP requests.
   - Invoke corresponding use case methods.

4. **Use Case Execution (usecase/data_tank_usecase.go):**
   - Implement business logic.
   - Interact with the repository for data storage.
   - Communicate with the serial port and deliver data via HTTPS.

5. **Repository Interaction (repository/data_tank_repository.go):**
   - Perform CRUD operations on the SQLite database.

6. **Serial Port Communication (serial/serial_port_impl.go):**
   - Connect to the specified serial port.
   - Read and write data.

7. **HTTPS Delivery (delivery/https_delivery.go):**
   - Send data to the configured HTTPS endpoint.

8. **Database Initialization (database/database.go):**
   - Initialize and migrate the SQLite database.

9. **Configuration Loading (configs/config_loader.go):**
   - Load the application configuration from a JSON file.

10. **Model Definitions (models/data_tank.go):**
   - Define data structures used in the application.

This flow represents the key components and their interactions within the ATGDataStreamer application.