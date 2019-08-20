#!/bin/bash
#****************************************************************#
# Create Date: 2019-02-02 22:16
#********************************* ******************************#

ROOTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

NAME="tekton-proxy"
GIT_COMMIT="$(git rev-parse --verify HEAD)"
GIT_BRANCH=`git branch | grep \* | cut -d ' ' -f2`
TAG="${GIT_BRANCH}_${GIT_COMMIT:0:8}-$(date +%Y''%m''%d''%H''%M''%S)"

docker build -t "${NAME}:${TAG}" -f ${ROOTDIR}/Dockerfile ${ROOTDIR}/../

#array=( registry.cn-beijing.aliyuncs.com  registry.cn-hangzhou.aliyuncs.com registry.cn-huhehaote.aliyuncs.com registry.cn-shanghai.aliyuncs.com registry.cn-shenzhen.aliyuncs.com  registry.cn-qingdao.aliyuncs.com registry.cn-zhangjiakou.aliyuncs.com registry.ap-southeast-2.aliyuncs.com )
array=( registry.cn-hangzhou.aliyuncs.com )
for registry in "${array[@]}"
do
    echo "push images to ${registry}/knative-sample/${NAME}:${TAG}"
    docker tag "${NAME}:${TAG}" "${registry}/knative-sample/${NAME}:${TAG}"
    docker push "${registry}/knative-sample/${NAME}:${TAG}"

    docker tag "${NAME}:${TAG}" "${registry}/knative-sample/${NAME}:latest"
    docker push "${registry}/knative-sample/${NAME}:latest"
done
