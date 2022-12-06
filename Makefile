CURRENT_DIR=$(shell pwd)

proto-gen:
	rm -rf genproto
	./gen-proto.sh ${CURRENT_DIR}