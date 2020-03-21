package cmd

import (
	"fmt"

	"github.com/dapperlabs/flow-go/cmd/bootstrap/run"
	"github.com/dapperlabs/flow-go/crypto"
	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/rs/zerolog/log"
)

var configFile string

type NodeConfig struct {
	Role    flow.Role
	Address string
	Stake   uint64
}

type NodeInfoPriv struct {
	Role           flow.Role
	Address        string
	NodeID         flow.Identifier
	NetworkPrivKey EncodableNetworkPrivKey
	StakingPrivKey EncodableStakingPrivKey
}

type NodeInfoPub struct {
	Role          flow.Role
	Address       string
	NodeID        flow.Identifier
	NetworkPubKey EncodableNetworkPubKey
	StakingPubKey EncodableStakingPubKey
	Stake         uint64
}

func genNetworkAndStakingKeys() ([]NodeInfoPub, []NodeInfoPriv) {
	var nodeConfigs []NodeConfig
	readJSON(configFile, &nodeConfigs)
	nodes := len(nodeConfigs)
	log.Info().Msgf("read %v node configurations", nodes)

	log.Debug().Msgf("will generate %v networking keys", nodes)
	networkKeys, err := run.GenerateNetworkingKeys(nodes, generateRandomSeeds(nodes))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot generate networking keys")
	}
	log.Info().Msgf("generated %v networking keys", nodes)

	log.Debug().Msgf("will generate %v staking keys", nodes)
	stakingKeys, err := run.GenerateStakingKeys(nodes, generateRandomSeeds(nodes))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot generate staking keys")
	}
	log.Info().Msgf("generated %v staking keys", nodes)

	nodeInfosPub := make([]NodeInfoPub, 0, nodes)
	nodeInfosPriv := make([]NodeInfoPriv, 0, nodes)
	for i, nodeConfig := range nodeConfigs {
		log.Debug().Int("i", i).Str("address", nodeConfig.Address).Msg("assembling node information")
		nodeInfoPriv, nodeInfoPub := assembleNodeInfo(nodeConfig, networkKeys[i], stakingKeys[i])
		nodeInfosPub = append(nodeInfosPub, nodeInfoPub)
		nodeInfosPriv = append(nodeInfosPriv, nodeInfoPriv)
		writeJSON(fmt.Sprintf("%v.node-info.priv.json", nodeInfoPriv.NodeID), nodeInfoPriv)
	}

	writeJSON("node-infos.pub.json", nodeInfosPub)

	return nodeInfosPub, nodeInfosPriv
}

func assembleNodeInfo(nodeConfig NodeConfig, networkKey, stakingKey crypto.PrivateKey) (NodeInfoPriv, NodeInfoPub) {
	nodeID, err := flow.PublicKeyToID(stakingKey.PublicKey())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot generate NodeID from PublicKey")
	}

	log.Debug().
		Str("networkPubKey", pubKeyToString(networkKey.PublicKey())).
		Str("stakingPubKey", pubKeyToString(stakingKey.PublicKey())).
		Msg("encoded public staking and network keys")

	nodeInfoPriv := NodeInfoPriv{
		Role:           nodeConfig.Role,
		Address:        nodeConfig.Address,
		NodeID:         nodeID,
		NetworkPrivKey: EncodableNetworkPrivKey{networkKey},
		StakingPrivKey: EncodableStakingPrivKey{stakingKey},
	}

	nodeInfoPub := NodeInfoPub{
		Role:          nodeConfig.Role,
		Address:       nodeConfig.Address,
		NodeID:        nodeID,
		NetworkPubKey: EncodableNetworkPubKey{networkKey.PublicKey()},
		StakingPubKey: EncodableStakingPubKey{stakingKey.PublicKey()},
		Stake:         nodeConfig.Stake,
	}

	return nodeInfoPriv, nodeInfoPub
}
