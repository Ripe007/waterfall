pragma solidity ^0.6.12;

interface IGetThePrice {

    function getLatestPrice() external returns(int, int);
}