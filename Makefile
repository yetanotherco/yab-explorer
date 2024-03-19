.PHONY: create_env clean deps build run

create_env:
	@echo "🚚 Creating environment variables files .env"
	cp .env.example .env
	echo "✅ Environment variables files created successfully!"

clean:
	@echo "🗑️ Cleaning..."
	rm -rf yab-explorer
	@echo "✅ Done"

deps:
	go install ./...

build: 
	@echo "🏗️ Building..."
	make clean 
	go build -o yab-explorer ./cmd

run:
	@echo "🧑‍💻 Starting API server on port $(API_PORT)..."
	make build 
	./yab-explorer