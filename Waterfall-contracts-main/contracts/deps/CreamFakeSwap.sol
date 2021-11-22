//SPDX-License-Identifier: MIT
pragma solidity >=0.4.18 <=0.6.12;

pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "hardhat/console.sol";

contract CreamFakeSwap is ERC20 {
    using SafeERC20 for IERC20;

    uint256 gains;
    address underlying;
    address alpacaToken;

    constructor(
        address _underlying,
        address _alpacaToken,
        uint256 _gains
    ) public ERC20("", "") {
        underlying = _underlying;
        alpacaToken = _alpacaToken;
        gains = _gains;
    }

    function mint(uint256 amount) public returns (uint256) {
        IERC20(underlying).safeTransferFrom(msg.sender, address(this), amount);
        _mint(msg.sender, amount);
        return amount;
    }

    function redeem(uint256 redeemTokens) public returns (uint256) {
        uint256 amount = (redeemTokens * (100 + gains)) / 100;
        IERC20(underlying).safeTransfer(msg.sender, amount);
        IERC20(alpacaToken).safeTransfer(msg.sender, amount);
        console.log(address(this));
        console.log(msg.sender);
        _burn(msg.sender, redeemTokens);
        return amount;
    }

    function balanceOf(address account) public view override returns (uint256) {
        return super.balanceOf(account);
    }
}
