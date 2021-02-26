APP_NAME ?= gin-basic
APP_PORT ?= 8080

build:
	docker build -t ${APP_NAME} .

start-build:
	docker run -p ${APP_PORT}:${APP_PORT} \
		${APP_NAME}

start:
	go mod tidy && go run .

publish:
	git push heroku
