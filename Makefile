DIRECTORIES_TO_BUILD := "./testpb"

proto:
	docker build -t dev:proto-build -f Dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/genproto.sh "$(DIRECTORIES_TO_BUILD)"
	./scripts/patch.sh $(CURDIR)

.PHONY: proto fastreflection