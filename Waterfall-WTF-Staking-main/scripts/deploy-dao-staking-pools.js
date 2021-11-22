// Process environment
require('dotenv').config();
import { BigNumber } from 'ethers'; 

const busd = '0xe9e7cea3dedca5984780bafc599bd69add087d56'; // mainnet busd
const wtf = '0x000000';
const vewtf = '0x00000';
const communityfund = '0x00000';

const d18 = ethers.BigNumber.from("10").pow("18"); // check how it works
const wtfTotalRewards = d18.mul('1000000');
const wtfStartRewardBlock = 77799999999; // configure here
const wtfEndRewardBlock = 9923949999234; // configure here
const wtfRewardPerBlock = wtfTotalRewards.div(wtfEndRewardBlock - wtfStartRewardBlock);

async function main() {

	const WTFRewards = await ethers.getContractFactory("WTFRewards");
	const FeeRewards = await ethers.getContractFactory("FeeRewards");
	const wtfrewards = await this.WTFRewards.deploy(vewtf, 
			                                        wtf,
			                                        communityfund,
			                                        wtfRewardPerBlock,
			                                        wtfStartRewardBlock,
			                                        wtfEndRewardBlock);
	await wtfrewards.deployed();

	// Deploy fee rewards 

	const feerewards = await this.FeeRewards.deploy(vewtf, busd);
	await feeRewards.deployed();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });