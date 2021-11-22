const wtf = "0x0000000";

async function main() {
	const VotingEscrow = await ethers.getContractFactory("VotingEscrow");
	const votingescrow = await VotingEscrow.deploy(wtf);
	await votingescrow.deployed();
	console.log(`Voting Escrow (VeWTF) is deployed to : ${votingescrow.address}`);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });