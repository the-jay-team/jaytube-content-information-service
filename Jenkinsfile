pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS=credentials('thejayteam-docker-hub-credentials')
    }

    stages {
        stage('Test') {
            tools { go '1.19' }
            steps {
                sh 'go test ./test/...'
            }
        }
        stage('Build stage') {
            when {
                branch 'develop'
            }
            steps {
                sh 'docker build -t thejayteam/content-information-service:stage .'
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh 'docker push thejayteam/content-information-service:stage'
            }
        }
        stage('Build productive') {
            when {
                branch 'main'
            }
            steps {
                script {
                    env.VERSION_TAG = sh(returnStdout: true, script: 'git tag --points-at HEAD')
                    if (env.VERSION_TAG == "") {
                        env.VERSION_TAG = sh(returnStdout: true, script: 'git rev-parse HEAD').substring(0, 15)
                    } else if (env.VERSION_TAG.startsWith("v")) {
                        env.VERSION_TAG = env.VERSION_TAG.substring(1)
                    }
                }
                sh 'docker build -t thejayteam/content-information-service:latest -t thejayteam/content-information-service:$VERSION_TAG .'
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh 'docker push thejayteam/content-information-service:latest'
                sh 'docker push thejayteam/content-information-service:$VERSION_TAG'
            }
        }
    }
}