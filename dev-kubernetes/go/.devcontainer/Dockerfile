# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Python version: 3, 3.9, 3.8, 3.7, 3.6
ARG VARIANT="1"
FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_go:0-${VARIANT}


# 直接从dockerhub下载kubectl kubectx kubens
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectx /usr/local/bin/kubectx
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubens /usr/local/bin/kubens
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectl /usr/local/bin/kubectl

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


# RUN apt update --allow-insecure-repositories
# RUN curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add - 
# RUN apt-get update && apt-get install -y apt-transport-https
# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get -o Acquire::Check-Valid-Until=false update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends kubectl kubectx

# aliases
RUN echo alias kl='kubectl' >> ~/.bashrc && \
    echo alias kx='kubectx' >> ~/.bashrc && \
    echo alias kn='kubens' >> ~/.bashrc


# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1