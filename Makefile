.PHONY: db api

db:
	@command -v systemctl >/dev/null && (systemctl is-active --quiet docker || sudo systemctl start docker) || true
	@if docker info >/dev/null 2>&1; then \
		docker compose up -d db; \
	else \
		sudo docker compose up -d db; \
	fi

api:
	@cd backend && set -a && . ../.env && set +a && go run ./cmd/api