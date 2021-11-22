// // SPDX-License-Identifier: Apache-2.0
// pragma solidity ^0.8.0;
// import "@openzeppelin/contracts/utils/math/SafeMath.sol";
// import "@openzeppelin/contracts/access/Ownable.sol";
// import "@openzeppelin/contracts/token/ERC20/IERC20.sol";



// contract FeesFunds is Ownable {

// 	IERC20 feeToken;
// 	bytes32 public constant REWARD_COLLECTOR_ROLE = keccak256("REWARD_COLLECTOR_ROLE");

// 	constructor(address rewardCollector, address _feeToken) public {

// 		_setupRole(REWARD_COLLECTOR_ROLE, rewardCollector);
// 		feeToken = IERC20(_feeToken);
// 	}

// 	function pullToken() external {
// 		require(hasRole(REWARD_COLLECTOR_ROLE, _msgSender()), "Cannot access the fees fund");
// 		uint256 availableFunds = feeToken.balanceOf(address(this));
// 		require(availableFunds > 0, 'There are no rewards in the fund');
// 		feeToken.transfer(msg.sender, availableFunds );
// 	}

// }

