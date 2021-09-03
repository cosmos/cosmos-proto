DIRECTORIES_TO_BUILD := "./testpb"

proto:
	docker build -t dev:proto-build -f Dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/genproto.sh "$(DIRECTORIES_TO_BUILD)"

fast:
	docker build -t dev:proto-build -f Dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/fastreflect.sh "$(DIRECTORIES_TO_BUILD)"

.PHONY: proto fastreflection