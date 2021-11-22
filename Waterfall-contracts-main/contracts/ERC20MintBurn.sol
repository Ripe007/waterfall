// SPDX-License-Identifier: MIT
pragma solidity >=0.6.0 <0.7.0;

// import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
// import "@openzeppelin/contracts/token/ERC20/ERC20Capped.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/GSN/Context.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract ERC20MintBurn is ERC20, Ownable {
    constructor(string memory name, string memory symbol)
        public
        ERC20(name, symbol)
    {}

    function mint(address u, uint256 amount) public onlyOwner {
        _mint(u, amount);
    }

    function burnFrom(address u, uint256 amount) public onlyOwner {
        _burn(u, amount);
    }
}
