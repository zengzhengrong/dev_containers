FROM cirrusci/flutter:latest
ENV FLUTTER_STORAGE_BASE_URL=https://mirrors.tuna.tsinghua.edu.cn/flutter
ENV PUB_HOSTED_URL=https://mirrors.tuna.tsinghua.edu.cn/dart-pub

# buildkit
RUN  sed -i "s@http://.*archive.ubuntu.com@http://mirrors.huaweicloud.com@g" /etc/apt/sources.list && \
     sed -i "s@http://.*security.ubuntu.com@http://mirrors.huaweicloud.com@g" /etc/apt/sources.list 

# buildkit
RUN apt-get update &&     \
    apt-get autoremove -y &&    \
    apt-get clean -y &&     \
    rm -rf /var/lib/apt/lists/*

RUN sdkmanager --install  --no_https --proxy=http --proxy_host=192.168.1.152 --proxy_port=4780 "platforms;android-29"
RUN sdkmanager --install  --no_https --proxy=http --proxy_host=192.168.1.152 --proxy_port=4780 "build-tools;29.0.2"