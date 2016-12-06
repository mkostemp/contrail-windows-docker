# This dockerfile will build a remote driver for Windows.
# nativebuildimage is required for this (has all dependencies resolved).
# See Dockerfile.windows in official docker repo to see how
# to obtain it.
#
# Steps to build (similar to instructions in Dockerfile.windows):
# >> git clone https://github.com/codilime/contrail-windows-docker c:\go\src\github.com\codilime/contrail-windows-docker
# >> cd c:\go\src\github.com\codilime/contrail-windows-docker
# >> docker build -t driverbuildimage -f Dockerfile .
# >> docker run --name remotedriver driverbuildimage sh -c 'cd /c/go/bin; go build github.com/codilime/contrail-windows-docker'
# >> docker cp remotedriver:c:\go\bin\docker-windows-driver.exe OUTPUT_PATH

FROM nativebuildimage
COPY . /go/src/github.com/codilime/contrail-windows-docker
