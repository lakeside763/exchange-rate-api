# Exchange Rate API

A simple Go API for retrieving currency exchange rates with Redis-backed storage and rate limiting.

## Features

- Get exchange rates for currency pairs (e.g., USD to EUR)
- Redis-based storage for fast access
- Rate limiting middleware to prevent abuse
- API key authentication

## Setup

1. **Install dependencies:**
   ```
   go mod tidy
   ```

2. **Start Redis:**  
   Make sure you have a Redis server running locally.

3. **Run the server:**
   ```
   go run main.go
   ```

## Usage

- **Endpoint:** `GET /api/v1/rates?base=USD&target=EUR`
- **Headers:**  
  - `X-API-Key: <your-api-key>`

**Example request:**
```
curl -H "X-API-Key: test123" "http://localhost:5200/api/v1/rates?base=USD&target=EUR"
```

## Testing

Run all tests:
```
go test ./... -v
```

## License

MIT