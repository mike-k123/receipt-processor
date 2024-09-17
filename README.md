# Receipt Processor API

A simple receipt processing API built with Go and the Gin framework. It allows users to submit receipts for processing and retrieve points based on specific rules.

## Table of Contents

- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
  - [Submit Receipt](#submit-receipt)
  - [Get Points for Receipt](#get-points-for-receipt)


## Getting Started

### Prerequisites

- Go 1.23+
- Postman or cURL for testing the API

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/mike-k123/receipt-processor.git
   cd receipt-processor
   ```

### Install Dependencies

Run the following commands to install the necessary packages:

```bash
go get github.com/gin-gonic/gin
go get github.com/google/uuid
```
or use 

```bash
go mod download
```

### Running the API

1. Run the server locally
```bash
go run main.go
```
Or through docker
```bash
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

2. Server will be running live on `http://localhost:8080`

### API Endpoints

1. Create a receipt - `POST /receipts/process`
##### Request body example
```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-02",
  "purchaseTime": "13:13",
  "total": "1.25",
  "items": [
    {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
  ]
}
```
##### Response example
```json
{
  "id": "adb6b560-0eef-42bc-9d16-df48f30e89b2"
}
```


2. Get a receipt - `GET /receipts/{id}/points`
##### Response example
```json
{
  "points": 32
}
```




