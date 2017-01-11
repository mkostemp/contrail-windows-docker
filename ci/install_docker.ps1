# DOCKER INSTALLATION

# according to:
# https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick-start/quick-start-windows-10

echo "Downloading docker"
Invoke-WebRequest "https://test.docker.com/builds/Windows/x86_64/docker-1.13.0-rc4.zip" -OutFile "$env:TEMP\docker-1.13.0-rc4.zip" -UseBasicParsing

echo "Installing docker"
Expand-Archive -Path "$env:TEMP\docker-1.13.0-rc4.zip" -DestinationPath $env:ProgramFiles

echo "Adding ProgramFiles\Docker to env"
$Env:Path += ";$env:ProgramFiles\Docker"
$p = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
$p = $p.split(';') | sort -unique
[system.String]::Join(";", $p)
[Environment]::SetEnvironmentVariable("Path", $p+";$env:ProgramFiles\Docker", [EnvironmentVariableTarget]::Machine)

echo "Starting docker"
dockerd --register-service
Start-Service Docker

echo "Pulling test image (nanoserver)"
docker pull microsoft/nanoserver
