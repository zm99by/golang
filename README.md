# golang
for Golang project <br>
## Install Jenkins on GKE with Helm<br><br>

git clone https://github.com/GoogleCloudPlatform/continuous-deployment-on-kubernetes.git<br>
helm install cd-jenkins -f jenkins/values.yaml jenkinsci/jenkins --wait<br>
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx\nhelm repo update<br>

## Install Ingress controller <br>
helm install nginx-ingress ingress-nginx/ingress-nginx<br>
kubectl get deployment nginx-ingress-ingress-nginx-controller\nkubectl get service nginx-ingress-ingress-nginx-controller<br>
