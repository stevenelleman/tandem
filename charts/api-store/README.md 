# API-Store 

## Error Resolution 
For a currently unknown reason sometimes the persistent volume or persistent volume claim is not created. These commands can be used to manually create them: 

```
kubectl -n default create -f - <<EOF
kind: PersistentVolume
apiVersion: v1
metadata:
  name: api-store-volume
  labels:
    type: local
    app: api-store
spec:
  storageClassName: manual
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteMany
  hostPath:
    path: "/data"
EOF
```

```
kubectl -n default create -f - <<EOF
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: api-store-claim
  labels:
    app: api-store
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
EOF
```