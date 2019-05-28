pragma solidity ^0.5.0;

contract deposit {
    address public owner;

    event Deposit(address indexed _from, uint256 indexed _value);
    event Withdrawal(address indexed _to, uint256 indexed _value);

    constructor() public {
        owner = msg.sender;
    }

    function withdrawal(uint256 _value) public {
        if (_value < address(this).balance) {
            msg.sender.transfer(_value);
        }
        emit Withdrawal(msg.sender, _value);
    }

    function() payable external {
        if (msg.value > 0){
            emit Deposit(msg.sender, msg.value);
        }
    }

    function Balance() public view returns(uint256) {
        return address(this).balance;
    }
}
