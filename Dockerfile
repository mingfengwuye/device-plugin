FROM ubuntu:16.04 as build

RUN apt-get update && apt-get install -y --no-install-recommends \
        g++ \
        ca-certificates \
        wget \
    rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.9.4
RUN wget -nv -O - https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz \
    | tar -C /usr/local -xz
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR /go/src/cambricon-device-plugin
COPY . .

RUN export CGO_LDFLAGS_ALLOW='-Wl,--unresolved-symbols=ignore-in-object-files' && \
    go install -ldflags="-s -w" -v cambricon-device-plugin


FROM debian:stretch-slim

COPY --from=build /go/bin/cambricon-device-plugin /usr/bin/cambricon-device-plugin

CMD ["cambricon-device-plugin"]
