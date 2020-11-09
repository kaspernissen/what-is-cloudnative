# Ingress 

Make sure the old cluster is gone before you continue. 

For the next demo, we need to spin up a new cluster with configuration for the nginx-ingress-controller. 

Have a look at the file `kind-ingress-multiple-nodes.yaml` to see the difference from the previous cluster.

Spin up the cluster
```
$ kind create cluster --config kind-ingress-multiple-nodes.yaml
```

Next, let's apply the NGINX ingress controller yaml. (yup, directly from the internet - don't do this for your production setup)
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/kind/deploy.yaml
```

## Let's add another endpoint to our webserver

This endpoint should respond with the hostname that the service is running on. 

You can check the diference between the `main.go` in this demo and the one in `01-containers`.

Let's build the new version, `v2`.

```
$ docker build -t helloworld:v2 .
```

Again, we need to load the image into kind:

```
$ kind load docker-image helloworld:v2
```

## Deploy the new version

Let's apply our kubernetes manifests (notice the difference in the image tag in the `Deployment`). Further this directory stores a resource of type `Ingress`. We will use this to create a way from outside the cluster through the ingress-controller to our webserver.

Let's apply the manifests and see what happens.
```
$ kubectl apply -f manifests/
```

Now we should be able to reach our service directly on localhost and our newly added endpont `/host`.

```
$ curl http://localhost/host
hello-world-699cd5cdf5-b9lzk
```
The endpoint will respond with the name of the pod running the webserver.

You can verify that it's the same host responding all the time, by running the curl in a loop.

```
while true; do sleep 0.1; curl http://localhost/host; echo "";done
```

## Scaling our service

That's all good. But running one instance is no good. What if that instance fails? 

Before continuing, try deleting the pod, with the `while`-loop running.

```
kubectl delete pod hello-world-699cd5cdf5-b9lzk
```

As the service is no longer responding, you should see some `Bad Requests`

```
<html>
<head><title>504 Gateway Time-out</title></head>
<body>
<center><h1>504 Gateway Time-out</h1></center>
<hr><center>nginx</center>
</body>
</html>
```
This is the nginx-ingress-controller responding that the service you are trying to reach is not available at the moment.

At some point you should see succesfull responses again, not coming from a new host. The pod has been restarted by Kubernetes. That's great. But we can do better!

Let's try to scale our service horizontally and add more instances.

Edit the file `manifests/deployment.yaml`.

Insted of `replicas: 1` change to `replicas: 3`.

Save the change, and apply:

```
$ kubectl apply -f manifests/
```

Now try to make multiple requests to our service again:

```
while true; do sleep 0.1; curl http://localhost/host; echo "";done
```

You should now see all three pods responding to the request.

Try to kill one of the three pods, instead of a Bad Gateway as before, you should continue to see succesfull responses!

Resiliency for the win!