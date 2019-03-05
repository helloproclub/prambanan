include .env
export

setup:
	go get -u github.com/satriahrh/mig
migration:
	mig create $(filter-out $@,$(MAKECMDGOALS)) -d $$GOPATH/src/github.com/helloproclub/prambanan/database/mysql/migrations
migrate:
	mig up $$MYSQL_DSN -d $$GOPATH/src/github.com/helloproclub/prambanan/database/mysql/migrations
rollback:
	mig down $$MYSQL_DSN -d $$GOPATH/src/github.com/helloproclub/prambanan/database/mysql/migrations
