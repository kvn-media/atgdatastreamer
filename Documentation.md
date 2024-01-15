# ATG Data Streamer Documentation

## Introduction

This application is designed to collect data from an Automatic Tank Gauge (ATG), parse the data, and store it in a database. The architecture consists of several components, each responsible for a specific task.

## Components

### 1. Configuration (config.go)

The `configs` package handles the loading and validation of the configuration from the `config.json` file. The configuration includes parameters such as ATG address, port, and data format.

- **Config Struct:**
  - `ATGAddress`: Address of the ATG.
  - `ATGPort`: Port number for ATG communication.
  - `ATGFormat`: Format of ATG data (currently supports only "json").

- **LoadConfig Function:**
  - Reads and parses the `config.json` file.
  - Validates the configuration parameters.
  - Returns a pointer to the `Config` struct.

### 2. ATG Manager (atg_manager.go)

The `managers` package handles the creation and management of the ATG connection.

- **ATGManager Struct:**
  - `Conn`: Exported field representing the connection to the ATG.

- **NewATGManager Function:**
  - Initializes a new `ATGManager` instance.
  - Opens a TCP connection to the ATG based on the configuration.
  - Returns a pointer to the `ATGManager` instance.

- **Close Method:**
  - Closes the connection to the ATG.

### 3. ATG Repository (atg_repo.go)

The `repositories` package handles interactions with the database for ATG data storage.

- **ATGRepository Struct:**
  - `db`: Database connection.

- **NewATGRepository Function:**
  - Initializes a new `ATGRepository` instance.
  - Accepts a pointer to a SQL database.
  - Returns a pointer to the `ATGRepository` instance.

- **SaveATGData Method:**
  - Inserts ATG data into the database.
  - Uses prepared statements for secure database operations.

### 4. ATG Data Model (atg.go)

The `models` package defines the structure of ATG data.

- **ATGData Struct:**
  - `Type`: Data type.
  - `ID`: Identifier.
  - `Value`: Numeric value.
  - `Timestamp`: Time of data collection.

### 5. ATG Services (atg.service.go)

The `services` package contains utility functions for handling ATG data.

- **ReadATGData Function:**
  - Reads data from the ATG connection.
  - Accepts a `net.Conn` parameter.
  - Returns the read data as a byte slice.

- **ParseATGData Function:**
  - Parses ATG data from JSON format.
  - Accepts a byte slice.
  - Returns a pointer to the `ATGData` struct.

- **PrintATGData Function:**
  - Prints ATG data to the console.
  - Accepts a pointer to the `ATGData` struct.

### 6. ATG Controller (atg_controller.go)

The `controllers` package orchestrates the data collection process.

- **CollectATGData Function:**
  - Continuously collects data from the ATG.
  - Reads data, parses it, saves it to the database, and optionally prints it.
  - Uses the `ATGManager`, `ATGRepository`, and `ATGServices` components.

### 7. Main Application (main.go)

The `main` package is the entry point of the application.

- **Init Function:**
  - Configures the maximum number of CPU cores.

- **Main Function:**
  - Loads configuration.
  - Connects to the database.
  - Creates instances of `ATGManager` and `ATGRepository`.
  - Launches the data collection process in a goroutine.
  - Blocks the main thread to keep the application running.
