# Database Migration
## Create Migration
```
migrate create -ext sql -dir migrations -seq create_account_table
```
## Execute Migration
```
migrate -path migrations -database "postgresql://admin:example@localhost:5432/example_db?sslmode=disable" up
```