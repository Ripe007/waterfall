package main

import (
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/waterfall/types"
	wattypes "github.com/waterfall/types"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"waterfall-contract-api/handle"
)

func main() {
	configs := new(types.AppConfig)
	fmt.Println("msg")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	bs, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Panic(err)
	}

	if err := yaml.Unmarshal(bs, &configs); err != nil {
		log.Panic(err)
	}
	fmt.Print(configs.Port)

	rpcs, err := types.NewRPCClientRoundRobin(configs.RPCs)
	if err != nil {
		log.Panic(err)
	}
	platformHandle := &handle.Platform{
		RPCClient:                rpcs,
		PlatformAddress:          ethcommon.HexToAddress(configs.PlatformAddress),
		CampaignContinuousCycles: ethcommon.HexToAddress(configs.CampaignContinuousCyclesAddress),
		Tokenlist: []wattypes.TokenInfo{
			{
				"BUSD",
				"0x5cbAb697c2440E3dCe5bebe48eC7Ef49baE1E699",
				18,
				"BUSD",
			},
		},
	}

	e.GET("/balance", platformHandle.Balance)
	e.GET("/mint", platformHandle.Mint)
	e.GET("/nextCycle", platformHandle.NextCycle)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world")
	})
	e.Logger.Fatal(e.Start(":" + configs.Port))

}
