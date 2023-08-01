FROM golang:1.20-alpine as builder

WORKDIR /nebius-instance-upper

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /nebius-instance-upper-amd64 ./cmd/nebius-instance-upper

# ---

FROM gcr.io/distroless/static:nonroot-amd64

COPY --from=builder /nebius-instance-upper-amd64 /nebius-instance-upper

WORKDIR /
CMD ["/nebius-instance-upper"]
