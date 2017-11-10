export GOPATH = $(PWD)


all:
	go build ./...

run:
	go build ./src/main.go && ./main

test:
	go test ./... --cover

clean:
	rm -rf main *.jpg *.png *.prof
