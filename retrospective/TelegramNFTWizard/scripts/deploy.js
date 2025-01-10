// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const { ethers } = require("hardhat");
const { toWei, fromWei } = require("./lib.js");
//const { toWei, fromWei } = require("./lib.js");
const hre = require("hardhat");

async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile
  // manually to make sure everything is compiled
  // await hre.run('compile');

  //var _account = hre.ethers.getSigner;
  //console.log(_account.address);

  console.log(hre.network.name);

  
  let owner;
  owner = await hre.ethers.getSigner();
  console.log("owner address:", owner.address);
  const balance = await ethers.provider.getBalance(owner.address);
  console.log(ethers.utils.formatEther(balance), "ETH");
  

  // We get the contract to deploy
  const SingletonNFT = await hre.ethers.getContractFactory("SingletonNFT");
  const singleton = await SingletonNFT.deploy("Telegram Collection Singleton", "TGNFT");

  await singleton.deployed();
  console.log("Singleton Collection deployed to:", singleton.address);

  const FactoryNFT = await hre.ethers.getContractFactory("FactoryNFT");
  const factory = await FactoryNFT.deploy();
  await factory.deployed();
  console.log("Factory NFT deployed to:", factory.address);
}


// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
