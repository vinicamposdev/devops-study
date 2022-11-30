# kubernetes
Created: 2022-11-29 10:56

 - install docker
 - install kubectl
 - install kind

## Docker Basics

```sh
docker login 
docker build -t <DOCKER_HUB_USERNAME>/hello-go .
docker push <DOCKER_HUB_USERNAME>/hello-go
````

## Kind
 - create cluster
```sh
kind create cluster
```
 - connect to cluster
```sh
kubectl cluster-info --context kind-kind
```
 - verify nodes using k8s
```sh
kubectl get nodes 
```
 - verify all clusters created with kind
```sh
kind get clusters
```
 - delete cluster
```sh
kind delete clusters kind
```
 - to maintain resilience
	 - use multiple masters/control-plane nodes
 - list clusters
 - switch context

## Components
### Pod
 - smallest deployed unit
 - run a single container
 - pods that run multiple containers that need to work together
	 - encapsulate an application composed of multiple co-located
 - general:
	 - 1 container to 1 pod
```yml
apiVersion: v1
kind: Pod
metadata:
  name: "servername"
  labels:
    app: "servername"
spec:
  containers:
    - name: goserver
      image: "image"
```

 - create a pod
```sh
kubectl apply -f k8s/pod.yml 
```
 - get all pods
```sh
kubectl get pods
```
 - redirect pod port
```sh
kubectl port-forward pod/podname <PORT_FROM_REDIRECT>:<PORT_TO_REDIRECT>
```
- delete a pod
```sh
kubectl delete pod <PODNAME>
```
- logs
```sh
kubectl logs <PODNAME>
```

### Replica Set

```yml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: "goserver"
  labels:
    app: "goserver"
spec:
  selector:
    matchLabels:
      app: "goserver"
  replicas: 2
  template:
	  <CONFIGURATION_OF_POD>
```
 - create replica set
```sh
kubectl apply -f k8s/replicaset.yml
```
 - get replicasets
```sh
kubectl get replicaset
```
 - manage pods
	 - if a pod is deleted, it ensure replicas numbers creating other
 - delete replicaset
```sh
kubectl delete replicaset <NAME>
```

### Deployment
 - Hierarchy
	 - Deployment > Replicaset > Pod
 - Implementation:
```yml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: "goserver"
  labels:
    app: "goserver"
spec:
  selector:
    matchLabels:
      app: "goserver"
  replicas: 2
  template:
	  <CONFIGURATION_OF_POD>
```
 - create deployments
```sh
kubectl apply -f k8s/deployment.yml
```
 - get deployments
```sh
kubectl get deployments
```
 - descrive deployment
```sh
kubectl describe deployment <NAME>
```
 - Rollout
	 - Rowback image to previous version
```sh
kubectl rollout undo deployment <NAME>
```
 - Revision
	 - rowback image to given version
```sh
kubectl rollout undo deployment --to-revision <NAME>
```

### Service
 - act as a node balancer
 - manage what pod to process
 - not use a port forward in practice
 - use selector
	 - filter what pods will be used in association
 - types
	 - ClusterIP
	 - NodePort
	 - NodeBalancer
	 - Headless Server
```yml
apiVersion: v1
kind: Service
metadata:
  name: goserver-service
spec:
  selector:
    app: goserver
  type: ClusterIP
  ports:
    - name: goserver-service
      port: 80
      protocol: TCP
```

 - get services
```sh
kubectl get service
kubectl get svc
```
 - forward port
```sh
kubectl port-forward svc/goserver-service 8000:80
```
 - specs.ports.port is service port
 - specs.ports.targetPort is container port

### Kubernetes Apis
 - research more in reference guide
```sh
kubectl proxy --port:8080 
```

### NodePort
 - NodePort
	 - more rare to use
	 - range of ports that you can provide in the application
	 - export a port
### NodeBalancer
 - generate an external ip

### Configuration Objects
 - environments variables
 - sensitive data like passwords
 - configmap
	 - using envs in specifics files

## Probes
### Healthz
 - when container is rebooting
 - when container is deploying
### Liveness
 - restart container, if not health
```yml
spec:
	containers:
		- name: container-name
		  image: paht-to-image
		  livenessProbe:
			  httpGet:
				  path: /healthz
				  port: 8000
				periodSeconds: 5
				failureThreshold: 3
				timeoutSeconds: 1
				successThreshold: 1
				initialDelaySeconds: 10
```
 - failureThreshold is the limit of when has some failure, the service will restart
 - timeoutSeconds, is the period of time that is acceptable to have a health system -> strict 
 - successThreshold how many time he will test to make sure that the system is health
 - see if is working
```sh
kubectl apply -f k8s/service.yml && watch -n1 kubectl get pods
```
 - get history of pod
```sh
kubectl describe pod <PODNAME>
```
### Readness
 - remove trafic, if not ready
 - verify if the applcation is ready, if is started
```yml
				readinessProbe:
            httpGet:
              path: /healthz
              port: 8000
            periodSeconds: 3
            failureThreshold: 3
            initialDelaySeconds: 10
```
 - initialDelaySeconds is a time to wait until start to test
### Combine Liveness Readness
 - all time readness is start, the liveness will restart container
 - they have conflict in behavior
 - an option is to increase initialDelayTime to consiliate then
### Startup Probe
 - verify during a period of time (periodSeconds * failureThreshold)
 - after verify, liberate to liveness & readness
## Resources e HPA
 - what is the thresholds/limits of resources that an application have
 - after reaching limit, start to need more resources and scale more pods
### Metrics Server
 - metrics server will collect the resources used in the server
 - this data can be used in Prometheus/Grafana to monitoring
 - metrics server is scalable until 5000 pods
 - get metrics server in kind
```yml
get  https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```
 - add in args to work with kind
```yml
- --kubelet-insecure-tls
```
 - check if its working, if there is a metricservers in the list
```sh
kubectl get apiservices | metricserver
```
### Resources
 - request: reserve at least this number of resource to a pod
 - if it hasn't a machine with this resource, the pod will be in pending state
 - vCPU is a param that has 1000m (milicores)
 - they have share cores
	 - using specific amount of milicores e.g. 500m
	 - using percentage e.g. 0.5 
 - use benchmark and stress test to reach the best numbers
 - limits: the max of resource that a pod must use
 - the sum of all limits reach the max of resource of machine
```yml
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
```
 - get resources consume
```sh
kubectl top pod <PODNAME>
```
### HPA
 - provisioning new replicas, ensure a scalable application
 - in general purpose, the hpa is sufficient
```yml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: goserver-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: goserver
  minReplicas: 1
  maxReplicas: 5
  targetCPUUtilizationPercentage: 75
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 50

```
 - targetCPUUtilizationPercentage is the percentage when the hpa will start to scale
 - create hpa
```sh
kubectl apply -f k8s/hpa.yml 
```
 - get hpa 
```sh
kubectl get hpa
```
### Stress Test Fortio
 - [Fortio](https://github.com/fortio/fortio) is a command line stress test
 - will test the application using 800 requests in 120 seconds with 70 in same time
 - qps -> queries per second
 - t -> what time will run
 - c -> simultaneous threads/connections
```sh
kubectl run -it fortio --rm --image=fortio/fortio -- load -qps 800 -t 120s -c 70 "http://goserver-service/healthz"
```
## Statefulset and Persistent Volumes
 - [Doc Kubernetes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)
 - It's best use critic databases using specialist services to manage them, like AWS aurora, RDS. It's difficult to tuning database
 - best to store some data, files
 - stateless is something that don't store data in the disk
 - if the pod need storage, a volume can create and use the disk
	 - dinammic -> create pool of storage, kubernetes in premise
	 - DStorageClass -> AWS
			 - Claim -> StorageClass -> Enable space needed -> BlockStorage
	 - static -> using fixed volume
 - ReadWriteOnce -> can write and read if you are in same node
 - ReadWriteMany -> read and write in different nodes
```yml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv1
spec:
  capacity:
    storage: 50Gi
  volumeMode: Filesystem
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Recycle
  storageClassName: slow
  mountOptions:
    - hard
    - nfsvers=4.1
  nfs:
    path: /tmp
    server: 172.17.0.2
```
 - get storage
```sh
kubectl get storageclass
```
### Statefulset
 - applications that need volumes and persist states e.g Database
 - has the leader and followers (master slave topology)
 - scale in a sort way
 - downsize need to ensure that the master is not killed
 - and if a master is killed, an election starts
```yml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql
spec:
  selector:
    matchLabels:
      app: mysql
  serviceName: mysql-h
  replicas: 2
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: k8s.gcr.io/nginx-slim:0.8
        ports:
        - containerPort: 80
          name: web
        volumeMounts:
        - name: www
          mountPath: /usr/share/nginx/html
  volumeClaimTemplates:
  - metadata:
      name: www
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```
 - creation order matters
 - but can create in parallel to
	 - podManagementPolicy: Parallel
 - you can scale:
```sh
kubectl scale statefulset <NAME> --replicas=<NUMBER>
```

### Headless Service
 - creating replication read
 - just write in some pods and read in another
 - using the clusterIP: None, k8s knows what pod will deliver the message by serviceName of them
 - the statefulset serviceName must be the same of metadata.name of headless service
```yml
apiVersion: v1
kind: Service
metadata:
  name: mysql-h
spec:
  selector:
    app: mysql
  ports:
    - name: goserver-service
      port: 80
      targetPort: 8000
      protocol: TCP
  clusterIP: None
```

## Ingress
 - reverse proxy
 - load balance
 - api gateway
 - is a unique entry point and route to the correct service
 - most common adapter is nginx
 - manage dns

## CertManager
 - tls/ssl security to give https to servers
 - can use it with helm
 - use a namespace to the cert manager
### Installation
 - [Cert Manager Docs](https://cert-manager.io/docs/)
 - Install maniphest
### Export Cert 
## Namespace
 - logical separation
 - can alocate different resources
 - can give different type of access, security roles
 - namespace to developmet, production
```sh
kubectl get ns
```
 - create namespace
```sh
kubectl create ns <NAME>
```
 - in deployment, use the namespace to indicate what the scope