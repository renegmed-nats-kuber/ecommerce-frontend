pipeline {
   agent any

   environment {
     // You must set the following environment variables
     // ORGANIZATION_NAME
     // YOUR_DOCKERHUB_USERNAME (it doesn't matter if you don't have one)

     SERVICE_NAME = "ecommerce-frontend"
     REPOSITORY_TAG="${YOUR_DOCKERHUB_USERNAME}/${SERVICE_NAME}:${BUILD_ID}"
     registryCredential = ‘DockerHub’
   }

   stages {
      stage('Preparation') {
         steps {
            cleanWs()
            git credentialsId: 'GitHub', url: "https://github.com/${ORGANIZATION_NAME}/${SERVICE_NAME}"
         }
      }
      stage('Build') {
         steps {
           sh 'go mod vendor' 
	   sh 'docker image build -t ${REPOSITORY_TAG} .'
         }
      }

      stage('Push Image') {
         steps {
           // sh 'docker push ${REPOSITORY_TAG} .'
           docker.withRegistry('https://registry.hub.docker.com/v2', 'DockerHub') {
             app.push("${env.BUILD_ID}")           

	 }
      }

      stage('Deploy to Cluster') {
          steps {
            sh 'envsubst < ${WORKSPACE}/front-end-deploy.yaml | kubectl apply -f -'
          }
      }
   }
}
