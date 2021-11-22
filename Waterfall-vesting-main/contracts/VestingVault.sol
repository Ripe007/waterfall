pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";


contract VestingVault is Ownable {

    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    struct Grant {
        uint256 startTime;
        uint256 amount;
        uint256 vestingDuration;
        uint256 vestingCliff;
        uint256 lastEpochClaimed;
        uint256 totalClaimed;
        uint256 weeksClaimed;
        address recipient;
    }

    IERC20 public wtf;
    uint256 constant internal SECONDS_PER_WEEK = 604800;
    mapping (uint256 => Grant) public grants;
    mapping (address => uint[]) private activeGrants;
    uint256 public totalVestingCount;

    modifier onlyValidAddress(address _recipient) {
        require(_recipient != address(0) && _recipient != address(this) && _recipient != address(wtf), "Vesting: not valid _recipient");
        _;
    }


    event GrantAdded(address indexed recipient, uint256 vestingId);
    event GrantTokensClaimed(address indexed recipient, uint256 amountClaimed);
    event GrantRemoved(address recipient, uint256 amountVested, uint256 amountNotVested);
    event ChangedMultisig(address multisig);


    constructor(address _wtf, address _multisigOwner) public {
        require(address(_wtf) != address(0), "Vesting: zero address");
        require(address(_multisigOwner) != address(0),"Vesting: zero address");
        transferOwnership(_multisigOwner);
        wtf = IERC20(_wtf);
    }
    
    function addGrant(
        address _recipient,
        uint256 _startTime,
        uint256 _amount,
        uint256 _vestingDurationWeeks,
        uint256 _vestingCliffWeeks    
    ) 
        external
        onlyOwner()
    {
        require(_vestingCliffWeeks <= 100, "Vesting: Cliff is too large");
        require(_vestingDurationWeeks <= 1000, "Vesting: Grant duration is too large");
        require(_vestingDurationWeeks >= _vestingCliffWeeks, "Vesting: Cliff should not last longer than vesting");
        require(_recipient != address(0), "Vesting: recipient is zero address");
        
        uint256 amountVestedWeekly = _amount.div(_vestingDurationWeeks);
        require(amountVestedWeekly > 0, "Vesting: amount vested per week should not be zero");

        // Transfer the grant tokens under the control of the vesting contract
        
        wtf.safeTransferFrom(owner(), address(this), _amount);

        Grant memory grant = Grant({
            startTime: _startTime == 0 ? currentTime() : _startTime,
            amount: _amount,
            vestingDuration: _vestingDurationWeeks,
            vestingCliff: _vestingCliffWeeks,
            lastEpochClaimed: 0,
            weeksClaimed: 0,
            totalClaimed: 0,
            recipient: _recipient
        });
        grants[totalVestingCount] = grant;
        activeGrants[_recipient].push(totalVestingCount);
        emit GrantAdded(_recipient, totalVestingCount);
        totalVestingCount++;
    }

    function getActiveGrants(address _recipient) public view returns(uint256[] memory){
        return activeGrants[_recipient];
    }

    /* @notice Calculate the vested and unclaimed months and tokens available for `_grantId` to claim
     * Due to rounding errors once grant duration is reached, returns the entire left grant amount
     * Returns (0, 0) if cliff has not been reached
    */

    function calculateGrantClaim(uint256 _grantId) public view returns (uint256, uint256) {
        Grant storage grant = grants[_grantId];

        // For grants created with a future start date, that hasn't been reached, return 0, 0
        if (currentTime() < grant.startTime) {
            return (0, 0);
        }

        // Check cliff was reached
        uint256 elapsedTime = currentTime().sub(grant.startTime);
        uint256 elapsedWeeks = elapsedTime.div(SECONDS_PER_WEEK);
        
        if (elapsedWeeks < grant.vestingCliff) {
            return (elapsedWeeks, 0);
        }

        // If over vesting duration, all tokens vested

        if (elapsedWeeks > grant.vestingDuration) {
            uint256 remainingGrant = grant.amount.sub(grant.totalClaimed);
            return (grant.vestingDuration, remainingGrant);
        } else {
            uint256 weeksVested = elapsedWeeks.sub(grant.weeksClaimed);
            uint256 amountVestedPerWeek = grant.amount.div(grant.vestingDuration);
            uint256 amountVested = weeksVested.mul(amountVestedPerWeek);
            return (weeksVested, amountVested);
        }
    }

    /* @notice Allows a grant recipient to claim their vested tokens. Errors if no tokens have vested
     * It is advised recipients check they are entitled to claim via `calculateGrantClaim` before calling this
     */

    function claimVestedTokens(uint256 _grantId) external {
        uint256 weeksVested;
        uint256 amountVested;
        (weeksVested, amountVested) = calculateGrantClaim(_grantId);
        require(amountVested > 0, "Vesting: vested amount is 0");

        Grant storage grant = grants[_grantId];
        grant.weeksClaimed = grant.weeksClaimed.add(weeksVested);
        grant.totalClaimed = grant.totalClaimed.add(amountVested);
        
        wtf.safeTransfer(grant.recipient, amountVested);
        emit GrantTokensClaimed(grant.recipient, amountVested);
    }

    /* @notice Terminate token grant transferring all vested tokens to the `_grantId`
     * and returning all non-vested tokens to the V12 MultiSig
     * Secured to the V12 MultiSig only
     * @param _grantId grantId of the token grant recipient
    */

    function revokeGrant(uint256 _grantId) 
        external 
        onlyOwner()
    {
        Grant storage grant = grants[_grantId];
        address recipient = grant.recipient;
        uint256 weeksVested;
        uint256 amountVested;
        (weeksVested, amountVested) = calculateGrantClaim(_grantId);

        uint256 amountNotVested = (grant.amount.sub(grant.totalClaimed)).sub(amountVested);

        wtf.safeTransfer(recipient, amountVested);
        wtf.safeTransfer(owner(), amountNotVested);

        grant.startTime = 0;
        grant.amount = 0;
        grant.vestingDuration = 0;
        grant.vestingCliff = 0;
        grant.weeksClaimed = 0;
        grant.totalClaimed = 0;
        grant.recipient = address(0);

        emit GrantRemoved(recipient, amountVested, amountNotVested);
    }

    function currentTime() private view returns(uint256) {
        return block.timestamp;
    }

    function tokensVestedPerWeek(uint256 _grantId) public view returns(uint256) {
        Grant storage grant = grants[_grantId];
        return grant.amount.div(grant.vestingDuration);
    }

}