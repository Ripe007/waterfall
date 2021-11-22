const BigNumber = ethers.BigNumber;
const addresses = require("../addresses.js");
const initialSupply = BigNumber.from("10").pow("18").mul("100000000");

async function main() {
  const WTF = await ethers.getContractFactory("WTF");
  const wtf = await WTF.deploy(initialSupply, addresses.teamMultisig );
  await wtf.deployed();
  console.log(`WTF token is deployed to: ${wtf.address}`);
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
