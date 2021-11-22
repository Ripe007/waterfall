// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract WTFLPRewards is Ownable, ReentrancyGuard {

	using SafeMath for uint256;
    using SafeERC20 for IERC20;

    /***************************************************** STORAGE **********************************************/

	IERC20 public wtf;
	IERC20 public lpToken;
	address public communityFund;

	uint256 public constant PRECISION = 1e18;

	struct User {
		uint256 amount;
		uint256 rewardDebt;
	}

    // WTF Rewards

	struct Pool {
		uint256 totalRewards;
		uint256 accRewardPerShare;
		uint256 startRewardTs;
		uint256 endRewardTs;
		uint256 lastRewardTs;
		uint256 rewardDuration;
	}

	Pool public pool;

	mapping(address => User) users;

	// Events

	event Stake (address indexed user, uint256 amount);
	event Unstake(address indexed user, uint256 amount);
	event Claim(address indexed user, uint256 amount);
	event AdminTokenRecovery(address indexed tokenRecovered, uint256 amount);
	event NewStartAndEndTs(uint256 startRewardTs, uint256 endRewardTs);

	constructor (address _wtf, 
		         address _lpToken,
		         address _communityFund, 
		         uint256 _startRewardTs, 
		         uint256 _endRewardTs,
		         uint256 _totalRewards
		         ) {
		
		require(_startRewardTs >= block.timestamp, "WTF LP Staking: reward start should be in the future");
		require(_endRewardTs > _startRewardTs, "WTF LP Staking: end reward time should be greater than start reward time");
		require(_wtf != address(0) && _lpToken != address(0) && _communityFund != address(0), "WTL LP staking: zero address");

		wtf = IERC20(_wtf);
		lpToken = IERC20(_lpToken);
		communityFund = _communityFund;
		pool.totalRewards = _totalRewards;
		pool.startRewardTs = _startRewardTs;
		pool.endRewardTs = _endRewardTs;
		pool.rewardDuration = _endRewardTs.sub(_startRewardTs);
		pool.lastRewardTs = pool.startRewardTs;

	}

	function getRewardDebt(address account) external view returns(uint256){
		return users[account].rewardDebt;
	}

	function isPoolActive() public view returns(bool) {
		return (block.timestamp <= pool.endRewardTs);
	}

	function getAccountData(address account) external view returns(User memory user) {
		user = users[account];
	}

	function lastRewardTs() external view returns (uint256) {
		return pool.lastRewardTs;
	}

	function rewardPerShare() external view returns(uint256) {
		return pool.accRewardPerShare;
	}

	function stake(uint256 _amount) external nonReentrant {

		require(isPoolActive(), 'WTF LP Staking: Rewards pool is not active');
		require(_amount > 0, "WTF LP Staking: Amount is zero");
		
		User storage user = users[msg.sender];
		
		_updatePool();
		
		if (user.amount > 0) {
			uint256 reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
			if (reward > 0) {
				wtf.safeTransferFrom(communityFund, msg.sender, reward);
			}
		}

		lpToken.safeTransferFrom(msg.sender, address(this), _amount);
		user.amount = user.amount.add(_amount);
		user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

		emit Stake(msg.sender, _amount);

	}

	function unstake(uint256 _amount) external nonReentrant  {

		User storage user = users[msg.sender];

		require(user.amount >= _amount, "WTF LP Staking: Not enough tokens to withdraw" );
		require(_amount > 0, "WTF LP Staking: Can`t withdraw zero amount");

		_updatePool();

		uint256 reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
		
		if (reward > 0) {
			wtf.safeTransferFrom(communityFund, msg.sender, reward);
	    }

	    lpToken.safeTransfer(msg.sender, _amount);
	    user.amount = user.amount.sub(_amount);
	    user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

	    emit Unstake(msg.sender, _amount);
	}


	function claimRewards() external nonReentrant {

		User storage user = users[msg.sender];

		require(block.timestamp >= pool.startRewardTs, "WTF LP Staking: Staking has not started yet");

		_updatePool();

		uint256 reward;

		// Calculate any debt

		if (user.amount > 0) {
			reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
			if (reward > 0) {
				wtf.safeTransferFrom(communityFund, msg.sender, reward);	
	       }
	    }

	    user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

	    emit Claim(msg.sender, reward);

	}

    function pendingReward(address _user) public view returns (uint256 reward) {

        User memory user = users[_user];

        uint256 totalStaked = lpToken.balanceOf(address(this));

        if (block.timestamp > pool.lastRewardTs && totalStaked != 0) {
            uint256 timeDiff = getTimeDiff(pool.lastRewardTs, block.timestamp);
            uint256 accReward = timeDiff.mul(pool.totalRewards).div(pool.rewardDuration);
            uint256 adjustedTokenPerShare = pool.accRewardPerShare.add(accReward.mul(PRECISION).div(totalStaked));
            reward = user.amount.mul(adjustedTokenPerShare).div(PRECISION).sub(user.rewardDebt);

        } else {
            reward =  user.amount.mul(pool.accRewardPerShare).div(PRECISION).sub(user.rewardDebt);
        }

    }

    function _updatePool() internal {

    	uint256 totalStaked = lpToken.balanceOf(address(this));

		if (block.timestamp <= pool.lastRewardTs) {
            return;
        }

		if (totalStaked == 0) {
            pool.lastRewardTs = block.timestamp;
            return;
        }

		uint256 timeDiff = getTimeDiff(pool.lastRewardTs, block.timestamp);
		uint256 accReward = timeDiff.mul(pool.totalRewards).div(pool.rewardDuration);
		pool.accRewardPerShare = pool.accRewardPerShare.add(accReward.mul(PRECISION).div(totalStaked));
		pool.lastRewardTs = block.timestamp;
	}

    function updateStartAndEndTs(uint256 _startRewardTs, uint256 _endRewardTs) external onlyOwner {

        require(block.timestamp < pool.startRewardTs, "WTF LP Staking: Pool has started");
        require(_startRewardTs < _endRewardTs, "WTF LP Staking: New start timestamp must be lower than new end timestamp");
        require(block.timestamp < _startRewardTs, "WTF LP Staking: New start timestamp must be higher than current block timestamp");

        pool.startRewardTs = _startRewardTs;
        pool.endRewardTs = _endRewardTs;

        // Set the lastRewardBlock as the startBlock
        pool.lastRewardTs = _startRewardTs;
        pool.rewardDuration = _endRewardTs.sub(_startRewardTs);

        emit NewStartAndEndTs(_startRewardTs, _endRewardTs);
    }


	function getTimeDiff (uint256 _from, uint256 _to) internal view returns(uint256) {

		if (_from >= pool.endRewardTs) {
			return 0; 
		}
		else if (_to >= pool.endRewardTs) {
			return pool.endRewardTs.sub(_from);
		}
		else {
			return _to.sub(_from);
		}

	}

	/*
     * @notice Stop rewards
     * @dev Only callable by owner
     */
    function stopReward() external onlyOwner {
        pool.endRewardTs = block.timestamp;
    }
	
     /* 
      * @notice It allows the admin to recover wrong tokens sent to the contract
      * @param _tokenAddress: the address of the token to withdraw
      * @param _tokenAmount: the number of tokens to withdraw
      * @dev This function is only callable by admin.
      */

	function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) external onlyOwner {
        require(_tokenAddress != address(lpToken), "WTF LP Staking: Cannot be staked token");
        require(_tokenAddress != address(wtf), "WTF LP Staking: Cannot be reward token");

        IERC20(_tokenAddress).safeTransfer(address(msg.sender), _tokenAmount);

        emit AdminTokenRecovery(_tokenAddress, _tokenAmount);
    }

}