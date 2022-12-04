.PHONY: test-unit
test-unit: 
	@echo "[go test] running unit tests and collecting coverage metrics"
	@go test -v -tags unit -race -coverprofile=coverage_unit.txt -covermode=atomic ./...

.PHONY: test-intregration
test-intregration:
	@echo "[go test] running integration tests and collecting coverage metrics"
	@go test -v -tags integration -race -coverprofile=coverage_integration.txt -covermode=atomic ./...

.PHONY: test-all
test-all: test-unit test-integration
	@echo "[go test] testing all"

.PHONY: cover
	@go tool cover -html=coverage_integration.txt
	@go tool cover -html=coverage_unit.txt