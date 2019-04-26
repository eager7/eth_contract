pragma solidity ^0.5.0;

contract deposit {
    address public owner;
    uint256 public balance;

    event Deposit(address indexed _from, uint256 indexed _value);
    event Withdrawal(address indexed _to, uint256 indexed _value);

    constructor() public {
        owner = msg.sender;
    }

    function withdrawal(address _to, uint256 _value) public {
        msg.sender.transfer(balance);
        emit Withdrawal(_to, _value);
    }

    function() payable external {
        if (msg.value > 0){
            balance += msg.value;
            emit Deposit(msg.sender, msg.value);
        }
    }
}
