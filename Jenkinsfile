pipeline {
    agent any
    options { timestamps() 
              ansiColor ('xterm')}
    stages {
        stage('Build') {
            steps {
                sh 'docker-compose build'
                echo '\033[34mBuild complete!'
            }
        }
        stage('Deploy') {
            steps {
                sh 'docker-compose up -d'
                echo '\033[34mDeployment complete!'
            }

        }
    }
}

