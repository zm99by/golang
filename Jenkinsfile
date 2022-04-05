pipeline {
    agent any
    tools {
        go 'Go 1.18'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go mod init'
                sh 'go build'
            }
        }
    }
}
