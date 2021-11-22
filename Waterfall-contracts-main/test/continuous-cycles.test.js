const { expect } = require("chai");

const {
  getBlockTimestamp,
  setBlockTime,
  increaseTime,
  mineBlocks,
  sync: { getCurrentTimestamp },
} = require("./helpers.js");

describe("continueous cycles", function () {
  const D18 = ethers.BigNumber.from("1000000000000000000");
  const D8 = ethers.BigNumber.from("100000000");
  const D6 = ethers.BigNumber.from("1000000");
  const SUPPLY = D18.mul(1000000);
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
    this.TokenFactory = Token;

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

  it("next cycle too early", async function () {
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

    await camp.nextCycle();
    await increaseTime(86400 * 6);
    await mineBlocks(1);
    await expect(camp.nextCycle()).to.be.revertedWith("new cycle too early");
  });

  it("full exit after 1 cycle, farm 10% gains", async function () {
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
      D18.mul(1000000),
      86400 * 7,
      86400 * 1
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

    await (await this.busd.mint(other1.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other1).join(0, D18.mul(1000))).wait(1);

    await (await this.busd.mint(other2.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();

    await expect(
      camp.connect(other1).join(0, D18.mul(1000))
    ).to.be.revertedWith("missing the join window");

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    await (await camp.connect(other2).applyForExit(2)).wait(1);

    await increaseTime(86400 * 7);
    await mineBlocks(1);
    // await expect(
    // camp.connect(other1).join(0, D18.mul(1000))
    // ).to.be.revertedWith("missing the join window");

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.nextCycle();

    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(2)).wait(1);
    let other1Balance = await this.busd.balanceOf(other1.address);
    let other2Balance = await this.busd.balanceOf(other2.address);

    // fee
    let operator_balance = await camp.connect(deployer).producedFee();
    expect(
      D18.mul(2200).sub(other1Balance.add(other2Balance).add(operator_balance))
    ).lte(10000);
  });

  it("full exit after 1 cycle, farm 10% loss", async function () {
    const [deployer, zeus, other1, other2, other3] = await ethers.getSigners();

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
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();

    await expect(
      camp.connect(other1).join(0, D18.mul(1000))
    ).to.be.revertedWith("missing the join window");

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    await (await camp.connect(other2).applyForExit(2)).wait(1);

    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.nextCycle();

    let other1BalanceBeforeExit = await this.busd.balanceOf(other1.address);
    let other2BalanceBeforeExit = await this.busd.balanceOf(other2.address);
    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(2)).wait(1);
    let other1BalanceAfterExit = await this.busd.balanceOf(other1.address);
    let other2BalanceAfterExit = await this.busd.balanceOf(other2.address);

    expect(other1BalanceBeforeExit).lt(other1BalanceAfterExit);
    expect(other2BalanceBeforeExit).lt(other2BalanceAfterExit);
  });

  it("2 cycles, each cycle 10% gains", async function () {
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
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();

    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await camp.nextCycle();

    await increaseTime(86400 * 7);
    await mineBlocks(1);
    await (await camp.connect(other1).applyForExit(0)).wait(1);
    await (await camp.connect(other2).applyForExit(2)).wait(1);

    await camp.nextCycle();

    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(2)).wait(1);
    let producedFee = await camp.connect(deployer).producedFee();
    let other1Balance = await this.busd.balanceOf(other1.address);
    let other2Balance = await this.busd.balanceOf(other2.address);

    let oneCycle = ethers.BigNumber.from("2197273202771499238000");
    oneCycle = oneCycle.add(oneCycle.mul(10).div(100));
    expect(oneCycle.sub(other1Balance.add(other2Balance).add(producedFee))).lte(
      10000
    );
  });

  it("farm 50% loss, junior should be busted, cycle should be jammed", async function () {
    const [deployer, zeus, other1, other2, other3] = await ethers.getSigners();

    this.creamFarm = await this.CreamFarmLoss.deploy(this.busd.address, 50);
    this.venusFarm = await this.CreamFarmLoss.deploy(this.busd.address, 50);
    this.alpacaFarm = await this.AlpacaFarmLoss.deploy(this.busd.address, 50);
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
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();

    let [totalRaised, totalCapital] = await camp.tranchesTotal();
    expect(totalRaised).eq(D18.mul(2000));
    expect(totalCapital).eq(D18.mul(2000));

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    // await (await camp.connect(other2).applyForExit(2)).wait(1);

    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));
    await camp.nextCycle();
    await (await camp.connect(other1).exit(0)).wait(1);
    // await (await camp.connect(other2).exit(2)).wait(1);

    let other1Balance = await this.busd.balanceOf(other1.address);
    // expect(D18.mul(1000).sub(other1Balance)).lt(D18.div(1000));

    // expect(other1Balance.add(other2Balance)).eq(D18.mul(1800));
    [totalRaised, totalCapital] = await camp.tranchesTotal();
    expect(totalRaised).eq(0);
    expect(totalCapital).eq(0);

    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await expect(camp.nextCycle()).to.be.revertedWith("empty pocket");

    await expect(camp.connect(other2).applyForExit(2)).to.be.revertedWith(
      "not an investor"
    );
  });

  it("terminate right after next cycle should work", async function () {
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
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();
    expect(await camp.launched()).eq(true);

    let [totalRaised, totalCapital] = await camp.tranchesTotal();
    expect(totalRaised).eq(D18.mul(2000));
    expect(totalCapital).eq(D18.mul(2000));

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    // await (await camp.connect(other2).applyForExit(2)).wait(1);

    await increaseTime(86400 * 1);
    await mineBlocks(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.terminate();

    expect(await camp.terminated()).eq(true);
    expect(await camp.launched()).eq(false);

    let tranche0 = await camp.tranches(0);
    let tranche0Option = this.TokenFactory.attach(tranche0.token);
    let tranche0Reward = this.TokenFactory.attach(tranche0.reward);
    let balanceBefore = await tranche0Reward.balanceOf(other1.address);
    let currencyBefore = await this.busd.balanceOf(other1.address);
    await camp.connect(other1).exitAfterTerminated(0);
    let balanceAfter = await tranche0Option.balanceOf(other1.address);
    let currencyAfter = await this.busd.balanceOf(other1.address);
    expect(balanceBefore.sub(balanceAfter)).gt(0);
    expect(balanceAfter).eq(0);
    expect(currencyAfter.sub(currencyBefore)).gt(0);

    let tranche2 = await camp.tranches(2);
    let tranche2Option = this.TokenFactory.attach(tranche2.token);
    let tranche2Reward = this.TokenFactory.attach(tranche2.reward);
    balanceBefore = await tranche2Reward.balanceOf(other2.address);
    currencyBefore = await this.busd.balanceOf(other2.address);
    await camp.connect(other2).exitAfterTerminated(2);
    balanceAfter = await tranche2Option.balanceOf(other2.address);
    currencyAfter = await this.busd.balanceOf(other2.address);
    expect(balanceBefore.sub(balanceAfter)).gt(0);
    expect(balanceAfter).eq(0);
    expect(currencyAfter.sub(currencyBefore)).gt(0);

    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await expect(
      camp.connect(other2).join(2, D18.mul(1000))
    ).to.be.revertedWith("already terminated");

    await expect(
      camp.connect(other2).exitAfterTerminated(2)
    ).to.be.revertedWith("not in this tranche");
  });

  it("termination returns new joins should work", async function () {
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
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await (await this.busd.mint(other3.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other3).approve(camp.address, D18.mul(1000))
    ).wait(1);

    await camp.nextCycle();
    expect(await camp.launched()).eq(true);

    let [totalRaised, totalCapital] = await camp.tranchesTotal();
    expect(totalRaised).eq(D18.mul(2000));
    expect(totalCapital).eq(D18.mul(2000));

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    // await (await camp.connect(other2).applyForExit(2)).wait(1);

    await increaseTime(86400 * 6);
    await mineBlocks(1);

    await (await camp.connect(other3).join(1, D18.mul(1000))).wait(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.terminate();

    expect(await camp.terminated()).eq(true);
    expect(await camp.launched()).eq(false);

    let tranche1 = await camp.tranches(1);
    let tranche1Option = this.TokenFactory.attach(tranche1.token);
    let tranche1Reward = this.TokenFactory.attach(tranche1.reward);
    let balanceBefore = await tranche1Reward.balanceOf(other3.address);
    let currencyBefore = await this.busd.balanceOf(other3.address);
    await camp.connect(other3).exitAfterTerminated(1);
    let balanceAfter = await tranche1Option.balanceOf(other3.address);
    let currencyAfter = await this.busd.balanceOf(other3.address);
    expect(balanceBefore.sub(balanceAfter)).gt(0);
    expect(balanceAfter).eq(0);
    expect(currencyAfter.sub(currencyBefore)).eq(D18.mul(1000));

    await expect(
      camp.connect(other3).exitAfterTerminated(1)
    ).to.be.revertedWith("not in this tranche");
  });

  it("exits cycles before termination should work", async function () {
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
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    await camp.nextCycle();
    expect(await camp.launched()).eq(true);

    await (await camp.connect(other1).applyForExit(0)).wait(1);

    await increaseTime(86400 * 7);
    await mineBlocks(1);

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.nextCycle();
    await increaseTime(86400 * 2);
    await mineBlocks(1);

    await camp.terminate();

    expect(await camp.terminated()).eq(true);
    expect(await camp.launched()).eq(false);

    await expect(
      camp.connect(other1).exitAfterTerminated(0)
    ).to.be.revertedWith("account has unfinished exit apply");

    let tranche0 = await camp.tranches(0);
    let tranche0Option = this.TokenFactory.attach(tranche0.token);
    let tranche0Reward = this.TokenFactory.attach(tranche0.reward);
    let balanceBefore = await tranche0Reward.balanceOf(other1.address);
    let currencyBefore = await this.busd.balanceOf(other1.address);
    await camp.connect(other1).exit(0);
    let balanceAfter = await tranche0Option.balanceOf(other1.address);
    let currencyAfter = await this.busd.balanceOf(other1.address);
    expect(balanceBefore.sub(balanceAfter)).gt(0);
    expect(balanceAfter).eq(0);
    expect(currencyAfter.sub(currencyBefore)).gt(0);

    await expect(camp.connect(other1).exit(0)).to.be.revertedWith(
      "no exit applies"
    );
    await expect(
      camp.connect(other1).exitAfterTerminated(0)
    ).to.be.revertedWith("not in this tranche");
  });

  it("full exit after 1 cycle, farm 10% gains, reward pool", async function () {
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
      D18.mul(1000000),
      86400 * 7,
      86400 * 1
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

    await (await this.busd.mint(other1.address, D18.mul(2000))).wait(1);
    await (
      await this.busd.connect(other1).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other1).join(0, D18.mul(1000))).wait(1);

    await (await this.busd.mint(other2.address, D18.mul(1000))).wait(1);
    await (
      await this.busd.connect(other2).approve(camp.address, D18.mul(1000))
    ).wait(1);
    await (await camp.connect(other2).join(2, D18.mul(1000))).wait(1);

    let [trancheToken, trancheReward] =
      await camp.getTrancheTokenRewardAddresses(2);
    const Reward = await ethers.getContractFactory("Reward");
    let trancheRewardPool = Reward.attach(trancheReward);
    await (await this.WTF.transfer(trancheReward, D18.mul(1000))).wait(1);
    await trancheRewardPool.newPeriod(10, D18.mul(1000));

    await camp.nextCycle();

    await expect(
      camp.connect(other1).join(0, D18.mul(1000))
    ).to.be.revertedWith("missing the join window");

    await (await camp.connect(other1).applyForExit(0)).wait(1);
    await (await camp.connect(other2).applyForExit(2)).wait(1);

    await increaseTime(86400 * 7);
    await mineBlocks(1);
    // await expect(
    // camp.connect(other1).join(0, D18.mul(1000))
    // ).to.be.revertedWith("missing the join window");

    await this.busd.mint(this.creamFarm.address, D18.mul(10000000));
    await this.busd.mint(this.venusFarm.address, D18.mul(10000000));
    await this.busd.mint(this.alpacaFarm.address, D18.mul(10000000));

    await camp.nextCycle();

    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(2)).wait(1);
    let other1Balance = await this.busd.balanceOf(other1.address);
    let other2Balance = await this.busd.balanceOf(other2.address);

    expect(await this.WTF.balanceOf(other2.address)).equal(D18.mul(1000));
    expect(await this.WTF.balanceOf(other1.address)).equal(0);

    // fee
    let operator_balance = await camp.connect(deployer).producedFee();
    expect(
      D18.mul(2200).sub(other1Balance.add(other2Balance).add(operator_balance))
    ).lte(10000);
  });
});
