## btp-wallet
The command-line wallet of the Btp network


## download or build

you can download from [release](https://github.com/btceasypay/btp-desktop-wallet/releases)

if you have golang environment, you can build it by yourself

1. install node  
    [https://nodejs.org](https://nodejs.org)
    
    #### Note:
    If you met network connectivity issue, please config HTTP or NPM proxy. See [#37](https://github.com/btceasypay/btp-desktop-wallet/issues/37)
    ```shell script
   # HTTP proxy
   export HTTP_PROXY=http://proxy.example.com:8080
   
   # NPM proxy
   npm config set proxy http://proxy.example.com:8080
   ```

2. install go  
    [https://golang.org/](https://golang.org/)  
    make sure go env is ready
    ```bash
    go env GOPATH
    ```
   
   ```shell script
   go version
   
   # Recommend for debian based OS to upgrade to the latest
   # https://github.com/golang/go/wiki/Ubuntu
    sudo add-apt-repository ppa:longsleep/golang-backports
    sudo apt update
    sudo apt install golang-go
    ```
3. build
    ```bash
    git clone https://github.com/btceasypay/btp-desktop-wallet ~/github.com/btceasypay/btp-desktop-wallet
    cd ~/github.com/btceasypay/btp-desktop-wallet
    make
    ```

## Usage

1. bc
    ```shell script
    ./btp-wallet bc
    
   btp wallet command
   
   Usage:
     btp-wallet bc [command]
   
   Available Commands:
     create                create
     createnewaccount      create new account
     getaddressesbyaccount get addresses by account
     getbalance            getbalance
     getlisttxbyaddr       get all transactions for address
     getnewaddress         create new address by account
     gettx                 Access to transaction information
     gettxspendinfo        gettxspendinfo
     importprivkey         import priKey
     listaccountsbalance   list Accounts Balance
     sendtoaddress         send transaction
     setsyncedtonum         please use caution when specifying how many blocks to update from
     syncheight            Get the number of local synchronization blocks
     updateblock           Update local block data
   
   Flags:
     -a, --appdatadir string       wallet db path
     -c, --configfile string       config file (default "config.toml")
         --confirmations int       Number of block confirmations  (default 10)
         --create                  Create a new wallet
     -d, --debuglevel string       Logging level {trace, debug, info, warn, error, critical} (default "info")
     -h, --help                    help for bc
         --listeners stringArray   rpc listens (default [127.0.0.1:38130])
     -l, --logdir string           log data path
         --mintxfee int            The minimum transaction fee default 20000 (aka. 0.0002 MEER/KB) (default 200000)
     -n, --network string          network (default "testnet")
     -P, --pubwalletpass string    data encryption password (default "public")
         --bcert string            Certificate path
         --qnotls                  disable TLS (default true)
     -p, --qpass string            btp node password (default "123456")
     -s, --qserver string          btp node server (default "127.0.0.1:8030")
         --qtlsskipverify          skip TLS verification (default true)
     -u, --quser string            btp node user (default "admin")
          --rpcPass string          rpc pass,default by random (default "OkROTj7OdtUFm94DwAlzJ2Nm")
          --rpcUser string          rpc user,default by random (default "3tkpuiUE")
         --ui                      Start Wallet with RPC and webUI interface (default true)
   
   Use "btp-wallet bc [command] --help" for more information about a command.

    ```

2. bx
    ```shell script
   ./btp-wallet bx
   
    btp wallet bx util
    
    Usage:
      btp-wallet bc bx [command]
    
    Available Commands:
      generatemnemonic generate mnemonic
      mnemonictoaddr   mnemonic to address
      mnemonictoseed   mnemonic to seed
      pritoaddr        private key to address
      pritopub         private key to public key
      pubtoaddr        public key to address
      seedtoaddr       seed to address
      seedtopri        Seed private key
      wiftopri         WIF key to private key
    ```

## How to use btp-wallet console command model

1:  Rename sample-config.toml to config.toml and modify the configuration parameters

```toml
   #configFile="" #Your config.toml profile directory
   #appDataDir="" # Your DB storage path
   #logDir="" # log path
   #network="mainnet" #network mainnet,testnet,privnet default testnet
   network="testnet"
   #Btpd
   QServer="127.0.0.1:8131"
   QUser="admin"
   QPass="123456"
   QNoTLS=true
   QTLSSkipVerify=true
   WalletPass="public" #Wallet encryption code default public
   
   MinTxFee=200000   # The minimum transaction fee default 200000 (aka. 0.002 MEER/kB)
   
   #web model
   #listeners=["127.0.0.1:8130"]
   #rpcUser=""
   #rpcPass=""
   #ui=true

```

2:  create a  wallet

```shell script
    ./btp-wallet bc create 
    
    
    #output
    
     # Wallet Password
     Enter the private passphrase for your new wallet:
     Confirm passphrase:
     # Wallet data encryption password, default public
     Do you want to add an additional layer of encryption for public data? (n/no/y/yes) [no]:y
     Enter the public passphrase for your new wallet:
     Confirm passphrase:
     NOTE: Use the -- wallet pass option to configure your public passphrase.
     PubPass: 123
     # Whether to import wallet seeds
     Do you have an existing wallet seed you want to use? (n/no/y/yes) [no]: n
     Your wallet generation seed is:
     17e28af99e36ff4527c95f91e13d3ecd82349864d23b9ff2d4f9d446ea078291
     IMPORTANT: Keep the seed in a safe place as you
     Will NOT be able to restore your wallet without it.
     Please keep in mind that any who has access
     To the seed can also restore your wallet thus
     Give them access to all your funds, so it is
     Imperative that you keep it in a secure location.
     Once you have stored the seed in a safe and secure location, enter "OK" to continue: OK
     Creating the wallet...
     INFO [0021] Opened Wallet
     Pri: 6eb6bbcd7ded317abc4ed5e373c2c8630dc4ad069470ad7ae72f5fb854423006
     INFO [0022] Imported payment address TmmC1hUN5A2RJzX9ZWFZqHaDbKUf6NaA4D
     The wallet has been created successfully.
```

3:  update db (update to the specified block or update all,syncheight can view the current DB synchronization height)
```shell script
    ./btp-wallet bc updateblock
    ./btp-wallet bc updateblock 130 
```
    
4:  when creating a wallet, you can import seeds or import private keys using importprivkey.

```shell script
    ./btp-wallet bc importprivkey 6eb6bbcd7ded317abc4ed5e373c2c8630dc4ad069470ad7ae72f5fb854423006` 
    
    #output
    
    ImportPrivKey: OK

```

    
5:  must use updateback to see the balance change after the transfer transaction

```shell script
    ./btp-wallet bc sendtoaddress TmbsdsjwzuGboFQ9GcKg6EUmrr3tokzozyF 0.01 youpassword
    
    #output
    
    0e441ecee44defe28711103eef0cc3d01c187c257738150869c032fbbf96d4c9

    ./btp-wallet bc updateback

```

## Web client
```shell script
./btp-wallet web
```
![desktop wallet](assets/wallet-info.png)

## Interactive console
```shell script
./btp-wallet console

wallet-cli:help
Usage:
        <command> [arguments]
        The commands are:
        <createNewAccount> : Create a new account. Parameter: [account]
        <getBalance> : Query the specified address balance. Parameter: [address]
        <getListTxByAddr> : Gets all transaction records at the specified address. Parameter: [address] [stype:in,out,all]
        <getNewAddress> : Create a new address under the account. Parameter: [account]
        <getAddressesByAccount> : Check all addresses under the account. Parameter: [account]
        <getAccountByAddress> : Inquire about the account number of the address. Parameter: [address]
        <importPrivKey> : Import private key. Parameter: [priKey]
        <importWifPrivKey> : Import wif format private key. Parameter: [priKey]
        <dumpPrivKey> : Export wif format private key by address. Parameter: [address]
        <getAccountAndAddress> : Check all accounts and addresses. Parameter: []
        <sendToAddress> : Transfer transaction. Parameter: [address] [num]
        <updateblock> : Update Wallet Block. Parameter: []
        <syncheight> : Current Synchronized Data Height. Parameter: []
        <unlock> : Unlock Wallet. Parameter: [password]
        <help> : help
        <exit> : Exit command mode

wallet-cli:

```

