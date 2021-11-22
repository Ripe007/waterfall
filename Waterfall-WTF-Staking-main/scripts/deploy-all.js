import { BigNumber } from 'ethers'; 
const wtf = "0x00000";
const busd = '0xe9e7cea3dedca5984780bafc599bd69add087d56'; // mainnet busd

// Staking rewards config

const d18 = ethers.BigNumber.from("10").pow("18"); // check how it works
const wtfTotalRewards = d18.mul('1000000'); // 1 million in rewards
const wtfStartRewardBlock = 77799999999; // configure here
const wtfEndRewardBlock = 9923949999234; // configure here
const wtfRewardPerBlock = wtfTotalRewards.div(wtfEndRewardBlock - wtfStartRewardBlock);

async function main () {
	const VotingEscrow = await ethers.getContractFactory("VotingEscrow");
	const CommunityFund = await ethers.getContractFactory("CommunityFund");
	const WTFRewards = await ethers.getContractFactory("WTFRewards");
	const FeeRewards = await ethers.getContractFactory("FeeRewards");

	// Deploy community fund
	const cf = await CommunityFund.deploy(wtf);
	await cf.deployed();
	console.log(`Community fund is deployed to: ${cf.address}`);
	
	// Deploy voting escrow

	const votingescrow = await VotingEscrow.deploy(wtf);
	await votingescrow.deployed();
	console.log(`Voting Escrow (VeWTF) is deployed to : ${votingescrow.address}`);

	// Deploy Staking

	const wtfrewards = await this.WTFRewards.deploy(votingescrow.address, 
			                                        wtf,
			                                        cf.address,
			                                        wtfRewardPerBlock,
			                                        wtfStartRewardBlock,
			                                        wtfEndRewardBlock);
	await wtfrewards.deployed();

	// Deploy fee rewards 

	const feerewards = await this.FeeRewards.deploy(votingescrow.address, busd);
	await feeRewards.deployed();

	//Set staking

	// Set WTF staking
	await votingescrow.setStaking(wtfRewards.address, feeRewards.address);

	// set allowance of WTF rewards for community fund fees

	await cf.setAllowance(wtfRewards.address, wtfTotalRewards);

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });