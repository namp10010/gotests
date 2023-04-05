.phony:gen
gen:
	go generate ./templates/gen.go

.phony:install
install:
	go install ./...