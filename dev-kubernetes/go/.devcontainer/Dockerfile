# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Python version: 3, 3.9, 3.8, 3.7, 3.6
ARG VARIANT="1"
FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_go:0-${VARIANT}


# 直接从dockerhub下载kubectl kubectx kubens
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectx /usr/local/bin/kubectx
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubens /usr/local/bin/kubens
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectl /usr/local/bin/kubectl

# ADD go proxy
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

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

# Install Go tools
ARG GO_TOOLS_WITH_MODULES="\
    golang.org/x/tools/gopls \
    honnef.co/go/tools/... \
    golang.org/x/tools/cmd/gorename \
    golang.org/x/tools/cmd/goimports \
    golang.org/x/tools/cmd/guru \
    golang.org/x/lint/golint \
    github.com/mdempsky/gocode \
    github.com/cweill/gotests/... \
    github.com/haya14busa/goplay/cmd/goplay \
    github.com/sqs/goreturns \
    github.com/josharian/impl \
    github.com/davidrjenni/reftools/cmd/fillstruct \
    github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    github.com/ramya-rao-a/go-outline \
    github.com/acroca/go-symbols \
    github.com/godoctor/godoctor \
    github.com/rogpeppe/godef \
    github.com/zmb3/gogetdoc \
    github.com/fatih/gomodifytags \
    github.com/mgechev/revive \
    github.com/go-delve/delve/cmd/dlv"

RUN mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    && export GOPATH=/tmp/gotools \
    # Go tools w/module support
    && export GO111MODULE=on \
    && (echo "${GO_TOOLS_WITH_MODULES}" | xargs -n 1 go get -x )2>&1 \
    # gocode-gomod
    && export GO111MODULE=auto \
    # && go get -x -d github.com/stamblerre/gocode 2>&1 \
    # && go build -o gocode-gomod github.com/stamblerre/gocode \
    # golangci-lint
    # && sh /tmp/library-scripts/golangci-lint.sh -b /usr/local/bin 2>&1 \
    # Move Go tools into path and clean up
    && mv /tmp/gotools/bin/* /usr/local/bin/ \
    # && mv gocode-gomod /usr/local/bin/ \
    && rm -rf /tmp/gotools


# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1