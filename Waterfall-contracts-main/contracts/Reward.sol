//SPDX-License-Identifier: MIT

pragma solidity >=0.4.22 <0.8.0;

import "@openzeppelin/contracts/GSN/Context.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/math/Math.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";

contract Reward is Context, Ownable {
    using Address for address;
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    address public stakeToken;
    IERC20 public rewardToken;
    uint256 public duration; // making it not a constant is less gas efficient, but portable

    uint256 public periodFinish = 0;
    uint256 public rewardRate = 0;
    uint256 public lastUpdate;
    uint256 public rewardPerTokenStored;
    mapping(address => uint256) public userRewardPerTokenPaid;
    mapping(address => uint256) public rewards;

    bool public stopped = false;

    uint256 private _totalSupply;
    mapping(address => uint256) private _balances;

    address public admin;

    event NewPeriod(uint256 duration, uint256 reward);
    event Staked(address indexed user, uint256 amount);
    event Withdrawn(address indexed user, uint256 amount);
    event RewardPaid(address indexed user, uint256 reward);
    event RewardDenied(address indexed user, uint256 reward);

    modifier updateReward(address account) {
        rewardPerTokenStored = rewardPerToken();
        lastUpdate = lastTimeRewardApplicable();
        if (account != address(0)) {
            rewards[account] = earned(account);
            userRewardPerTokenPaid[account] = rewardPerTokenStored;
        }
        _;
    }

    modifier fromAdmin() {
        require(
            msg.sender == admin || msg.sender == owner(),
            "only admin or owner"
        );
        _;
    }

    modifier onlyStopped() {
        require(stopped == true, "not stopped");
        _;
    }

    constructor(
        address _admin,
        address _rewardToken,
        address _stakeToken
    ) public {
        rewardToken = IERC20(_rewardToken);
        stakeToken = _stakeToken;
        admin = _admin;
    }

    function lastTimeRewardApplicable() public view returns (uint256) {
        return Math.min(block.timestamp, periodFinish);
    }

    function rewardPerToken() public view returns (uint256) {
        if (totalSupply() == 0) {
            return rewardPerTokenStored;
        }
        return
            rewardPerTokenStored.add(
                lastTimeRewardApplicable()
                    .sub(lastUpdate)
                    .mul(rewardRate)
                    .mul(1e18)
                    .div(totalSupply())
            );
    }

    function earned(address account) public view returns (uint256) {
        return
            balanceOf(account)
                .mul(rewardPerToken().sub(userRewardPerTokenPaid[account]))
                .div(1e18)
                .add(rewards[account]);
    }

    function stakeFor(address who, uint256 amount)
        public
        fromAdmin
        updateReward(who)
    {
        require(amount > 0, "Cannot stake 0");
        _stake_(msg.sender, who, amount);
        emit Staked(who, amount);
    }

    function withdrawFor(address who, uint256 amount)
        public
        fromAdmin
        updateReward(who)
    {
        require(amount > 0, "Cannot withdraw 0");
        pushReward(who);
        _withdraw_(who, amount);
        emit Withdrawn(who, amount);
    }

    function exitFor(address who) external fromAdmin returns (uint256) {
        withdrawFor(who, balanceOf(who));
        return getRewardFor(who);
    }

    /// A push mechanism for accounts that have not claimed their rewards for a long time.
    /// The implementation is semantically analogous to getReward(), but uses a push pattern
    /// instead of pull pattern.
    function pushReward(address recipient)
        public
        updateReward(recipient)
        returns (uint256)
    {
        uint256 reward = earned(recipient);
        if (reward > 0) {
            rewards[recipient] = 0;
            rewardToken.safeTransfer(recipient, reward);
            emit RewardPaid(recipient, reward);
        }
        return reward;
    }

    function getRewardFor(address who)
        public
        fromAdmin
        updateReward(who)
        returns (uint256)
    {
        uint256 reward = earned(who);
        if (reward > 0) {
            rewards[who] = 0;
            rewardToken.safeTransfer(who, reward);
            emit RewardPaid(who, reward);
        }
        return reward;
    }

    function getReward() public updateReward(msg.sender) returns (uint256) {
        uint256 reward = earned(msg.sender);
        if (reward > 0) {
            rewards[msg.sender] = 0;
            rewardToken.safeTransfer(msg.sender, reward);
            emit RewardPaid(msg.sender, reward);
        }
        return reward;
    }

    function newPeriod(uint256 _duration, uint256 _reward)
        public
        fromAdmin
        updateReward(address(0))
    {
        if (block.timestamp >= periodFinish) {
            rewardRate = _reward.div(_duration);
        } else {
            uint256 remaining = periodFinish.sub(block.timestamp);
            uint256 leftover = remaining.mul(rewardRate);
            rewardRate = _reward.add(leftover).div(_duration);
        }
        lastUpdate = block.timestamp;
        periodFinish = block.timestamp.add(_duration);
        duration = _duration;

        emit NewPeriod(_duration, _reward);
    }

    function stop() public fromAdmin {
        stopped = true;
        if (periodFinish > block.timestamp) {
            periodFinish = block.timestamp;
        }
    }

    // this will only get your principal back, if there is any
    function evacuateOnEmergency() public onlyStopped {
        uint256 amount = _balances[msg.sender];
        _totalSupply = _totalSupply.sub(amount);
        delete _balances[msg.sender];
        IERC20(stakeToken).safeTransfer(msg.sender, amount);
    }

    function totalSupply() public view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    function _stake_(
        address _funder,
        address forWhom,
        uint256 amount
    ) internal virtual {
        _totalSupply = _totalSupply.add(amount);
        _balances[forWhom] = _balances[forWhom].add(amount);
        IERC20(stakeToken).safeTransferFrom(_funder, address(this), amount);
    }

    function _withdraw_(address who, uint256 amount) internal virtual {
        _totalSupply = _totalSupply.sub(amount);
        _balances[who] = _balances[who].sub(amount);
        IERC20(stakeToken).safeTransfer(who, amount);
    }

    function open() public fromAdmin {
        stopped = false;
    }
}
