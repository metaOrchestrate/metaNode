# stage 1 Generate celestia-appd Binary
FROM golang:1.18-alpine as builder
# hadolint ignore=DL3018
RUN apk update && apk --no-cache add make gcc musl-dev git
COPY ./app /celestia-app
WORKDIR /celestia-app
RUN make build

# stage 2
FROM alpine:3.17.1
# hadolint ignore=DL3018
RUN apk update && apk --no-cache add bash

COPY --from=builder /celestia-app/build/celestia-appd /bin/celestia-appd
COPY --from=builder /celestia-app/docker/entrypoint.sh /opt/entrypoint.sh

# Install Node Exporter for metrics
ENV NODE_EXPORTER_VERSION=1.7.0
RUN wget https://github.com/prometheus/node_exporter/releases/download/v${NODE_EXPORTER_VERSION}/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz \
    && tar xvfz node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz \
    && mv node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64/node_exporter /usr/local/bin \
    && rm -rf node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64

# p2p, rpc and prometheus port
EXPOSE 26656 26657 1317 9090 9100

ENV CELESTIA_HOME /opt

# Copy the start.sh script to the root directory
COPY start.sh /opt/start.sh
RUN chmod +x /opt/start.sh

# Execute the start.sh script when the container starts
CMD ["/opt/start.sh"]
