package params

import (
	"math/big"
	"testing"
)

func TestNetworkIDToChainConfigOrDefault(t *testing.T) {
	tests := []struct {
		name            string
		networkID       *big.Int
		wantChainConfig *ChainConfig
	}{
		{
			"taikoInternalL2ANetworkId",
			TaikoInternalL2ANetworkID,
			TaikoChainConfig,
		},
		{
			"taikoInternalL2BNetworkId",
			TaikoInternalL2BNetworkID,
			TaikoChainConfig,
		},
		{
			"snaefoll",
			SnaefellsjokullNetworkID,
			TaikoChainConfig,
		},
		{
			"askja",
			AskjaNetworkID,
			TaikoChainConfig,
		},
		{
			"grimsvotn",
			GrimsvotnNetworkID,
			TaikoChainConfig,
		},
		{
			"eldfellNetworkID",
			EldfellNetworkID,
			TaikoChainConfig,
		},
		{
			"jolnirNetworkID",
			JolnirNetworkID,
			TaikoChainConfig,
		},
		{
			"katlaNetworkID",
			KatlaNetworkID,
			TaikoChainConfig,
		},
		{
			"mainnet",
			MainnetChainConfig.ChainID,
			MainnetChainConfig,
		},
		{
			"sepolia",
			SepoliaChainConfig.ChainID,
			SepoliaChainConfig,
		},
		{
			"goerli",
			GoerliChainConfig.ChainID,
			GoerliChainConfig,
		},
		{
			"doesntExist",
			big.NewInt(89390218390),
			AllEthashProtocolChanges,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if config := NetworkIDToChainConfigOrDefault(tt.networkID); config != tt.wantChainConfig {
				t.Fatalf("expected %v, got %v", config, tt.wantChainConfig)
			}
		})
	}
}
