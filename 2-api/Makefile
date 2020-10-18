.PHONY: build clean deploy

build.movie:
	pushd movie-app/
		env GOOS=linux go build ./ -ldflags="-s -w" -o bin/main
	popd

build.api:
	pushd api-app/
		env GOOS=linux go build ./ -ldflags="-s -w" -o bin/main
	popd

build:
	make build.movie
	make build.api


test:
	go test ./... -cover

clean:
	rm -rf ./bin

docker:
	docker-compose up --build -d
