DIRECTORIES_TO_BUILD := "./testpb"

pulsar:
	docker build -t dev:proto-build -f Dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build ./scripts/fastreflect.sh "$(DIRECTORIES_TO_BUILD)"

cosmos-proto:
	docker build -t dev:proto-build -f Dockerfile .
	docker run -v "$(CURDIR):/genproto" -w /genproto dev:proto-build protoc -I=. -I=./third_party/proto --go_out=paths=source_relative:. cosmos.proto

.PHONY: pulsar