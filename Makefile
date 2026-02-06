up:
	docker compose up -d --build

ps:
	docker compose ps -a

open-grafana:
	xdg-open http://localhost:3000

down:
	docker compose down --volumes
