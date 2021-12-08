docker build -t worming/http-helloworld -t worming/http-helloworld:v1 -t worming/http-helloworld:v2 .
docker push worming/http-helloworld 
docker push worming/http-helloworld:v1
docker push worming/http-helloworld:v2