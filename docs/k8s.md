# k8s

Deploying and orchestrating with k8s

## Provisioning

### DigitalOcean managed cluster

Pressed the button to make cluster :ghost: :tada:

### Kops

Usually use kops with aws but should be able to use with 
digitalocean.

## Resources

### Deploy a workload to cluster

In Kubernetes there’s various types of workloads you can deploy. Below you can find 4 different example manifests that can be deployed to your cluster. Copy the example manifest to a file on your workstation and use kubectl to apply it.

	kubectl create -f ./my-manifest.yaml
	
### Create a Deployment

Deployments describe a set of identical Pods without unique identities. A Deployment will run multiple replicas of your application and will automatically replace instances that fail or become unresponsive.

	apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: nginx-deployment-example
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: nginx-deployment-example
      template:
        metadata:
          labels:
            app: nginx-deployment-example
        spec:
          containers:
            - name: nginx
              image: library/nginx
              
### Create a Cronjob

A Cron Job creates Jobs on a time-based schedule. One CronJob object is like one line of a crontab (cron table) file. 
It runs a job periodically on a given schedule, written in Cron format.

	apiVersion: batch/v1beta1
    kind: CronJob
    metadata:
      name: cronjob-example
    spec:
      schedule: '*/5 * * * *'
      jobTemplate:
        spec:
          template:
            spec:
              containers:
                - name: cronjob-example
                  image: busybox
                  args:
                    - /bin/sh
                    - '-c'
                    - echo This is an example cronjob running every five minutes
              restartPolicy: OnFailure

### Create a Pod

A Pod is the basic building block of Kubernetes–the smallest and simplest unit in the Kubernetes object model that you 
create or deploy. A Pod represents a running process on your cluster.

	apiVersion: v1
    kind: Pod
    metadata:
      name: nginx-pod-example
    spec:
      containers:
        - name: nginx-pod-example
          image: library/nginx

### Create a ReplicaSet

A ReplicaSet ensures that a specified number of pod replicas are running at any given time. 
You can specify how many replicas of the pod that should be running by editing the 'replicas' key in the example below:

	apiVersion: apps/v1
    kind: ReplicaSet
    metadata:
      name: nginx-replicaset-example
    spec:
      replicas: 1
      selector:
        matchLabels:
          app: nginx-replicaset-example
      template:
        metadata:
          labels:
            app: nginx-replicaset-example
        spec:
          containers:
            - name: nginx-replicaset-example
              image: library/nginx

### Add Block Storage Volumes to your cluster

When you need to write and access persistent data in a Kubernetes cluster, you can create and access DigitalOcean block storage volumes by creating a PersistentVolumeClaim as part of your deployment.

The claim can allow cluster workers to read and write database records, user-generated website content, log files, and other data that should persist after a process has completed.

### Create a Configuration File

The example configuration defines two types of objects:

The PersistentVolumeClaim called csi-pvc which is responsible for locating the block storage volume by name if it already exists and creating the volume if it does not.
The Pod named my-csi-app, which will create containers, then add a mountpoint to the first object and mount the volume there.
Continue on to define the Persistent Volume Claim.

### Add a Load Balancer to your cluster

The DigitalOcean Cloud Controller supports provisioning DigitalOcean Load Balancers in a cluster’s resource configuration file.

The example configuration will define a load balancer and create it if one with the same name does not already exist.

### Create a Configuration File

You can add an external load balancer to a cluster by creating a new configuration file or adding the following lines to your existing service config file. Note that both the type and ports values are required for type: LoadBalancer:

	spec:
      type: LoadBalancer
      ports:
        - port: 80
          targetPort: 3000
          name: http

## Setup k8s dashboard

 	$ kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/master/src/deploy/recommended/kubernetes-dashboard.yaml

Create User 

    $ kubectl apply -f dashboard-adminuser.yml
    $ kubectl apply -f admin-crb.yml

## Setup Helm and Tiller

	$ helm init
	$ helm init --upgrade
	$ kubectl create serviceaccount \
		--namespace kube-system tiller
	$ kubectl create clusterrolebinding tiller-cluster-rule \
		--clusterrole = cluster-admin --serviceaccount= \
		kube-system:tiller
	$ kubectl patch deploy --namespace kube-system \
		tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'

## Monitoring - Prometheus

	Setup custom values
    
        $ cat > custom-values.yaml <<EOF
        # Depending on which DNS solution you have installed in your cluster enable the right exporter
        coreDns:
            enabled: false
    
        kubeDns:
            enabled: true
    
        alertmanager:
            alertmanagerSpec:
                storage:
                volumeClaimTemplate:
                    spec:
                    accessModes: ["ReadWriteOnce"]
                    resources:
                        requests:
                        storage: 10Gi
    
        prometheus:
            prometheusSpec:
                storage:
                volumeClaimTemplate:
                    spec:
                    accessModes: ["ReadWriteOnce"]
                    resources:
                        requests:
                        storage: 10Gi
    
        grafana:
            adminPassword: "YourPass123#"
            ingress:
                enabled: true
                annotations:
                kubernetes.io/ingress.class: nginx
                kubernetes.io/tls-acme: "true"
                hosts:
                - grafana.test.akomljen.com
                tls:
                - secretName: grafana-tls
                    hosts:
                    - grafana.test.akomljen.com
            persistence:
                enabled: true
                accessModes: ["ReadWriteOnce"]
                size: 10Gi
        EOF

Install with helm

	$ helm install --name prom --namespace monitoring -f custom-prom-values.yml stable/prometheus-operator
    $ # if this doesnt work try the next cmd
    $ # helm upgrade b/c there is a bug when deploying
    $ # Upgrade is a workaround
    $ helm upgrade --install --force prom --namespace monitoring -f custom-prom-values.yml stable/prometheus-operator


## Logging - EFL