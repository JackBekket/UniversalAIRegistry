# Telegram NFT Wizard Contracts
https://telegram-nft-wizard.vercel.app
Smart contracts for Telegram NFT-Wizard bot

Wizard accepts 'file_id' obtained from telegram bot and can:
1. Create NFT as a single record inside SingletonNFT Collection (gas efficient and cheap)
2. Create the whole new NFT collection (deploy new erc721 contract from FactoryNFT.sol) and pollute it with file_ids[] items

Assuming that this repo is using alongside with Telegram NFT Wizard bot, so the whole how-to look like this:
1. User start telegram bot
2. Telegram bot await input media files from user, upload it to telegram server and obtain unique file_id (string)
3. Telegram bot generates link like https://telegram-nft-wizard.vercel.app/?file_id="someuniquestring"
4. User click button create NFT
5. ????
6. PROFIT!

It's also assuming further integration into NFT marketplace and Telegram-bot-NFT-Marketplace




Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```

## Generating ABI
```
solc --abi --bin ./contracts/FactoryNFT.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,
solc --abi --bin ./contracts/SampleNFT.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,
solc --abi --bin ./contracts/SingletonNFT.sol -o build ..=.. --overwrite --allow-paths *,/node_modules/,


```


## Generating GO
```
abigen --abi="build/FactoryNFT.abi" --pkg=FactoryNFT --out="./go/FactoryNFT/FactoryNFT.go"
abigen --abi="build/SingletonNFT.abi" --pkg=SingletonNFT --out="./go/SingletonNFT/SingletonNFT.go"
```
