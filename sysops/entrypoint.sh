wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# Watch changes in go files
CompileDaemon --build="go build -o main main.go" --command=./main