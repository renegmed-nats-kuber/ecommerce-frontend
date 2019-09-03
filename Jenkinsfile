pipeline {
   agent any

   environment {
     // You must set the following environment variables
     // ORGANIZATION_NAME
     // YOUR_DOCKERHUB_USERNAME (it doesn't matter if you don't have one)

     SERVICE_NAME = "ecommerce-frontend"
     REPOSITORY_TAG="${YOUR_DOCKERHUB_USERNAME}/${SERVICE_NAME}:${BUILD_ID}"
     dockerImage = ""
   }

   stages {
      stage('Preparation') {
         steps {
            cleanWs()
            //git credentialsId: 'GitHub', url: "https://github.com/${ORGANIZATION_NAME}/${SERVICE_NAME}"
            checkout scm
         }
      }

      stage('Build image') {
         steps {
            sh 'go mod vendor' 
	         //sh 'docker image build -t ${REPOSITORY_TAG} .'
            script {
              dockerImage = docker.build REPOSITORY_TAG
            }
         }
      }

      stage('Push Image') {
         steps {
            script{
              // sh 'docker push ${REPOSITORY_TAG} .'
              docker.withRegistry('https://registry.hub.docker.com', 'DockerHub') {
                 dockerImage.push()
                 dockerImage.push("latest")   
              }
            }         
	      } 
      }

      stage('Remove Unused docker image') {
        steps {
          sh "docker rmi ${REPOSITORY_TAG}"
        }
      }

      stage('Modify deploy file') {
          steps {
            sh 'envsubst < ${WORKSPACE}/front-end-deploy.yaml'            
          }
      }

      stage('Deploy to Cluster') {
          steps {
            //sh 'envsubst < ${WORKSPACE}/front-end-deploy.yaml'
            sh 'kubectl apply -f .'
          }
      }

   } // stages

}
