FROM centos
  
MAINTAINER langwei "langwei@tencent.com"

WORKDIR /data/

ADD log /data/

ENTRYPOINT ["./log"]

ENV aa=ll
