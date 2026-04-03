.PHONY: infra-up infra-down seed-db middleware-rebuild middleware-build show-redis compose-start compose-stop

middleware-build:
	docker build -t asigna-tenant-middleware:0.0.1 -f core/tenant-middleware/Dockerfile .

infra-up:
	docker compose -f deployments/local/docker-compose.yml -p asigna-development up -d

infra-down:
	docker compose -f deployments/local/docker-compose.yml -p asigna-development down -v

seed-db:
	@echo "Insertando datos en Tenants Central..."
	docker exec -i asigna-central-db psql -U asigna_admin -d asigna_db < deployments/local/init-scripts/01_init_catalog.sql

middleware-rebuild:
	docker compose -f deployments/local/docker-compose.yml up -d --build tenant-middleware

show-redis:
	docker exec -it asigna-redis redis-cli keys "*"

compose-start:
	docker compose -p asigna-development start

compose-stop:
	docker compose -p asigna-development stop