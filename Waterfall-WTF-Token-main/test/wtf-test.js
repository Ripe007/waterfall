const { expect } = require("chai");
const { BigNumber } = require("ethers");
const addresses = require("../addresses.js");
const d18 = BigNumber.from("10").pow("18");
const initialSupply = BigNumber.from("10").pow("18").mul("100000000");

describe("WTF token", function() {
  it("Should mint initial supply to deployer", async function() {
  	const [deployer, user1] = await ethers.getSigners();
    const WTF = await ethers.getContractFactory("WTF");
    const wtf = await WTF.deploy(initialSupply, addresses.teamMultisig );
    await wtf.deployed();
    expect(await wtf.balanceOf(deployer.address)).to.equal(initialSupply);
  });
  it("Should transfer ownership to team multisig", async function() {
  	const [deployer,user1] = await ethers.getSigners();
    const WTF = await ethers.getContractFactory("WTF");
    const wtf = await WTF.deploy(initialSupply, addresses.teamMultisig );
    await wtf.deployed();
    expect(await wtf.owner()).to.equal(addresses.teamMultisig);
  });
  it("Should not mint if the caller is not the contract owner", async function() {
    const [deployer,user1,user2] = await ethers.getSigners();
    const WTF = await ethers.getContractFactory("WTF");
    const wtf = await WTF.deploy(initialSupply, deployer.address );
    await wtf.deployed();
    expect(await wtf.owner()).to.equal(deployer.address);
    await expect(wtf.connect(user1).mint(user2.address, d18.mul("1000"))).to.be.revertedWith("Ownable: caller is not the owner");
  });
  it("Should mint if the call is made by the owner", async function() {
    const [deployer,user1,user2] = await ethers.getSigners();
    const WTF = await ethers.getContractFactory("WTF");
    const wtf = await WTF.deploy(initialSupply, deployer.address );
    await wtf.deployed();
    expect(await wtf.owner()).to.equal(deployer.address);
    await wtf.connect(deployer).mint(user2.address, d18.mul("1000"));
    expect(await wtf.balanceOf(user2.address)).to.equal(d18.mul("1000"));
    expect(await wtf.balanceOf(deployer.address)).to.equal(initialSupply);
  });


  
});
