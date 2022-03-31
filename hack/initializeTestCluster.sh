
minikube addons enable olm
sleep 30
kubectl apply -f https://operatorhub.io/install/postgresql.yaml
sleep 30
kubectl create ns testing || true
sleep 10
kubectl apply -f https://redhat-developer.github.io/service-binding-operator/userguide/getting-started/_attachments/pgcluster-deployment.yaml -n testing
sleep 30
kubectl port-forward svc/hippo-pgbouncer 5432:5432 -n testing 
echo `kubectl get secret hippo-pguser-hippo --template='{{index .data "pgbouncer-uri" }}' -n testing | base64 -d |cut -d "@" -f 1`"@localhost:5432/hippo" > bindings/postgresql/pgbouncer-uri
psql $(cat bindings/postgresql/pgbouncer-uri) -f db/schema.sql
 