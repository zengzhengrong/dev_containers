# [Choice] Node.js version: 14, 12, 10
ARG VARIANT=14

FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_javascript_node:0-${VARIANT}
ADD sources.list /etc/apt/sources.list
RUN sudo -u node npm install --registry=https://registry.npm.taobao.org -g http-server @vue/cli @vue/cli-service-global
WORKDIR /app

EXPOSE 8080

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends git

# git proxy
# git config --global https.https://github.com.proxy https://localhost:4780
# git config --global http.https://github.com.proxy http://localhost:4780