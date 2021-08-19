FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod go.sum /build
RUN go mod download
COPY . /build
RUN CGO_ENABLED=0 go build -o /owrtb --ldflags '-extldflags "-static"' ./cmd/frontend

FROM scratch
COPY --from=builder /owrtb /bin/owrtb
ENTRYPOINT ["/bin/owrtb"]