PROGRAM_NAME := r36s-hardware-test
DEPLOY_PATH := /userdata/roms/native
IP := 192.168.0.107
USN := root
PWD := linux

all: clean docker deploy

clean:
	rm bin/${PROGRAM_NAME}.exec -f

docker:
	#docker run -d --name arkos-sdk -c 1024 -it --volume=/home/vitaly/GolandProjects/:/work/ --workdir=/work/ arkos-sdk
	docker exec arkos-sdk /bin/bash -c 'cd ${PROGRAM_NAME} && make build'

build:
	go build -o bin/${PROGRAM_NAME}.exec ${PROGRAM_NAME}/src/

deploy:
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "/etc/init.d/S31emulationstation stop"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "rm ${DEPLOY_PATH}/${PROGRAM_NAME}.exec -f"
	sshpass -p ${PWD} scp -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null bin/${PROGRAM_NAME}.exec ${USN}@${IP}:${DEPLOY_PATH}
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "chmod 777 ${DEPLOY_PATH}/${PROGRAM_NAME}.exec"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "pkill -f ${PROGRAM_NAME}.exec"
	sshpass -p ${PWD} ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null ${USN}@${IP} "sh -c 'cd /tmp; ${DEPLOY_PATH}/${PROGRAM_NAME}.exec'" &


