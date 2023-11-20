# purchase-transaction
Repository for purchase transaction.

This project uses [decimal](https://github.com/shopspring/decimal) for doing monetary calculation with precision.

## How to run (Docker)
1. Clone the repository
2. Update .env file for db credentials (default one's can also be used)
3. Run `docker-compose up -d --build` in the root folder

## How to run (Without Docker)
1. Clone the repository
2. Run postgresSQL locally
3. Update .env file
4. This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)
5. Run command at root of repo : `go run main.go`

## How to run unit test 
`go test ./...`

## API
There are 2 API's :
1. Store transaction
   - **URL:** `/v1/api/transaction`
   - **Method:** `POST`
   - **Request Body:**
     - description (string)
     - transaction_date (string)
     - us_dollar_amount (string)
   - **Response:**
     - 200 OK - Success
        ```json
       {
       "transaction_id": "253b647e-4ba7-4fe2-9305-66b04bb4afb4"
       }
        ```
     - 500 Internal Server Error
        ```json
       {
       "error": "internal error"
       }
        ```
       Other standard HTTP response codes are also possible outcomes !
   - Example curl :
   `curl --location 'http://127.0.0.1:8080/v1/api/transaction' \
     --header 'Content-Type: application/json' \
     --data '{
     "description" : "transaction 4322",
     "transaction_date" : "2023-10-12",
     "us_dollar_amount" : "1.86"
     }'`

2. Retrieve transaction
   - **URL:** `/v1/api/transaction?transaction_id={transaction_id}&country={country}`
   - **Method:** `GET`
   - **Request Parameter:**
      - transaction_id (string)
      - country (string)
   - **Response:**
       - 200 OK - Success
          ```json
         {
         "transaction_id": "253b647e-4ba7-4fe2-9305-66b04bb4afb4",
         "transaction_date": "2022-10-12",
         "us_dollar_amount": "1.86",
         "exchange_rate": "83.071",
         "converted_amount": "154.51"
         }
          ```
       - 404 Not found
          ```json
         {
         "error": "trx not found"
         }
          ```
         Other standard HTTP response codes are also possible outcomes !
   - Example curl :
   `curl --location 'http://127.0.0.1:8080/v1/api/transaction?transaction_id=fd23d3b9-216f-4f2a-ab46-6e4192459721&country=Zimbabwe'`

Postman Collection can be found at root of project