async function main() {
	const WTF = await ethers.getContractFactory("WTF");
	const wtf = await WTF.deploy();
	await wtf.deployed();
	console.log(`WTF token deployed to: ${wtf.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });