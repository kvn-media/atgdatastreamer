# ATG Data Streamer

> An application for collecting, parsing, and storing data from Automatic Tank Gauges (ATGs).

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](license)

ATG Data Streamer is a Go application that continuously collects data from ATGs, parses it, and stores it in a database. This README provides an overview of the code structure, components, usage, and other relevant information.

## Table of Contents

- [ATG Data Streamer](#atg-data-streamer)
  - [Table of Contents](#table-of-contents)
  - [Components](#components)
  - [Usage](#usage)
  - [Dependencies](#dependencies)
  - [Configuration](#configuration)
  - [License](#license)
  - [Contributing](#contributing)
  - [Acknowledgments](#acknowledgments)

## Components

1. [Configuration](#1-configuration-configgo)
2. [ATG Manager](#2-atg-manager-atg_managergo)
3. [ATG Repository](#3-atg-repository-atg_repogo)
4. [ATG Data Model](#4-atg-data-model-atggo)
5. [ATG Services](#5-atg-services-atgservicego)
6. [ATG Controller](#6-atg-controller-atg_controllergo)
7. [Main Application](#7-main-application-maingo)

## Usage

1. Ensure you have a valid `config.json` file with ATG configuration.
2. Run the application:

    ```bash
    go run main.go
    ```

3. The application will continuously collect data from the ATG, parse it, and store it in the configured database.

## Dependencies

- The application uses a SQLite database for data storage.

## Configuration

Edit the `config.json` file to customize ATG and database settings.

```json
{
  "ATGAddress": "127.0.0.1",
  "ATGPort": 1234,
  "ATGFormat": "json"
}
```

## License

This project is licensed under the [MIT License](license).

## Contributing

Feel free to contribute to the project by creating issues or submitting pull requests.

## Acknowledgments

- This project was inspired by the need to efficiently collect and manage data from Automatic Tank Gauges.
