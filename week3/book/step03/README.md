kubectl apply -f \                  
https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.16.2/deploy/mandatory.yaml

kubectl apply -f \
https://raw.githubusercontent.com/kubernetes/ingress-nginx/nginx-0.16.2/deploy/provider/cloud-generic.yaml

kubectl -n ingress-nginx get service,pod

kubectl apply -f simple-service.yaml 
kubectl apply -f simple-ingress.yaml 

curl http://localhost -H 'Host: ch05.gihyo.local'