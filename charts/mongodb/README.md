# Ensure Secrets Exist Prior to mdbc CRD Creation

When using External Secrets with ArgoCD, apply the sync wave annotation with a value of at least 3 to the mdbc resource:
```yaml
  annotations:
    argocd.argoproj.io/sync-wave: "3"
```

It's necessary for the External Secrets addon to sync secrets before creating the mdbc resource.

For deployments without using External Secrets addon, make sure that all secrets (such as user passwords and monitoring passwords) are created BEFORE running the helm upgrade --install command.
