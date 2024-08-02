# Makefile for Air Hot Reloading

# Variables
AIR_CMD=air

# Help
help:
	@echo "Usage:"
	@echo "    make install          Install air tool"
	@echo "    make init             Init .air.toml"
	@echo "    make run              Run air"

# Targets
install:
	@echo "Installing air..."
	go install github.com/cosmtrek/air@v1.50.0

init:
	@echo "Creating .air.toml"
	$(AIR_CMD) init

run:
	@echo "Running air"
	$(AIR_CMD)
