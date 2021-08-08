dockerImage = ''

pipeline{

    agent any
    
    tools {
        go 'go-1.16'
    }

    environment {
        GO111MODULE = 'on'
        registry = "someshdokerbox/golang-okta-token-handler"
        registryCredential = 'docker-hub' // Create Jenkins Crediential ID named 'docker-hub' (or anything you want and refrenece here)
    }

    stages{
        stage("Build"){
            steps {
                script {
                    sh "make build"
                }
            }
        }
        stage("Build Docker Image") {
            steps {
                script {
                    dockerImage = docker.build registry + ":latest"
                }
            }
        }
        stage("Push Image") {
            steps {
                script {
                        docker.withRegistry( '', registryCredential ) {
                        dockerImage.push()
                    }
                }
            }
        }
    }
}