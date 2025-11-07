APP := docker compose -f docker-compose.app.yaml -f docker-compose.observability.yaml -p demo
MON := docker compose -f docker-compose.observability.yaml -p demo

.PHONY: db-up
db-up:
	$(APP) up -d postgres

.PHONY: app-up
app-up:
	$(APP) up -d

.PHONY: app-build
app-build:
	$(APP) up -d --build

.PHONY: app-stop
app-stop:
	$(APP) stop

.PHONY: app-down
app-down:
	$(APP) down

.PHONY: mon-up
mon-up:
	$(MON) up -d

.PHONY: logs
logs:
	$(APP) logs -f

.PHONY: ps
ps:
	$(APP) ps

.PHONY: clean
clean:
	$(APP) down -v --remove-orphans