# This how we want to name the binary output
BINARY=tool

# These are the values we want to pass for VERSION and BUILD
VERSION=1.0.1
BUILD=`git rev-parse HEAD`
BUILDTIME=`date "+%Y-%m-%d"`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X tool/cmd.ToolVer=${VERSION} -X tool/cmd.Build=${BUILD} -X tool/cmd.BuildTime=${BUILDTIME}"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}-${VERSION}

# Installs our project: copies binaries
install:
	rm -f "${GOPATH}/bin/${BINARY}"
	go install ${LDFLAGS}
# Cleans our project: deletes binaries
clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install