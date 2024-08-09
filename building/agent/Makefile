tag = cookbook/agent

all:
	$(error please pick a target)

build:
	docker build -t $(tag) .

shell:
	docker run --rm -it $(tag) /bin/sh

