#!/usr/bin/bash
minikube addons enable olm
sleep 60
kubectl rollout status deployment packageserver -n olm
kubectl rollout status deployment olm-operator -n olm
kubectl rollout status deployment catalog-operator -n olm
kubectl create ns testing || true
kubectl apply -f https://operatorhub.io/install/postgresql.yaml
sleep 400
kubectl wait subscriptions my-postgresql --for=jsonpath='{.status.currentCSV}'=postgresoperator.v5.0.5 -n operators
kubectl get subscriptions my-postgresql -n operators -o yaml
kubectl get installplans -n operators -o yaml
sleep 60
kubectl wait csv postgresoperator.v5.0.5 --for=jsonpath='{.status.phase}'=Succeeded  -n operators
sleep 60
kubectl rollout status deployment pgo -n operators
kubectl apply -f https://redhat-developer.github.io/service-binding-operator/userguide/getting-started/_attachments/pgcluster-deployment.yaml -n testing
sleep 60
kubectl rollout status deployment hippo-pgbouncer -n testing
kubectl port-forward svc/hippo-pgbouncer 5432:5432 -n testing &
echo `kubectl get secret hippo-pguser-hippo --template='{{index .data "pgbouncer-uri" }}' -n testing | base64 -d |cut -d "@" -f 1`"@localhost:5432/hippo" > bindings/postgresql/pgbouncer-uri
psql $(cat bindings/postgresql/pgbouncer-uri) -f db/schema.sql
