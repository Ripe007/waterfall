// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract VestingFundsAllocator is Ownable {

	using SafeERC20 for IERC20;

    address[] public vestingBeneficiaries;
    uint[] public vestingAmounts;
    uint public lastAllocatedAddress;
    IERC20 private wtf;

    constructor (address[] memory _vestingBeneficiaries, uint[] memory _vestingAmounts, address _wtf, address _multisigOwner) public {
        vestingBeneficiaries = _vestingBeneficiaries;
        vestingAmounts = _vestingAmounts;
        wtf = IERC20(_wtf);
        transferOwnership(_multisigOwner);
    }

    function allocateVestingFunds () public onlyOwner {
        for (uint i = lastAllocatedAddress; i < vestingBeneficiaries.length; i++) {
            if (wtf.balanceOf(address(this)) < vestingAmounts[i] || gasleft() < 20000) {
                break;
            }
            lastAllocatedAddress++;
            wtf.safeTransfer(vestingBeneficiaries[i], vestingAmounts[i]);
        }
    }

    fallback () external { allocateVestingFunds(); }
}
