# GOLANG INSTALLATION

# based on:
# https://gist.github.com/andrewkroh/2c93f8a5953f6093a505

Function Download-File([string] $source, [string] $target)  { $wc = New-Object net.webclient; $wc.Downloadfile($source, $target) }

$goroot  = "C:\go"
$gopath  = "C:\go_workspace"

echo "Downloading go"
$version = 1.7.4
Download-File "https://storage.googleapis.com/golang/go1.7.4.windows-amd64.zip" C:\go.zip

echo "Expanding go archive"
Expand-Archive C:\go.zip -DestinationPath C:\ -Force

echo "Setting goroot"
$Env:GOROOT = $goroot
[Environment]::SetEnvironmentVariable("GOROOT", "$goroot", [EnvironmentVariableTarget]::Machine)

echo "Setting gopath"
$Env:GOPATH = $gopath
[Environment]::SetEnvironmentVariable("GOPATH", "$gopath", [EnvironmentVariableTarget]::Machine)

echo "Setting up go workspace in C:\go_workspace"
New-Item -ItemType directory -Path C:\go_workspace
New-Item -ItemType directory -Path C:\go_workspace\bin
New-Item -ItemType directory -Path C:\go_workspace\src
New-Item -ItemType directory -Path C:\go_workspace\pkg

echo "Adding goroot\bin and gopath\bin to path"
$Env:Path += ";$goroot\bin;$gopath\bin"

$p = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
$p = $p.split(';') | sort -unique
[system.String]::Join(";", $p)

[Environment]::SetEnvironmentVariable("Path", $p+";$goroot\bin;$gopath\bin", [EnvironmentVariableTarget]::Machine)
