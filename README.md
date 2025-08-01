# Pack Calculator (J Software Engineering Challenge)

A Go-based REST API and simple web UI to solve the pack calculation challenge. The project is designed to calculate the optimal number of packs to ship for a given order size based on a defined set of rules.

The implementation focuses on a clean architecture, clear API endpoints, and containerization for easy deployment and testing.

## Approach

The backend is a web server using the Echo framework. The core calculation logic is implemented using a dynamic programming (DP) approach to find the optimal solution. The pack sizes data is managed in memory for simplicity.

Key decisions:
- Modular packages (`server`, `handler`, `calculator`, `config`).
- In-memory state management for pack sizes, configurable via the API.
- A simple vanilla HTML and JavaScript frontend to interact with the API.
- Containerized with Docker.
- Unit tests for the core calculation logic.

---

## Core Dependencies

- **`github.com/labstack/echo/v4`**: The web framework used for routing and handling HTTP requests.
- **`testing` & `reflect`**: Library packages used for the unit test to ensure the calculator works as expected.

---

## File Structure

```
.
├── app
│   └── handler
│       ├── handler.go
│       └── order
│           └── order.go
├── cmd
│   └── server
│       └── main.go
├── config
│   ├── config.go
│   └── constants
│       └── messages
│           └── en.go
├── pkg
│   └── calculator
│       ├── calculator.go
│       └── calculator_test.go
├── public
│   └── index.html
├── server
│   ├── routes.go
│   └── server.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

---

## API Endpoints

| Method | Path                | Description                                   |
|--------|---------------------|-----------------------------------------------|
| `GET`  | `/api/pack-sizes`   | Returns the current list of pack sizes.       |
| `POST` | `/api/pack-sizes`   | Sets a new list of pack sizes.                |
| `POST` | `/api/calculate-packs` | Calculates the pack combination for an order. |

**POST `/api/pack-sizes` Body:**
```json
{
    "pack_sizes": [250, 500, 1000]
}
```

**POST `/api/calculate-packs` Body:**
```json
{
    "order_size": 1001
}
```

---

## How to Run

### Using Docker (Recommended)

The simplest way to get the application running.

```bash
# Start the server
docker-compose up
```

The application will be available at `http://localhost:8180`.

### Locally

If you have Go installed on your machine.

```bash
# Run the server
go run cmd/server/main.go
```

### Running Tests

To run the unit tests for the calculator package:

```bash
go test ./pkg/calculator/... -v
```

---

## Weaknesses & Potential Improvements

- **In-Memory State**: The pack sizes are stored in memory and will reset if the server restarts. A persistence layer (like a simple file or a database) could be added to make the configuration permanent.
- **Input Mutation**: The calculator currently sorts the `packSizes` slice in-place. While this is simple, in a larger application it would be safer to sort a copy to prevent unexpected side effects in other parts of the program.
- **Basic UI**: The user interface is minimal and designed only for basic interaction with the API.

---

## Comments

This project fulfills the requirements of the coding challenge. It provides a working API and UI, is containerized, and includes unit tests for the core logic. The code is structured to be readable and maintainable, demonstrating a practical approach to building a small Go web service.
