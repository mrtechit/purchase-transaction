# purchase-transaction
Repository for purchase transaction

## How to run (Without Docker)
1. Clone the repository
2. Run postgresSQL locally
3. Update .env file
4. This project is implemented using Go. If Go is not yet installed, please download and install from [here](https://golang.org/doc/install)
5. Run command at root of repo : `go run main.go`

## API
There are 2 API's :
1. Store transaction
   `curl --location 'http://127.0.0.1:8080/v1/api/transaction' \
   --header 'Content-Type: text/plain' \
   --data '{
   "description" : "transaction 3",
   "transaction_date" : "2023-10-12",
   "us_dollar_amount" : "1.86"
   }'`
2. Retrieve transaction
   `curl --location 'http://127.0.0.1:8080/v1/api/transaction?transaction_id=fd23d3b9-216f-4f2a-ab46-6e4192459721&country=Zimbabwe'`

