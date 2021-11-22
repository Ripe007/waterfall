// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract BUSDMock is ERC20, Ownable {

	constructor()
        public
        ERC20( "BUSD",  "BUSD")
    {
        
        mint(msg.sender, 1000000e18);
    }

    function mint(address to, uint256 amount) public onlyOwner {
    	_mint(to,amount);
    }

}