FROM alpine:latest

WORKDIR /workspace

COPY kube-version /workspace

ENTRYPOINT [ "./kube-version", "-k8s-version" ]