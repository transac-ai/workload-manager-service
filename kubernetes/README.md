# Deploying WMS to Google Kubernetes Engine (GKE)

This guide provides instructions for deploying the WMS service to Google Kubernetes Engine (GKE).

## Steps

1. Setup [gcloud](https://cloud.google.com/sdk/docs/install) CLI.

```bash
gcloud init
```

2. Confirm project name.

```bash
gcloud config get-value project
```

3. Create repository in Google Cloud Artifact Registry.

```bash
gcloud artifacts repositories create wms-repo \
    --project=transac-ai \
    --repository-format=docker \
    --location=us-east1 \
    --description="Transac AI WMS Repository"
```

4. Build and push the image to the repository.

```bash
gcloud builds submit --tag us-east1-docker.pkg.dev/transac-ai/wms-repo/transac-ai-wms-gke:1.0.2 .
```

5. Create a GKE cluster.

```bash
gcloud container clusters create-auto transac-ai-gke --location us-east1
```

Output may look something like this:

```bash
NAME                LOCATION  MASTER_VERSION      MASTER_IP      MACHINE_TYPE  NODE_VERSION        NUM_NODES  STATUS
transac-ai-gke      us-east1  1.30.5-gke.1014001  34.73.***.***  e2-small      1.30.5-gke.1014001  3          RUNNING
```

6. Verify cluster creation.

```bash
gcloud container clusters list

or

kubectl get nodes
```

7. Set secrets for environment variables to be available throughout the cluster.

```bash
kubectl create secret generic transac-ai-wms-secrets \
--from-literal=transac-ai-wms-api-key='' \
--from-literal=transac-ai-rss-api-key='' \
--from-literal=transac-ai-igs-api-key='' \
--from-literal=rss-url='' \
--from-literal=igs-url=''
```

8. Deploy deployment to GKE.

```bash
kubectl apply -f kubernetes/deployment.yaml
```

9. Check the status of the deployment.

```bash
kubectl get deployments
```

Output:

```bash
NAME                 READY   UP-TO-DATE   AVAILABLE   AGE
transac-ai-wms-gke   2/2     2            2           4h11m
```

10. Create and deploy load balancer service to access the deployment.

```bash
kubectl apply -f kubernetes/service.yaml
```

11. Check the status of the service, and get the external IP.

```bash
kubectl get services
```

Output:

```bash
NAME                     TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)        AGE
transac-ai-wms-service   LoadBalancer   **.***.***.***   **.***.***.***   80:31664/TCP   3h41m
```

12. Test the service.

```
curl \
 --header "Content-Type: application/json" \
 --header "Authorization: Bearer <WMS API KEY>" \
 --data '{"clientId":"test_client","promptId":2,"recordsSourceId":"SUPABASE","promptTemplatesSourceId":"SUPABASE","fromTime":"2019-12-29T06:39:22Z","toTime":"2019-12-29T23:49:22Z"}' \
  0.0.0.0:80/wms.v1.WMSService/GenerateInsights
```

## Commands

### Patching Secrets

For example, for changing the address of PBS or ISS services.

```bash
kubectl patch secret transac-ai-wms-secrets \
--type='json' \
-p='[{"op": "replace", "path": "/data/igs-url", "value": "'$(echo -n "<value>" | base64)'"}]'
```

### Restarting Deployment

For initiating a rolling restart of the deployment.

```bash
kubectl rollout restart deployment transac-ai-wms-gke
```

### Checking Image Version Pod is Running

```bash
kubectl describe pod <pod id or name> | grep -i image
```
