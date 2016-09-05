FROM busybox

ADD bin/azurectl-linux64-static /usr/local/bin/azurectl

VOLUME /tmp
WORKDIR /tmp
