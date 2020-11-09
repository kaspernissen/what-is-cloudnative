# Initial deploy of our webserver

In the previous section we spun up a Kubernetes cluster, next let's run our dockerized webser on it.

Normally we will push our docker image to a docker registry. For simplicity we will load the image directly into kind instead. We can do that with the following command:

```
$ kind load docker-image helloworld:v1
```

## Deploy the webserver

Next we can deploy our webserver using the manifest files stored in `manifests`. The directory contains a `Deployment` and a `Service`.

```
$ kubectl apply -f manifests/
```

Verify that the application is running in the cluster as a single pod:

```
kubectl get pods
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-5b459f67c8-5g22d   1/1     Running   0          5s
```

We verify the service is running by forwarding the containers port to our `localhost` using `kubectl port-forward`.

```
$ kubectl port-forward svc/hello-world 8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
```

In another terminal or in a browser, visit `localhost:8080/hello?name=world``

You should see the response; `Hello, world` once again.

Great job so far!

Let's close down this cluster, and continue to the next demo.

```
$ kind delete clusters kind
```

