pipeline {
    agent any
    tools {
        go 'Go 1.18'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Init') {
            steps {
                sh 'go mod init'
            }
        }
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
    }
}
