const wtf = '0x00000';
async function main() {
	const CommunityFund = await ethers.getContractFactory("CommunityFund");
	const cf = await CommunityFund.deploy(wtf);
	await cf.deployed();
	console.log(`Community fund is deployed to: ${cf.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });