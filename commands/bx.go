package commands

import (
	"fmt"
	"github.com/btceasypay/bitcoinpay/bx"
	"github.com/spf13/cobra"
)

var BxCmd = &cobra.Command{
	Use:   "bx",
	Short: "bx util",
	Long:  `btp wallet bx util`,
}

func AddQxCommand() {
	BxCmd.AddCommand(generatemnemonicCmd)
	BxCmd.AddCommand(mnemonictoseedCmd)
	BxCmd.AddCommand(seedtopriCmd)
	BxCmd.AddCommand(pritopubCmd)
	BxCmd.AddCommand(mnemonictoaddrCmd)
	BxCmd.AddCommand(seedtoaddrCmd)
	BxCmd.AddCommand(pritoaddrCmd)
	BxCmd.AddCommand(pubtoaddrCmd)
	BxCmd.AddCommand(wifToPriCmd)
}

var generatemnemonicCmd = &cobra.Command{
	Use:   "generatemnemonic",
	Short: "generate mnemonic",
	Example: `
		generatemnemonic
		`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		msg, err := bx.NewEntropy(32)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		bx.MnemonicNew(msg)
	},
}
var mnemonictoseedCmd = &cobra.Command{
	Use:   "mnemonictoseed",
	Short: "mnemonic to seed",
	Example: `
		mnemonictoseed "mnemonic"
		`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bx.MnemonicToSeed("", args[0])
	},
}
var seedtopriCmd = &cobra.Command{
	Use:   "seedtopri",
	Short: "Seed private key",
	Example: `
		seedtopri "seed"
		`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		msg, err := bx.EcNew("secp256k1", args[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
	},
}
var pritopubCmd = &cobra.Command{
	Use:   "pritopub {private key} {bool,uncompressed,defalut false}",
	Short: "private key to public key",
	Example: `
		pritoaddr pri 
		pritoaddr pri false
		`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b := false
		if len(args) > 1 {
			uncompressed := args[1]
			if uncompressed == "true" {
				b = true
			}
		}
		msg, err := priToPub(args[0], b)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
	},
}
var pubtoaddrCmd = &cobra.Command{
	Use:   "pubtoaddr {public key} {string,network value: mainnet,privnet,testnet}",
	Short: "public key to address",
	Example: `
		pubtoaddr "pub" "testnet"
		`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[1] != "mainnet" && args[1] != "privnet" && args[1] != "testnet" {
			fmt.Println("Wrong network type")
			return
		}
		msg, err := pubToAddr(args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
	},
}

var mnemonictoaddrCmd = &cobra.Command{
	Use:   "mnemonictoaddr {mnemonic} {string,network value: mainnet,privnet,testnet}",
	Short: "mnemonic to address",
	Example: `
		mnemonictoaddr "mnemonic" "testnet"
		`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[1] != "mainnet" && args[1] != "privnet" && args[1] != "testnet" {
			fmt.Println("Wrong network type")
			return
		}
		msg, err := mnemonicToAddr(args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
	},
}
var seedtoaddrCmd = &cobra.Command{
	Use:   "seedtoaddr {seed} {string,network value: mainnet,privnet,testnet}",
	Short: "seed to address",
	Example: `
		seedtoaddr "seed" "testnet"
		`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[1] != "mainnet" && args[1] != "privnet" && args[1] != "testnet" {
			fmt.Println("Wrong network type")
			return
		}
		msg, err := seedToAddr(args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
	},
}

var pritoaddrCmd = &cobra.Command{
	Use:   "pritoaddr {pri} {string,network value: mainnet,privnet,testnet}",
	Short: "private key to address",
	Example: `
		pritoaddr "pri" "testnet"
		`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if args[1] != "mainnet" && args[1] != "privnet" && args[1] != "testnet" {
			fmt.Println("Wrong network type")
			return
		}
		msg, err := priToAddr(args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(msg)
	},
}

var wifToPriCmd = &cobra.Command{
	Use:   "wiftopri {string, wif}",
	Short: "WIF key to private key",
	Example: `
		wiftopri  5HueCGU8rMjxEXxiPuD5BDku4MkFqeZyd4dZ1jvhTVqvbTLvyTJ
		`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bx.WifToEcPrivateKey(args[0])
	},
}
