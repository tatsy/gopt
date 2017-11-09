all:
	go build src/main.go

run:
	go run src/main.go

test:
	go test ./src/... --count 10 --cover

clean:
	rm -rf main image.jpg
