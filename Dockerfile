FROM busybox:glibc
ADD crd-controller /
ENTRYPOINT ["/crd-controller"]
