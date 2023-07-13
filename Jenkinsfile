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
                 dir('/home/ruconnext/ruconnext-dev') {
                    sh 'cd /home/ruconnext/ruconnext-dev'
                    sh 'ls /home/ruconnext/ruconnext-dev -a'
                    sh 'docker-compose down'
                    sh 'cp /home/ruconnext/ruconnext-dev/config.yaml /home/ruconnext/jenkins_agent/workspace/${JOB_NAME}/environments'
                    sh 'ls -la /home/ruconnext/jenkins_agent/workspace/${JOB_NAME}/environments'
                    sh 'docker rm $(docker ps -a -q -f status=exited)'
                    sh 'docker rmi $(docker images -f "dangling=true" -q)'
                    sh 'docker build -t ru-connext-api .'
                    sh 'cd /home/ruconnext/ruconnext-dev'
                    sh 'docker-compose up -d'
                    sh 'docker-compose up --scale ru-connext-api=4 -d'
                 }
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
