# -*- mode: Python -*-

# Once working, add in
# For more on Extensions, see: https://docs.tilt.dev/extensions.html
# load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')

# Environment Mode:
# mode = 0 # Dev
mode = 1 # Staging
# mode = 2 # Prod

env = ''
if mode == 0:
    env = 'dev'
if mode == 1:
    env = 'staging'
if mode == 2:
    env = 'prod'

# 1. Build Images
services = {
    'api-store': False,
    'grpc': True,
    'public-api': True,
    'web-frontend': True,
    'ingress': False,
    'ingress-nginx-controller': False,
    'cert-manager': False,
}

# Apply `external-dns` resources in Staging and Production
if mode == 1 or mode == 2:
    services['external-dns'] = False

# 2. Load Scoped Credential Info
if mode == 1:
    k8s_yaml('./env/sensitive/staging.yaml')

# TODO: change name to match the source directory, local images should take precedence
registryRoot = 'docker.io/selleman/web-microservice-shell_'
for s in services.keys():
    if services[s]:
        # TODO: run `make build-${svc}` to reupdate vendored packages, should have an explicit filewatch on libraries, and run `make build-all` on changes
        docker_build('{root}{service}'.format(root = registryRoot, service = s), context='services/{svc}'.format(svc = s))

# 3. Load k8s Charts
chartVars = './env/charts/values-%s.yaml' % (env)
charts = ['./charts/{svc}/'.format(svc = s) for s in services.keys()]
for svc in charts:
    k8s_yaml(helm(svc, values=[chartVars]))

# 4. Create k8s Resources (w/ Dependencies)
k8s_resource('api-store')
k8s_resource('grpc')
k8s_resource('public-api', resource_deps=['api-store', 'grpc'])
k8s_resource('web-frontend', resource_deps=['public-api'])
k8s_resource('ingress-nginx-controller')
k8s_resource(objects=['internet-facing-services:ingress'], new_name='ingress', resource_deps=['web-frontend', 'public-api', 'ingress-nginx-controller'])

if mode == 1 or mode == 2:
    k8s_resource('external-dns', resource_deps=['ingress'])
    k8s_resource('cert-manager', resource_deps=['external-dns', 'ingress-nginx-controller'])

