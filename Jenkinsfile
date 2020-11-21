pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                    echo 'Running test'
                    sh 'go test -v ./...'
                }
            }
        }

        stage('Approval') {
            timeout(time:3, unit:'DAYS') {
                input 'Do I have your approval for deployment?'
        }

        stage('Deploy') {
            steps {
                echo '> Deploying the application to VM'
                // install galaxy roles
                sh "ansible-galaxy install -r /home/vu/Documents/ansible-project/go-todo-app/requirements.yml"   
                ansiblePlaybook(
                    vaultCredentialsId: 'AnsibleVault',
                    inventory: '/home/vu/Documents/ansible-project/go-todo-app/inventories/vagrant/inventory',
                    playbook: '/home/vu/Documents/ansible-project/go-todo-app/playbooks/service/main.yaml'
                )
            }
        }
    }
    post {
        always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"
            
        }
    }  
}