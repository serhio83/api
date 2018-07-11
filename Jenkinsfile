pipeline {
  agent any
  stages {
    stage('stage12') {
      parallel {
        stage('stage1') {
          steps {
            sh 'pwd'
          }
        }
        stage('parallel stage2') {
          steps {
            echo 'It works in parallel'
          }
        }
      }
    }
    stage('stage #2') {
      steps {
        sh 'ls -la'
      }
    }
  }
}
