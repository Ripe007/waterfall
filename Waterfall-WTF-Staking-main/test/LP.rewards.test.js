const BN = require("bignumber.js");
const { expect } = require("chai");
var moment = require('moment');
const {
  getBlockTimestamp,
  setBlockTime,
  increaseTime,
  mineBlocks,
  sync: { getCurrentTimestamp },
} = require("./helpers.js");

describe ('WTF Staking Rewards', function(){
	let snapshotId;
	const d18 = ethers.BigNumber.from("1000000000000000000");
	const year = 31536000;
    const month = 2592000;
    const WTF_total_reward = d18.mul(800000);
	
	beforeEach(async function(){
		snapshotId = await ethers.provider.send('evm_snapshot')
		const WTF = await ethers.getContractFactory("WTFMock");
		const LPFake = await ethers.getContractFactory("LPFake");
		const CommunityFund = await ethers.getContractFactory("CommunityFund");
		this.LPRewards = await ethers.getContractFactory("WTFLPRewards");
		this.wtf = await WTF.deploy();
		await this.wtf.deployed()
		this.lpfake = await LPFake.deploy();
		await this.lpfake.deployed();
		this.communityfund = await CommunityFund.deploy(this.wtf.address);
		await this.communityfund.deployed();
	})
	afterEach(async function () {
        await ethers.provider.send('evm_revert', [snapshotId])
    })

	it('should give LP staking rewards', async function(){

		const [deployer, other1, other2, other3] = await ethers.getSigners();
		const _currTimestamp = await getCurrentTimestamp();
		const _startRewardTs = ethers.BigNumber.from(moment().add(2,'hours').unix());
		const _endRewardTs = _startRewardTs.add(year); 
	
		let PRECISION = ethers.BigNumber.from('1000000000000');

		let LPRewards = await this.LPRewards.deploy( this.wtf.address,
			                                         this.lpfake.address,
			                                         this.communityfund.address,
			                                         _startRewardTs, 
			                                         _endRewardTs,
			                                         WTF_total_reward
			                                         );
		await LPRewards.deployed();

		// Mint to community fund

		await this.wtf.connect(deployer).mint(this.communityfund.address, WTF_total_reward);
		await this.wtf.connect(deployer).transfer(this.communityfund.address, WTF_total_reward);
		await this.communityfund.connect(deployer).setAllowance(LPRewards.address, WTF_total_reward);

		// First mint to Other1

		await(await this.lpfake.connect(deployer).mint(other1.address, d18.mul(10000))).wait(1);

		await (await this.lpfake.connect(other1).approve(LPRewards.address, d18.mul(10000))).wait(1)

		await(await this.lpfake.connect(deployer).mint(other2.address, d18.mul(10000))).wait(1);

		await (await this.lpfake.connect(other2).approve(LPRewards.address, d18.mul(10000))).wait(1);

		await LPRewards.connect(other1).stake(d18.mul(10000));
		await LPRewards.connect(other2).stake(d18.mul(10000));

		await increaseTime(month*3);
		await mineBlocks(1);

		let pendingReward1 = await LPRewards.pendingReward(other1.address);
		let pendingReward2 = await LPRewards.pendingReward(other2.address);

		expect(pendingReward1).to.be.equal(pendingReward2);

		await increaseTime(month);
		await mineBlocks(1);

		pendingReward1 = await LPRewards.pendingReward(other1.address);
		pendingReward2 = await LPRewards.pendingReward(other2.address);

		console.log(`Pending1: ${pendingReward1.toString()}`)
		console.log(`Pending2: ${pendingReward2.toString()}`)

		expect(pendingReward1).to.be.equal(pendingReward2);

		await increaseTime(year);
		await mineBlocks(1);

		// claim rewards

		await LPRewards.connect(other1).claimRewards();

		await LPRewards.connect(other2).claimRewards();

		let balother1 = await this.wtf.balanceOf(other1.address);
		let balother2 = await this.wtf.balanceOf(other2.address);

		console.log(`Balance other1: ${balother1}`);
		console.log(`Balance other2: ${balother2}`);

		expect (balother1.add(balother2)).to.be.equal(WTF_total_reward);

	})

})

