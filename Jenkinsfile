pipeline {
    agent any
    
    /*
    environment {
        CI='true'
    }
    */

    tools {
        go 'go-1.12.1'
    }

stages {
  stage('Test') {
    steps {
      // for rice
      withEnv(["PATH+EXTRA=${HOME}/go/bin"]){
        sh 'go get github.com/GeertJohan/go.rice/rice'
        sh 'make build'
      }
    }
  }
}
  }
