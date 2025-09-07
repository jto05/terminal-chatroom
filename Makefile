build:
	go build -o ./cmd/client_cmd.exe ./cmd/client/client_cmd.go
	go build -o ./cmd/server_cmd.exe ./cmd/server/server_cmd.go

clean:
	find cmd -type f -name "*.exe" -exec rm -f {} +

