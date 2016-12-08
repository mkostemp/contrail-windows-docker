echo "Preparing the build..."
cp -Recurse C:\go\src\github.com\codilime\contrail-windows-docker C:\go\src\github.com\codilime\contrail-windows-docker-build
echo "Building!"
go build github.com/codilime/contrail-windows-docker-build
echo "Built finished, copying result"
cp C:\contrail-windows-docker-build.exe C:\output\contrail-windows-docker.exe
echo "Result is ready!"