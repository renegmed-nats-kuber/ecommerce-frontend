.PHONY: build
build:
	go build -o ecommerce-frontend .

.PHONY: docker
docker:
	docker build -t ecommerce-frontend .

.PHONY: run  
run:
	docker run --name ecommerce-frontend --net host -d -p 8282:8080 -e PORT="8080" ecommerce-frontend

.PHONY: tag
tag:
	docker tag ecommerce-frontend:latest renegmedal/ecommerce-frontend:1.0.1

.PHONY: push
push:
	docker push renegmedal/ecommerce-frontend:1.0.1

.PHONY: up
up:
	docker-compose up --build -d 
