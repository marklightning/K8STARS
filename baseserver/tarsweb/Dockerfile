FROM centos/nodejs-8-centos7:8

COPY web /web 
ENV PATH /opt/rh/rh-nodejs8/root/usr/bin:/opt/app-root/src/node_modules/.bin/:/opt/app-root/src/.npm-global/bin/:/opt/app-root/src/bin:/opt/app-root/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
USER root

RUN npm config set registry http://mirrors.cloud.tencent.com/npm/
RUN cd /web/demo && npm install
RUN cd /web && npm install

CMD cd /web/demo && sh -c 'npm run start &' && cd /web && npm run start
