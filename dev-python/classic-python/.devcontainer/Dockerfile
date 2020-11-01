# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.140.1/containers/python-3/.devcontainer/base.Dockerfile

# [Choice] Python version: 3, 3.8, 3.7, 3.6
ARG VARIANT="3"
FROM newbe36524/vscode_python:0-${VARIANT}


COPY ./requirements.txt /tmp/pip-tmp/
RUN pip3 --disable-pip-version-check --no-cache-dir install -i https://pypi.tuna.tsinghua.edu.cn/simple -r /tmp/pip-tmp/requirements.txt \
   && rm -rf /tmp/pip-tmp


RUN echo 192.30.255.112	gist.github.com >> /etc/hosts \
&& echo 192.30.255.112	github.com >> /etc/hosts \ 
&& echo 192.30.255.112	www.github.com >> /etc/hosts \
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

# [Option] Install Node.js
# ARG INSTALL_NODE="false"
# ARG NODE_VERSION="lts/*"
# RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1