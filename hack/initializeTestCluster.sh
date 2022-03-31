#!/usr/bin/bash
kubectl create ns testing || true
kubectl apply -f hack/postgres-deployment.yaml -n testing
kubectl apply -f hack/postgres-svc.yaml -n testing
sleep 30
kubectl rollout status deployment postgres-deployment -n testing
kubectl port-forward svc/postgres-svc 5432:5432 -n testing &
sleep 10
psql $(cat bindings/postgresql/pgbouncer-uri) -f db/schema.sql
npm install -g @angular/cli
cd web && npm install && ng build && cd ..
