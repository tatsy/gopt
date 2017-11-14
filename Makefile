all:
	go build ./...

run:
	go build ./main.go && ./main ${ARGS}

test:
	go test ./... --cover -v

clean:
	rm -rf main *.jpg *.png *.prof
