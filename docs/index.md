# metaNode

metaNode is an orchestration stack for deploying nodes for the `celestia` modular blockchain ecosystem.

There are several types of nodes that can be deployed to participate in the network.

- `Full Consensus Node`: This type of node plays a crucial role in the consensus process by participating in block production and voting. It helps maintain the integrity and security of the Celestia blockchain.
- `Validator Node`: Validator nodes are a specific type of full consensus node that can be set up to participate in the consensus process by producing and voting on blocks. They contribute to the decentralization and security of the network.
- `Bridge Node`: Bridge nodes serve as intermediaries between the Data Availability network and the Consensus network. They facilitate the transfer of blocks between these two networks, ensuring data availability and consistency.
- `Full Storage Node`: Full storage nodes store all the data of the Celestia blockchain but do not actively participate in the consensus process. They provide redundancy and help ensure the availability of historical blockchain data.
- `Light Node`: Light nodes are lightweight clients that conduct data availability sampling on the Data Availability network. They help verify the availability of data without the need to store the entire blockchain history.

### Getting started

> The stack will deploy multiple services:

1. **Prometheus**   -> `metrics`
2. **Grafana**      -> `visibility`
3. **Celestia**     -> `blockchain`
4. **Grafana Loki** -> `logs`
5. **Promtail**     -> `logs agent`

**Grafana** default password is `admin:admin`

### Deploy a development stack

> Clone the repository
```bash
git clone git@github.com:metaOrchestrate/metaNode.git
cd metaNode
```

> Add environment Variables example:

```bash
# Modify ".env-dev" file acordingly 

NODE_VERSION=v0.12.1
NETWORK=mocha-4
NODE_TYPE=light
RPC_URL=full.consensus.mocha-4.celestia-mocha.com
```

```bash
# 1. Start a dev stack
make dev

# 2. Stop the stack
make dev-down
```

### Environment variables
Change the `.env-dev` file with your own values