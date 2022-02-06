wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# Watch changes in grpc files
CompileDaemon --build="go build -o ./grpc/main ./grpc/main.go" --command=./grpc/main