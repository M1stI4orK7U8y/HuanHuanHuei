package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/client/rpc"
	"gitlab.com/packtumi9722/huanhuanhuei/src/huanhuancore/config"
)

// SendToAddress send eth to address
func SendToAddress(address string, amount string) (string, error) {

	client := rpc.GetEthInstance()
	privateKey, err := crypto.HexToECDSA(config.ETHSecret()) // 要求來源的私鑰
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(address)
	var data []byte
	sendValue, _ := new(big.Int).SetString(amount, 10) // send ether
	tx := types.NewTransaction(nonce, toAddress, sendValue, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

// GetBalance get eth balance
func GetBalance(_ethaddr string) (string, error) {
	client := rpc.GetEthInstance()

	account := common.HexToAddress(_ethaddr)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", err
	}
	return balance.String(), nil
}

// GetTxDetail get btc tx detail
func GetTxDetail(txid string) (*types.Transaction, bool, error) {
	return rpc.GetEthInstance().TransactionByHash(context.Background(), common.HexToHash(txid))
}
