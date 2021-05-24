# Environment Variables 

## VSources 
1. `shared.yaml`: variables that all environments share, but may want to be changed on a per-project basis.
2. `dev.yaml`: variables for local development environment. 
3. `prod.yaml`: variables for external production environment. 

## Examples 

Sample `shared.yaml`:
```
apiVersion: v1
kind: Secret
metadata:
  name: shared-env-vars
type: Opaque
data:
  aws-access-key-id: < base64-encoded value >
  aws-secret-access-key: < base64-encoded value >
```

Sample `dev.yaml` / `prod.yaml`: 
```
apiVersion: v1
kind: Secret
metadata:
  name: env-vars
type: Opaque
data:
  aws-hosted-zone-id: < base64-encoded value >
  domain-filter: < base64-encoded value >
```