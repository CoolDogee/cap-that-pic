gcloud container clusters get-credentials $GKE_CLUSTER --zone=$GKE_CLUSTER_ZONE
kubectl create -f ./
watch kubectl get all