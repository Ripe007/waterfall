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
	const d18 = ethers.BigNumber.from("1000000000000000000");
	let snapshotId;
	beforeEach(async function(){
		snapshotId = await ethers.provider.send('evm_snapshot');
		const WTF = await ethers.getContractFactory("WTFMock");
		const BUSD = await ethers.getContractFactory("BUSDMock");
		const CommunityFund = await ethers.getContractFactory("CommunityFund");
		const [deployer, other1, other2, other3] = await ethers.getSigners();
		this.wtf = await WTF.deploy();
		await this.wtf.deployed()
		this.busd = await BUSD.deploy();
		await this.busd.deployed();
		this.communityfund = await CommunityFund.deploy(this.wtf.address);
		await this.communityfund.deployed();
		this.VotingEscrow = await ethers.getContractFactory("VotingEscrow");
		this.WTFStakingRewards = await ethers.getContractFactory("WTFRewards");
		this.FeeStakingRewards = await ethers.getContractFactory("FeeRewards");
		this.year = 31536000;
		this.month = 2592000;
		this.WTF_total_reward = d18.mul(1000000);
		this.blocks = 1000;
		this.rewardPerBlock = this.WTF_total_reward.div(this.blocks);
	})

	afterEach(async function () {
        await ethers.provider.send('evm_revert', [snapshotId])
    })

	it('Extends lock time and mints VeWTF', async function(){
		const [deployer, other1, other2, other3] = await ethers.getSigners();
		let userWTFAmount = d18.mul(10000);
		let _currentBlock = await web3.eth.getBlockNumber();
		let _endRewardBlock = _currentBlock + this.blocks;
		
		let lockDuration = ethers.BigNumber.from(this.year).div(2); // half-year

		// Deploy Voting Escrow

		let votingescrow = await this.VotingEscrow.deploy(this.wtf.address);
		await votingescrow.deployed();

		// Deploy WTF Staking Rewards

		let wtfRewards = await this.WTFStakingRewards.deploy(votingescrow.address, 
			                                              this.wtf.address,
			                                              this.communityfund.address,
			                                              this.rewardPerBlock, 
			                                              _currentBlock,
			                                              _endRewardBlock
			                                               );
		await wtfRewards.deployed();

		// Deploy fee rewards 

		let feeRewards = await this.FeeStakingRewards.deploy(votingescrow.address, this.busd.address);

		await feeRewards.deployed();

		// Set staking rewards contract addresses

		await votingescrow.connect(deployer).setStaking(wtfRewards.address, feeRewards.address);
		await this.wtf.connect(deployer).transfer(this.communityfund.address, this.WTF_total_reward);
		await this.communityfund.setAllowance(wtfRewards.address, this.WTF_total_reward);

		// Lock some WTF to Vesting Escrow

		// First mint to User

		await(await this.wtf.connect(deployer).mint(other1.address, userWTFAmount)).wait(1);

		// approve for Voting Escrow

		await this.wtf.connect(other1).approve(votingescrow.address, userWTFAmount);

		// Create lock

		await votingescrow.connect(other1).createLock(d18.mul(10000), lockDuration);

		// Check locked amount

		let wtfBal = await votingescrow.connect(other1).getLockedAmount(other1.address);

		expect(wtfBal).eq(d18.mul(10000));

		let vewtfBal = await votingescrow.balanceOf(other1.address);

		console.log(`Percent of total staked: ${vewtfBal.mul(100).div(userWTFAmount).toNumber()}`)

		// Since lock duration is 6 month and max duration is 1 year we expect vewtf balance = wtfStaked / 2 

		expect(vewtfBal).eq(userWTFAmount.div(2));

		// Make lock duration two times bigger than now (~ 6 month)

		let lock = await votingescrow.connect(other1).getLockData(other1.address);

		let currentDuration = lock.expiryTimestamp.sub(lock.startTimestamp);

		let newExpiry = currentDuration.add(lock.expiryTimestamp);

		let expiryDiff = newExpiry.sub(lock.expiryTimestamp);

		// Check balance 

		let vewtfBalBefore = await votingescrow.connect(other1).balanceOf(other1.address);

		await votingescrow.connect(other1).increaseTimeAndAmount(0, newExpiry);

		// Check balance 

		let vewtfBalAfter = await votingescrow.connect(other1).balanceOf(other1.address);

		// Balance should currentBalance * 2 because currentDuration = previousDuration * 2

		console.log(`VeWTF balance before lock extension: ${vewtfBalBefore}`);
		console.log(`VeWTF balance after lock extension: ${vewtfBalAfter}`);


		expect(vewtfBalAfter).eq(vewtfBalBefore.mul(2));

	})

	it('Unlocks and sends back locked WTF amount', async function(){

		const [deployer, other1, other2, other3] = await ethers.getSigners();

		let wtfUserAmount = d18.mul(10000);
		let _currentBlock = await web3.eth.getBlockNumber();
		let _endRewardBlock = _currentBlock + this.blocks;
		let lockDuration = ethers.BigNumber.from(this.year).div(2); // half-year

		// Deploy Voting Escrow

		let votingescrow = await this.VotingEscrow.deploy(this.wtf.address);
		await votingescrow.deployed();

		// Deploy WTF Staking Rewards

		let wtfRewards = await this.WTFStakingRewards.deploy(votingescrow.address, 
			                                              this.wtf.address,
			                                              this.communityfund.address,
			                                              this.rewardPerBlock, 
			                                              _currentBlock,
			                                              _endRewardBlock
			                                               );
		await wtfRewards.deployed();

		// Deploy fee rewards 

		let feeRewards = await this.FeeStakingRewards.deploy(votingescrow.address, this.busd.address);

		await feeRewards.deployed();

		// Set staking rewards contract addresses

		await votingescrow.connect(deployer).setStaking(wtfRewards.address, feeRewards.address);
		await this.wtf.connect(deployer).transfer(this.communityfund.address, this.WTF_total_reward);
		await this.communityfund.setAllowance(wtfRewards.address, this.WTF_total_reward);

		// Lock some WTF to Vesting Escrow

		// First mint to User

		await(await this.wtf.connect(deployer).mint(other1.address, wtfUserAmount)).wait(1);

		// approve for Voting Escrow

		await this.wtf.connect(other1).approve(votingescrow.address, wtfUserAmount);

		// Create lock

		await votingescrow.connect(other1).createLock(d18.mul(10000), lockDuration);

		// Try unlock before expiry

		await increaseTime(this.month);

		await expect(votingescrow.connect(other1).unlock()).to.be.revertedWith("WTF Voting Escrow: Cannot unlock tokens before expiry");

		// Fast forward to the unlock period

		let wtfBalBefore = await this.wtf.balanceOf(other1.address);

		expect(wtfBalBefore).eq(0);

		await increaseTime(this.year);

		let vewtfBalBefore = await votingescrow.balanceOf(other1.address);

		await votingescrow.connect(other1).unlock();

		let wtfBalAfter = await this.wtf.balanceOf(other1.address);
		let vewtfBalAfter = await votingescrow.balanceOf(other1.address);

		console.log(`VEWTF balance before: ${vewtfBalBefore}`);
		console.log(`VEWTF balance after: ${vewtfBalAfter}`);

		console.log(`WTF before: ${wtfBalBefore}`);
		console.log(`WTF after: ${wtfBalAfter}`);

		expect(wtfBalAfter).gte(wtfUserAmount);

	})

	})