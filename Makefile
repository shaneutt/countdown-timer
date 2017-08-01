all:
	@go build

clean:
	@rm -f ./countdown-timer

container: all
	@docker build -t countdown-timer .

.PHONY: all container
