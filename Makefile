.phony:gen
gen:
	go generate ./...

.phony:install
install:
	go install ./...