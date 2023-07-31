redis:
	docker run --name redis-url -p 6379:6379 -d redis:7.2-rc3-alpine

start:
	docker container start redis-url

.PHONY: redis start