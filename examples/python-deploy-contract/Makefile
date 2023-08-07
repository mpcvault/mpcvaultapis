# Makefile

SHELL := /bin/bash
PYTHONPATH := ./
PROJECT_DIR := ./mpcvaultapis
REMOTE_REPO := https://github.com/mpcvault/mpcvaultapis.git
BRANCH = main
PROTO_FILE := mpcvaultapis/mpcvault/platform/v1/api.proto

.PHONY: install repo compile-proto

all: install repo compile-proto

install:
	@echo "------ Installing python dependencies ------"
	pip install -r requirements.txt

repo:
	@echo "------ Cloning the repository ------"
	@if [ -d "$(PROJECT_DIR)" ]; then echo "$(PROJECT_DIR) already exists"; else git clone -b $(BRANCH) $(REMOTE_REPO) $(PROJECT_DIR); fi

compile-proto:
	@echo "------ Compiling the proto files ------"
	python3 -m grpc_tools.protoc --proto_path=. ./$(PROTO_FILE) --python_out=. --grpc_python_out=.