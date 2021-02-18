# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Python version: 3, 3.9, 3.8, 3.7, 3.6
ARG VARIANT="1"
ARG COMPOSE_VERISON=1.27.4
FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_go:0-${VARIANT}


ENV TZ=Asia/Shanghai

# download kubectl kubectx kubens from dockerhub
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectx /usr/local/bin/kubectx
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubens /usr/local/bin/kubens
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectl /usr/local/bin/kubectl
# download kind from dockerhub
COPY --from=abdennour/kind-cli /usr/local/bin/kind /usr/local/bin/kind

# download helm from dockerhub


# Config sources list
ADD sources.list /etc/apt/sources.list
ADD config.yaml /etc/kind/config.yaml
# [Option] Enable non-root Docker access in container
ARG ENABLE_NONROOT_DOCKER="true"
ARG SOURCE_SOCKET=/var/run/docker-host.sock
ARG TARGET_SOCKET=/var/run/docker.sock

COPY library-scripts/*.sh /tmp/library-scripts/
RUN curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
RUN apt-get -o Acquire::Check-Valid-Until=false update --fix-missing\
    && /bin/bash /tmp/library-scripts/docker-debian.sh "${ENABLE_NONROOT_DOCKER}" "${SOURCE_SOCKET}" "${TARGET_SOCKET}" "root" "${COMPOSE_VERISON}" \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts/
# ADD go proxy
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct




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

# RUN apt-get update && apt-get install -y apt-transport-https
# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get -o Acquire::Check-Valid-Until=false update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends kubectl kubectx

# aliases
RUN echo alias kl='kubectl' >> ~/.bashrc && \
    echo alias kx='kubectx' >> ~/.bashrc && \
    echo alias kn='kubens' >> ~/.bashrc && \
    echo alias kd='kind' >> ~/.bashrc

ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1