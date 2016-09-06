FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ADD bin/azurectl-linux64-static /usr/local/bin/azurectl

ENTRYPOINT ["/usr/local/bin/azurectl"]
