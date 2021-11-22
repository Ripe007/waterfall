// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";


contract FeeRewards is Ownable, ReentrancyGuard {

	using SafeMath for uint256;
    using SafeERC20 for IERC20;

    /***************************************************** STORAGE **********************************************/

    address public vewtf; 

	IERC20 public rewardToken;

	uint256 public constant PRECISION = 1e18;

	struct User {
		uint256 amount;
		uint256 rewardDebt;
	}

	struct Pool {
		bool isActive;
		uint256 creationTS;
		uint256 accRewardPerShare;
		uint256 totalStaked;
	}

	Pool public pool; 

	mapping(address => User) public users; 

	// Events

	event Stake (address indexed user, uint256 amount);
	event Unstake(address indexed user, uint256 amount);
	event Claim(address indexed user, uint256 reward);
	event AdminTokenRecovery(address indexed tokenRecovered, uint256 amount);

	constructor (address _vewtf, address _rewardToken) {

		vewtf = _vewtf;
		rewardToken = IERC20(_rewardToken);
		pool.creationTS = block.timestamp;
		pool.isActive = true; 

	}

	modifier onlyVotingEscrow(){
		require(msg.sender == address(vewtf), 'VeWTF Staking: Not authorized');
		_;
	}

	function isPoolActive() external view returns(bool) {
		return pool.isActive;
	}

	/**
	 * @notice sendRewards can be used by fee collector to send fee rewards to this contract
	 */

	function sendRewards(uint256 _amount) external  {
		
		// Allow sending rewards when pool is active

		require(pool.isActive, "VeWTF Staking: Pool is closed" );

		// Rewards can be sent if there are stakes

        if (pool.totalStaked > 0) {
            rewardToken.safeTransferFrom(msg.sender, address(this), _amount);
            _updatePool(_amount);
        } 
        else { return; }
    }


	function stake(address account, uint256 _amount) external nonReentrant onlyVotingEscrow {

		require(pool.isActive, "VeWTF Staking: Pool is closed" );
		require(_amount > 0, "VeWTF Staking: Stake amount should be positive");

		User storage user = users[account];

		// Send pending rewards

		if (user.amount > 0) {

			uint256 reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
			if (reward > 0) {
				rewardToken.safeTransfer(account, reward);
			}
		}

		user.amount = user.amount.add(_amount);
		pool.totalStaked = pool.totalStaked.add(_amount);

		user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

		emit Stake(account, _amount);
	}


	function unstake(address account, uint256 _amount) external nonReentrant onlyVotingEscrow {

		User storage user = users[account];

		require(_amount > 0, "VeWTF Staking: cannot withdraw zero amount");
		require (_amount <= user.amount, "VeWTF Staking: not enough tokens to withdraw");

		// Send rewards

		uint256 reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
		
		if (reward > 0) {
			rewardToken.safeTransfer(account, reward);
	    }

	    user.amount = user.amount.sub(_amount);
	    pool.totalStaked = pool.totalStaked.sub(_amount);

	    user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

	    emit Unstake(account, _amount);

	}

	function claimRewards() external nonReentrant {

		User storage user = users[msg.sender];
		uint256 reward;

		if (user.amount > 0) {
			reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
			if (reward > 0) {
				rewardToken.safeTransfer(msg.sender, reward);
	       }
	    }
	    user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);
	    emit Claim(msg.sender, reward);
	}

	function _updatePool(uint256 _amount) internal {

    	pool.accRewardPerShare = pool.accRewardPerShare.add(_amount.mul(PRECISION).div(pool.totalStaked));
    }

	/*
     * @notice View function to see pending reward on frontend.
     * @param _user: user address
     * @return Pending reward for a given user
     */
    function pendingRewardOf(address _user) external view returns (uint256 reward) {

    	User memory user = users[_user];
        reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
    }


    function setRewardToken(address _rewardToken) external onlyOwner {
    	require(_rewardToken != address(0), "VeWTF Staking: zero address");
    	rewardToken = IERC20(_rewardToken);
    }

    function closePool() external onlyOwner {
    	pool.isActive = false;
    }

    /*
     * @notice Stop rewards
     * @dev Only callable by owner. Needs to be for emergency.
     */
    function emergencyRewardWithdraw(uint256 _amount) external onlyOwner {
        rewardToken.safeTransfer(address(msg.sender), _amount);
    }

    function evacuateETH(address recv) public onlyOwner{
        payable(recv).transfer(address(this).balance);
    }

     /* 
      * @notice It allows the admin to recover wrong tokens sent to the contract
      * @param _tokenAddress: the address of the token to withdraw
      * @param _tokenAmount: the number of tokens to withdraw
      * @dev This function is only callable by admin.
      */

	function recoverWrongTokens(address _tokenAddress, uint256 _tokenAmount) external onlyOwner {
        require(_tokenAddress != address(vewtf), "VeWTF Staking: Cannot be staked token");
        require(_tokenAddress != address(rewardToken), "VeWTF Staking: Cannot be reward token");

        IERC20(_tokenAddress).safeTransfer(address(msg.sender), _tokenAmount);

        emit AdminTokenRecovery(_tokenAddress, _tokenAmount);
    }

}

