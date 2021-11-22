pragma solidity >=0.4.18 <=0.6.12;

interface ICakeToken {

    function transfer(address recipient, uint256 amount) external returns (bool);
       function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

}
