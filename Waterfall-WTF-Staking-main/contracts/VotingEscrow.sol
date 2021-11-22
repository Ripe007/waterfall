// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./interfaces/IWTF.sol";
import "./VeWTF.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

interface IRewards {

	function stake(address account, uint256 _amount) external;
	function unstake(address account, uint256 _amount) external;
	function isPoolActive() external view returns(bool);
}

contract VotingEscrow is Ownable, ReentrancyGuard, VEWTF {

	using SafeMath for uint256;
	using SafeERC20 for IERC20;

	IERC20 public wtf; 
	IRewards public wtfRewards;
	IRewards public feeRewards;

	uint256 public constant MAX_LOCK_TIME = 31536000; // 365 common days
	uint256 public constant MIN_LOCK_TIME = 2592000; // 30 days

	struct Lock {
		uint startTimestamp;
		uint expiryTimestamp;
		uint amount;
	}

	mapping(address => Lock) internal locks;

	// Events

	event LockCreated(address indexed account, uint256 amount, uint256 startTimestamp, uint256 expiryTimestamp);
    event LockAmountIncreased(address indexed account, uint256 increasedAmount);
    event LockExtended(address indexed account, uint256 amount, uint256 newExpiryTimestamp);
    event Unlocked(address indexed account, uint256 amount);
    event StakingSet (address wtfRewards, address feeRewards);

	constructor(address _wtf) VEWTF("WTF Voting Escrow", "VeWTF") {
		wtf = IERC20(_wtf);
	}

	function setStaking (address _wtfRewards, address _feeRewards) external onlyOwner {

		require(_wtfRewards != address(0) && _feeRewards != address(0), "WTF Voting Escrow: staking zero address");
		
		wtfRewards = IRewards(_wtfRewards);
		feeRewards = IRewards(_feeRewards);

		emit StakingSet(_wtfRewards, _feeRewards);
	}

	/* 
	 * @notice: Creates a lock for WTF tokens if it does not exist yet
	 * @param: _amount: amount of token to lock
	 * @param: _duration: lock duration
	 */

	function createLock(uint256 _amount, uint256 duration )  external nonReentrant {

		_assertNotContract();

		Lock storage lock = locks[msg.sender];

		require (_amount > 0, "WTF Voting Escrow: Lock amount is zero" );
		require(lock.amount == 0, "WTF Voting Escrow: Lock already exists" );
		require(duration >= MIN_LOCK_TIME, "WTF Voting Escrow: Minimum lock time is 30 days");
		require(duration <= MAX_LOCK_TIME, "WTF Voting Escrow: Maximum lock time is 365 days");

		uint256 mint = _amount.mul(duration).div(MAX_LOCK_TIME);

		lock.startTimestamp = block.timestamp;
		lock.expiryTimestamp = lock.startTimestamp.add(duration);
		lock.amount = _amount;

		wtf.safeTransferFrom(msg.sender, address(this), _amount);

		_mint(msg.sender, mint);
		_stake(msg.sender,mint);

		emit LockCreated(msg.sender, lock.amount, lock.startTimestamp, lock.expiryTimestamp );
	}


	/*
	 * @notice: Allows extending duration of existing locks and increasing amount of tokens locked
	 * @param _amount: amount of tokens
	 * @param _newExpiryTimestamp: when lock expires
	 */

	function increaseTimeAndAmount(uint256 _amount, uint256 _newExpiryTimestamp) external nonReentrant {
		
		Lock storage lock = locks[msg.sender];
		
		require(lock.amount > 0, "WTF Voting Escrow: Lock does not exist");

		// Lock should not be expired

		require(!_isLockExpired(lock.expiryTimestamp), "WTF Voting Escrow: Lock is expired");

		require(_newExpiryTimestamp >= lock.expiryTimestamp.add(MIN_LOCK_TIME), "WTF Voting Escrow: Extended duration is less than MIN_LOCK_TIME");

		require(_newExpiryTimestamp.sub(lock.startTimestamp) <= MAX_LOCK_TIME, "WTF Voting Escrow: New lock duration is greater than MAX_LOCK_TIME" );

		uint256 expiryDiff = _newExpiryTimestamp.sub(lock.expiryTimestamp);

		uint256 mint; 

		// If new deposit is added to the lock

		if (_amount > 0) {

			uint256 timeToNewExpiry = _newExpiryTimestamp.sub(block.timestamp);
			uint256 mintForNewAmount = _amount.mul(timeToNewExpiry).div(MAX_LOCK_TIME);
			uint256 mintForOldAmount = lock.amount.mul(expiryDiff).div(MAX_LOCK_TIME);
			mint = mintForOldAmount.add(mintForNewAmount);
			lock.amount = lock.amount.add(_amount);
			wtf.safeTransferFrom(msg.sender,address(this), _amount);
			_mint(msg.sender, mint);
			_stake(msg.sender, mint);
		} else {
			mint = lock.amount.mul(expiryDiff).div(MAX_LOCK_TIME);
			_mint(msg.sender, mint);
			_stake(msg.sender,mint);
			
		}

		lock.expiryTimestamp = _newExpiryTimestamp;

		emit LockExtended(msg.sender, _amount, lock.expiryTimestamp);

	}


	/*
	 * @notice: Add tokens to existing lock
	 */

	function increaseAmount (uint256 _amount) external nonReentrant {

		Lock storage lock = locks[msg.sender];

		require(_amount > 0, "WTF Voting Escrow: Amount should be positive");
		require(lock.amount > 0, "WTF Voting Escrow: Lock does not exist");

		// Lock should not be expired

		require(!_isLockExpired(lock.expiryTimestamp), "WTF Voting Escrow: Lock is expired");

		uint256 timeToExpiry = lock.expiryTimestamp.sub(block.timestamp);
		uint256 mint = _amount.mul(timeToExpiry).div(MAX_LOCK_TIME);
		lock.amount = lock.amount.add(_amount);
		
		wtf.safeTransferFrom(msg.sender, address(this),_amount);
		_mint(msg.sender, mint);
		_stake(msg.sender, mint);
		emit LockAmountIncreased(msg.sender,_amount);
	}

	/*
	 * @notice Unlocks all tokens after lock expiry
	 */

	function unlock() external nonReentrant  {

		Lock storage lock = locks[msg.sender];
		uint256 wtfAmount = lock.amount;
		
		require(_isLockExpired(lock.expiryTimestamp), "WTF Voting Escrow: Cannot unlock tokens before expiry" );
		
		uint256 vewtfBal = balanceOf(msg.sender);
		require(vewtfBal > 0, "WTF Voting Escrow: VeWTF balance is zero");
		
		_burn(msg.sender, vewtfBal);
		_unstake(msg.sender, vewtfBal);
		wtf.safeTransfer(msg.sender, lock.amount);

		// Set lock to default values

		lock.amount = 0;
		lock.startTimestamp = 0;
		lock.expiryTimestamp = 0;
		
		emit Unlocked(msg.sender, wtfAmount);
	}

	function _stake(address account, uint256 amount) internal {

		if (address(wtfRewards) != address(0)) {
			if (wtfRewards.isPoolActive()) {
				wtfRewards.stake(account, amount);
			}
		}

		if (address(feeRewards) != address(0)) {
			if (feeRewards.isPoolActive()){
				feeRewards.stake(account, amount);
			}
		}
	}

	function _unstake(address account, uint256 amount) internal {
		
		if (address(wtfRewards) != address(0)) {
			wtfRewards.unstake(account, amount);
		}

		if (address(feeRewards) != address(0)) {
			feeRewards.unstake(account, amount);
		}
	}


	function _isLockExpired (uint256 expiryTimestamp ) internal view returns(bool){
		return (block.timestamp > expiryTimestamp);
	}

	function _assertNotContract() private view {
        if (msg.sender != tx.origin) {
            revert("Smart contract depositors not allowed");
        } else {
        	return;
        }
    }

    function getLockedAmount(address account) external view returns (uint256) {
        return locks[account].amount;
    }

    function totalLocked() external view  returns (uint256) {
        return wtf.balanceOf(address(this));
    }

    function getLockData(address account)
        external
        view
        returns (Lock memory lock)
    {
        return locks[account];
    }

}