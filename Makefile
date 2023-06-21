PLATFORM=$(shell uname -m)
DATETIME=$(shell date "+%Y%m%d%H%M%S")

# mac环境编译
#onlyId:
#	go build -o cmd/onlyId cmd/main.go
#	chmod +x cmd/onlyId

# linux交叉编译
onlyId:
	GOOS=linux GOARCH=amd64 go build -o cmd/onlyId cmd/main.go

clean:
	$(RM) tmp/* $(TARGET)
# ./onlyId -conf ./only_id.toml