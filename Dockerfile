FROM golang:1.20-alpine as builder

WORKDIR /nebius-instance-upper

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /nebius-instance-upper-amd64 ./cmd/nebius-instance-upper

# ---

FROM alpine:3.18

ARG KUBECTL_ARCH="amd64"
ENV KUBECTL_ARCH=$KUBECTL_ARCH

ARG KUBECTL_VERSION="v1.27.4"
ENV KUBCTL_VERSION=$KUBCTL_VERSION

RUN apk --no-cache add \
      bash \
      curl \
      jq \
    && wget -O /usr/bin/kubectl "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" \
    && chmod +x /usr/bin/kubectl

COPY --from=builder /nebius-instance-upper-amd64 /nebius-instance-upper

WORKDIR /
CMD ["/nebius-instance-upper"]
