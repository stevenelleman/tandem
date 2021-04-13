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
    'external-dns': False
}

# TODO: change name to match the source directory, local images should take precedence
registryRoot = ''
for s in services.keys():
    if services[s]:
        docker_build('{root}{service}-image'.format(root = registryRoot, service = s), context='services/{svc}'.format(svc = s))

# 2. Load Environment Variable yamls
environment = 'dev'
yamls = ['./deployment/env/{src}.yaml'.format(src = s) for s in [environment, 'shared']]
k8s_yaml(yamls)

# 3. Load k8s Resource yamls
yamls = ['./services/{svc}/k8s.yaml'.format(svc = s) for s in services.keys()]
k8s_yaml(yamls)

# 4. Create k8s Resources (w/ Dependencies)
k8s_resource('api-store')
k8s_resource('grpc')
k8s_resource('public-api', resource_deps=['api-store', 'grpc'])
k8s_resource('web-frontend', resource_deps=['public-api'])
k8s_resource('external-dns', resource_deps=['web-frontend'])

