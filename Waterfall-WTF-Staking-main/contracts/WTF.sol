// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract WTF is ERC20, Ownable {

	constructor(
        uint256 initialSupply,
        address teamMultisig
    )
        public
        ERC20(
            "Waterfall Governance Token",       
            "WTF"                         
        )
    {
        _mint(
            msg.sender,                               
            initialSupply                                  
        );

        transferOwnership(teamMultisig);
    }

}
