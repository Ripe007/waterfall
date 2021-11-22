const { expect } = require("chai");

const {
  getBlockTimestamp,
  setBlockTime,
  increaseTime,
  mineBlocks,
  sync: { getCurrentTimestamp },
} = require("./helpers.js");

describe("extra token swap to busd", function () {
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
    this.CreamFarm = await ethers.getContractFactory("CreamFakeSwap");
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

  it("Extra token swap(venus and cream(alpaca))full exit after 1 cycle, farm 10% gain", async function () {
    const [deployer, zeus, other1, other2, other3] = await ethers.getSigners();
    this.creamFarm = await this.CreamFarm.deploy(
      this.busd.address,
      this.alpacaToken.address,
      10
    );
    this.venusFarm = await this.CreamFarm.deploy(
      this.busd.address,
      this.alpacaToken.address,
      10
    );
    this.alpacaFarm = await this.AlpacaFarm.deploy(this.busd.address, 10);
    let Uniswap = await ethers.getContractFactory("UniswapV2Router02");
    let uniswap = await Uniswap.deploy(this.busd.address);
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
    await uniswap.deployed();

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

    await camp.setPancakeRouter(uniswap.address);

    let uniswapAddress = await camp.PancakeRouter();

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

    await (
      await this.alpacaToken.mint(this.creamFarm.address, D18.mul(10000))
    ).wait(1);
    await (
      await this.alpacaToken.mint(this.venusFarm.address, D18.mul(10000))
    ).wait(1);
    await (await this.busd.mint(uniswapAddress, D18.mul(10000))).wait(1);
    //alpacaToken
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
    await (await camp.connect(other1).exit(0)).wait(1);
    await (await camp.connect(other2).exit(2)).wait(1);
    let other1Balance = await this.busd.balanceOf(other1.address);
    let other2Balance = await this.busd.balanceOf(other2.address);
    // fee
    let operator_balance = await camp.connect(deployer).producedFee();

    expect(
      D18.mul(3960).sub(other1Balance.add(other2Balance).add(operator_balance))
    ).lte(10000);
  });
});
