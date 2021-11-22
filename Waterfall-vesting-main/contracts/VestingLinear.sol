// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract VestingLinear is Ownable, ReentrancyGuard {

    using SafeMath for uint256;
    using SafeERC20 for IERC20;
    uint public constant NUMBER_OF_EPOCHS = 104;
    uint public constant EPOCH_DURATION = 604800; // 1 week duration
    IERC20 public wtf;

    uint256 public lastClaimedEpoch;
    uint256 private startTime;
    uint256 public totalVested;

    // Should we disallow contracts from becoming beneficiaries

    constructor(address _beneficiary, address _wtf, uint _startTime, uint _totalVested) public {
        transferOwnership(_beneficiary);
        wtf = IERC20(_wtf);
        startTime = _startTime;
        totalVested = _totalVested;

    }

    function claim () public nonReentrant {
        _claim(owner());
    }

    function _claim (address to) internal {
        uint256 balance;
        uint256 currentEpoch = getCurrentEpoch();

        if (currentEpoch > NUMBER_OF_EPOCHS + 1) {
            lastClaimedEpoch = NUMBER_OF_EPOCHS;
            wtf.safeTransfer(to, wtf.balanceOf(address (this)));
            return;
        }

        if (currentEpoch > lastClaimedEpoch) {
            balance = (currentEpoch.sub(1).sub(lastClaimedEpoch)).mul(totalVested).div(NUMBER_OF_EPOCHS);
            if (balance > 0) {
               wtf.safeTransfer(to, balance);
               lastClaimedEpoch = currentEpoch - 1;
          }
 
        }
        
    }

    function balance () public view returns (uint){
        return wtf.balanceOf(address (this));
    }

    function getCurrentEpoch () public view returns (uint){
        if (block.timestamp < startTime) {
            return 0;
        }
        return block.timestamp.sub(startTime).div(EPOCH_DURATION).add(1);
    }
    // default
    fallback() external { claim(); }
}
