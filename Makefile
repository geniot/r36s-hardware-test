PROGRAM_NAME := r36s-hardware-test
DEPLOY_PATH := /roms/native

all: clean build deploy

clean:
	rm bin/${PROGRAM_NAME}.exec -f

build:
	go build -o bin/${PROGRAM_NAME}.exec ${PROGRAM_NAME}/src/

deploy:
	sudo rm ${DEPLOY_PATH}/${PROGRAM_NAME}.exec -f
	sudo cp bin/${PROGRAM_NAME}.exec ${DEPLOY_PATH}
	sudo pkill -f ${PROGRAM_NAME}.exec
	sh -c 'cd /tmp; ${DEPLOY_PATH}/${PROGRAM_NAME}.exec'