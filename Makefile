all:
	export GOPATH=`pwd`
	go build ./...

run:
	go run src/main.go

test:
	go test ./... -v --count 10 --cover

clean:
	rm -rf main *.jpg *.png
