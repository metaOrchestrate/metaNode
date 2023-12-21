#!/bin/sh

# Start Node Exporter in the background
/usr/local/bin/node_exporter &

# Start the service
celestia light init --p2p.network $NETWORK
celestia $NODE_TYPE start --core.ip $RPC_URL --p2p.network $NETWORK