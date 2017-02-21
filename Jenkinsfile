#!/usr/bin/env groovy



def powershell(script, returnStatus) { 
  pshell = "powershell.exe"
  base64script = script.bytes.encodeBase64().toString()
  bat(script: "${pshell} -NoProfile  -NonInteractive -NoLogo -ExecutionPolicy Bypass -EncodedCommand ${base64script}", returnStatus: $returnStatus)
}

def gotool(tool, args) {
  
}

def getBranch() {
    env.BRANCH_NAME.tokenize("/").last()
}

@NonCPS
def getWorkspace(buildType) {
  def directory_name = pwd().tokenize("\\").last()

  pwd().replace("%2F", "_") + buildType
}

node('windows-server-2016') {
  ws(getWorkspace("")){
    timestamps{
      try {
        def branch = getBranch()
        withEnv(["GOPATH=${pwd()}",
                 "PATH+GOPATH=${pwd()}\\bin"]){
          dir('src'){
            dir('github.com'){
              dir('codilime'){
                dir("contrail-windows-docker"){
                  stage 'checkout'
                  checkout scm
                  stage 'prepare deps'
                  bat script: "go get -u github.com/docker/docker/client"
                  bat script: "go get -u github.com/docker/docker/api/types"
                  bat script: "go get -u github.com/onsi/ginkgo/ginkgo"
                  bat script: "go get -u github.com/onsi/gomega"
                  bat script: "go get -u github.com/onsi/ginkgo/extensions/table"
                  stage 'clean'
                  bat script: "go clean -i -r "
                  stage 'build'
                  bat script: "go build -i -v"
                  stage 'test'
                  echo 'gingko.exe -r .'
                }

              }
            }
          }
        }
        stage 'archive'
        archiveArtifacts artifacts: 'bin/**/*', fingerprint: true
        stage 'cleanup'
        step([$class: 'WsCleanup'])
      }
      catch (error) {
        throw error
      }
    } 
  }
}
