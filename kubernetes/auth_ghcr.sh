kubectl create secret generic regcred --from-file=.dockerconfigjson=/data/mike/.docker/config.json --type=kubernetes.io/dockerconfigjson
