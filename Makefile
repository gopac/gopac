BIN_PATH=bin/gopac
INSTALL_PATH=/usr/local/bin/gopac
LIBGIT_BUILD_PATH=vendor/git2go/script

all: gopac

dependencies: vendor/
	git submodule init
	git submodule update
	git submodule foreach git pull origin master

# add to gopac and move actual commands
# into this file instead of an external
libgit2: ${LIBGIT_BUILD_PATH}/build-libgit2.sh
	cd ${LIBGIT_BUILD_PATH} && ./build-libgit2.sh
	

gopac: dependencies gopac.go
	go build -o ${BIN_PATH}

install:

clean: ${BIN_PATH}
	rm -f ${BIN_PATH}