# Local Kubernetes with kind

In this section I will demonstrate how you can get started with Kubernetes on your local machine. For this demo, we use Kubernetes in Docker (KinD)

## Spinning up the Kubernetes cluster

We will create a cluster with 1 master node and three worker nodes.

In order to setup multiple nodes, we need to create the cluster using a config file; `kind-multiple-nodes.yaml`.

This file contains a simple definition of a cluster:

```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
- role: worker
````
To spin up the cluster, run the following command:
```
$ kind create cluster --config kind-multiple-nodes.yaml
```
Verify that the cluster is up and running with

```
$ kubectl get nodes
NAME                 STATUS   ROLES    AGE   VERSION
kind-control-plane   Ready    master   71s   v1.19.1
kind-worker          Ready    <none>   42s   v1.19.1
kind-worker2         Ready    <none>   42s   v1.19.1
kind-worker3         Ready    <none>   42s   v1.19.1
```

