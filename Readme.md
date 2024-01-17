# ATG Data Streamer

ATG Data Streamer is a Go application for streaming and managing data from ATG tanks.

## Features

- **Data Management:** Create, read, update, and delete data from ATG tanks.
- **Serial Communication:** Read data from the serial port connected to the ATG tank.
- **HTTPS Delivery:** Send tank data securely over HTTPS.

## Prerequisites

Before running the application, make sure you have the following installed:

- Go (Golang)
- SQLite

## Installation

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

Configure the application using the `configs/config.json` file. Update the following parameters:

- `SerialPortName`: The name of the serial port connected to the ATG tank.
- `SerialPortBaud`: The baud rate for serial communication.
- `HTTPSEndpoint`: The endpoint for secure data delivery.

## Usage

1. Start the application.

2. Access the API endpoints:

   - Create DataTank: `POST /data-tank`
   - Get All DataTanks: `GET /data-tank`
   - Update DataTank: `PUT /data-tank/{id}`
   - Delete DataTank: `DELETE /data-tank/{id}`
   - Read from Serial: `GET /read-serial`

## Database Migration

The application uses SQLite. To perform database migration, run:

```bash
./atgdatastreamer migrate
```

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](license) file for details.
```
