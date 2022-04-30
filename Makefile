NAME = cloud-native-go:1.1.0

default: build

docker:
	docker build -t $(NAME) .

build: 
	CGO_ENABLED=0 go build -v ./...

install: 
	CGO_ENABLED=0 go install -v ./...

clean:
	rm cloud-native-go