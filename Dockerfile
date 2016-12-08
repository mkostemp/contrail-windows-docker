# This dockerfile will build a remote driver for Windows.
# nativebuildimage is required for this (has all dependencies resolved).
# See Dockerfile.windows in official docker repo to see how
# to obtain it.
#
# Steps to build (similar to instructions in Dockerfile.windows):
# >> git clone https://github.com/codilime/contrail-windows-docker c:\go\src\github.com\codilime/contrail-windows-docker
# >> cd c:\go\src\github.com\codilime/contrail-windows-docker
# >> docker build -t driverbuildimage .
# >> docker run --entrypoint cmd --rm -v <OUTPUT_DIR>:C:\output -v <SOURCE_DIR>:C:\go\src\github.com\codilime\contrail-windows-docker driverbuildimage /c powershell C:\generate.ps1

FROM nativebuildimage

RUN mkdir C:\output

COPY generate.ps1 C:/generate.ps1

VOLUME 'C:/go/src/github.com/codilime/contrail-windows-docker'
VOLUME 'C:/output'

ENTRYPOINT "cmd"
# Workaround, because GO Panics when following symlinks...
CMD ['/c', 'powershell', 'C:\generate.ps1']