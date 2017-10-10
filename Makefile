.PHONY: clean compile web-dev server-dev

SERVERDIR = server/pkg/cmd/molamola-server
SERVERBIN = molamola
CLIDIR = server/pkg/cmd/molamola-cli
CLIBIN = molamolactl

clean:
	rm -f ${SERVERDIR}/${SERVERBIN}
	rm -f ${CLIDIR}/${CLIBIN}
	${MAKE} -C web clean

compile-cli:
	cd ${CLIDIR} && \
		go build -o ${CLIBIN} .

compile-server:
	cd ${SERVERDIR} && \
		go build -o ${SERVERBIN} .

web-dev:
	${MAKE} -C web dev

server-dev: compile-server compile-cli
	mkdir -p /tmp/molamola
	rm -rf /tmp/molamola/*
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user add mola --password mola
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user remove mola
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user add mola --password mola
	./${CLIDIR}/${CLIBIN} --data /tmp/molamola user
	./${SERVERDIR}/${SERVERBIN} --data /tmp/molamola
