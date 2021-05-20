FROM golang:1.16-alpine as dev
RUN apk add --no-cache git ca-certificates make
RUN adduser -D appuser
COPY . /src/
WORKDIR /src

ENV GO111MODULE=auto
RUN --mount=type=cache,sharing=locked,id=gomod,target=/go/pkg/mod/cache \
    --mount=type=cache,sharing=locked,id=goroot,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -o rack-detect .

FROM busybox
# Add Certificates into the image, for anything that does API calls
COPY --from=dev /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# Add rack-detect binary
COPY --from=dev /src/rack-detect /
ENTRYPOINT ["/rack-detect"]
