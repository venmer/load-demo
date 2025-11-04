app-run:
	docker compose -f docker-compose.app.yaml up -d

app-build:
	docker compose -f docker-compose.app.yaml up -d --build

app-stop:
	docker compose -f docker-compose.app.yaml stop

app-down:
	docker compose -f docker-compose.app.yaml down