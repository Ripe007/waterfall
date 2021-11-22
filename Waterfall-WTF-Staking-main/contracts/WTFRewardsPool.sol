// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";


contract WTFRewards is Ownable, ReentrancyGuard {

	using SafeMath for uint256;
    using SafeERC20 for IERC20;

    /***************************************************** STORAGE **********************************************/

    address public vewtf; 

	IERC20 public wtf;

	address public communityFund; 

	uint256 public constant PRECISION = 1e18;

	struct User {
		uint256 amount;
		uint256 rewardDebt;
	}

	struct Pool {
		uint256 accRewardPerShare;
		uint256 startRewardBlock;
		uint256 endRewardBlock;
		uint256 lastRewardBlock;
		uint256 rewardPerBlock;
		uint256 totalStaked;
	}

	Pool public pool;

	mapping(address => User) public users; 

	// Events

	event Stake (address indexed user, uint256 amount);
	event Unstake(address indexed user, uint256 amount);
	event Claim(address indexed user, uint256 amount);
	event AdminTokenRecovery(address indexed tokenRecovered, uint256 amount);
	event NewStartAndEndBlocks(uint256 startRewardBlock, uint256 endRewardBlock);

	constructor (address _vewtf, 
		         address _wtf,
		         address _communityFund,
		         uint256 _rewardPerBlock, 
		         uint256 _startRewardBlock, 
		         uint256 _endRewardBlock
    ) {
		vewtf = _vewtf;
		wtf = IERC20(_wtf);
		communityFund = _communityFund;
		pool.rewardPerBlock = _rewardPerBlock;
		pool.startRewardBlock = _startRewardBlock;
		pool.endRewardBlock = _endRewardBlock;
		pool.lastRewardBlock = _startRewardBlock;
	}

	/** 
	 * @dev calls to stake and unstake are allowed only from the Voting Escrow contract that automatically registers
	 * deposits and withdrawals
	 */
	
	modifier onlyVotingEscrow(){
		require(msg.sender == address(vewtf), 'VeWTF Staking: Not authorized');
		_;
	}

	function getRewardDebt(address account) external view returns(uint256){
		return users[account].rewardDebt;
	}

	function totalStaked() external view returns(uint256) {
		return pool.totalStaked;
	}

	function getAccountData(address account) external view returns(User memory user) {

		user = users[account];
	}

	function isPoolActive() public view returns(bool) {
		return (block.number < pool.endRewardBlock);
	}

	function lastRewardBlock() external view returns (uint256) {
		return pool.lastRewardBlock;
	}

	function rewardPerShare() external view returns(uint256) {
		return pool.accRewardPerShare;
	}

	function stake(address account, uint256 _amount) external onlyVotingEscrow {

		require(isPoolActive(), 'VeWTF Staking: Rewards pool is not active');
		require(_amount > 0, "VeWTF Staking: Amount is zero");
		
		User storage user = users[account];
		
		_updatePool();
		
		if (user.amount > 0) {
			uint256 reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
			if (reward > 0) {
				wtf.safeTransferFrom(communityFund, account, reward);
			}
		}

		user.amount = user.amount.add(_amount);
		pool.totalStaked = pool.totalStaked.add(_amount);
		user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

		emit Stake(account, _amount);

	}

	function unstake(address account, uint256 _amount) external nonReentrant onlyVotingEscrow {

		User storage user = users[account];

		require(user.amount >= _amount, "VeWTF Staking: Not enough tokens to withdraw" );
		require(_amount > 0, "VeWTF Staking: Can`t withdraw zero amount");

		_updatePool();

		uint256 reward = pool.accRewardPerShare.mul(user.amount).div(PRECISION).sub(user.rewardDebt);
		
		if (reward > 0) {
			wtf.safeTransferFrom(communityFund, account, reward);
	    }

	    user.amount = user.amount.sub(_amount);
	    pool.totalStaked = pool.totalStaked.sub(_amount);

	    user.rewardDebt = pool.accRewardPerShare.mul(user.amount).div(PRECISION);

	    emit Unstake(account, _amount);
	}


	function claimRewards() external nonReentrant {

		User storage user = users[msg.sender];

		require(block.number >= pool.startRewardBlock, "VeWTF Staking: Staking has not started yet");

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

        if (block.number > pool.lastRewardBlock && pool.totalStaked != 0) {
            
            uint256 blockDiff = getBlockDiff(pool.lastRewardBlock, block.number);
            uint256 accReward = blockDiff.mul(pool.rewardPerBlock);
            uint256 adjustedTokenPerShare = pool.accRewardPerShare.add(accReward.mul(PRECISION).div(pool.totalStaked));
            reward = user.amount.mul(adjustedTokenPerShare).div(PRECISION).sub(user.rewardDebt);

        } else {
            reward =  user.amount.mul(pool.accRewardPerShare).div(PRECISION).sub(user.rewardDebt);
        }

    }

    function _updatePool() internal {

		if (block.number <= pool.lastRewardBlock) {
            return;
        }

		if (pool.totalStaked == 0) {
            pool.lastRewardBlock = block.number;
            return;
        }

		uint256 multiplier = getBlockDiff(pool.lastRewardBlock, block.number);
		uint256 accReward = multiplier.mul(pool.rewardPerBlock);
		pool.accRewardPerShare = pool.accRewardPerShare.add(accReward.mul(PRECISION).div(pool.totalStaked));
		pool.lastRewardBlock = block.number;
	}

    function updateStartAndEndBlocks(uint256 _startRewardBlock, uint256 _endRewardBlock) external onlyOwner {

        require(block.number < pool.startRewardBlock, "VeWTF Staking: Pool has started");
        require(_startRewardBlock < _endRewardBlock, "VeWTF Staking: New startBlock must be lower than new endBlock");
        require(block.number < _startRewardBlock, "VeWTF Staking: New startBlock must be higher than current block");

        pool.startRewardBlock = _startRewardBlock;
        pool.endRewardBlock = _endRewardBlock;

        // Set the lastRewardBlock as the startBlock
        pool.lastRewardBlock = _startRewardBlock;

        emit NewStartAndEndBlocks(_startRewardBlock, _endRewardBlock);
    }


	function getBlockDiff (uint256 _from, uint256 _to) internal view returns(uint256) {

		if (_from >= pool.endRewardBlock) {
			return 0; 
		}
		else if (_to >= pool.endRewardBlock) {
			return pool.endRewardBlock.sub(_from);
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
        pool.endRewardBlock = block.number;
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
        require(_tokenAddress != address(wtf), "VeWTF Staking: Cannot be reward token");

        IERC20(_tokenAddress).safeTransfer(address(msg.sender), _tokenAmount);

        emit AdminTokenRecovery(_tokenAddress, _tokenAmount);
    }

}