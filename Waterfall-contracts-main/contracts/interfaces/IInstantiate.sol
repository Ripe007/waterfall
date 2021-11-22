// SPDX-License-Identifier: MIT

pragma solidity >=0.4.18 <=0.6.12;

interface IInstantiate {
    function instantiate(bytes calldata initdata) external returns (address);
}
