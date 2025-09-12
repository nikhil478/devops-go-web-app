for listing available regions
doctl kubernetes options regions

for creating cluster
doctl kubernetes cluster create learn-k8s --region nyc1 --size s-1vcpu-2gb --count 1

for deleting that cluster
doctl kubernetes cluster delete learn-k8s

kubectl apply -f k8s/manifests/deployment.yaml
