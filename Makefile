.PHONY: clean compile web-dev server-dev

SERVERDIR = server/pkg/cmd/molamola-server
SERVERBIN = molamola

clean:
	rm -f ${SERVERDIR}/${SERVERBIN}
	${MAKE} -C web clean

compile-server:
	cd ${SERVERDIR} && \
		go build -o ${SERVERBIN} .

web-dev:
	${MAKE} -C web dev

server-dev: compile-server
	mkdir -p /tmp/molamola
	rm -rf /tmp/molamola/*
	./${SERVERDIR}/${SERVERBIN} -data /tmp/molamola
