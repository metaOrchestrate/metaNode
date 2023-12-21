FROM ghcr.io/celestiaorg/celestia-node:v0.12.1 as node

USER root

# Install curl and other necessary packages like wget and tar for downloading and unpacking
RUN apk update && apk add --no-cache curl wget tar

# Install Node Exporter for metrics
ENV NODE_EXPORTER_VERSION=1.7.0
RUN wget https://github.com/prometheus/node_exporter/releases/download/v${NODE_EXPORTER_VERSION}/node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz \
    && tar xvfz node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz \
    && mv node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64/node_exporter /usr/local/bin \
    && rm -rf node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64.tar.gz node_exporter-${NODE_EXPORTER_VERSION}.linux-amd64

# Node Exporter
EXPOSE 2121 9100 26658

# Copy the init script into the image
COPY start.sh /start.sh
RUN chmod +x /start.sh
