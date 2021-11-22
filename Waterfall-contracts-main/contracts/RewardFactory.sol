//SPDX-License-Identifier: MIT
pragma solidity >=0.4.18 <=0.6.12;

import "./Reward.sol";

contract RewardFactory {
    function deploy(
        address admin,
        address rewardToken,
        address stakeToken
    ) public returns (address) {
        Reward r = new Reward(admin, rewardToken, stakeToken);
        r.transferOwnership(msg.sender);
        return address(r);
    }
}
