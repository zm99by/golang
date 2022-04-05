pipeline {
    agent any
    tools {
        go 'Go 1.11'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Compile') {
            steps {
                sh 'go build'
            }
        }
    }
}
