run:
	go run cmd/server.go

snap:
	@curl localhost:8080/snap
