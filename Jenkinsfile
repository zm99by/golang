pipeline {
    agent any
    tools {
        go 'Go 1.18'
    }
    options {
        // This is required if you want to clean before build
        skipDefaultCheckout(true)
    }
    environment {
        GO111MODULE = 'on'
    }

    stages {
        stage('Build') {
            steps {
                // Clean before build
                cleanWs()
                // We need to explicitly checkout from SCM here
                checkout scm
                echo "Building ${env.JOB_NAME}..."
            }
        }
        stage('Init') {
            steps {
                sh 'go mod init example.com/m'
            }
        }
        stage('Compile') {
            steps {
                sh 'go build main.go'
            }
        }

    }

}
