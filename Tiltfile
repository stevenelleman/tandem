# -*- mode: Python -*-

# Once working, add in
# For more on Extensions, see: https://docs.tilt.dev/extensions.html
# load('ext://restart_process', 'docker_build_with_restart')

# 1. Build Images
services = {
    'api-store': False,
    'grpc': True,
    'public-api': True,
    'web-frontend': True,
    'ingress': False,
}

# TODO: change name to match the source directory, local images should take precedence
registryRoot = 'docker.io/selleman/web-microservice-shell_'
for s in services.keys():
    if services[s]:
        docker_build('{root}{service}'.format(root = registryRoot, service = s), context='services/{svc}'.format(svc = s))

# 2. Load Environment Variable yamls
k8s_yaml('./env/sensitive/dev.yaml')

# 3. Load k8s Charts
charts = ['./charts/{svc}/'.format(svc = s) for s in services.keys()]
for svc in charts:
    k8s_yaml(helm(svc, values=['./env/charts/values-dev.yaml']))

# 4. Create k8s Resources (w/ Dependencies)
k8s_resource('api-store')
k8s_resource('grpc')
k8s_resource('public-api', resource_deps=['api-store', 'grpc'])
k8s_resource('web-frontend', resource_deps=['public-api'])

# Name notation: https://docs.tilt.dev/tiltfile_concepts.html#kubernetes-object-selectors
k8s_resource(objects=['public-api-rule:ingress', 'web-frontend-rule:ingress'], new_name='ingress', resource_deps=['web-frontend', 'public-api'])

if mode == 1 or mode == 2:
    k8s_resource('external-dns', resource_deps=['ingress'])

