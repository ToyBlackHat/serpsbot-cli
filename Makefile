serpsbot-cli: serpsbot.go
	go build -o serpsbot-cli serpsbot.go

deploy: serpsbot-cli
	cp serpsbot-cli ~/bin/serpsbot-cli
	@echo "serpsbot-cli installed in ~/bin/serpsbot-cli"
