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
	beforeEach(async function(){
		snapshotId = await ethers.provider.send('evm_snapshot');
		this.year = 31536000;
		this.month = 2592000;
		this.WTF_total_reward = d18.mul(800000);
		this.blocks = 400;
		this.rewardPerBlock = this.WTF_total_reward.div(this.blocks);

		const WTF = await ethers.getContractFactory("WTFMock");
		const BUSD = await ethers.getContractFactory("BUSDMock");
		const CommunityFund = await ethers.getContractFactory("CommunityFund");
		this.VotingEscrow = await ethers.getContractFactory("VotingEscrow");
		this.WTFStakingRewards = await ethers.getContractFactory("WTFRewards");
		this.FeeStakingRewards = await ethers.getContractFactory("FeeRewards");
		
		const [deployer, other1, other2, other3] = await ethers.getSigners();
		this.wtf = await WTF.deploy();
		await this.wtf.deployed()
		this.busd = await BUSD.deploy();
		await this.busd.deployed();
		this.communityfund = await CommunityFund.deploy(this.wtf.address);
		await this.communityfund.deployed();
	
	})

	afterEach(async function () {
        await ethers.provider.send('evm_revert', [snapshotId])
    })

	it('Allow lock creation for EOA and disallow for contracts', async function(){
		
		const [deployer, other1, other2, other3] = await ethers.getSigners();
		let _currentBlock = await web3.eth.getBlockNumber();
		let _endRewardBlock = _currentBlock + this.blocks;
		
		let duration = ethers.BigNumber.from(this.year).div(2);

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
		await this.communityfund.setAllowance(wtfRewards.address, this.WTF_total_reward);


		// Lock some WTF to Vesting Escrow

		// First mint to User

		await(await this.wtf.connect(deployer).mint(other1.address, d18.mul(10000))).wait(1);

		// approve for Voting Escrow

		await this.wtf.connect(other1).approve(votingescrow.address, d18.mul(10000));

		// Create lock

		await votingescrow.connect(other1).createLock(d18.mul(10000), duration);

		// Check locked amount

		let wtfBal = await votingescrow.connect(other1).getLockedAmount(other1.address);

		expect(wtfBal).eq(d18.mul(10000));
		// Try to lock with contract. Should revert

		let TestContract = await ethers.getContractFactory("TestContractAccessToVEscrow");
		let testcontract = await TestContract.deploy(votingescrow.address);
		
		await testcontract.deployed();

		await expect(testcontract.testLock(d18.mul(10000),duration)).to.be.revertedWith("Smart contract depositors not allowed");

	})

	it('Stake and get WTF rewards', async function(){
		// Deploy staking (general with locking feature)
        const [deployer, other1, other2, other3] = await ethers.getSigners();
		let _currentBlock = await web3.eth.getBlockNumber();
		let _startRewardBlock = _currentBlock + 20;
		let _endRewardBlock = _startRewardBlock + this.blocks;
		
		let duration = ethers.BigNumber.from(this.year).div(4);
		let PRECISION = ethers.BigNumber.from('1000000000000');

		let votingescrow = await this.VotingEscrow.deploy(this.wtf.address);

		await votingescrow.deployed();


		let wtfRewards = await this.WTFStakingRewards.deploy(votingescrow.address, 
			                                              this.wtf.address,
			                                              this.communityfund.address,
			                                              this.rewardPerBlock, 
			                                              _startRewardBlock,
			                                              _endRewardBlock
			                                               );
		await wtfRewards.deployed();

		// Deploy fee rewards 

		let feeRewards = await this.FeeStakingRewards.deploy(votingescrow.address, this.busd.address);

		await feeRewards.deployed();

		await votingescrow.connect(deployer).setStaking(wtfRewards.address, feeRewards.address);
		
		// Lock some WTF to Vesting Escrow

		// Mint to community fund

		await(await this.wtf.connect(deployer).mint(this.communityfund.address, this.WTF_total_reward)).wait(1);
		await this.wtf.connect(deployer).transfer(this.communityfund.address, this.WTF_total_reward);
		await this.communityfund.connect(deployer).setAllowance(wtfRewards.address, this.WTF_total_reward);

		// First mint to Other1

		await(await this.wtf.connect(deployer).mint(other1.address, d18.mul(10000))).wait(1);

		await (await this.wtf.connect(other1).approve(votingescrow.address, d18.mul(10000))).wait(1)

		await(await this.wtf.connect(deployer).mint(other2.address, d18.mul(10000))).wait(1);

		await (await this.wtf.connect(other2).approve(votingescrow.address, d18.mul(10000))).wait(1)

		// Create lock

	    await votingescrow.connect(other1).createLock(d18.mul(10000), duration);

		await votingescrow.connect(other2).createLock(d18.mul(10000), duration);


		// Check locked amount

		let wtfBal1 = await votingescrow.connect(other1).getLockedAmount(other1.address);

		let wtfBal2 = await votingescrow.connect(other2).getLockedAmount(other1.address);

		expect(wtfBal1.toString()).eq(d18.mul(10000).toString());
		expect(wtfBal2.toString()).eq(d18.mul(10000).toString());

		// Claim rewards after the end of all blocks. Should accumulated all rewards

		await mineBlocks(100);

		let currBlock = await web3.eth.getBlockNumber();

		let rewardPeriod = currBlock - _startRewardBlock;

		console.log(rewardPeriod)

		let vewtf_totalSupply = await votingescrow.totalSupply();
		
		console.log(`total supply vewtf`, vewtf_totalSupply.div(d18).toString());
		
		let vewtfBalOther1 = await votingescrow.balanceOf(other1.address);
		let vewtfBalOther2 = await votingescrow.balanceOf(other2.address);

		console.log(`vewtf balance other1:${vewtfBalOther1.div(d18).toString()}`);
		console.log(`vewtf balance other2:${vewtfBalOther2.div(d18).toString()}`);

		let wtfBalBeforeOther1 = await this.wtf.balanceOf(other1.address);
		let wtfBalBeforeOther2 = await this.wtf.balanceOf(other2.address);

		let pendingRewardOther1 = await wtfRewards.pendingReward(other1.address);
		let pendingRewardOther2 = await wtfRewards.pendingReward(other2.address);

		console.log(`Pending reward Other1: ${pendingRewardOther1.div(d18).toString()}`);
		console.log(`Pending reward Other 2: ${pendingRewardOther2.div(d18).toString()}`);

		await wtfRewards.connect(other1).claimRewards();
		await wtfRewards.connect(other2).claimRewards();

		let wtfBalAfterOther1 = await this.wtf.balanceOf(other1.address);
		let wtfBalAfterOther2 = await this.wtf.balanceOf(other2.address);

		console.log(wtfBalBeforeOther1.div(d18).toString());
		console.log(wtfBalBeforeOther2.div(d18).toString());
		console.log(wtfBalAfterOther1.div(d18).toString());
		console.log(wtfBalAfterOther2.div(d18).toString());

		expect(wtfBalAfterOther1).to.be.equal(this.WTF_total_reward.div(400).mul(rewardPeriod + 1).div(2));
		expect(wtfBalAfterOther2).to.be.equal(this.WTF_total_reward.div(400).mul(rewardPeriod + 2).div(2));

	})

    it('distributes fee rewards immediately after they are sent on pro rata basis', async function(){
		
		// Deploy staking (general with locking feature)
        const [deployer, other1, other2, other3] = await ethers.getSigners();
		let _currentBlock = await web3.eth.getBlockNumber();
		let _startRewardBlock = _currentBlock + 20;
		let _endRewardBlock = _currentBlock + this.blocks;
		
		let duration = ethers.BigNumber.from(this.year).div(4);
		let PRECISION = ethers.BigNumber.from('1000000000000');

		let votingescrow = await this.VotingEscrow.deploy(this.wtf.address);

		await votingescrow.deployed();

		let wtfRewards = await this.WTFStakingRewards.deploy(votingescrow.address, 
			                                              this.wtf.address,
			                                              this.communityfund.address,
			                                              this.rewardPerBlock, 
			                                              _startRewardBlock,
			                                              _endRewardBlock
			                                               );
		await wtfRewards.deployed();

		// Deploy fee rewards 

		let feeRewards = await this.FeeStakingRewards.deploy(votingescrow.address, this.busd.address);

		await feeRewards.deployed();

		// Send first installment of fee rewards to fee rewards contract
		// Can't send rewards yet because nothing is staked yet

		await this.busd.connect(deployer).approve(feeRewards.address, d18.mul(1000000));

		// Set staking

		await votingescrow.connect(deployer).setStaking(wtfRewards.address, feeRewards.address);
		await this.wtf.connect(deployer).transfer(this.communityfund.address, this.WTF_total_reward);
		await this.communityfund.setAllowance(wtfRewards.address, this.WTF_total_reward);

        // Send WTF to staking contract

		await(await this.wtf.connect(deployer).mint(other1.address, d18.mul(10000))).wait(1);

		await (await this.wtf.connect(other1).approve(votingescrow.address, d18.mul(10000))).wait(1)

		await(await this.wtf.connect(deployer).mint(other2.address, d18.mul(10000))).wait(1);

		await (await this.wtf.connect(other2).approve(votingescrow.address, d18.mul(10000))).wait(1)


		// Create lock

	    await votingescrow.connect(other1).createLock(d18.mul(10000), duration);

		await votingescrow.connect(other2).createLock(d18.mul(10000), duration);

		// Now as staking is automatically enabled we can send fee rewards to the fee rewards contract

		let feeBatch1 = d18.mul('100000');
		let feeBatch2 = d18.mul('50299');

		await feeRewards.connect(deployer).sendRewards(feeBatch1);

		// Check pending reward debt for both users. It should be equal 50000 for each because all rewards sent automatically split pro rata


		let feeRewardPending1 = await feeRewards.pendingRewardOf(other1.address);
		let feeRewardPending2 = await feeRewards.pendingRewardOf(other2.address);

		console.log(feeRewardPending1)

		expect(feeRewardPending1).to.be.equal(feeBatch1.div(2));
		expect(feeRewardPending2).to.be.equal(feeBatch1.div(2));

		// add second reward batch 

		await feeRewards.connect(deployer).sendRewards(feeBatch2);

		feeRewardPending1 = await feeRewards.pendingRewardOf(other1.address);
		feeRewardPending2 = await feeRewards.pendingRewardOf(other2.address);

		expect(feeRewardPending1).to.be.equal(feeBatch1.add(feeBatch2).div(2));
		expect(feeRewardPending2).to.be.equal(feeBatch1.add(feeBatch2).div(2));

		await feeRewards.connect(other1).claimRewards();
		await feeRewards.connect(other2).claimRewards();

		let busdBal1 = await this.busd.balanceOf(other1.address);
		let busdBal2 = await this.busd.balanceOf(other2.address);

		expect(busdBal1).to.be.equal(busdBal2);
		expect(busdBal1.add(busdBal2)).to.be.equal(feeRewardPending1.add(feeRewardPending2));

		expect(busdBal1.add(busdBal2)).to.be.equal(feeBatch2.add(feeBatch1));


	})


})

