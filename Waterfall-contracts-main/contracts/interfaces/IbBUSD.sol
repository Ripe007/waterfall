pragma solidity >=0.4.18 <=0.6.12;

interface IbBUSD {
    function deposit(uint256 amountToken) external;

    function withdraw(uint256 share) external;
}
