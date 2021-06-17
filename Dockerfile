FROM ubuntu:focal

ADD ./moex /usr/bin/moex

RUN chmod +x /usr/bin/moex

RUN \
  apt-get update && \
  DEBIAN_FRONTEND='noninteractive' apt-get install ca-certificates -y && \
  apt-get clean

COPY ./html /var/html
COPY ./.vars.yml /etc/vars.yml

ENTRYPOINT [ "/usr/bin/moex" ]