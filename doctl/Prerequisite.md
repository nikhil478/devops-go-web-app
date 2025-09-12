Install Kubectl

->     brew install kubernetes-cli

Install and configure doctl Command Line Interface (CLI)

-> brew install doctl
-> doctl auth init --context <NAME>
-> doctl auth list
-> doctl auth switch --context <NAME>
-> doctl account get

kubectl apply -f k8s/manifests/deployment.yaml
kubectl get pods


kubectl apply -f k8s/manifests/service.yaml
kubectl apply -f k8s/manifests/ingress.yaml

kubectl get ing
