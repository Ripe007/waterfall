const { expect } = require('chai')

describe('Vesting', function () {
    const distributedAmount = ethers.BigNumber.from(500000).mul(ethers.BigNumber.from(10).pow(18))
    let snapshotId;
    const epochDuration = 604800
    const epoch1Start = getCurrentUnix() + 1000;

    beforeEach(async function () {
        const [creator, beneficiary1, beneficiary2] = await ethers.getSigners();
        snapshotId = await ethers.provider.send('evm_snapshot')
        this.creator = creator;
        this.beneficiary1 = beneficiary1;
        this.beneficiary2 = beneficiary2;
        const Vesting = await ethers.getContractFactory('VestingLinear');
        const WTFMock = await ethers.getContractFactory('WTFMock');
        this.wtf = await WTFMock.deploy();
        await this.wtf.deployed();
        this.vesting = await Vesting.deploy(this.beneficiary1.address, this.wtf.address, epoch1Start, distributedAmount);
        await this.vesting.deployed();
    })

    afterEach(async function () {
        await ethers.provider.send('evm_revert', [snapshotId])
    })

    describe('General Contract checks', function () {
        it('should be deployed', async function () {
            expect(this.vesting.address).to.not.equal(0);
            expect(this.wtf.address).to.not.equal(0);
        })
        it('should have the owner set as beneficiary1', async function () {
            expect(await this.vesting.owner()).to.be.equal(this.beneficiary1.address);
        })
        it('should have current epoch 0', async function () {
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(0);
            await moveAtEpoch(-10);
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(0);
        })
        it('should have current epoch 1', async function () {
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(0);
            await moveAtEpoch(1);
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(1);
        })
        it('should have last claimed epoch 0', async function () {
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(0);
        })
        it('should have wtf balance 0', async function () {
            expect(await this.vesting.balance()).to.be.equal(0);
        })
        it('should have totalDistributedBalance', async function () {
            expect(await this.vesting.totalVested()).to.be.equal(distributedAmount);
        })
        it('should have claim function callable by anyone', async function () {
            await expect(this.vesting.connect(this.beneficiary2).claim()).to.not.be.revertedWith('Ownable: caller is not the owner');
        })
        it('should have the epoch1', async function () {
            await moveAtEpoch(1)
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(1)
        })
        it('should have the epoch 0', async function () {
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(0)
        })
    })

    describe('Contract Tests', function () {
        it('should deposit some tokens', async function () {
            await this.wtf.connect(this.creator).mint(this.vesting.address, distributedAmount);
            expect(await this.vesting.balance()).to.be.equal(distributedAmount);
        })
        it('should have 0 WTF beneficiar1 balance during 1 epoch', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount); // set tokens
            await moveAtEpoch(1)
            await this.vesting.connect(this.beneficiary1).claim();
            expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(0);
            expect(await this.vesting.balance()).to.be.equal(distributedAmount);
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(0);
        })
        it('should mint for 1 week', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount); // set tokens
            await moveAtEpoch(2)
            await this.vesting.connect(this.beneficiary1).claim();
            expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount.div(104));
            expect(await this.vesting.balance()).to.be.equal(distributedAmount.sub(distributedAmount.div(104)));
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(1);
        })
        it('should mint with the fallback function', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount) // set tokens
            await moveAtEpoch(3)
            await this.beneficiary2.sendTransaction({
                to: this.vesting.address,
            })
            expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount.div(104).mul(2))
            expect(await this.vesting.balance()).to.be.equal(distributedAmount.sub(distributedAmount.div(104).mul(2)))
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(2)
        })
        it('should mint with any user calling claim', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount) // set tokens
            await moveAtEpoch(3)
            await this.vesting.connect(this.beneficiary1).claim()
            expect(await this.wtf.balanceOf(this.beneficiary2.address)).to.be.equal(0)
            expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal((distributedAmount.div(104)).mul(2))
            expect(await this.vesting.balance()).to.be.equal(distributedAmount.sub(distributedAmount.div(104).mul(2)))
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(2)
        })
        it('should mint with any user sending 0 ETH', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount) // set tokens
            await moveAtEpoch(6)
            await this.creator.sendTransaction({
                to: this.vesting.address,
            })
            expect(await this.wtf.balanceOf(this.creator.address)).to.be.equal('1000000000000000000000000');
            expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal((distributedAmount.div(104)).mul(5));
            expect(await this.vesting.balance()).to.be.equal(distributedAmount.sub(distributedAmount.div(104).mul(5)));
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(5)
        })
        it('should mint for 104 week', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount.add(1)) // set tokens
            await moveAtEpoch(104)
            expect(await this.vesting.getCurrentEpoch()).to.be.equal(104);
            await this.vesting.connect(this.beneficiary1).claim();
            expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount.add(1).sub(distributedAmount.add(1).div(104)))
            expect(await this.vesting.balance()).to.be.equal(0)
            expect(await this.vesting.lastClaimedEpoch()).to.be.equal(104)
        })
        it('should emit', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount) // set tokens
            await moveAtEpoch(59)
            expect(this.vesting.connect(this.beneficiary1).claim()).to.emit(this.wtf, 'Transfer')
        })
        it('should not emit', async function () {
            await this.wtf.mint(this.vesting.address, distributedAmount) // set tokens
            await moveAtEpoch(60)
            await this.vesting.connect(this.beneficiary1).claim()
            expect(this.vesting.connect(this.beneficiary1).claim()).to.not.emit(this.wtf, 'Transfer')
        })
    })

    function getCurrentUnix () {
        return Math.floor(Date.now() / 1000)
    }

    async function setNextBlockTimestamp (timestamp) {
        const block = await ethers.provider.send('eth_getBlockByNumber', ['latest', false])
        const currentTs = parseInt(block.timestamp)
        const diff = timestamp - currentTs
        await ethers.provider.send('evm_increaseTime', [diff])
    }

    async function moveAtEpoch (epoch) {
        await setNextBlockTimestamp(epoch1Start + epochDuration * (epoch - 1))
        await ethers.provider.send('evm_mine')
    }

})

