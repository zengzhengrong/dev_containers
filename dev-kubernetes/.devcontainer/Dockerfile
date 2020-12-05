# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Python version: 3, 3.9, 3.8, 3.7, 3.6
ARG VARIANT="3"
FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_python:0-${VARIANT}

# Config sources list
ADD sources.list /etc/apt/sources.list

# Conifg kubernetes
ADD config /root/.kube/

# Install helm
RUN wget https://mirrors.huaweicloud.com/helm/v3.4.0/helm-v3.4.0-linux-amd64.tar.gz \ 
    && tar -zxvf helm-v3.4.0-linux-amd64.tar.gz \
    && mv linux-amd64/helm /usr/local/bin/helm \
    && rm -rf linux-amd64

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] If your pip requirements rarely change, uncomment this section to add them to the image.
# COPY requirements.txt /tmp/pip-tmp/
# RUN pip3 --disable-pip-version-check --no-cache-dir install -r /tmp/pip-tmp/requirements.txt \
#    && rm -rf /tmp/pip-tmp


RUN curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 
RUN apt-get update && apt-get install -y apt-transport-https
# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get -o Acquire::Check-Valid-Until=false update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends kubectl kubectx


# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1