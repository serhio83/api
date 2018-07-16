pipeline {
  agent any
  
  stages {
    stage('stage1') {
      steps {
        sh 'echo 1 > /1.txt'
      }
    }
    stage('stage #2') {
      steps {
        sh 'ls -la'
      }
    }
    stage('Back-end') {
      agent {
          docker { image 'jenkinsci/slave' }
      }
      steps {
          sh 'uname -a && ls -la; echo done > /done.txt'
      }
    }
  }
}
