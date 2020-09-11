# -*- mode: Python -*-

# Once working, add in
# For more on Extensions, see: https://docs.tilt.dev/extensions.html
# load('ext://restart_process', 'docker_build_with_restart')

# 1. Build Images
services = {
    "api-store": False,
    "sg": True,
    "public-api": True,
    "web-frontend": True
}

for s in services.keys():
    if services[s]:
        docker_build('sg/{svc}-image'.format(svc = s), context='services/{svc}'.format(svc = s))

# 2. Load yaml Configs
yamls = ["./services/{svc}/tilt.yaml".format(svc = s) for s in services.keys()]
k8s_yaml(yamls)

# 3. Create k8s Resources (w/ Dependencies)
k8s_resource('sg-api-store')
k8s_resource('sg-sg')
k8s_resource('sg-public-api', resource_deps=['sg-api-store', 'sg-sg'])
k8s_resource('sg-web-frontend', resource_deps=['sg-public-api'])


