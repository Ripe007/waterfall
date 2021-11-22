const { expect } = require("chai");

const {
  getBlockTimestamp,
  setBlockTime,
  increaseTime,
  minnBlocks,
  sync: { getCurrentTimestamp },
  mineBlocks,
} = require("./helpers.js");

describe("Many people join the same of trache test", function () {
  const D18 = ethers.BigNumber.from("1000000000000000000");
  const D8 = ethers.BigNumber.from("100000000");
  const D6 = ethers.BigNumber.from("1000000");
  const SUPPY = D18.mul(1000000);
  const USDT_TOTAL = ethers.BigNumber.from("100000000000").mul(D8);
  const INIT_SUPPLY = D18.mul(1000000);
  const sell_price = ethers.BigNumber.from(65).mul(D6).div(10);
  const recycle_price = ethers.BigNumber.from(5).mul(D6);

  beforeEach(async function () {
    const [deployer, zeus, other1, other2, other3] = await ethers.getSigners();
    const Token = await ethers.getContractFactory("Token");
    this.CreamFarm = await ethers.getContractFactory("CreamFake");
    this.AlpacaFarm = await ethers.getContractFactory("AlpacaFake");
    this.CreamFarmLoss = await ethers.getContractFactory("CreamFakeLoss");
    this.AlpacaFarmLoss = await ethers.getContractFactory("AlpacaFakeLoss");

    const WTF = await ethers.getContractFactory("WTF");
    const RewardFactory = await ethers.getContractFactory("RewardFactory");
    const TrancheTokenFactory = await ethers.getContractFactory(
      "TrancheTokenFactory"
    );

    this.busd = await Token.deploy(
      "BUSD",
      "BUSD",
      18,
      D18.mul(1000000000),
      D18.mul(1000000000)
    );
    await this.busd.deployed();

    this.alpacaToken = await Token.deploy(
      "alpaca",
      "alpaca",
      18,
      D18.mul(10000000),
      D18.mul(10000000)
    );
    this.venusToken = await Token.deploy(
      "venus",
      "venus",
      18,
      D18.mul(10000000),
      D18.mul(10000000)
    );

    this.WTF = await WTF.deploy(
      "WTF",
      "WTF",
      deployer.address,
      D18.mul(1000000000)
    );
    this.RewardFactory = await RewardFactory.deploy();
    this.TrancheTokenFactory = await TrancheTokenFactory.deploy();
  });

  it("Two join same tranche2 (full exit after 1 cycle. farm 10% gains)", async function () {
    const [deployer, zeus, other1, other2, other3, other4, other5, other6] =
      await ethers.getSigners();
    this.creamFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.venusFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.alpacaFarm = await this.AlpacaFarm.deploy(this.busd.address, 10);

    const Camp = await ethers.getContractFactory("CampaignContinuousCycles");
    let camp = await Camp.deploy(
      deployer.address,
      0,
      "",
      this.WTF.address,
      this.RewardFactory.address,
      this.TrancheTokenFactory.address,
      this.busd.address,
      D18.mul(1000000),
      86400 * 7,
      86400 * 3
    );

    await camp.deployed();

    await camp.setTranches([
      { apy: 10000, fee: 33, percentage: 40000 },
      { apy: 30000, fee: 50, percentage: 30000 },
      { apy: 0, fee: 200, percentage: 30000 },
    ]);
    await camp.setAlpacaToken(this.alpacaToken.address);
    await camp.setVenusToken(this.venusToken.address);
    await camp.setFarms(
      this.creamFarm.address,
      40000,
      this.venusFarm.address,
      40000,
      this.alpacaFarm.address,
      20000
    );
    //set tranche fee
    await camp.connect(deployer).setTranchesFee(0, 33);
    await camp.connect(deployer).setTranchesFee(1, 50);
    await camp.connect(deployer).setTranchesFee(2, 200);

    // person 1-4 :
    await (await this.busd.mint(other1.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1000))
    ).wait(1);

    await (await camp.connect(other1).join(0, D18.mul(1000))).wait(1);
    // person 4 :
    await (await this.busd.mint(other4.address, D18.mul(4000000))).wait(1);
    await (
      await this.busd.connect(other4).approve(camp.address, D18.mul(4000000))
    ).wait(1);
    /**
     * test join() not enough quota
     * */
    await expect(
      camp.connect(other4).join(0, D18.mul(4000000))
    ).to.be.revertedWith("not enough quota");

    //person 2-5 :
    await (await this.busd.mint(other2.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(1, D18.mul(1000))).wait(1);

    await (await this.busd.mint(other5.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other5).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other5).join(1, D18.mul(1000))).wait(1);

    //person 3:
    await (await this.busd.mint(other3.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other3).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other3).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    // applyForExit
    await (await camp.connect(other2).applyForExit(1)).wait(1);

    await (await camp.connect(other5).applyForExit(1)).wait(1);
    await (await camp.connect(other3).applyForExit(2)).wait(1);
    /**
     * test: user not in a tranche, whether can call applyForExit()
     * */
    await expect(camp.connect(other3).applyForExit(1)).to.be.revertedWith(
      "not an investor"
    );
    await increaseTime(86400 * 7);
    await mineBlocks(1);
    /*
     * now cycle
     */
    let cycleNumber = await camp.cycle();

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));
    /*
     * Test: early exit
     * */

    await expect(camp.connect(other1).exit(0)).to.be.revertedWith(
      "allow to exit from next cycle"
    );

    await camp.nextCycle();
    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(1)).wait(1);
    await (await camp.connect(other5).exit(1)).wait(1);
    /*
     * Test: applyForExit id different exit id
     */
    await expect(camp.connect(other3).exit(1)).to.be.revertedWith(
      "no exit applies"
    );

    await (await camp.connect(other3).exit(2)).wait(1);
    let other2Balance = await this.busd.balanceOf(other2.address);
    let other1Balance = await this.busd.balanceOf(other1.address);
    let other5Balance = await this.busd.balanceOf(other5.address);
    let other3Balance = await this.busd.balanceOf(other3.address);
    let operatorBalance = await camp.producedFee();
    expect(
      D18.mul(4400).sub(
        other1Balance
          .add(other2Balance)
          .add(other3Balance)
          .add(other5Balance)
          .add(operatorBalance)
      )
    ).lte(10000);
    expect(other2Balance).eq(other5Balance);
  });

  it("Two join same tranche2 (users full exit afer 1 cycle, farm 10% loss)", async function () {
    const [deployer, zeus, other1, other2, other3, other4, other5, other6] =
      await ethers.getSigners();

    this.creamFarm = await this.CreamFarmLoss.deploy(this.busd.address, 10);
    this.venusFarm = await this.CreamFarmLoss.deploy(this.busd.address, 10);
    this.alpacaFarm = await this.AlpacaFarmLoss.deploy(this.busd.address, 10);

    const Camp = await ethers.getContractFactory("CampaignContinuousCycles");
    let camp = await Camp.deploy(
      deployer.address,
      0,
      "",
      this.WTF.address,
      this.RewardFactory.address,
      this.TrancheTokenFactory.address,
      this.busd.address,
      D18.mul(1000000),
      86400 * 7,
      86400 * 3
    );
    await camp.deployed();

    await camp.setTranches([
      { apy: 10000, fee: 33, percentage: 40000 },
      { apy: 30000, fee: 50, percentage: 30000 },
      { apy: 0, fee: 200, percentage: 30000 },
    ]);
    await camp.setAlpacaToken(this.alpacaToken.address);
    await camp.setVenusToken(this.venusToken.address);
    await camp.setFarms(
      this.creamFarm.address,
      40000,
      this.venusFarm.address,
      40000,
      this.alpacaFarm.address,
      20000
    );

    await (await this.busd.mint(other1.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other1).join(0, D18.mul(1000))).wait(1);

    await (await this.busd.mint(other2.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(1, D18.mul(1000))).wait(1);
    await (await this.busd.mint(other3.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other3).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other3).join(1, D18.mul(1000))).wait(1);
    await (await this.busd.mint(other4.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other4).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other4).join(2, D18.mul(1000))).wait(1);
    await camp.nextCycle();

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    await (await camp.connect(other2).applyForExit(1)).wait(1);
    await (await camp.connect(other3).applyForExit(1)).wait(1);
    await (await camp.connect(other4).applyForExit(2)).wait(1);
    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));

    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));

    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.nextCycle();
    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(1)).wait(1);
    await (await camp.connect(other3).exit(1)).wait(1);
    await (await camp.connect(other4).exit(2)).wait(1);

    let other1Balance = await this.busd.balanceOf(other1.address);

    let other2Balance = await this.busd.balanceOf(other2.address);

    let other3Balance = await this.busd.balanceOf(other3.address);

    let other4Balance = await this.busd.balanceOf(other4.address);

    let operatorBalance = await camp.producedFee();
    expect(
      D18.mul(3600).sub(
        other1Balance
          .add(other2Balance)
          .add(other3Balance)
          .add(other4Balance)
          .add(operatorBalance)
      )
    ).lte(10000);

    expect(other2Balance).eq(other3Balance);
  });

  it("Two join same tranche2 (users 2 cycles, each cycle 10% gains)", async function () {
    const [deployer, zeus, other1, other2, other3, other4, other5, other6] =
      await ethers.getSigners();

    this.creamFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.venusFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.alpacaFarm = await this.AlpacaFarm.deploy(this.busd.address, 10);

    const Camp = await ethers.getContractFactory("CampaignContinuousCycles");
    let camp = await Camp.deploy(
      deployer.address,
      0,
      "",
      this.WTF.address,
      this.RewardFactory.address,
      this.TrancheTokenFactory.address,
      this.busd.address,
      D18.mul(1000000),
      86400 * 7,
      86400 * 3
    );
    await camp.deployed();

    await camp.setTranches([
      { apy: 10000, fee: 33, percentage: 40000 },
      { apy: 30000, fee: 50, percentage: 30000 },
      { apy: 0, fee: 200, percentage: 30000 },
    ]);
    await camp.setAlpacaToken(this.alpacaToken.address);
    await camp.setVenusToken(this.venusToken.address);
    await camp.setFarms(
      this.creamFarm.address,
      40000,
      this.venusFarm.address,
      40000,
      this.alpacaFarm.address,
      20000
    );

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await (await this.busd.mint(other1.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other1).join(0, D18.mul(1000))).wait(1);

    await (await this.busd.mint(other2.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(1, D18.mul(1000))).wait(1);
    await (await this.busd.mint(other3.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other3).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other3).join(1, D18.mul(1000))).wait(1);
    await (await this.busd.mint(other4.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other4).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other4).join(2, D18.mul(1000))).wait(1);
    let contractBalance = await this.busd.balanceOf(camp.address);
    await camp.nextCycle();
    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await camp.nextCycle();
    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    await (await camp.connect(other2).applyForExit(1)).wait(1);
    await (await camp.connect(other3).applyForExit(1)).wait(1);
    await (await camp.connect(other4).applyForExit(2)).wait(1);
    await camp.nextCycle();
    contractBalance = await this.busd.balanceOf(camp.address);
    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(1)).wait(1);
    await (await camp.connect(other3).exit(1)).wait(1);
    await (await camp.connect(other4).exit(2)).wait(1);
    let other1Balance = await this.busd.balanceOf(other1.address);

    let other2Balance = await this.busd.balanceOf(other2.address);

    let other3Balance = await this.busd.balanceOf(other3.address);

    let other4Balance = await this.busd.balanceOf(other4.address);
    let operatorBalance = await camp.producedFee();

    let totalBalance = other1Balance
      .add(other2Balance)
      .add(other3Balance)
      .add(other4Balance)
      .add(operatorBalance);
    contractBalance = await this.busd.balanceOf(camp.address);
    let oneCycle = ethers.BigNumber.from("4395890463284373413000");
    oneCycle = oneCycle.mul(110).div(100);

    expect(oneCycle.sub(totalBalance)).lte(10000);
  });

  it("Two join same tranche2 (overflow target full exit after 1 cycle, farm 10% gains)", async function () {
    const [deployer, zeus, other1, other2, other3] = await ethers.getSigners();

    this.creamFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.venusFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.alpacaFarm = await this.AlpacaFarm.deploy(this.busd.address, 10);

    const Camp = await ethers.getContractFactory("CampaignContinuousCycles");
    let camp = await Camp.deploy(
      deployer.address,
      0,
      "",
      this.WTF.address,
      this.RewardFactory.address,
      this.TrancheTokenFactory.address,
      this.busd.address,
      D18.mul(2000), // target ovorflow test
      86400 * 7,
      86400 * 3
    );
    await camp.deployed();

    await camp.setTranches([
      { apy: 10000, fee: 33, percentage: 40000 },
      { apy: 30000, fee: 50, percentage: 30000 },
      { apy: 0, fee: 200, percentage: 30000 },
    ]);
    await camp.setAlpacaToken(this.alpacaToken.address);
    await camp.setVenusToken(this.venusToken.address);
    await camp.setFarms(
      this.creamFarm.address,
      40000,
      this.venusFarm.address,
      40000,
      this.alpacaFarm.address,
      20000
    );

    // set campaign target
    await camp.setCampaignTarget(D18.mul(3000));

    await (await this.busd.mint(other1.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1200))
    ).wait(1);
    await (await camp.connect(other1).join(0, D18.mul(1200))).wait(1);

    await (await this.busd.mint(other2.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(2, D18.mul(900))).wait(1);
    await (await this.busd.mint(other3.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other3).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other3).join(0, D18.mul(900))).wait(1);

    await camp.nextCycle();
    await increaseTime(86400 * 7);
    await mineBlocks(1);
    // await camp.connect(deployer).terminate().wait(1);
    // await (await camp.connect(other3).join(0, D18.mul(600))).wait(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));
    await camp.nextCycle();

    let target = await camp.target();

    // fee
    await increaseTime(86400 * 7);
    await mineBlocks(1);
    await camp.nextCycle();
    let b = await camp.tranchesTotal();
    expect(target).lte(b[1]);
  });

  it("Whether apply exit (full exit after 1 cycle, farm 10% gains)", async function () {
    const [deployer, zeus, other1, other2, other3] = await ethers.getSigners();

    this.creamFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.venusFarm = await this.CreamFarm.deploy(this.busd.address, 10);
    this.alpacaFarm = await this.AlpacaFarm.deploy(this.busd.address, 10);

    const Camp = await ethers.getContractFactory("CampaignContinuousCycles");
    let camp = await Camp.deploy(
      deployer.address,
      0,
      "",
      this.WTF.address,
      this.RewardFactory.address,
      this.TrancheTokenFactory.address,
      this.busd.address,
      D18.mul(2000), // target ovorflow test
      86400 * 7,
      86400 * 3
    );
    await camp.deployed();

    await camp.setTranches([
      { apy: 10000, fee: 33, percentage: 40000 },
      { apy: 30000, fee: 50, percentage: 30000 },
      { apy: 0, fee: 200, percentage: 30000 },
    ]);
    await camp.setAlpacaToken(this.alpacaToken.address);
    await camp.setVenusToken(this.venusToken.address);
    await camp.setFarms(
      this.creamFarm.address,
      40000,
      this.venusFarm.address,
      40000,
      this.alpacaFarm.address,
      20000
    );

    // set campaign target
    await camp.setCampaignTarget(D18.mul(3000));

    await (await this.busd.mint(other1.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1200))
    ).wait(1);
    await (await camp.connect(other1).join(0, D18.mul(1200))).wait(1);

    await (await this.busd.mint(other2.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(2, D18.mul(900))).wait(1);
    await (await this.busd.mint(other3.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other3).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other3).join(0, D18.mul(900))).wait(1);

    await camp.nextCycle();
    await increaseTime(86400 * 7);
    await mineBlocks(1);
    // await camp.connect(deployer).terminate().wait(1);
    // await (await camp.connect(other3).join(0, D18.mul(600))).wait(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));
    await camp.nextCycle();
    let isApplyExitOther1 = await camp.connect(other1).isApplyExit(0);
    // let is_app = isApplyExit.wait(1);
    expect(isApplyExitOther1.toString()).eq("false");

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    isApplyExitOther1 = await camp.connect(other1).isApplyExit(0);
    expect(isApplyExitOther1.toString()).eq("true");
  });
});
