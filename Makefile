PROGRAM_NAME := r36s-hardware-test
DEPLOY_PATH := /roms/native
IP := 192.168.0.105
USN := ark
PWD := ark

all: clean docker deploy

clean:
	rm bin/${PROGRAM_NAME}.exec -f

docker:
	#docker run -d --name arkos-sdk -c 1024 -it --volume=/home/vitaly/GolandProjects/r36s-hardware-test/:/work/ --workdir=/work/ --rm arkos-sdk
	docker exec arkos-sdk make build

build:
	go build -o bin/${PROGRAM_NAME}.exec ${PROGRAM_NAME}/src/

deploy:
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "sudo rm ${DEPLOY_PATH}/${PROGRAM_NAME}.exec -f"
	sshpass -p ${PWD} scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null bin/${PROGRAM_NAME}.exec ${USN}@${IP}:${DEPLOY_PATH}
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "sudo pkill -f ${PROGRAM_NAME}.exec"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "sh -c 'cd /tmp; ${DEPLOY_PATH}/${PROGRAM_NAME}.exec'" &


