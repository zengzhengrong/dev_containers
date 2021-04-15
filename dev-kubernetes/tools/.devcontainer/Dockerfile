# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.163.1/containers/alpine/.devcontainer/base.Dockerfile

# [Choice] Alpine version: 3.13, 3.12, 3.11, 3.10
ARG VARIANT="3.13"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-alpine-${VARIANT}
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectx /usr/local/bin/kubectx
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubens /usr/local/bin/kubens
COPY --from=binbash/k8s-kubectx /usr/local/bin/kubectl /usr/local/bin/kubectl
COPY --from=zengzhengrong889/helm:3 /usr/bin/helm /usr/local/bin/helm 
COPY krew.sh .
RUN /bin/bash ./krew.sh && \
    mv ~/.krew/bin/kubectl-krew /usr/local/bin/
# aliases
# RUN echo alias kl='kubectl' >> ~/.bashrc && \
#     echo alias kx='kubectx' >> ~/.bashrc && \
#     echo alias kn='kubens' >> ~/.bashrc && \
# ** [Optional] Uncomment this section to install additional packages. **
# RUN apk update \
#     && apk add --no-cache <your-package-list-here>