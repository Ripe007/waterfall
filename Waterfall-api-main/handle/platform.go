package handle

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	"github.com/waterfall/contracts"
	wattypes "github.com/waterfall/types"
	apierr "github.com/waterfall/types/error"
	"golang.org/x/crypto/sha3"
)

// WTF token address 0x80dE8e844F152DfAe30B89289f0f953C60C75b1D:

type Response = wattypes.AppResponse

// solidity:RationScale
const PercentageParamScal = 10000
const PercentageScale = 1000000000000000000

type Platform struct {
	RPCClient                *wattypes.RPCClientRoundRobin
	PlatformAddress          ethcommon.Address
	CampaignContinuousCycles ethcommon.Address
	Tokenlist                []wattypes.TokenInfo
}

func (pltf *Platform) Mint(c echo.Context) error {
	rpc := "https://data-seed-prebsc-1-s3.binance.org:8545"
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA("fa41515dfb14f2baa9832377201043d66e2ec70ee193710f6d33f3577d11b782")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := ethcommon.HexToAddress("0x5cbAb697c2440E3dCe5bebe48eC7Ef49baE1E699")
	// instance, err := contracts.NewToken((contractAddress), client)
	// if err != nil {
	// return err
	// }
	toAddress := common.HexToAddress("0x72e738ae000410Bcb9bA88A7CEB476B714bf7143")

	mintfnSignature := []byte("mint(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(mintfnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))

	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))

	var data []byte

	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	gasLimit = uint64(3000000)

	fmt.Println(gasLimit)
	tx := types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return err
}

// waterfall function
func (pltf *Platform) Balance(c echo.Context) error {
	input := &wattypes.BalanceInput{}
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.BadParams,
			Msg:  err.Error(),
		})
	}
	token, err := pltf.getTokenInfo(ethcommon.HexToAddress(input.Token))
	if err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.UnknownToken,
			Msg:  err.Error(),
		})
	}
	erc20, err := contracts.NewERC20(ethcommon.HexToAddress(input.Token), pltf.RPCClient.Next())
	if err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.Contract,
			Msg:  err.Error(),
		})
	}
	balance, err := erc20.BalanceOf(nil, ethcommon.HexToAddress(input.User))
	if err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.Contract,
			Msg:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		Code: apierr.OK,
		Data: pltf.prettifyAmount(balance, big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(token.Decimals)), nil)),
	})

}

func (pltf *Platform) NextCycle(c echo.Context) error {
	input := &wattypes.FarmInput{}
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.BadParams,
			Msg:  err.Error(),
		})
	}
	privateKey, err := crypto.HexToECDSA(input.User)
	if err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.BadParams,
			Msg:  err.Error(),
		})

	}
	publicKey := privateKey.Public()
	publicKeyESDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return c.JSON(
			http.StatusOK, Response{
				Code: apierr.BadParams,
				Msg:  "cannot assert type: publicKey is not of type *ecdsa.PublicKey",
			})
	}
	// fromAddress := crypto.PubkeyToAddress(*publicKeyESDSA)
	// nonce, err := pltf.RPCClient.n

	address := common.HexToAddress(input.Address)
	instance, err := contracts.NewCampaignContinuousCycles(address, pltf.RPCClient.Next())
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code: apierr.BadParams,
			Msg:  err.Error(),
		})
	}
	fmt.Println("contract is loaded")
	operaterAddress, err := instance.AlpacaBUSD(nil)
	if err != nil {
		return c.JSON(http.StatusOK, Response{
			Code: apierr.BadParams,
			Msg:  err.Error(),
		})

	}
	fmt.Println("operatorAddress:", operaterAddress)
	return c.JSON(http.StatusOK, Response{
		Code: apierr.OK,
		Msg:  operaterAddress.String(),
	})
}
func (pltf *Platform) getTokenInfo(address ethcommon.Address) (wattypes.TokenInfo, error) {
	for _, t := range pltf.Tokenlist {
		if address == ethcommon.HexToAddress(t.Address) {
			return t, nil
		}
	}
	return wattypes.TokenInfo{}, fmt.Errorf(apierr.UnknownTokenMsg)
}
func (pltf *Platform) prettifyAmount(p *big.Int, d *big.Int) string {
	pricef := new(big.Float).SetInt(p)
	scalef := new(big.Float).SetInt(d)
	pricef.Quo(pricef, scalef)
	return pricef.Text('f', 8)
}
