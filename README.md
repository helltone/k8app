# k8app
Helm chart for easily manage gitops service

## Get started

Add this repository to Helm. (GitHub)

```
helm repo add k8app https://pin-up-global.github.io/k8app
```
Add this repository to Helm (GitLab)
Authenticate to the Helm repository
To authenticate to the Helm repository, you need either:
  A deploy token with the scope set to read_package_registry
  --username <token_name>
  --password <deploy_token>
```
helm repo add --username <token_name> --password <deploy_token> k8app https://gitlab.com/api/v4/projects/58057527/packages/helm/k8app
```

Install an example.

```
helm install app k8app/app
```

## 