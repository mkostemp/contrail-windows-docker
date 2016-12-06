cp -Recurse C:\go\src\github.com\codilime\contrail-windows-docker C:\go\src\github.com\codilime\contrail-windows-docker-build
go build github.com/codilime/contrail-windows-docker-build
cp C:\contrail-windows-docker-build.exe C:\output\contrail-windows-docker-build.exe