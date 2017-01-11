# INSTALL GIT

# according to:
# https://docs.docker.com/opensource/project/software-req-win/
echo ""
Write-Host "Checking git version        "  -nonewline
Try {
  $version = git version
  Write-Host "  [ $version ]" -f green
}
Catch
{
  Write-Host "  [ MISSING ]" -f red
  echo "Start installing git"
  echo "  Downloading git"
  Invoke-Webrequest "https://github.com/git-for-windows/git/releases/download/v2.7.2.windows.1/Git-2.7.2-64-bit.exe" -OutFile C:\git.exe -UseBasicParsing

  echo "  Running git installer"
  Start-Process C:\git.exe -ArgumentList '/VERYSILENT /SUPPRESSMSGBOXES /CLOSEAPPLICATIONS /DIR=c:\git\' -Wait

  echo "  Adding C:\git\cmd to path"
  $Env:Path += ";C:\git\cmd"

  $p = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine)
  $p = $p.split(';') | sort -unique
  [system.String]::Join(";", $p)

  [Environment]::SetEnvironmentVariable("Path", $p+";C:\git\cmd", [EnvironmentVariableTarget]::Machine)
}
