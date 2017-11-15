all:
	go build ./...

run:
	go run ./main.go ${ARGS}

test:
	go test ./... -cover -v
	
clean:
	go clean
	rm -rf main *.jpg *.png *.prof
