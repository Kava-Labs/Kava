package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EvmTxRequest struct {
	Tx   *ethtypes.Transaction
	Data interface{}
}
type EvmTxResponse struct {
	Request EvmTxRequest
	TxHash  common.Hash
	Err     error
}

type ErrEvmFailedToSign struct{ Err error }

func (e ErrEvmFailedToSign) Error() string {
	return fmt.Sprintf("failed to sign tx: %s", e.Err)
}

type ErrEvmFailedToBroadcast struct{ Err error }

func (e ErrEvmFailedToBroadcast) Error() string {
	return fmt.Sprintf("failed to broadcast tx: %s", e.Err)
}

// EvmSigner manages signing and broadcasting requests to transfer Erc20 tokens
// Will work for calling all contracts that have func signature `transfer(address,uint256)`
type EvmSigner struct {
	signerAddress common.Address
	Auth          *bind.TransactOpts
	EvmClient     *ethclient.Client
}

func NewEvmSigner(
	evmClient *ethclient.Client,
	privKey *ecdsa.PrivateKey,
	chainId *big.Int,
) (*EvmSigner, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainId)
	if err != nil {
		return &EvmSigner{}, err
	}

	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return &EvmSigner{}, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return &EvmSigner{
		Auth:          auth,
		signerAddress: crypto.PubkeyToAddress(*publicKeyECDSA),
		EvmClient:     evmClient,
	}, nil
}

func (s *EvmSigner) Run(requests <-chan EvmTxRequest) <-chan EvmTxResponse {
	responses := make(chan EvmTxResponse)

	// receive tx requests, sign & broadcast them.
	// Responses are sent once the tx is added to the pending tx pool.
	// To see result, use TransactionReceipt after tx has been included in a block.
	go func() {
		for {
			// wait for incoming request
			req := <-requests

			signedTx, err := s.Auth.Signer(s.signerAddress, req.Tx)
			if err != nil {
				err = ErrEvmFailedToSign{Err: err}
			} else {
				err = s.EvmClient.SendTransaction(context.Background(), signedTx)
				if err != nil {
					err = ErrEvmFailedToBroadcast{Err: err}
				}
			}

			responses <- EvmTxResponse{
				Request: req,
				TxHash:  signedTx.Hash(),
				Err:     err,
			}
		}
	}()

	return responses
}

func (s *EvmSigner) Address() common.Address {
	return s.signerAddress
}
