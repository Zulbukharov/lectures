minimock -i post/internal/app.Repository -o ./internal/repository/repository_mock.go

migrate create -ext sql -dir migrations -seq create_posts_table

docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=demo -p 5432:5432 -d postgres:14.2


export POSTGRESQL_URL='postgres://postgres:postgres@localhost:5432/demo?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path migrations up
