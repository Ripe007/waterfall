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
		snapshotId = await ethers.provider.send('evm_snapshot')
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
		this.WTF_total_reward = d18.mul(950000);
		this.blocks = 1000;
		this.rewardPerBlock = this.WTF_total_reward.div(this.blocks);
	
	})

	afterEach(async function () {
        await ethers.provider.send('evm_revert', [snapshotId])
    })

	it('Check delegation', async function(){
		// Deploy staking (general with locking feature)
        const [deployer, other1, other2, other3] = await ethers.getSigners();
		let _currentBlock = await web3.eth.getBlockNumber();
		let _endRewardBlock = _currentBlock + this.blocks;
		
		let duration = ethers.BigNumber.from(this.year).div(4);
		let PRECISION = ethers.BigNumber.from('1000000000000');

		let votingescrow = await this.VotingEscrow.deploy(this.wtf.address);

		await votingescrow.deployed();


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

		await votingescrow.connect(deployer).setStaking(wtfRewards.address, feeRewards.address);

		await this.wtf.connect(deployer).transfer(this.communityfund.address, this.WTF_total_reward);

		await this.communityfund.setAllowance(wtfRewards.address, this.WTF_total_reward);

		// Lock some WTF to Vesting Escrow

		// First mint to Other1

		await(await this.wtf.connect(deployer).mint(other1.address, d18.mul(10000))).wait(1);

		await (await this.wtf.connect(other1).approve(votingescrow.address, d18.mul(10000))).wait(1)

		await(await this.wtf.connect(deployer).mint(other2.address, d18.mul(10000))).wait(1);

		await (await this.wtf.connect(other2).approve(votingescrow.address, d18.mul(10000))).wait(1)


		// Create lock

	    await votingescrow.connect(other1).createLock(d18.mul(10000), duration);

		await votingescrow.connect(other2).createLock(d18.mul(10000), duration);

		// Current votes for other1 should equal 0 because there was no self-delegation

		expect(await votingescrow.getCurrentVotes(other1.address)).to.be.equal(0);

		// Self-delegate

		await votingescrow.connect(other1).delegate(other1.address);

		// Votes should now equal the vewtf balance of other1

		expect(await votingescrow.getCurrentVotes(other1.address)).to.be.equal(await votingescrow.balanceOf(other1.address));

		// Delegate votes to other2

		await votingescrow.connect(other1).delegate(other2.address);

		// Votes are zero for other1 again

		expect(await votingescrow.getCurrentVotes(other1.address)).to.be.equal(0);

		// Votes of other2 now equal other vewtf balance

		expect(await votingescrow.getCurrentVotes(other2.address)).to.be.equal(await votingescrow.balanceOf(other1.address));

		// Self-delegate for other2

		await votingescrow.connect(other2).delegate(other2.address);

		// votes for other2 = vewtf balance * 2 because of other1 delegation

		let other2bal = await votingescrow.balanceOf(other2.address);

		expect(await votingescrow.getCurrentVotes(other2.address)).to.be.equal(other2bal.mul(2));

		// record current block

		let blockRecord = await web3.eth.getBlockNumber();

		// mine 100 blocks

		await mineBlocks(100);

		// now check the past votes of other2 at the block number recorded. They should be equal to other2bal * 2

		expect (await votingescrow.getPriorVotes(other2.address, blockRecord)).to.be.equal(other2bal.mul(2));

		// they should be equal to other2bal at the block before blockrecord

		expect (await votingescrow.getPriorVotes(other2.address, blockRecord - 1)).to.be.equal(other2bal);

		// Now let's increase lock time for other2 to get more vewtf

		let lock = await votingescrow.getLockData(other2.address);

		other2bal = await votingescrow.balanceOf(other2.address);

		await votingescrow.connect(other2).increaseTimeAndAmount(0, lock.startTimestamp.add(duration.mul(3))); // half a year

		// expect other2 vewtf balance to be half of the deposited wtf

		expect (await votingescrow.balanceOf(other2.address)).to.be.equal(other2bal.mul(3));

		// Since other2 is self-delegated the balance should move to voting power

		let other1bal = await votingescrow.balanceOf(other1.address); 

		other2bal = await votingescrow.balanceOf(other2.address); 

		expect (await votingescrow.getCurrentVotes(other2.address)).to.be.equal(other2bal.add(other1bal));

		// Get to lock expiration for user2

		await increaseTime(Number(lock.expiryTimestamp));

		// unlock his tokens

		await votingescrow.connect(other2).unlock();

		// after unlock, vewtf is burned, other2 self-delegates, so we expect the voting power to decrease by the total vewtf balance burned.
		// the total voting power is now equal to vewtf delegated by other1 => other1 balance

		expect (await votingescrow.getCurrentVotes(other2.address)).to.be.equal(other1bal);


	})


})

