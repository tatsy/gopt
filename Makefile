export GOPATH = $(PWD)

all:
	go build ./...

run:
	go build ./src/main.go && ./main ${ARGS}

test:
	go test ./... --cover -v

clean:
	rm -rf main *.jpg *.png *.prof
