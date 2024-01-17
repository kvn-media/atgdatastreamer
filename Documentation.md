# ATG Data Streamer Documentation

## Overview

ATG Data Streamer is a Go application designed for streaming and managing data from ATG tanks. This documentation provides information on the application's features, usage, code structure, and deployment.

## Features

- **Data Management:** Create, read, update, and delete data from ATG tanks.
- **Serial Communication:** Read data from the serial port connected to the ATG tank.
- **HTTPS Delivery:** Send tank data securely over HTTPS.

## Table of Contents

- [ATG Data Streamer Documentation](#atg-data-streamer-documentation)
  - [Overview](#overview)
  - [Features](#features)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Usage](#usage)
    - [API Endpoints](#api-endpoints)
    - [Postman Usage](#postman-usage)
  - [Database Migration](#database-migration)
  - [Code Structure](#code-structure)
  - [Testing Scenarios](#testing-scenarios)
  - [Future Improvements](#future-improvements)

## Prerequisites

Before using the application, ensure you have the following installed:

- Go (Golang)
- SQLite

## Installation

Follow these steps to install the application:

1. Clone the repository:

   ```bash
   git clone https://github.com/kvn-media/atgdatastreamer.git
   ```

2. Change to the project directory:

   ```bash
   cd atgdatastreamer
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Build the application:

   ```bash
   go build
   ```

5. Run the application:

   ```bash
   ./atgdatastreamer
   ```

## Configuration

Configure the application using the `configs/config.json` file. Update parameters such as `SerialPortName`, `SerialPortBaud`, and `HTTPSEndpoint`.

1. Copy the example configuration file:
   ```bash
   cp configs/config.example.json configs/config.json
   ```
2. Modify configs/config.json with your specific configuration settings.

## Usage

### API Endpoints

The application exposes the following API endpoints:

- Create DataTank: `POST /data-tank`
- Get All DataTanks: `GET /data-tank`
- Update DataTank: `PUT /data-tank/{id}`
- Delete DataTank: `DELETE /data-tank/{id}`
- Read from Serial: `GET /read-serial`

### Postman Usage

Use [Postman](https://www.postman.com/) for convenient API testing. Import the provided [Postman collection](postman_collection.json) to quickly test the API endpoints.

## Database Migration

To perform database migration, run:

```bash
./atgdatastreamer migrate
```

## Code Structure

The code is structured into several packages, including `application`, `controllers`, `database`, `delivery`, `repository`, `serial`, and `usecase`. Each package serves a specific purpose in the application's architecture.

## Testing Scenarios
- Daily Data Collection: Simulate daily data collection from the ATG system and observe successful HTTPS transfers.
- Invalid API Requests: Test the application's response to invalid API requests.
- Database Migration: Verify that the database migration process completes successfully.


## Future Improvements
- Enhanced Logging
- API Authentication
- HTTPS Configuration
- Data Validation and Sanitization
- Unit Tests and Test Coverage
- Swagger/OpenAPI Documentation
- CORS Handling
- Graceful Shutdown
- Containerization
- Monitoring and Metrics
- Continuous Integration/Continuous Deployment (CI/CD)
- Error Handling