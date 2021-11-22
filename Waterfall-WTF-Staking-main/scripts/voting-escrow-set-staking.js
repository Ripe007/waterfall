const vewtf = "0x00000....";
const wtfRewards = "0x00000";
const feeRewards = "0x00000";
const communityfund = "0x0000";

const d18 = ethers.BigNumber.from("10").pow("18"); // check how it works
const wtfTotalRewards = d18.mul('1000000');

async function main() {
	const VotingEscrow = await ethers.getContractFactory("VotingEscrow");
	const CommunityFund = await ethers.getContractFactory("CommunityFund");
	const votingescrow = await VotingEscrow.attach(vewtf);
	const cf = await CommunityFund.attach(communityfund);
	
	// Set WTF staking
	await votingescrow.setStaking(wtfRewards, feeRewards);

	// set allowance of WTF rewards for community fund fees

	await cf.setAllowance(wtfRewards, wtfTotalRewards);
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });