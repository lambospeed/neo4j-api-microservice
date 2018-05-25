DOCKER=docker
KUBECTL=kubectl
IMAGE_LOCAL_PREFIX=kobylyanskiy
DEPLOYMENT=dgraph-api
VERSION=$1
PROJECT_ID=spy-crowd

${DOCKER} build -t ${IMAGE_LOCAL_PREFIX}/${DEPLOYMENT}:${VERSION} .

echo $GCLOUD_SERVICE_KEY_PRD | base64 --decode -i > ${HOME}/gcloud-service-key.json
gcloud auth activate-service-account --key-file ${HOME}/gcloud-service-key.json

gcloud --quiet config set project $PROJECT_NAME_PRD
gcloud --quiet config set container/cluster $CLUSTER_NAME_PRD
gcloud --quiet config set compute/zone ${CLOUDSDK_COMPUTE_ZONE}
gcloud --quiet container clusters get-credentials $CLUSTER_NAME_PRD

${DOCKER} tag ${IMAGE_LOCAL_PREFIX}/${DEPLOYMENT}:${VERSION} gcr.io/${PROJECT_ID}/${DEPLOYMENT}:${VERSION}

docker login -u _json_key --password-stdin https://gcr.io < ${HOME}/gcloud-service-key.json

${DOCKER} push gcr.io/${PROJECT_ID}/${DEPLOYMENT}:${VERSION}

${KUBECTL} set image deployment/${DEPLOYMENT}-deployment ${DEPLOYMENT}=gcr.io/${PROJECT_ID}/${DEPLOYMENT}:${VERSION}
${KUBECTL} rollout status deployment/${DEPLOYMENT}-deployment
