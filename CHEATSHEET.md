## Tips + Useful Commands 
- `TF_LOG=debug terraform plan` is very handy.
- `kubectl get po --all-namespaces --field-selector 'status.phase!=Running' -o json | kubectl delete -f -`: to delete all non-running pods. 
    - if connection is refused and there are tons of evicted pods following `kubectl get pods --all-namespaces` this cmd may help
- `docker system prune -a --volumes`: to purge unused docker resources 
- `kubectl logs ingress-nginx-controller-57cb5bf694-kc6jr -n ingress-nginx --follow`
- `kubectl logs -l app=public-api --tail=1500`