package types

import "github.com/mpjunior92/eigensdk-go/crypto/bls"

type TestOperator struct {
	OperatorId     OperatorId
	StakePerQuorum map[QuorumNum]StakeAmount
	BlsKeypair     *bls.KeyPair
}
