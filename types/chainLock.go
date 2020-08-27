package types

import (
	"errors"
	"fmt"
	"github.com/quantumexplorer/tendermint/crypto"
	tmproto "github.com/quantumexplorer/tendermint/proto/tendermint/types"
)

type ChainLock struct {
	CoreBlockHeight    uint32          `json:"core_block_height,string,omitempty"`   // height of Chain Lock.
	CoreBlockHash      []byte         `json:"core_block_hash,string,omitempty"`     // hash of Chain Lock.
	Signature          []byte         `json:"signature,string,omitempty"`           // signature.
}

// ToProto converts Header to protobuf
func (cl *ChainLock) ToProto() *tmproto.ChainLock {
	if cl == nil {
		return nil
	}

	return &tmproto.ChainLock{
		CoreBlockHeight:         cl.CoreBlockHeight,
		CoreBlockHash:           cl.CoreBlockHash,
		Signature:               cl.Signature,
	}
}

// FromProto sets a protobuf Header to the given pointer.
// It returns an error if the header is invalid.
func ChainLockFromProto(clp *tmproto.ChainLock) (ChainLock, error) {
	if clp == nil {
		return ChainLock{}, errors.New("nil Chain Lock")
	}

	cl := new(ChainLock)
	cl.CoreBlockHeight = clp.CoreBlockHeight
	cl.CoreBlockHash = clp.CoreBlockHash
	cl.Signature = clp.Signature

	return *cl, cl.ValidateBasic()
}

func (cl *ChainLock) PopulateFromChainLockParams(clp tmproto.ChainLockParams) error {

	cl.CoreBlockHeight = clp.ChainLockHeight
	cl.CoreBlockHash = clp.ChainLockHash
	cl.Signature = clp.Signature

	return cl.ValidateBasic()
}

// ValidateBasic performs stateless validation on a Chain Lock returning an error
// if any validation fails.
// It does not verify the signature
func (cl ChainLock) ValidateBasic() error {

	if cl.CoreBlockHeight == 0 {
		return errors.New("zero Chain Lock Height")
	}

	if len(cl.CoreBlockHash) == 0 {
		return fmt.Errorf("chain lock block hash is empty")
	}

	if err := ValidateHash(cl.CoreBlockHash); err != nil {
		return fmt.Errorf("chain lock block hash is wrong size")
	}

	if len(cl.Signature) == 0 {
		return fmt.Errorf("chain lock signature is empty")
	}

	if err := ValidateSignatureSize(crypto.BLS12381, cl.Signature); err != nil {
		return fmt.Errorf("chain lock signature is wrong size")
	}

	return nil
}

func NewMockChainLock() ChainLock {
	return ChainLock{
		CoreBlockHeight: 1,
		CoreBlockHash: []uint8{0x72, 0x3e, 0x3, 0x4e, 0x19, 0x53, 0x35, 0x75, 0xe7,
			0x0, 0x1d, 0xfe, 0x14, 0x34, 0x17, 0xcf, 0x61, 0x72, 0xf3, 0xf3, 0xd6,
			0x5, 0x8d, 0xdd, 0x60, 0xf7, 0x20, 0x99, 0x3, 0x28, 0xce, 0xde},
		Signature: []uint8{0xa3, 0x10, 0xf8, 0x2c, 0x3, 0xf7, 0x39, 0xa0, 0x1, 0x84,
			0x91, 0x77, 0x4d, 0xe1, 0x87, 0xd9, 0x9a, 0x12, 0x11, 0x4d, 0xee, 0xb5,
			0x5e, 0x4, 0x3c, 0x10, 0xd2, 0x91, 0x4e, 0xe5, 0x7e, 0xb8, 0xff, 0x10,
			0x82, 0xef, 0x89, 0x8e, 0x56, 0x5f, 0xbc, 0x24, 0xdc, 0x5f, 0xe2, 0x3,
			0xce, 0x75, 0x2d, 0xdc, 0x50, 0x1b, 0x1d, 0x34, 0xe1, 0xad, 0x7a, 0x82,
			0x1a, 0x8d, 0x7d, 0xe9, 0x6c, 0xcb, 0xe, 0xe3, 0x82, 0x80, 0xe4, 0xe3,
			0xb5, 0x48, 0x8d, 0x1f, 0x59, 0x10, 0xfc, 0x4d, 0xd8, 0x97, 0x58, 0x38,
			0x97, 0xd0, 0x91, 0xc1, 0x50, 0x4d, 0x69, 0x4a, 0xdd, 0x77, 0xd8, 0x88,
			0x2b, 0xdf},
	}
}
