FROM cirrusci/flutter:latest


RUN sed -i "s@http://.*archive.ubuntu.com@http://mirrors.huaweicloud.com@g" /etc/apt/sources.list &&\
    sed -i "s@http://.*security.ubuntu.com@http://mirrors.huaweicloud.com@g" /etc/apt/sources.list 
    # buildkit
RUN apt-get update &&\
    apt-get autoremove -y &&\
    apt-get clean -y &&\
    rm -rf /var/lib/apt/lists/* 
    # buildkit

ENV PUB_HOSTED_URL=https://mirrors.tuna.tsinghua.edu.cn/dart-pub
ENV FLUTTER_STORAGE_BASE_URL=https://mirrors.tuna.tsinghua.edu.cn/flutter