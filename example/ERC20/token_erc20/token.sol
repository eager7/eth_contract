pragma solidity >=0.4.21 <0.6.0;

interface tokenRecipient {
    function receiveApproval(address _from, uint256 _value, address _token, bytes calldata _extraData) external;
}

contract Token {
    address public owner;
    uint public last_completed_token;

    string public name;
    string public symbol;
    uint8 constant public decimals = 18;
    uint256 public totalSupply;

    mapping (address => uint256) public _balances;
    mapping (address => mapping(address => uint256)) public _allowance;

    constructor(string memory tokenName, string memory tokenSymbol, uint256 initSupply) public {
        owner = msg.sender;
        name = tokenName;
        symbol = tokenSymbol;
        totalSupply = initSupply * 10**uint256(decimals);
    }

    modifier restricted() {
        if (msg.sender == owner) _;
    }

    function setCompleted(uint completed) public restricted {
        last_completed_token = completed;
    }

    function upgrade(address new_address) public restricted {
        Token upgraded = Token(new_address);
        upgraded.setCompleted(last_completed_token);
    }

    event Transfer(address indexed _from, address indexed _to, uint256 _value);
    event Approval(address indexed _owner, address indexed _spender, uint256 _value);
    event Burn(address indexed _from, uint256 _value);

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return _balances[_owner];
    }
    function transfer(address _to, uint256 _value) public returns (bool success){
        require(_to != address (0x0));
        require(_balances[msg.sender] >= _value);

        _balances[msg.sender] -= _value;
        _balances[_to] += _value;

        emit Transfer(msg.sender, _to, _value);

        return true;
    }
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        require(_value < _balances[_from]);
        require(_value < _allowance[_from][msg.sender]);
        require(_to != address (0x0));

        _balances[_from] -= _value;
        _balances[_to] += _value;
        _allowance[_from][msg.sender] -= _value;

        emit Transfer(_from, _to, _value);


    }
    function approve(address _spender, uint256 _value) public returns (bool success) {
        _allowance[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }
    function allowance(address _owner, address _spender) public view returns (uint256) {
        return _allowance[_owner][_spender];
    }

    function approveAndCall(address _spender, uint256 _value, bytes memory _extraData) public returns (bool success) {
        tokenRecipient spender = tokenRecipient(_spender);
        if (approve(_spender, _value)) {
            spender.receiveApproval(msg.sender, _value, address(this), _extraData);
            return true;
        }
        return false;
    }
    function burn(uint256 _value) public returns (bool success) {
        require(_balances[msg.sender] >= _value);

        _balances[msg.sender] -= _value;
        totalSupply -= _value;

        emit Burn(msg.sender, _value);

        return true;
    }
    function burnFrom(address _from, uint256 _value) public returns (bool success) {
        require(_balances[_from] >= _value);
        require(_value <= _allowance[_from][msg.sender]);

        _balances[_from] -= _value;
        _allowance[_from][msg.sender] -= _value;
        totalSupply -= _value;

        emit Burn(_from, _value);

        return true;
    }
}
