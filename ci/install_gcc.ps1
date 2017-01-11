# GCC INSTALLATION

# based on:
# https://github.com/docker/docker/blob/master/Dockerfile.windows

Function Download-File([string] $source, [string] $target)  { $wc = New-Object net.webclient; $wc.Downloadfile($source, $target) }

echo "Downloading gcc components"
Download-File https://raw.githubusercontent.com/jhowardmsft/docker-tdmgcc/master/gcc.zip C:\gcc.zip
Download-File https://raw.githubusercontent.com/jhowardmsft/docker-tdmgcc/master/runtime.zip C:\runtime.zip
Download-File https://raw.githubusercontent.com/jhowardmsft/docker-tdmgcc/master/binutils.zip C:\binutils.zip

echo "Installing gcc components"
Expand-Archive C:\gcc.zip -DestinationPath C:\gcc -Force
Expand-Archive C:\runtime.zip -DestinationPath C:\gcc -Force
Expand-Archive C:\binutils.zip -DestinationPath C:\gcc -Force

echo "Adding C:\gcc\bin to path"
$Env:Path += ";C:\gcc\bin"
$p = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
$p = $p.split(';') | sort -unique
[system.String]::Join(";", $p)
[Environment]::SetEnvironmentVariable("Path", $p+";C:\gcc\bin", [EnvironmentVariableTarget]::Machine)
