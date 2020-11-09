# local kubernetes with ingress

```
kind create cluster --config kind-ingress-multiple-nodes.yaml
```

Apply the NGINX ingress controller
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
```

Now we can reach our service directly on localhost. 

```
while true; do sleep 0.1; curl http://localhost/host; echo "";done
```

You should now see all three pods responding to the request.