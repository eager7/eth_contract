pragma solidity ^0.5.0;

contract plainchant {
    string public name = "plainchant";
    string public symbol = "PCT";
    uint8 constant public decimals = 18;

    mapping (address => uint256) public _balances;

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    event Burn(address indexed _from, uint256 _value);

    function transfer(address _to, uint256 _value) public returns (bool success) {
        require(_to != address(0x0));
        require(_balances[_to] >= _value);

        _balances[msg.sender] -= _value;
        _balances[_to] += _value;

        emit Transfer(msg.sender, _to, _value);

        return true;
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return _balances[_owner];
    }


}
