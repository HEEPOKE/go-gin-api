pipeline {
    agent any

    stages {
        stage('Pull Image') {
            steps {
                script {
                    def imageName = 'heepoke/go-app:latest'
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub_credentials') {
                        def app = docker.image(imageName)
                        app.pull()
                    }
                }
            }
        }
        stage('Remove Existing Container') {
            steps {
                script {
                    sh 'docker rm -f go-gin || true'
                }
            }
        }
        stage('Run Container') {
            steps {
                script {
                    def imageName = 'heepoke/go-app:latest'
                    def app = docker.image(imageName)
                    app.run('--name go-gin --network app_network -p 6476:6476 -d')
                }
            }
        }
        stage('Apply Migrations') {
            steps {
                script {
                    sh 'docker exec -i mysql-db mysql -u root -p go-api < ./path/to/migrations.sql'
                }
            }
        }
    }
}