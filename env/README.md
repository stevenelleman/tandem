# Sensitive Environment Variables 

## Sources 
1. `./sensitive/dev.yaml`: variables for local development environment. 
2. `./sensitive/prod.yaml`: variables for external production environment. 

## Example
Sample `dev.yaml` / `prod.yaml`: 
```
apiVersion: v1
kind: Secret
metadata:
  name: env-vars
type: Opaque
data:
  aws-access-key-id: < belonging to configured AWS user > 
  aws-secret-access-key: < same > 
```