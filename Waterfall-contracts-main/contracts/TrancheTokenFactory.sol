//SPDX-License-Identifier: MIT
pragma solidity >=0.4.18 <=0.6.12;

import "./ERC20MintBurn.sol";

contract TrancheTokenFactory {
    function deploy(string memory name, string memory symbol)
        public
        returns (address)
    {
        ERC20MintBurn token = new ERC20MintBurn(name, symbol);
        token.transferOwnership(msg.sender);
        return address(token);
    }
}
