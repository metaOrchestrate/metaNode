#!/bin/sh

#debug mode
set -x

# Start Node Exporter in the background
/usr/local/bin/node_exporter &

# Get public IP
#IP=$(curl ifconfig.io)

# only create the priv_validator_state.json if it doesn't exist and the command is start
if [[ $1 == "start" && ! -f ${CELESTIA_HOME}/data/priv_validator_state.json ]]
then
    mkdir -p ${CELESTIA_HOME}/data
    cat <<EOF > ${CELESTIA_HOME}/data/priv_validator_state.json
{
  "height": "0",
  "round": 0, 
  "step": 0
}
EOF
fi

# Start constructing the command
CMD="/bin/celestia-appd --home ${CELESTIA_HOME} $@"

# Execute the command
echo "Executing: $CMD"
eval $CMD
