# Web app monitoring with Prometheus

This repository is an example set up of a web application written in Go
that exposes its metrics for [Prometheus](https://prometheus.io/) monitoring toolkit.
The application and Prometheus run on [Kubernetes](http://kubernetes.io/) cluster.

## Hello app

There are three versions of the hello app which can be pulled from
[Docker Hub](https://hub.docker.com/r/marselester/prom-on-k8s/).

```bash
$ sudo docker run --rm -p 8000:8000 marselester/prom-on-k8s:v3
$ curl localhost:8000/hello
Hello, World!
$ curl localhost:8000/metrics
...
# HELP hello_request_duration_seconds Histogram of the /hello request duration.
# TYPE hello_request_duration_seconds histogram
hello_request_duration_seconds_bucket{le="0.01"} 0
hello_request_duration_seconds_bucket{le="0.025"} 0
hello_request_duration_seconds_bucket{le="0.05"} 0
hello_request_duration_seconds_bucket{le="0.1"} 1
hello_request_duration_seconds_bucket{le="0.25"} 1
hello_request_duration_seconds_bucket{le="0.5"} 1
hello_request_duration_seconds_bucket{le="1"} 1
hello_request_duration_seconds_bucket{le="2.5"} 1
hello_request_duration_seconds_bucket{le="5"} 1
hello_request_duration_seconds_bucket{le="10"} 1
hello_request_duration_seconds_bucket{le="+Inf"} 1
hello_request_duration_seconds_sum 0.083953974
hello_request_duration_seconds_count 1
# HELP hello_requests_total Total number of /hello requests.
# TYPE hello_requests_total counter
hello_requests_total{status="500"} 1
```

If you want to build the hello app from source code, you should
clone the repository into your `GOPATH`, otherwise vendoring won't work
(we use [Glide](https://glide.sh/)).

```bash
$ git clone git@github.com:marselester/prometheus-on-kubernetes.git \
	$GOPATH/src/github.com/marselester/prometheus-on-kubernetes
$ cd $GOPATH/src/github.com/marselester/prometheus-on-kubernetes/hello-app
$ glide install
$ make build
```

## Kubernetes manifests

Kubernetes manifests have excessive names, e.g., `name: prometheus-deployment`
instead of just `name: prometheus`. That is done for demonstration purpose.

- [prometheus][kube-prometheus]
- [hello app][kube-hello]

[kube-prometheus]: ./kube/prometheus
[kube-hello]: ./kube/hello
