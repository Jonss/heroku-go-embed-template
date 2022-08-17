new-project:
	grep -ilr 'heroku-go-embed-template' * | xargs -I@ sed -i '' 's/heroku-go-embed-template/$(name)/g' @

build:
	CGO_ENABLED=0 GOOS=linux go build -o app/bin main.go

test:
	go test -v --cover ./pkg/...

build-ui:
	cd ui && yarn build

run: build-ui
	go run main.go

env-up:
	docker-compose up --build -d

env-down:
	docker-compose down --remove-orphans
