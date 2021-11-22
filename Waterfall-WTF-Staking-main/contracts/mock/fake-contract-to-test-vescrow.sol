// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IVotingEscrow {

	function createLock(uint256 _amount, uint256 expiryTimestamp) external;
}


contract TestContractAccessToVEscrow {

	IVotingEscrow internal escrow; 

	constructor (address _ve) public {
		escrow = IVotingEscrow(_ve);
	}

	function testLock  (uint256 _amount, uint256 expiryTimestamp) external {
		escrow.createLock(_amount, expiryTimestamp);
	}
}