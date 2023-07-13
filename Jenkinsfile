pipeline {
    agent none

    stages {
        stage("docker build") {
            agent {
                node {
                    label 'ruconnext-dev'   
                }
            }
            steps {
                echo 'building...'
            }
        }

        stage("testing") {
            agent {
                node {
                    label 'ruconnext-uat'   
                }
            }
            steps {
                 echo 'testing...'
            }
        }

       stage("deploy") {
            agent {
                node {
                    label 'ruconnext-uat'   
                }
            }
            steps {
                 echo 'deploying...'
            }
        }
    }
    
    post {
        success {
            echo "Release Success"
        }
        failure {
            echo "Release Failed"
        }
    }
}
