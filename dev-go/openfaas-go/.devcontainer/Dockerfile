# Update the VARIANT arg in devcontainer.json to pick an Go version
ARG VARIANT=1
ARG COMPOSE_VERISON=1.27.4
# Get openfaas-cli
FROM openfaas/faas-cli:latest as build

FROM golang:${VARIANT}

# [Option] Enable non-root Docker access in container
ARG ENABLE_NONROOT_DOCKER="true"
ARG SOURCE_SOCKET=/var/run/docker-host.sock
ARG TARGET_SOCKET=/var/run/docker.sock

ENV GOPROXY=https://goproxy.cn
# Options for setup script
ARG INSTALL_ZSH="false"
ARG UPGRADE_PACKAGES="false"
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ADD sources.list /etc/apt/sources.list
# ADD faas-cli /usr/local/bin/
ADD config /root/.kube/


RUN echo 140.82.112.4	gist.github.com >> /etc/hosts \
&& echo 140.82.112.4	github.com >> /etc/hosts \ 
&& echo 140.82.112.4	www.github.com >> /etc/hosts \
&& echo 151.101.56.133	avatars0.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars1.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars2.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars3.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars4.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars5.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars6.githubusercontent.com >> /etc/hosts \
&& echo 51.101.56.133	avatars7.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	avatars8.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	camo.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	cloud.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	gist.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	marketplace-screenshots.githubusercontent.com >> /etc/hosts \
&& echo 151.101.228.133 raw.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	repository-images.githubusercontent.com >> /etc/hosts \
&& echo 151.101.56.133	user-images.githubusercontent.com >> /etc/hosts
# Install needed packages and setup non-root user. Use a separate RUN statement to add your own dependencies.
COPY library-scripts/*.sh /tmp/library-scripts/
COPY library-scripts/.oh-my-zsh /root/
COPY faas-scripts/* /tmp/faas-scripts/
# COPY kubernetes-client.tar.gz /root/
## copy faas-cli
COPY --from=build /home/app /usr/bin/

RUN chmod -R 777 /tmp/library-scripts/

# Move login shell to usr/local/bin
RUN mv /tmp/faas-scripts/faas-login.sh /usr/local/bin
RUN curl -s https://mirrors.huaweicloud.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
RUN apt-get update --fix-missing\
    && /bin/bash /tmp/library-scripts/common-debian.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" \
    && /bin/bash /tmp/library-scripts/docker-debian.sh "${ENABLE_NONROOT_DOCKER}" "${SOURCE_SOCKET}" "${TARGET_SOCKET}" "${USERNAME}" "${COMPOSE_VERISON}" \
    && /bin/bash /tmp/library-scripts/install-kubectl.sh \
    && /bin/bash /tmp/library-scripts/install-helm.sh \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts/

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

ENV GO111MODULE=auto


ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]

# [Optional] Uncomment the next line to use go get to install anything else you need
# RUN go get -x <your-dependency-or-tool>

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>