# Clones docker network driver repo and runs tests

echo "Cloning code repo github.com/codilime/contrail-windows-docker "
New-Item -ItemType directory -Path C:\go_workspace\src\github.com
New-Item -ItemType directory -Path C:\go_workspace\src\github.com\codilime
git clone https://github.com/codilime/contrail-windows-docker -b vendors C:\go_workspace\src\github.com\codilime\contrail-windows-docker

echo "Building docker network driver"
cd C:\go_workspace\bin
go build github.com/codilime/contrail-windows-docker

echo "Downloading test framework"
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

echo "Running tests"
cd C:\go_workspace\src\github.com\codilime\contrail-windows-docker
C:\go_workspace\bin\ginkgo.exe -r .
