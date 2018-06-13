APP=go
VERSION=`date +'%Y-%m-%d_%H:%M:%S'`
FLAGS="-s -w -X main.Version=${VERSION}"
build:
	go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS}
deploy-test:
	GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS}
	scp athena admin@myathena.cn:/tmp
	ssh admin@myathena.cn "cd /var/local/athena/test && docker-compose stop athena athena-sms-worker && docker-compose rm -f && cp /tmp/athena ./ && docker-compose up -d"
deploy-prod:
	GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS}
	scp athena admin@myathena.cn:/tmp
	ssh admin@myathena.cn "cd /var/local/athena/prod && docker-compose stop athena athena-sms-worker && docker-compose rm -f && cp /tmp/athena ./ && docker-compose up -d"
run-server:
	go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS} && ./${APP} server
run-worker:
	go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS} && ./${APP} worker
clean:
	@rm ${APP}
