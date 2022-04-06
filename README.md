# golang
for Golang project <br>
## Install Jenkins on GKE with Helm

git clone https://github.com/GoogleCloudPlatform/continuous-deployment-on-kubernetes.git
helm install cd-jenkins -f jenkins/values.yaml jenkinsci/jenkins --wait
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx\nhelm repo update

## Install Ingress controller 
helm install nginx-ingress ingress-nginx/ingress-nginx
kubectl get deployment nginx-ingress-ingress-nginx-controller\nkubectl get service nginx-ingress-ingress-nginx-controller