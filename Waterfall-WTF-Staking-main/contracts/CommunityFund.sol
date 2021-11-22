// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract CommunityFund is Ownable {

    IERC20 public wtf;

    constructor (address _wtf) {
        wtf = IERC20(_wtf);
    }

    event SetAllowance(address indexed caller, address indexed spender, uint256 amount);

    function setAllowance(address spender, uint amount) public onlyOwner {
        wtf.approve(spender, amount);

        emit SetAllowance(msg.sender, spender, amount);
    }
}