SHELL := /bin/bash

.PHONY: help
help:
	@echo "Usage make <TARGET>"
	@echo ""
	@echo "  Targets:"
	@echo "	   build				Build the todo executable"
	@echo "	   run					Run the todo program from code"
	@echo "	   run-bin				Run the todo executable"
	@echo "	   load-db				Add sample data via curl"
	@echo "	   get-by-voterid		Get a todo by id pass id=<id> on command line"
	@echo "	   get-polls		    Get a todo by id pass id=<id> on command line"
	@echo "	   get-by-pollid		Get a todo by id pass id=<id> on command line"
	@echo "	   get-all				Get all todos"
	@echo "	   update-by-voterid	Update record 2, pass a new title in using title=<title> on command line"
	@echo "	   update-by-pollid		Update record 2, pass a new title in using title=<title> on command line"
	@echo "	   delete-all			Delete all todos"
	@echo "	   delete-by-voterid	Delete a todo by id pass id=<id> on command line"
	@echo "	   delete-by-voterid	Delete a todo by id pass id=<id> on command line"
	@echo "	   get-v2				Get all todos by done status pass done=<true|false> on command line"
	@echo "	   get-v2-all			Get all todos using version 2"
	@echo "	   build-amd64-linux	Build amd64/Linux executable"
	@echo "	   build-arm64-linux	Build arm64/Linux executable"





.PHONY: build
build:
	go build .

.PHONY: build-amd64-linux
build-amd64-linux:
	GOOS=linux GOARCH=amd64 go build -o ./todo-linux-amd64 .

.PHONY: build-arm64-linux
build-arm64-linux:
	GOOS=linux GOARCH=arm64 go build -o ./todo-linux-arm64 .

	
.PHONY: run
run:
	go run main.go

.PHONY: run-bin
run-bin:
	./todo

.PHONY: restore-db
restore-db:
	(cp ./data/todo.json.bak ./data/todo.json)

.PHONY: restore-db-windows
restore-db-windows:
	(copy.\data\todo.json.bak .\data\todo.json)

.PHONY: load-db
load-db:
	curl -d '{"voter_id":1,"name":"John Doe-1","email":"johnd1@example.com","vote_history":[{"poll_id":10,"vote_id":100,"vote_date":"2024-02-25T09:00:00Z"},{"poll_id":20,"vote_id":21,"vote_date":"2024-02-26T10:00:00Z"},{"poll_id":30,"vote_id":32,"vote_date":"2024-02-26T10:00:00Z"}]}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
	curl -d '{"voter_id":2,"name":"John Doe-2","email":"johnd2@example.com","vote_history":[{"poll_id":10,"vote_id":200,"vote_date":"2024-02-25T09:00:00Z"},{"poll_id":30,"vote_id":31,"vote_date":"2024-02-26T10:00:00Z"}]}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
	curl -d '{"voter_id":3,"name":"John Doe-3","email":"johnd3@example.com","vote_history":[{"poll_id":10,"vote_id":300,"vote_date":"2024-02-25T09:00:00Z"},{"poll_id":40,"vote_id":41,"vote_date":"2024-02-26T10:00:00Z"}]}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
	curl -d '{"voter_id":4,"name":"John Doe-4","email":"johnd4@example.com","vote_history":[{"poll_id":10,"vote_id":400,"vote_date":"2024-02-25T09:00:00Z"},{"poll_id":50,"vote_id":51,"vote_date":"2024-02-26T10:00:00Z"}]}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
	curl -d '{"voter_id":5,"name":"John Doe-5","email":"johnd5@example.com","vote_history":[{"poll_id":10,"vote_id":500,"vote_date":"2024-02-25T09:00:00Z"},{"poll_id":60,"vote_id":61,"vote_date":"2024-02-26T10:00:00Z"}]}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
	curl -d '{"voter_id":6,"name":"John Doe-6","email":"johnd6@example.com","vote_history":[{"poll_id":10,"vote_id":600,"vote_date":"2024-02-25T09:00:00Z"},{"poll_id":70,"vote_id":71,"vote_date":"2024-02-26T10:00:00Z"}]}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
	

.PHONY: update-by-voterid
update-by-voterid:
	curl -d '{ "voter_id":$(id),"name": "updated name", "email": "updated@email.com","vote_history":[{"poll_id":77,"vote_id":88,"vote_date":"2024-02-26T00:00:00Z"},{"poll_id":78,"vote_id":89,"vote_date":"2024-02-26T00:00:00Z"}]}' -H "Content-Type: application/json" -X PUT http://localhost:1080/voters/$(id) 

.PHONY: update-by-pollid
update-by-pollid:
	curl -d '{ "poll_id":$(pollid),"vote_id":88,"vote_date":"2099-12-31T23:23:59Z"}' -H "Content-Type: application/json" -X PUT http://localhost:1080/voters/$(id)/polls/$(pollid)


.PHONY: get-by-voterid
get-by-voterid:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/voters/$(id) 

.PHONY: get-polls
get-polls:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/voters/$(id)/polls 

.PHONY: get-by-pollid
get-by-pollid:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/voters/$(id)/polls/$(pollid) 

.PHONY: get-all
get-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/voters 

.PHONY: delete-all
delete-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1080/voters 

.PHONY: delete-by-voterid
delete-by-voterid:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1080/voters/$(id) 

.PHONY: delete-by-pollid
delete-by-pollid:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X DELETE http://localhost:1080/voters/$(id)/polls/$(pollid) 	

.PHONY: get-v2
get-v2:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/v2/todo?done=$(done) 

.PHONY: get-v2-all
get-v2-all:
	curl -w "HTTP Status: %{http_code}\n" -H "Content-Type: application/json" -X GET http://localhost:1080/v2/todo
