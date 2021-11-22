## Waterfall Smart Contract

Table of Contents
=================

   * [Waterfall Smart Contract](#waterfall-smart-contract)
     * [1. Code](#1-code)
     * [2. Address](#2-address)
     * [3. Events](#3-events)
       * [Join](#join)
       * [ApplyForExit](#applyforexit)
       * [NewTranche](#newtranche)
       * [TrancheSettlement](#tranchesettlement)
       * [NextCycle](#nextcycle)
       * [WithDrawFee](#withdrawfee)
       * [Terminate](#terminate)
     * [4. Read-Only Functions](#4-read-only-functions)
       * [Below params are public state variable](#below-params-are-public-state-variable)
       * [tranchesTotal](#tranchestotal)
       * [isAppayExit](#isappayexit)
     * [5. Write-Only Functions](#5-write-only-functions)
       * [setFarms](#setfarms)
       * [setPancakeRouter](#setpancakerouter)
       * [setTranches](#settranches)
       * [setTranchesFee](#settranchesfee)
       * [setCampaignTarget](#setcampaigntarget)
       * [nextCycle](#nextcycle-1)
       * [withdrawFee](#withdrawfee-1)
       * [terminate](#terminate-1)
       * [join](#join-1)
       * [applyForExit](#applyforexit-1)
       * [exit](#exit)
       * [exitAfterTerminated](#exitafterterminated)
     * [6. interface](#6-interface)
       * [Farm interface Information](#farm-interface-information)
       * [Other Interface Information](#other-interface-information)
     * [7. Test](#7-test)
       * [7.1 Environment](#71-environment)
       * [7.2 Test process](#72-test-process)


### 1. Code

```shell
CampaignContinuousCycles.sol
```

### 2. Address

BSC_TestNet has deployed our test contract, and his information you can find here:

[https://github.com/Waterfall-protocol/Waterfall-contracts/tree/feature/v2Test]: https://github.com/Waterfall-protocol/Waterfall-contracts/tree/feature/v2Test

|     Param     |                     Decription                     |
| :-----------: | :------------------------------------------------: |
| Contract path |     0xD5A258b370337245c631f5A589b3a4AB2B8d38E1     |
|     Path      |    v2Test/contract/CampaignContinusCyscleTest01    |
| contract name |            CampaignContinusCyscleTest01            |
|    network    |                   Biance_testnet                   |
| abi file path | v2Test/abis |
|               |                                                    |

**Table 1: Waterfall information** 

Note: current contract is not verity and public!

### 3. Events

#### Join

Users joint their money well trigger this trigger

```
event Join(
    address indexed user,
    uint256 indexed trancheID,
    uint256 cycle,
    uint256 amount,
    uint256 optionBalance
);
```

- **user** is user address;
- **trancheID**: is tranche(role) id. The id is start 0;
- **cycle**:  campaign cycle now;
- **amount** : invest token amount;
- **optionBalance**: invest amount can received;

#### ApplyForExit

user Apply exit campaign will trigger this event;

```
event ApplyForExit(
    address indexed user,
    uint256 indexed trancheID,
    uint256 cycle,
    uint256 optionBalance
);
```

- **user**: user address;
- **trancheID**: wants to apply exit tranche id;
- **cycle**: apply exit cycle ;
- **optionBalance** : can get Token in this cycle;

#### NewTranche

```
 event NewTranche(uint256 id, uint256 target, address token);
 
```

- **id** :Tranche id 
- **target**:new tranche target invest amount
- **token**: BUSD address

#### TrancheSettlement 

```
 event TrancheSettlement(
        uint256 indexed trancheID,
        uint256 indexed cycle,
        uint256 principal,
        uint256 capital,
        uint256 exchangerate
    );
```

- **trancheID**: tranche id
- **cycle**: current cycle 
- **principal**: current token total amount
- **capital**: current invested amount
- **exchangerate**: two token exchangerate

#### NextCycle

next cycle event

```
event NextCycle(uint256 cycle);
```

watcht it , you can get  the number of cycle;

- ***cycle**: cycle number now;

#### WithDrawFee

operator withdraw her/his amount will trigger this event;

```
event WithdrawFee(address operator, uint256 amount);
```

- **operator**: operator address;
- **amount**: operator withdraw fee ;

#### Terminate

The campaign stop event

```
event Terminate(uint256 cycle);
```

- **cycle**:  stop campaign cycle

### 4. Read-Only Functions

#### Below params are public state variable

you can call them by web3.js and other types.

```
    address public CreamFarm;   // cream farm address
    address public VenusFarm;   // venus farm address
    address public VenusToken;  // venus token address
    address public AlpacaFarm;	// alpaca farm address
    address public AlpacaToken; // alpaca token address
    address public PancakeFactory; // pacake factory address
    // Pancake router v2 address
    address public PancakeRouter;

    // Authorization
    address public operator;   // current contract operator
    uint256 public producedFee;// campaign platform fee

    // Campaign properties
    bool public terminated;    // campaign whether stop
    bool public launched;      // campaign whether start
    uint256 public id;  	   // campaign id
    string public campaignName;// campaign name
    address public currency;   // BUSD address
    uint256 public target;   // campaign target amount
    uint256 public joinWindow; // user join campaign window period
    uint256 public duration;  // lock time 
    uint256 public nextCycleAt; // next cycle start time 
    uint256 public actualStartAt; // actual campaign start timestamp
    uint256 public actualEndAt; // actual campaign stop timestamp
    uint256 public cycle; // current cycle

    // This is only for those joins which are within the same cycle as termination.
    mapping(uint256 => JoinDuring[]) trancheJoinList;

    // Tranche
    // There can only be one junior tranche.
    // The last tranche should always be junior.
    Tranche[] public tranches; // tranche (senior,...,junior)
    // user account information: userAddress->trancheID->UserAccount
    mapping(address => mapping(uint256 => Account)) public accounts;
    // every cycle, different exchange rate
    // cycle->trancheID->exchangeRate
    mapping(uint256 => mapping(uint256 => uint256))
        public exchangeRateCheckpoint;
    // tranche exit amount
    // trancheID ->exitAmount
    mapping(uint256 => uint256) public trancheExitBuf;
    // tranche join amount
    // trancheID -> joinAmount
    mapping(uint256 => uint256) public trancheJoinBuf;
    // collext fee set
    // tranche id => fee
    mapping(uint256 => uint256) public trancheFeeInf;

    // porfolio farms
    // farms: 0->cream;1->venus;2-> alpaca
    mapping(uint256 => address) public farms;
    mapping(uint256 => uint256) public farmShares;


```



#### tranchesTotal

call the function you can get all tranche principal  and  capital

```
function tranchesTotal() public view returns (uint256 principal, uint256 capital)

```

**return**

current campaign principal and capital

#### isAppayExit

this is  a public function can get an account whether apply exit

```
function isApplyExit(uint256 trancheID) public view returns (bool) {
        Account storage account = accounts[msg.sender][trancheID];
        if (account.exitCheckpoint[trancheID].token != address(0)) {
            return true;
        } else {
            return false;
        }
    }

```



### 5. Write-Only Functions

#### setFarms 

NOTE: The function should be called after calling constructor.

set the farm address , id and fee, and only Operator can call via the function.

**Note** must be careful , id ->address

```
function setFarms(
        address cream,
        uint256 s1,
        address venus,
        uint256 s2,
        address alpaca,
        uint256 s3
    ) public onlyOperator 
```

#### setPancakeRouter

set pancake router , if you want to change pancake router, and  only Operator can call via the function

```
function setPancakeRouter(address a) public onlyOperator
```

#### setTranches

```solidity
 function setTranches(TrancheParams[] memory _tranches) public onlyOperator 
```

this function can set tranches params like:

```solidity
 await camp.setTranches([
      { apy: 10000, fee: 33, percentage: 40000 },
      { apy: 30000, fee: 50, percentage: 30000 },
      { apy: 0, fee: 200, percentage: 30000 },
    ]);
```



#### setTranchesFee

set every tranche fee, and only Operator can call 

**Note** the function is one tranche id, one fee , if you call it twice , data will be override;

```
function setTranchesFee(uint256 trancheID, uint256 fee) public  onlyOperator
```

**Params**

- **trancheID**: is every tranche (senior, ..., junior) id;
- **fee**: ist every tranche fee. **MUST BE CAREFUL**: 33-->0.033% 

#### setCampaignTarget

this function will set new campaign invest target, and can call onlyOperator. All tranche target will be change.

```
function setCampaignTarget(uint256 _target) public onlyOperator 
```

**Params**

- **target**: set new campaign's target

#### nextCycle

After 3 window period, will be called by operator to invest all collect money into 3 farms, and call be called by operator

```
function nextCycle() public onlyOperator notTerminated 
```

#### withdrawFee

Operator withdraw his fee, and can be called by operator

```
function withdrawFee(uint256 amount) public onlyOperator
```

**params**

- **amount** withdraw amount

#### terminate 

stop the campaign, and  call be called by operator

```
function terminate() public onlyOperator
```

#### join

join function is a vital function. and can be called by user. If you want to invest your money, here is a entry.

```
 function join(uint256 trancheID, uint256 amount)
        public
        notTerminated
        checkTranches
        checkTrancheID(trancheID)

```

 **Param**

- **trancheID**: is user want to joint trancheID, and amount is invest the amount of money
- **amount**: is invest amount

#### applyForExit 

This is a apply for exit the campaign function, when you apply, if can exit in the next cycle start.

```solidity
function applyForExit(uint256 trancheID) public checkTrancheID(trancheID) 
```

**Param**

- **trancheID** is a your join tranche id.

#### exit 

Exit and take all back if already applied in previous cycles.  Everything needed for exit is already set in application, so we don't access tranche info in exit anymore.

```
exit(uint256 trancheID) public checkTrancheID(trancheID)
```

**Param**

- **trancheID** is a your join tranche id

#### exitAfterTerminated

so if the msg.sender has an exit application before the termination, we don't allow this user just exiting directly now. he must finished the former exit apply first.

New exit application will be turned down because of former application.

```
 function exitAfterTerminated(uint256 trancheID)
        public
        checkTrancheID(trancheID)
        requireTerminated
```

**Param**

- **trancheID** is a your join tranche id

### 6. interface

#### Farm interface Information

|     Name      | Description   |
| :-----------: | ------------- |
| Venus & Cream | ICompoundLike |
|    Alpaca     | IAlpacaLike   |

####  Other Interface Information

|       Name        | Description            |
| :---------------: | ---------------------- |
|        ERC        | IERC20.sol             |
|   UniswapV2Pair   | IUniswapV2Pair.sol     |
| UniswapV2Router02 | IUniswapV2Router02.sol |



### 7. Test

#### 7.1 Environment

- git  (clone contrac project (feature/v2));
- nodejs(npm) (install depend package);

#### 7.2 Test process

```
step 1: git clone git@github.com:Waterfall-protocol/Waterfall-contracts.git  (must checkout feature/v2) 
step 2: cd Waterfall-contracts && npm install 
step 3: npx hardhat node(open a terminal and input, like below) 
```

<img src="./images/Screenshot from 2021-07-28 13-57-38.png" style="zoom:50%;" />

```
step 4: npx hardhat test test/continuous-cycles.test.js (open other terminal input, like below)
```

<img src="./images/Screenshot from 2021-07-28 14-02-59.png" style="zoom:50%;" />

END.
