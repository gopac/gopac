BIN_PATH=bin/gopac
INSTALL_PATH=/usr/local/bin/gopac

all: gopac

dependencies: vendor/
	git submodule init
	git submodule update
	git submodule foreach git pull origin master

gopac: dependencies gopac.go
	go build -o ${BIN_PATH}

install:
	cp bin/gopac /usr/local/bin/gopac
	&echo "Gopac has been installed!"

clean: ${BIN_PATH}
	rm -f ${BIN_PATH}