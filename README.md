# devops-go-web-app

go build -o main . 
./main


docker build -t nikhilmatta5/go-web-app:v1
docker run -p 8080:8080 -it nikhil478/go-web-app:v1
docker push nikhil478/go-web-app:v1

