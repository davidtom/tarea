APP_NAME := tarea

build:
	go build -o ${APP_NAME} main.go

clean:
	rm -f ${APP_NAME}

bin:
	go install

# TODO: add other tasks from go-t repo (ie lint and fmt)