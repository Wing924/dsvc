all: test build

test:
	go list ./... | xargs -n1 go test -race

cover:
	go list ./... | xargs -n1 go test -cover

build:
	go build github.com/Wing924/dsvc/cmd/dsvc

clean:
	rm -f dsvc