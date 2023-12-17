BIN_DIR = bin
PROTO_DIR = Partus_users/api
SERVER_DIR = server
CLIENT_DIR = client

ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SHELL_VERSION = $(shell (Get-Host | Select-Object Version | Format-Table -HideTableHeaders | Out-String).Trim())
	SYS = $(shell "{0} {1}" -f "windows", (Get-ComputerInfo -Property OsVersion, OsArchitecture | Format-Table -HideTableHeaders | Out-String).Trim())
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
	CHECK_DIR_CMD = if (!(Test-Path $@)) { $$e = [char]27; Write-Error "$$e[31mDirectory $@ doesn't exist$${e}[0m" }
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
	RM_RF_CMD = ${RM_F_CMD} -Recurse
	SERVER_BIN = ${SERVER_DIR}.exe
	CLIENT_BIN = ${CLIENT_DIR}.exe
else
	SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		SYS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		SYS = linux ${VERSION_AND_ARCH}
	else
		$(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
	CHECK_DIR_CMD = test -d $@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	SERVER_BIN = ${SERVER_DIR}
	CLIENT_BIN = ${CLIENT_DIR}
endif

.DEFAULT_GOAL := help
.PHONY: partus_users clean-partus_users run-server-partus_users

partus_users: ## Generate Go code from .proto files for partus_users
	@${CHECK_DIR_CMD}
	protoc -I${PROTO_DIR} --go_out=${PROTO_DIR} --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:${PROTO_DIR} --go-grpc_opt=paths=source_relative ${PROTO_DIR}/*.proto
	cd Partus_users; if ($$?) { go build -o ./bin/server.exe ./cmd/server/main.go }

run-server-partus_users: partus_users ## Run the server for partus_users
ifeq ($(OS),Linux)
	./Partus_users/${BIN_DIR}/${SERVER_BIN}
else
	./Partus_users/${BIN_DIR}/${SERVER_BIN}
endif

clean-partus_users: ## Clean generated files for partus_users
	${RM_F_CMD} ${PROTO_DIR}/*.pb.go
	${RM_F_CMD} ${BIN_DIR}/${SERVER_BIN}
	${RM_F_CMD} ${BIN_DIR}/${CLIENT_BIN}

help: ## Show this help
	@${HELP_CMD}

up-partus_users: ## Start the partus_users service in development
	docker-compose -f ../PARTUS_PROJECT/docker-compose.yml up partus_users

down: ## Stop the service
	docker-compose -f ../PARTUS_PROJECT/docker-compose.yml down