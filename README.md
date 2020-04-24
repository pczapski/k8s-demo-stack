# k8s Starter Pack

## Istio + Knative + helloworld apps inside

#### Requirements

- [taskfile](https://taskfile.dev/#/)
- [kind](https://github.com/kubernetes-sigs/kind)
- [direnv](https://direnv.net/)
- [helmfile](https://github.com/roboll/helmfile)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [skaffold](https://skaffold.dev/)

#### How To

1. Check `.envrc.dist` variables and if no changes required
2. Configure local domain (`/etc/hosts`)
3. Run `task init-env`

(Optional)

1. Run `task direnv:init` and adjust `.envrc` file
2. Configure local domain (`/etc/hosts`)
3. Run `task init-env`

#### Test

1. Run `task test` to check helloworlds apps

### TODO

- download deps scripts
- app deployment example ([skaffold](https://skaffold.dev/))
