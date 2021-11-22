//SPDX-License-Identifier: MIT
pragma solidity >=0.4.18 <=0.6.12;

pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/SafeERC20.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";

import "hardhat/console.sol";

interface ITrancheTokenFactory {
    function deploy(string memory name, string memory symbol)
        external
        returns (address);
}

interface IERC20MintBurn {
    function mint(address, uint256) external;

    function burnFrom(address, uint256) external;
}

// venus & cream
interface ICompoundLike {
    function mint(uint256 amount) external returns (uint256);

    function redeem(uint256 redeemTokens) external returns (uint256);

    function balanceOf(address account) external returns (uint256);

    function balanceOfUnderlying(address account) external returns (uint256);
}

// alpaca
interface IAlpacaLike {
    function deposit(uint256 amount) external;

    function withdraw(uint256 share) external;

    function balanceOf(address amount) external;

    /// @dev Return the pending interest that will be accrued in the next call.
    /// @param value Balcance value to subtract off address(this).balance when called from payable functions.
    function pendingInterest(uint256 value) external returns (uint256);
}

interface ITrancheReward {
    function stakeFor(address who, uint256 amount) external;

    function withdrawFor(address who, uint256 amount) external;

    function exitFor(address who) external;
}

interface IRewardFactory {
    function deploy(
        address admin,
        address rewardToken,
        address stakeToken
    ) external returns (address);
}

contract CampaignContinuousCycles is Ownable {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    struct TrancheParams {
        uint256 apy;
        uint256 fee;
        uint256 percentage;
    }

    struct Tranche {
        uint256 target;
        uint256 principal;
        uint256 capital;
        uint256 percentage;
        uint256 apy;
        uint256 fee;
        address token;
        address reward;
    }

    struct TrancheExit {
        address token;
        uint256 cycle;
        uint256 balance;
        uint256 investment;
    }

    struct Account {
        uint256 investment;
        uint256 joinCycle;
        mapping(uint256 => TrancheExit) exitCheckpoint;
    }

    struct JoinDuring {
        address addr;
        uint256 amount;
        uint256 opt;
        bool evac;
    }

    // Constant declaration
    uint256 public constant PercentageParamScale = 1e5;
    uint256 public constant PercentageScale = 1e18;

    address public WTF;
    address public RewardFactory;
    address public TrancheTokenFactory;
    address public CreamFarm;
    address public VenusFarm;
    address public VenusToken;
    address public AlpacaFarm;
    address public AlpacaToken;
    address public PancakeFactory;

    // Pancake router v2
    address public PancakeRouter;

    // Authorization
    address public operator;
    uint256 public producedFee;

    // Campaign properties
    bool public terminated;
    bool public launched;
    uint256 public id;
    string public campaignName;
    address public currency;
    uint256 public target;
    uint256 public joinWindow;
    uint256 public duration;
    uint256 public nextCycleAt;
    uint256 public actualStartAt;
    uint256 public actualEndAt;
    uint256 public cycle;

    // This is only for those joins which are within the same cycle as termination.
    mapping(uint256 => JoinDuring[]) trancheJoinList;

    // Tranche
    // There can only be one junior tranche.
    // The last tranche should always be junior.
    Tranche[] public tranches;
    mapping(address => mapping(uint256 => Account)) public accounts;
    mapping(uint256 => mapping(uint256 => uint256))
        public exchangeRateCheckpoint;
    mapping(uint256 => uint256) public trancheExitBuf;
    mapping(uint256 => uint256) public trancheJoinBuf;

    // tranche id => fee
    mapping(uint256 => uint256) public trancheFeeInf;

    // profolio farms
    // farms: 0->cream;1->venus;2-> alpaca
    mapping(uint256 => address) public farms;
    mapping(uint256 => uint256) public farmShares;

    modifier checkTranches {
        require(tranches.length > 1, "tranches is incomplete");
        require(
            tranches[tranches.length - 1].apy == 0,
            "the last tranche must carry zero apy"
        );
        _;
    }

    modifier checkTrancheID(uint256 tid) {
        require(tid < tranches.length, "invalid tranche id");
        _;
    }

    modifier requireLaunched {
        require(launched, "not launched");
        _;
    }

    modifier notLaunched {
        require(!launched, "launched");
        _;
    }

    modifier notTerminated {
        require(!terminated, "already terminated");
        _;
    }

    modifier requireTerminated {
        require(terminated, "not terminated");
        _;
    }

    modifier onlyOperator {
        require(msg.sender == operator, "not operator");
        _;
    }

    event Join(
        address indexed user,
        uint256 indexed trancheID,
        uint256 cycle,
        uint256 amount,
        uint256 optionBalance
    );
    event ApplyForExit(
        address indexed user,
        uint256 indexed trancheID,
        uint256 cycle,
        uint256 optionBalance
    );
    event Exit(
        address indexed user,
        uint256 indexed trancheID,
        uint256 cycle,
        uint256 amount
    );
    event NewTranche(uint256 id, uint256 target, address token, address reward);
    event TrancheSettlement(
        uint256 indexed trancheID,
        uint256 indexed cycle,
        uint256 principal,
        uint256 capital,
        uint256 exchangerate
    );
    event NextCycle(uint256 cycle, uint256 totalP, uint256 totalC);
    event Terminate(uint256 cycle);
    event WithdrawFee(address operator, uint256 amount);

    constructor(
        address _operator,
        uint256 _id,
        string memory _campaignName,
        address _WTF,
        address _rewardFactory,
        address _trancheTokenFactory,
        address _currency,
        uint256 _target,
        uint256 _duration,
        uint256 _joinwindow
    ) public {
        operator = _operator;
        terminated = false;
        launched = false;

        id = _id;
        campaignName = _campaignName;
        currency = _currency;
        target = _target;
        duration = _duration;
        joinWindow = _joinwindow;
        WTF = _WTF;
        RewardFactory = _rewardFactory;
        TrancheTokenFactory = _trancheTokenFactory;

        producedFee = 0;

        // default addresses on BSC mainnet
        CreamFarm = 0xb316F4F692d3bc53B79C83c97fDD45bC94255F53;
        VenusFarm = 0x95c78222B3D6e262426483D42CfA53685A67Ab9D;
        VenusToken = 0xcF6BB5389c92Bdda8a3747Ddb454cB7a64626C63;
        AlpacaFarm = 0x7C9e73d4C71dae564d41F78d56439bB4ba87592f;
        AlpacaToken = 0x8F0528cE5eF7B51152A59745bEfDD91D97091d2F;
        PancakeFactory = 0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73;
        PancakeRouter = 0x10ED43C718714eb63d5aA57B78B54704E256024E;

        farms[0] = CreamFarm;
        farms[1] = VenusFarm;
        farms[2] = AlpacaFarm;
    }

    // override all tranches, so called with caution
    function setTranches(TrancheParams[] memory _tranches) public onlyOperator {
        require(
            _tranches[_tranches.length - 1].apy == 0,
            "the last tranche must carry zero apy"
        );
        for (uint256 i = 0; i < _tranches.length; i++) {
            address tt = ITrancheTokenFactory(TrancheTokenFactory).deploy(
                "WTF Tranche",
                "WTFT"
            );
            address tr = IRewardFactory(RewardFactory).deploy(
                operator,
                WTF,
                tt
            );
            tranches.push(
                Tranche({
                    target: target.mul(_tranches[i].percentage).div(
                        PercentageParamScale
                    ),
                    percentage: _tranches[i]
                        .percentage
                        .mul(PercentageScale)
                        .div(PercentageParamScale),
                    apy: _tranches[i].apy.mul(PercentageScale).div(
                        PercentageParamScale
                    ),
                    fee: _tranches[i].fee,
                    capital: 0,
                    principal: 0,
                    token: tt,
                    reward: tr
                })
            );
            trancheFeeInf[i] = _tranches[i].fee;
            emit NewTranche(
                i,
                tranches[i].target,
                tranches[i].token,
                tranches[i].reward
            );
        }
    }

    function _resetTranche(uint256 trancheID) internal {
        tranches[trancheID].capital = 0;
        tranches[trancheID].principal = 0;
        tranches[trancheID].token = ITrancheTokenFactory(TrancheTokenFactory)
        .deploy("WTF Tranche", "WTFT");
        tranches[trancheID].reward = IRewardFactory(RewardFactory).deploy(
            operator,
            WTF,
            tranches[trancheID].token
        );
    }

    function getTrancheTokenRewardAddresses(uint256 trancheID)
        public
        view
        returns (address token, address reward)
    {
        return (tranches[trancheID].token, tranches[trancheID].reward);
    }

    function tranchesTotal()
        public
        view
        returns (uint256 principal, uint256 capital)
    {
        for (uint256 i = 0; i < tranches.length; i++) {
            principal += tranches[i].principal;
            capital += tranches[i].capital;
        }
        return (principal, capital);
    }

    function changeOperator(address o) public {
        require(
            msg.sender == owner() || msg.sender == operator,
            "neither owner nor operator"
        );
        operator = o;
    }

    function setAlpacaToken(address a) public onlyOperator {
        AlpacaToken = a;
    }

    function setVenusToken(address a) public onlyOperator {
        VenusToken = a;
    }

    function setTrancheAPY(uint256 trancheID, uint256 apy) public onlyOperator {
        require(apy < PercentageParamScale);
        tranches[trancheID].apy = apy.mul(PercentageScale).div(
            PercentageParamScale
        );
    }

    function setFarms(
        address cream,
        uint256 s1,
        address venus,
        uint256 s2,
        address alpaca,
        uint256 s3
    ) public onlyOperator {
        CreamFarm = cream;
        VenusFarm = venus;
        AlpacaFarm = alpaca;
        farms[0] = cream;
        farms[1] = venus;
        farms[2] = alpaca;
        farmShares[0] = s1.mul(PercentageScale).div(PercentageParamScale);
        farmShares[1] = s2.mul(PercentageScale).div(PercentageParamScale);
        farmShares[2] = s3.mul(PercentageScale).div(PercentageParamScale);
    }

    function setPancakeRouter(address a) public onlyOperator {
        PancakeRouter = a;
    }

    // collect fee set
    function setTranchesFee(uint256 trancheID, uint256 fee)
        public
        onlyOperator
    {
        trancheFeeInf[trancheID] = fee;
    }

    function setCampaignTarget(uint256 _target) public onlyOperator {
        target = _target;
        for (uint256 i = 0; i < tranches.length; i++) {
            tranches[i].target = tranches[i].percentage.mul(target);
        }
    }

    function setTrancheTarget(uint256 trancheID, uint256 t)
        public
        onlyOperator
    {
        tranches[trancheID].target = t;
    }

    // Don't ever touch accounts in a universal transition.
    // Otherwise you get DEFI wrong.
    function nextCycle() public onlyOperator notTerminated {
        require(
            nextCycleAt == 0 || block.timestamp >= nextCycleAt,
            "new cycle too early"
        );

        uint256 p;
        uint256 c;
        uint256 _totalCapital = 0;
        if (cycle == 0) {
            _processJoins();
        } else {
            (p, c) = tranchesTotal();
            if (c == 0) {
                revert("empty pocket");
            }
            _totalCapital = _redeem();
            _cycleExchangeRateCheckpoint(_totalCapital);
            _processExits();
            _processJoins();
        }

        actualStartAt = block.timestamp;
        nextCycleAt = block.timestamp + duration;
        cycle++;

        _launch();

        (p, c) = tranchesTotal();
        emit NextCycle(cycle, p, c);
    }

    // operator without fee from contract
    function withdrawFee(uint256 amount) public onlyOperator {
        require(amount <= producedFee, "not enough balance for fee");
        producedFee = producedFee.sub(amount);
        IERC20(currency).safeTransfer(msg.sender, amount);
        emit WithdrawFee(msg.sender, amount);
    }

    function _cycleExchangeRateCheckpoint(uint256 totalCapital) internal {
        uint256 restCapital = totalCapital;
        // accounting for each tranche,
        // this is where tranches' status to be updated.
        uint256 interestShouldBe;
        uint256 cycleExchangeRate;

        // senior capital
        for (uint256 i = 0; i < tranches.length - 1; i++) {
            Tranche storage senior = tranches[i];
            if (restCapital > 0) {
                // calculate tranche interest from last cycle.
                interestShouldBe = senior
                .capital
                .mul(senior.apy)
                .mul(block.timestamp - actualStartAt)
                .div(365)
                .div(86400)
                .div(PercentageScale);

                bool chargeFee = true;
                if (restCapital < senior.capital) {
                    chargeFee = false;
                }

                // calculate the new capital for this tranche.
                if (restCapital <= senior.capital + interestShouldBe) {
                    senior.capital = restCapital;
                    restCapital = 0;
                } else {
                    senior.capital += interestShouldBe;
                    restCapital = restCapital.sub(senior.capital);
                }

                if (chargeFee) {
                    // collect fee
                    uint256 fee = senior.capital.mul(trancheFeeInf[i]).div(
                        PercentageParamScale
                    );
                    producedFee = producedFee.add(fee);
                    senior.capital = senior.capital.sub(fee);
                }

                cycleExchangeRate = _calculateExchangeRate(
                    senior.capital,
                    senior.principal
                );
                exchangeRateCheckpoint[cycle][i] = cycleExchangeRate;

                emit TrancheSettlement(
                    i,
                    cycle,
                    senior.principal,
                    senior.capital,
                    cycleExchangeRate
                );
            } else {
                // forced clear
                _resetTranche(i);
                exchangeRateCheckpoint[cycle][i] = 0;

                emit NewTranche(
                    i,
                    tranches[i].target,
                    tranches[i].token,
                    tranches[i].reward
                );
            }
        }

        // junior tranche should be treated separately
        uint256 juniorIndex = tranches.length - 1;
        Tranche storage junior = tranches[juniorIndex];
        if (restCapital > 0) {
            junior.capital = restCapital;
            uint256 fee = junior.capital.mul(trancheFeeInf[juniorIndex]).div(
                PercentageParamScale
            );
            producedFee = producedFee.add(fee);
            junior.capital = junior.capital.sub(fee);
            cycleExchangeRate = _calculateExchangeRate(
                junior.capital,
                junior.principal
            );
            exchangeRateCheckpoint[cycle][juniorIndex] = cycleExchangeRate;

            emit TrancheSettlement(
                juniorIndex,
                cycle,
                junior.principal,
                junior.capital,
                cycleExchangeRate
            );
        } else {
            _resetTranche(juniorIndex);
            exchangeRateCheckpoint[cycle][juniorIndex] = 0;

            emit NewTranche(
                juniorIndex,
                tranches[juniorIndex].target,
                tranches[juniorIndex].token,
                tranches[juniorIndex].reward
            );
        }
    }

    // Must be called after function _cycleExchangeRateCheckpoint.
    // At the end of each cycle, processing exit is where most of the works are.
    // In processing exit, we update every thing about a tranch except for new joins.
    function _processExits() internal {
        uint256 cycleExchangeRate;
        uint256 principalToDecrease;
        // senior and junior capital
        for (uint256 i = 0; i < tranches.length - 1; i++) {
            Tranche storage senior = tranches[i];
            cycleExchangeRate = exchangeRateCheckpoint[cycle][i];

            if (cycleExchangeRate > 0) {
                principalToDecrease = _fromOptionToCurrency(
                    trancheExitBuf[i],
                    cycleExchangeRate
                );
                require(
                    IERC20(senior.token).totalSupply() >= trancheExitBuf[i],
                    "too much exit"
                );
                if (IERC20(senior.token).totalSupply() == trancheExitBuf[i]) {
                    // normal full exit
                    senior.principal = 0;
                    senior.capital = 0;
                } else {
                    senior.principal = senior.principal.sub(
                        principalToDecrease
                    );
                    senior.capital = senior.capital.sub(principalToDecrease);
                }
            }
        }

        uint256 juniorIndex = tranches.length - 1;
        Tranche storage junior = tranches[juniorIndex];
        cycleExchangeRate = exchangeRateCheckpoint[cycle][juniorIndex];
        if (cycleExchangeRate > 0) {
            principalToDecrease = _fromOptionToCurrency(
                trancheExitBuf[juniorIndex],
                cycleExchangeRate
            );
            if (
                IERC20(junior.token).totalSupply() ==
                trancheExitBuf[juniorIndex]
            ) {
                junior.capital = 0;
                junior.principal = 0;
            } else {
                junior.principal = junior.principal.sub(principalToDecrease);
                junior.capital = junior.capital.sub(principalToDecrease);
            }
        }

        for (uint256 i = 0; i < tranches.length; i++) {
            trancheExitBuf[i] = 0;
        }
    }

    // Just move buffered total joins for this cycle into corresponding tranche.
    function _processJoins() internal {
        for (uint256 i = 0; i < tranches.length; i++) {
            tranches[i].principal += trancheJoinBuf[i];
            tranches[i].capital += trancheJoinBuf[i];
            trancheJoinBuf[i] = 0;

            // empty the join list every new cycle
            delete (trancheJoinList[i]);
        }
    }

    // Just in case I mess up with mul and div.
    function _fromOptionToCurrency(uint256 balance, uint256 rate)
        internal
        pure
        returns (uint256)
    {
        return balance.mul(rate).div(PercentageScale);
    }

    // Just in case I mess up with mul and div.
    function _fromCurrencyToOption(uint256 amount, uint256 rate)
        internal
        pure
        returns (uint256)
    {
        return rate == 0 ? 0 : amount.mul(PercentageScale).div(rate);
    }

    function _calculateExchangeRate(uint256 current, uint256 base)
        internal
        pure
        returns (uint256)
    {
        if (current == base) {
            return PercentageScale;
        } else if (current > base) {
            return
                PercentageScale.add(
                    (current - base).mul(PercentageScale).div(base)
                );
        } else {
            return
                PercentageScale.sub(
                    (base - current).mul(PercentageScale).div(base)
                );
        }
    }

    // Totaly stop the protocol, and there is no turning back.
    // All assets should be returned from farming.
    function terminate() public onlyOperator requireLaunched {
        (uint256 p, uint256 c) = tranchesTotal();
        if (c == 0) {
            revert("empty pocket");
        }
        uint256 _totalCapital = _redeem();
        _cycleExchangeRateCheckpoint(_totalCapital);

        // exit applies with in this cycle should be removed.
        // joins should be returned without process.
        // account's exit apply will be removed when calling exitAfterTerminate.
        for (uint256 i = 0; i < tranches.length; i++) {
            trancheExitBuf[i] = 0;
            trancheJoinBuf[i] = 0;
        }

        actualEndAt = block.timestamp;
        terminated = true;
        launched = false;

        emit Terminate(cycle);
    }

    // Call join will make amount of token been locked up in the protocol for the rest of current cycle with no gains.
    // But from the next cycle on, the amount of your token will start farming.
    function join(uint256 trancheID, uint256 amount)
        public
        notTerminated
        checkTranches
        checkTrancheID(trancheID)
    {
        require(
            actualStartAt == 0 ||
                actualStartAt + duration - block.timestamp < joinWindow,
            "missing the join window"
        );
        Tranche storage trch = tranches[trancheID];
        require(trch.target >= trch.principal.add(amount), "not enough quota");
        Account storage account = accounts[msg.sender][trancheID];
        account.joinCycle = cycle;
        uint256 optionBalance = 0;
        if (cycle == 0) {
            optionBalance = amount;
        } else {
            uint256 cycleExchangeRate = exchangeRateCheckpoint[cycle - 1][
                trancheID
            ];
            if (cycleExchangeRate == 0) {
                optionBalance = amount;
            } else {
                optionBalance = _fromCurrencyToOption(
                    amount,
                    cycleExchangeRate
                );
            }
        }
        account.investment += amount;
        trancheJoinBuf[trancheID] += amount;
        IERC20(currency).safeTransferFrom(msg.sender, address(this), amount);

        // IERC20MintBurn(trch.token).mint(msg.sender, optionBalance);
        IERC20MintBurn(trch.token).mint(address(this), optionBalance);
        IERC20(trch.token).approve(trch.reward, optionBalance);
        ITrancheReward(trch.reward).stakeFor(msg.sender, optionBalance);

        if (launched) {
            trancheJoinList[trancheID].push(
                JoinDuring({
                    addr: msg.sender,
                    amount: amount,
                    opt: optionBalance,
                    evac: false
                })
            );
        }

        emit Join(msg.sender, trancheID, cycle, amount, optionBalance);
    }

    // New exit application will be turned down because of former application.
    function applyForExit(uint256 trancheID) public checkTrancheID(trancheID) {
        Tranche storage trch = tranches[trancheID];
        Account storage account = accounts[msg.sender][trancheID];
        require(
            account.exitCheckpoint[trancheID].token == address(0),
            "pending exit application"
        );
        uint256 trancheTokenBalance = IERC20(trch.token).balanceOf(msg.sender);
        trancheTokenBalance = trancheTokenBalance.add(
            IERC20(trch.reward).balanceOf(msg.sender)
        );
        require(trancheTokenBalance > 0, "not an investor");
        require(cycle >= account.joinCycle, "apply exit from next cycle");
        account.exitCheckpoint[trancheID] = TrancheExit({
            token: trch.token,
            cycle: cycle,
            balance: trancheTokenBalance,
            investment: account.investment
        });
        trancheExitBuf[trancheID] += trancheTokenBalance;

        // logic lock the tranche token balance of msg.sender

        emit ApplyForExit(msg.sender, trancheID, cycle, trancheTokenBalance);
    }

    // Whether apply for exit
    function isApplyExit(uint256 trancheID) public view returns (bool) {
        Account storage account = accounts[msg.sender][trancheID];
        if (account.exitCheckpoint[trancheID].token != address(0)) {
            return true;
        } else {
            return false;
        }
    }

    // Exit and take all back if already applied in previous cycles.
    // Everything needed for exit is already set in application,
    // so we don't access tranche info in exit anymore.
    function exit(uint256 trancheID) public checkTrancheID(trancheID) {
        Tranche storage trch = tranches[trancheID];
        Account storage account = accounts[msg.sender][trancheID];
        TrancheExit storage exitpoint = account.exitCheckpoint[trancheID];
        require(exitpoint.token != address(0), "no exit applies");
        uint256 exitCycle = account.exitCheckpoint[trancheID].cycle;
        require(
            exitCycle > 0 && cycle > exitCycle,
            "allow to exit from next cycle"
        );
        uint256 trancheTokenBalance = IERC20(trch.token).balanceOf(msg.sender);
        if (trancheTokenBalance < exitpoint.balance) {
            ITrancheReward(trch.reward).withdrawFor(
                msg.sender,
                exitpoint.balance - trancheTokenBalance
            );
        }
        require(IERC20(trch.token).balanceOf(msg.sender) >= exitpoint.balance);

        uint256 exchangeRate = exchangeRateCheckpoint[exitpoint.cycle][
            trancheID
        ];
        uint256 exitInvestment = _fromOptionToCurrency(
            exitpoint.balance,
            exchangeRate
        );
        IERC20MintBurn(exitpoint.token).burnFrom(msg.sender, exitpoint.balance);
        if (exitInvestment > 0) {
            IERC20(currency).safeTransfer(msg.sender, exitInvestment);
        }

        account.investment -= exitpoint.investment;
        delete account.exitCheckpoint[trancheID];

        emit Exit(msg.sender, trancheID, exitCycle, exitInvestment);
    }

    // so if the msg.sender has an exit application before the termination,
    // we don't allow this user just exiting directly now.
    // he must finished the former exit apply first.
    function exitAfterTerminated(uint256 trancheID)
        public
        checkTrancheID(trancheID)
        requireTerminated
    {
        Tranche storage trch = tranches[trancheID];
        uint256 trancheTokenBalance = IERC20(trch.token).balanceOf(msg.sender);
        uint256 trancheRewardBalance = IERC20(trch.reward).balanceOf(
            msg.sender
        );
        require(
            trancheTokenBalance + trancheRewardBalance > 0,
            "not in this tranche"
        );
        Account storage account = accounts[msg.sender][trancheID];
        require(account.investment > 0, "account has no investment");
        TrancheExit storage exitpoint = account.exitCheckpoint[trancheID];

        // exit all of user's tranche token from the tranche reward pool
        ITrancheReward(trch.reward).exitFor(msg.sender);

        // joins yet to take effects
        for (uint256 i = 0; i < trancheJoinList[trancheID].length; i++) {
            if (
                trancheJoinList[trancheID][i].addr == msg.sender &&
                trancheJoinList[trancheID][i].evac == false
            ) {
                trancheJoinList[trancheID][i].evac = true;
                IERC20MintBurn(trch.token).burnFrom(
                    msg.sender,
                    trancheJoinList[trancheID][i].opt
                );
                IERC20(currency).transfer(
                    msg.sender,
                    trancheJoinList[trancheID][i].amount
                );
            }
        }
        // must update the tranche token balance, we don't wanna steal from our beloved users.
        trancheTokenBalance = IERC20(trch.token).balanceOf(msg.sender);

        require(
            (exitpoint.token != address(0) && exitpoint.cycle == cycle) ||
                (exitpoint.token == address(0)),
            "account has unfinished exit apply"
        );

        uint256 exchangeRate;
        uint256 exitInvestment;
        if (trancheTokenBalance > 0) {
            exchangeRate = exchangeRateCheckpoint[cycle][trancheID];
            exitInvestment = _fromOptionToCurrency(
                trancheTokenBalance,
                exchangeRate
            );
            IERC20MintBurn(trch.token).burnFrom(
                msg.sender,
                trancheTokenBalance
            );
            account.investment = 0;
            if (exitInvestment > 0) {
                IERC20(currency).safeTransfer(msg.sender, exitInvestment);
            }
        }

        delete account.exitCheckpoint[trancheID];
        delete accounts[msg.sender][trancheID];

        emit Exit(msg.sender, trancheID, cycle, exitInvestment);
    }

    function _getCurrentTotalCapital() internal view returns (uint256) {
        uint256 total = 0;
        for (uint256 i = 0; i < tranches.length; i++) {
            total += tranches[i].capital;
        }
        return total;
    }

    // At the begining of each cycle, stake capital into 3rd farms.
    // If total capital is 0, _launch will just silently shortcut 3rd-farm-operations.
    // The reason we don't throw an error is that we might need leave an empty cycle
    // for users to exit their former investments.
    function _launch() internal {
        if (!launched) {
            launched = true;
        }

        uint256 totalCapital = _getCurrentTotalCapital();
        if (totalCapital == 0) {
            return;
        }
        uint256 used = 0;
        uint256 amount = 0;
        // cream mint
        amount = totalCapital.mul(farmShares[0]).div(PercentageScale);
        IERC20(currency).approve(farms[0], amount);
        ICompoundLike(farms[0]).mint(amount);
        used += amount;

        // venus mint
        amount = totalCapital.mul(farmShares[1]).div(PercentageScale);
        IERC20(currency).approve(farms[1], amount);
        ICompoundLike(farms[1]).mint(amount);
        used += amount;

        // alpaca deposit
        amount = totalCapital - used;
        IERC20(currency).approve(farms[2], amount);
        IAlpacaLike(farms[2]).deposit(amount);
    }

    // redeem from 3 farms
    function _redeem() internal returns (uint256) {
        uint256 formerBalance = IERC20(currency).balanceOf(address(this));
        uint256 alpacaBalance = IERC20(AlpacaToken).balanceOf(address(this));
        uint256 venusBalance = IERC20(VenusToken).balanceOf(address(this));

        uint256 share = IERC20(farms[0]).balanceOf(address(this));
        if (share > 0) {
            ICompoundLike(farms[0]).redeem(share);
        }

        share = IERC20(farms[1]).balanceOf(address(this));
        if (share > 0) {
            ICompoundLike(farms[1]).redeem(share);
        }

        share = IERC20(farms[2]).balanceOf(address(this));
        if (share > 0) {
            IAlpacaLike(farms[2]).withdraw(share);
        }
        share = IERC20(AlpacaToken).balanceOf(address(this));
        if (share.sub(alpacaBalance) > 0) {
            _swapForBUSDFromPancake(
                AlpacaToken,
                currency,
                share.sub(alpacaBalance)
            );
        }

        share = IERC20(VenusToken).balanceOf(address(this));
        if (share.sub(venusBalance) > 0) {
            _swapForBUSDFromPancake(
                VenusToken,
                currency,
                share.sub(venusBalance)
            );
        }
        return IERC20(currency).balanceOf(address(this)).sub(formerBalance);
    }

    // swap all alpaca to BUSD
    // venue and alpaca
    function _swapForBUSDFromPancake(
        address _fromToken,
        address _toToken,
        uint256 amount
    ) private returns (uint256) {
        address[] memory swapPath2 = new address[](2);
        swapPath2[0] = _fromToken;
        swapPath2[1] = _toToken;
        uint256[] memory busdAmount = IUniswapV2Router02(PancakeRouter)
        .swapExactTokensForTokens(
            amount,
            1,
            swapPath2,
            address(this),
            block.timestamp + 200
        );
        return busdAmount[busdAmount.length - 1];
    }
}
