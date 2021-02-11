FROM registry.cn-hangzhou.aliyuncs.com/newbe36524/vscode_base:0-focal

ADD sources.list /etc/apt/sources.list

ADD entrypoint.sh /

RUN apt-get update \
    && apt-get -y upgrade \
    && apt-get install -y openssh-server  && apt-get clean -y \
    && (echo "root" && echo "root") | passwd \
    && sed -ie 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/g'  /etc/ssh/sshd_config
   

ENTRYPOINT [ "/bin/bash","entrypoint.sh" ]