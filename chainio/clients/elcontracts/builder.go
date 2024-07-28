package elcontracts

import (
	"github.com/mpjunior92/eigensdk-go/chainio/clients/eth"
	"github.com/mpjunior92/eigensdk-go/chainio/txmgr"
	"github.com/mpjunior92/eigensdk-go/logging"
	"github.com/mpjunior92/eigensdk-go/metrics"
)

func BuildClients(
	config Config,
	client eth.HttpBackend,
	txMgr txmgr.TxManager,
	logger logging.Logger,
	eigenMetrics *metrics.EigenMetrics,
) (*ChainReader, *ChainWriter, *ContractBindings, error) {
	elContractBindings, err := NewBindingsFromConfig(
		config,
		client,
		logger,
	)
	if err != nil {
		return nil, nil, nil, err
	}

	elChainReader := NewChainReader(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.AvsDirectory,
		elContractBindings.RewardsCoordinator,
		logger,
		client,
	)

	elChainWriter := NewChainWriter(
		elContractBindings.Slasher,
		elContractBindings.DelegationManager,
		elContractBindings.StrategyManager,
		elContractBindings.RewardsCoordinator,
		elContractBindings.StrategyManagerAddr,
		elChainReader,
		client,
		logger,
		eigenMetrics,
		txMgr,
	)

	return elChainReader, elChainWriter, elContractBindings, nil
}
