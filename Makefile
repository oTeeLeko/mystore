mariadb:
	docker run --name mariadb -p 3306:3306 -e MARIADB_USER=root -e MARIADB_PASSWORD=password -e MARIADB_ROOT_PASSWORD=password -d mariadb

createdb:
	docker exec -it mariadb mariadb -uroot -ppassword -e "CREATE DATABASE mystore;"

dropdb:
	docker exec -it mariadb mariadb -uroot -ppassword -e "DROP DATABASE mystore;"

migrateup:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/mystore?parseTime=true" -verbose up

migratedown:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/mystore?parseTime=true" -verbose down

migrateup1:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/mystore?parseTime=true" -verbose up 1

migratedown1:
	migrate -path db/migrations -database "mysql://root:password@tcp(127.0.0.1:3306)/mystore?parseTime=true" -verbose down 1

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: mariadb createdb dropdb migrateup migratedown sqlc server migrateup1 migratedown1