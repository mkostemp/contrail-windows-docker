# Clones docker network driver repo and runs tests

$branch=$args[0]

$repoExists = Test-Path C:\go_workspace\src\github.com\codilime\contrail-windows-docker
if($repoExists -eq $True) {
    echo "Checking out $branch from github.com/codilime/contrail-windows-docker"
    cd C:\go_workspace\src\github.com\codilime\contrail-windows-docker
    git pull origin $branch
    git checkout $branch
} else {
    echo "Cloning code repo github.com/codilime/contrail-windows-docker, branch $branch"
    New-Item -ItemType directory -Path C:\go_workspace\src\github.com
    New-Item -ItemType directory -Path C:\go_workspace\src\github.com\codilime
    git clone https://github.com/codilime/contrail-windows-docker -b $branch C:\go_workspace\src\github.com\codilime\contrail-windows-docker
}

echo "Building docker network driver"
cd C:\go_workspace\bin
go build github.com/codilime/contrail-windows-docker

echo "Downloading test framework"
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega

echo "Running tests"
cd C:\go_workspace\src\github.com\codilime\contrail-windows-docker

Start-Transcript -path C:\testresults.txt
C:\go_workspace\bin\ginkgo.exe -r .
Stop-Transcript
