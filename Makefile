APP_NAME ?= gin-basic
APP_TAG ?= master
APP_PORT ?= 8080

build:
	docker build -t ${APP_NAME} .

start:
	docker run -p ${APP_PORT}:${APP_PORT} \
		${APP_NAME}


