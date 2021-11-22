//SPDX-License-Identifier: MIT
pragma solidity >=0.4.18 <=0.6.12;

pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract CreamFake is ERC20 {
    using SafeERC20 for IERC20;

    uint256 gains;
    address underlying;

    constructor(address _underlying, uint256 _gains) public ERC20("", "") {
        underlying = _underlying;
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
        _burn(msg.sender, redeemTokens);
        return amount;
    }

    function balanceOf(address account) public view override returns (uint256) {
        return super.balanceOf(account);
    }
}

contract CreamFakeLoss is ERC20 {
    using SafeERC20 for IERC20;

    uint256 loss;
    address underlying;

    constructor(address _underlying, uint256 _loss) public ERC20("", "") {
        underlying = _underlying;
        loss = _loss;
    }

    function mint(uint256 amount) public returns (uint256) {
        IERC20(underlying).safeTransferFrom(msg.sender, address(this), amount);
        _mint(msg.sender, amount);
        return amount;
    }

    function redeem(uint256 redeemTokens) public returns (uint256) {
        uint256 amount = (redeemTokens * (100 - loss)) / 100;

        IERC20(underlying).safeTransfer(msg.sender, amount);
        _burn(msg.sender, redeemTokens);
        return amount;
    }

    function balanceOf(address account) public view override returns (uint256) {
        return super.balanceOf(account);
    }
}
