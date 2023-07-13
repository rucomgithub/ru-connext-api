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
                    sh 'ls -a'
                    sh 'ls /home/ruconnext/ruconnext-dev -a'
                    sh 'docker-compose down'
                    sh 'cp /home/ruconnext/ruconnext-dev/config.yaml /home/ruconnext/jenkins_agent/workspace/${JOB_NAME}/environments'
                 }
                sh 'ls -la /home/ruconnext/jenkins_agent/workspace/${JOB_NAME}/environments'
                sh 'docker build -t ru-connext-api .'
                 dir('/home/ruconnext/ruconnext-dev') {
                    sh 'ls -a'
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
                dir('/home/ruconnext/docker/k6') {
                    sh 'ls -a'
                    sh 'docker-compose run --rm k6 run --summary-export=/scripts/results.json /scripts/testapi/test_spec.js'
                }
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
                 dir('/home/ruconnext/ruconnext-prod') {
                    sh 'ls -a'
                    sh 'docker-compose down'
                    sh 'cp /home/ruconnext/ruconnext-prod/config.yaml /home/ruconnext/jenkins_agent/workspace/${JOB_NAME}/environments'
                 }
                sh 'ls -la /home/ruconnext/jenkins_agent/workspace/${JOB_NAME}/environments'
                sh 'docker build -t ru-connext-api .'
                 dir('/home/ruconnext/ruconnext-prod') {
                    sh 'ls -a'
                    sh 'docker-compose up -d'
                    sh 'docker-compose up --scale ru-connext-api=4 -d'
                 }
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
