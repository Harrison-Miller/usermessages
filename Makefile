build:
	go build .
	cp usermessages server

docker:
	docker build -t usermessages .