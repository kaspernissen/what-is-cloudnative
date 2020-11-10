# Service Mesh with linkerd

In this section we will deploy a service mesh to our cluster. 

Install linkerd following the official instructions.

## Deploy linkerd service mesh to the cluster

```
$ linkerd install | kubectl apply -f -
``` 

A lot of resources will be installed in the namespace `linkerd`.

You can check it out with the following command:

```
$ kubectl get pods -n linkerd 
NAME                                      READY   STATUS    RESTARTS   AGE
linkerd-controller-856dd4c67-vnlbd        2/2     Running   0          2m48s
linkerd-destination-54448ff8f4-p2q48      2/2     Running   0          2m48s
linkerd-grafana-84c6d9fd47-b2bqb          2/2     Running   0          2m47s
linkerd-identity-5b8f7d5999-9929t         2/2     Running   0          2m48s
linkerd-prometheus-6f77b564c9-v9c72       2/2     Running   0          2m48s
linkerd-proxy-injector-8566b47778-brz9m   2/2     Running   0          2m48s
linkerd-sp-validator-84dffcbd69-7c65d     2/2     Running   0          2m47s
linkerd-tap-686cc9fcb-gmk9g               2/2     Running   0          2m47s
linkerd-web-78479646b-vg5vg               2/2     Running   0          2m48s
```

We can verify that linkerd is installed correctly by running:

```
$ linkerd check
...
Status check results are √
```

## Explore the linkerd dashboard

Linkerd comes with a cool frontned where you can get a lot of information about the mesh and pods in the mesh.

```
$ linkerd dashboard
```

## Let's add our service to the mesh

You can set up automatic proxy injection, however for the purpose of this demo, we are going to do it "manually" with the linkerd cli.

```
$ kubectl get deploy hello-world -o yaml | linkerd inject - | kubectl apply -f -
```
