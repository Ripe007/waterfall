//SPDX-License-Identifier: MIT

pragma solidity >=0.4.22 <0.8.0;

pragma experimental ABIEncoderV2;

import "./ERC20.sol";

contract WTF is ERC20Capped {
    constructor(
        string memory name_,
        string memory symbol_,
        address faucet,
        uint256 cap_
    ) public ERC20(name_, symbol_) ERC20Capped(cap_) {
        _mint(faucet, cap_);
    }
}
