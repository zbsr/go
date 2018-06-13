APP=go
VERSION=`date +'%Y-%m-%d_%H:%M:%S'`
FLAGS="-s -w -X main.Version=${VERSION}"
build:
	go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS}
deploy-test:
	GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS}
	scp quiz admin@myathena.cn:/tmp
	ssh admin@myathena.cn "cd /var/local/quiz/test && docker-compose stop quiz quiz-sms-worker && docker-compose rm -f && cp /tmp/quiz ./ && docker-compose up -d"
deploy-prod:
	GOOS=linux GOARCH=amd64 go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS}
	scp quiz admin@myathena.cn:/tmp
	ssh admin@myathena.cn "cd /var/local/quiz/prod && docker-compose stop quiz quiz-sms-worker && docker-compose rm -f && cp /tmp/quiz ./ && docker-compose up -d"
run-server:
	go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS} && ./${APP} server
run-worker:
	go build -tags=jsoniter -o ${APP} -ldflags ${FLAGS} && ./${APP} worker
clean:
	@rm ${APP}
