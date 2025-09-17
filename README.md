# devops-go-web-app

What i covered : 
    Containerization -> Dockerfile
    Kubernetes -> deployment, ingest, service
    CI -> github actions
    CD -> argoCD
    Kubernetes -> DOKS
    Helm ->
        : dev
        : qa
        : prod
    IngressController -> Load Balancer -> DOKS
    Load Balanacer -> DNS


go build -o main . 
./main


docker build -t nikhilmatta5/go-web-app:v1
docker run -p 8080:8080 -it nikhil478/go-web-app:v1
docker push nikhil478/go-web-app:v1


https://docs.digitalocean.com/reference/doctl/how-to/install/