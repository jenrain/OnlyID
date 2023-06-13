PLATFORM=$(shell uname -m)
DATETIME=$(shell date "+%Y%m%d%H%M%S")

onlyId:
	go build -o cmd/onlyId cmd/main.go
	chmod +x cmd/onlyId
clean:
	$(RM) tmp/* $(TARGET)
# ./onlyId -conf ./only_id.toml