# -*- mode: Python -*-

# Once working, add in
# For more on Extensions, see: https://docs.tilt.dev/extensions.html
# load('ext://restart_process', 'docker_build_with_restart')

# docker_compose('docker-compose.yml')

docker_build('sg/public-api-image', context='services/public-api')
docker_build('sg/web-frontend-image', context='services/web-frontend')

k8s_yaml("services/api-store/tilt.yaml")
k8s_yaml("services/public-api/tilt.yaml")
k8s_yaml("services/web-frontend/tilt.yaml")

k8s_resource('sg-api-store')
k8s_resource('sg-public-api', port_forwards='9001', resource_deps=['sg-api-store'])
k8s_resource('sg-web-frontend', port_forwards='9002', resource_deps=['sg-public-api'])

# 1. Docker Images

# 2. Kubernetes Service YAMLs
# services = ['api-store','public-api','web-frontend',]
# kubes = ["./services/{service}/tilt.yaml".format(service = s) for s in services]
# k8s_yaml(kubes)

# 3. Tilt Resources
# Group 1: Public API


