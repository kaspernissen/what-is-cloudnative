# Monitoring

In this section we will explore monitoring a bit. To get up and running quickly we will use the official helm chart for the prometheus-operator. 

This requires helm. Helm is a package manager for kubernetes.

```
$ helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
$ helm repo add stable https://charts.helm.sh/stable
$ helm repo update
```

## Install the Prometheus monitoring stack

```
$ helm install prometheus prometheus-community/kube-prometheus-stack
```

## Explore Grafana and the built-in dashboards

To visit grafana, you can do a `kubectl port-forward` similar to what we've done previously.

Locate the `prometheus-grafana-*` pod with `kubectl get pods`.

Next do a port-forward to port 3000.

```
$ kubectl port-forward prometheus-grafana-8db9dd95-wjxv2 3000
```

You can now access grafana in your browser: http://localhost:3000

Username: admin
Password: prom-operator