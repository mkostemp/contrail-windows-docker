# Sets up environment from scratch and runs tests.

Invoke-Expression "./install_git.ps1"
Invoke-Expression "./install_golang.ps1"
Invoke-Expression "./install_gcc.ps1"
Invoke-Expression "./install_docker.ps1"
Invoke-Expression "./clone_and_test.ps1"
