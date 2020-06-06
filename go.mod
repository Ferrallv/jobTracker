module jobtracker

go 1.14

require (
	github.com/jackc/pgx/v4 v4.6.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	jobtracker/models v0.0.0
)

replace jobtracker/models => ./models

