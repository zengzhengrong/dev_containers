# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.158.0/containers/alpine/.devcontainer/base.Dockerfile

# [Choice] Alpine version: 3.12, 3.11, 3.10
ARG VARIANT="3.12"
FROM zengzhengrong889/dev-alpine:${VARIANT}
ADD config /root/.kube/
COPY --from=civo/cli /usr/local/bin/civo /usr/local/bin/civo 
COPY --from=civo/cli /usr/local/bin/kubectl /usr/local/bin/kubectl 
COPY --from=zengzhengrong889/helm:3 /usr/bin/helm /usr/local/bin/helm 

# ** [Optional] Uncomment this section to install additional packages. **
# RUN apk update \
#     && apk add --no-cache <your-package-list-here>
