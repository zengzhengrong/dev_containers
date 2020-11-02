# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1"
FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_go:0-${VARIANT}

ENV GOPROXY=https://goproxy.cn
ADD sources.list /etc/apt/sources.list
# [Option] Install Node.js
# ARG INSTALL_NODE="true"
# ARG NODE_VERSION="lts/*"
# RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends git

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1