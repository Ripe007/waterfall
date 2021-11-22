const { expect } = require('chai')

describe('Vesting', function () {
    const distributedAmount = ethers.BigNumber.from(500000).mul(ethers.BigNumber.from(10).pow(18))
    let snapshotId;
    const epochDuration = 604800 // seconds in one week
    const grantStart = getCurrentUnix() + 1000;
    const grantDuration = 104; // 104 weeks
    const grantCliff = 24; // 24 weeks 

    beforeEach(async function () {
        const [creator, beneficiary1, multisigFake] = await ethers.getSigners();
        snapshotId = await ethers.provider.send('evm_snapshot')
        this.creator = creator;
        this.beneficiary1 = beneficiary1;
        this.multisigFake = multisigFake;
        const Vesting = await ethers.getContractFactory('VestingVault');
        const WTFMock = await ethers.getContractFactory('WTFMock');
        this.wtf = await WTFMock.deploy();
        await this.wtf.deployed();
        this.vesting = await Vesting.deploy(this.wtf.address, multisigFake.address);
        await this.vesting.deployed();
    })

    afterEach(async function () {
        await ethers.provider.send('evm_revert', [snapshotId])
    })

    describe('General VestingVault checks', function () {
        it('should be deployed', async function () {
            expect(this.vesting.address).to.not.equal(0);
            expect(this.wtf.address).to.not.equal(0);
        })
        it('should have the owner set as multisigFake', async function () {
            expect(await this.vesting.owner()).to.be.equal(this.multisigFake.address);
        })
        it('should disallow other accounts except owner to crerate grants', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await expect(this.vesting.connect(this.beneficiary1).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff)).to.be.revertedWith('Ownable: caller is not the owner');

            
        })

        it('should create a grant for beneficiary1', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await this.vesting.connect(this.multisigFake).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff);

            expect(await this.vesting.totalVestingCount()).to.be.equal(1);

        })
        it('should not vest before the cliff end', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await this.vesting.connect(this.multisigFake).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff);

           let grantID = await this.vesting.totalVestingCount();

           expect(await this.wtf.balanceOf(this.vesting.address)).to.be.equal(distributedAmount);

           await moveAtEpoch(23);

           await expect(this.vesting.claimVestedTokens(grantID)).to.be.revertedWith("Vesting: vested amount is 0");

        })

         it('should distribute amount accumulated during the cliff after the cliff ends', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await this.vesting.connect(this.multisigFake).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff);


           let activeGrants = await this.vesting.getActiveGrants(this.beneficiary1.address);

           let grantId = activeGrants[0].toString();

           await moveAtEpoch(24);

           await this.vesting.claimVestedTokens(grantId);

           expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount.div(grantDuration).mul(grantCliff));

        })

        it('should mint for 30 weeks', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await this.vesting.connect(this.multisigFake).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff);


           let activeGrants = await this.vesting.getActiveGrants(this.beneficiary1.address);

           let grantId = activeGrants[0].toString();

           await moveAtEpoch(30);

           await this.vesting.claimVestedTokens(grantId);

           expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount.div(grantDuration).mul(30));

        })

         it('should mint all amount after grant duration', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await this.vesting.connect(this.multisigFake).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff);


           let activeGrants = await this.vesting.getActiveGrants(this.beneficiary1.address);

           let grantId = activeGrants[0].toString();

           await moveAtEpoch(105);

           await this.vesting.claimVestedTokens(grantId);

           expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount);

        })

        it('should revoke the grant and send all vested amount to beneficiary and non-vested to multisig', async function () {

            await this.wtf.connect(this.creator).mint(this.multisigFake.address, distributedAmount);
            await this.wtf.connect(this.multisigFake).approve(this.vesting.address, distributedAmount);

            await this.vesting.connect(this.multisigFake).addGrant(this.beneficiary1.address, 
                                                                   grantStart, 
                                                                   distributedAmount, 
                                                                   grantDuration,
                                                                   grantCliff);


           let activeGrants = await this.vesting.getActiveGrants(this.beneficiary1.address);

           let grantId = activeGrants[0].toString();

           await moveAtEpoch(28);

           await this.vesting.claimVestedTokens(grantId);

           await moveAtEpoch(30);

           await this.vesting.connect(this.multisigFake).revokeGrant(grantId);

           expect(await this.wtf.balanceOf(this.beneficiary1.address)).to.be.equal(distributedAmount.div(grantDuration).mul(30));

           expect(await this.wtf.balanceOf(this.multisigFake.address)).to.be.equal(distributedAmount.sub(distributedAmount.div(grantDuration).mul(30)));

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
        await setNextBlockTimestamp(grantStart + (epochDuration * epoch));
        await ethers.provider.send('evm_mine')
    }

})

