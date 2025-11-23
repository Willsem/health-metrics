run:
	@dotenv -e ./config/local.env -- go run ./cmd/service/main.go

ent_generate:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./internal/generated/ent/schema
