pipeline {
    agent any
    tools {
        go 'go-1.18'
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
