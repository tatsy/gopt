all:
	export GOPATH=`pwd`
	go build ./...

run:
	export GOPATH=`pwd`
	go build ./src/main.go && ./main

test:
	go test ./... --cover

clean:
	rm -rf main *.jpg *.png *.prof
