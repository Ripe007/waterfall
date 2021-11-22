require("@nomiclabs/hardhat-waffle");
require('@openzeppelin/hardhat-upgrades');
require("@nomiclabs/hardhat-web3");
const infuraKey = "6871a10ec4a44296b572197203b48017";
const privateKeys = ["fa41515dfb14f2baa9832377201043d66e2ec70ee193710f6d33f3577d11b782",
    "9ac35efc636a1132fd2f6fe26a9159582d69330b41ecec4193c7c437151b23ea",
    "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
    "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
];

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async () => {
  const accounts = await ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});



// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  defaultNetwork: "localhost",
  networks: {
    hardhat: {},
    binance_testnet: {
      chainId: 97,
      url: `https://data-seed-prebsc-1-s3.binance.org:8545/`,
      accounts: privateKeys,
      timeout: 200000,
    },
    binance_mainnet: {
      chainId: 56,
      url: `https://bsc-dataseed.binance.org/`,
      accounts: privateKeys,
      timeout: 200000,
    }
  },
  solidity: {
    version: "0.8.0",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  }
};

