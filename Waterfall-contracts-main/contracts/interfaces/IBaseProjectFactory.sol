// SPDX-License-Identifier: MIT

pragma solidity >=0.4.18 <=0.6.12;

interface IBaseProjectFactory {
    function instantiate(
        bytes32 _pid, // project id
        address _platform,
        address _recv,
        uint256 _collect_start,
        uint256 _collect_end
    )
        external
        returns (
            // uint256 _actual_collect,
            // uint256 _target_amount,
            // uint256 _profit_rate_senior,
            // uint256 _profit_rate_junior
            address
        );
}
