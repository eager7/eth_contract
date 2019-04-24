pragma solidity ^0.4.24;

contract ERC20Interface {
    function totalSupply() public constant returns (uint);
    function balanceOf(address tokenOwner) public constant returns (uint balance);
    function allowance(address tokenOwner, address spender) public constant returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens) public returns (bool success);
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}

contract SafeMath {
    function safeAdd(uint a, uint b) public pure returns (uint c) {
        c = a + b;
        require(c >= a);
    }
    function safeSub(uint a, uint b) public pure returns (uint c) {
        require(b <= a);
        c = a - b;
    }
    function safeMul(uint a, uint b) public pure returns (uint c) {
        c = a * b;
        require(a == 0 || c / a == b);
    }
    function safeDiv(uint a, uint b) public pure returns (uint c) {
        require(b > 0);
        c = a / b;
    }
}

contract _Base {
    address internal owner;
    address internal walletCenter;

    modifier onlyOwner {
        require(owner == msg.sender);
        _;
    }
    modifier onlyWallet(address _addr) {
        require(WalletCenter(walletCenter).isWallet(_addr));
        _;
    }
    modifier onlyShop(address _addr) {
        require(WalletCenter(walletCenter).isShop(_addr));
        _;
    }

    function balanceOf(address _erc20) public constant returns (uint256 balance) {
        if(_erc20==address(0))
            return address(this).balance;
        return ERC20Interface(_erc20).balanceOf(this);
    }

    function transfer(address _to, address _erc20, uint256 _value) internal returns (bool success) {
        require((_erc20==address(0)?address(this).balance:ERC20Interface(_erc20).balanceOf(this))>=_value);
        if(_erc20==address(0))
            _to.transfer(_value);
        else
            ERC20Interface(_erc20).approve(_to,_value);
        return true;
    }

    function withdrawal(address _erc20, uint256 _value) public returns (bool success);

    event Pay(address indexed _who, uint256 indexed _item, uint256 indexed _value);
    event Refund(address indexed _who, uint256 indexed _item, uint256 indexed _value);
    event Prize(address indexed _who, uint256 indexed _item, uint256 indexed _value);
}

contract _Wallet is _Base {
    constructor(address _who) public {
        owner           = _who;
        walletCenter    = msg.sender;
    }

    function pay(address _shop, uint256 _item) private {
        require(_Shop(_shop).canBuy(this,_item));

        address _erc20;
        uint256 _value;
        (_erc20,_value) = _Shop(_shop).price(_item);

        transfer(_shop,_erc20,_value);
        _Shop(_shop).pay(_item);
        emit Pay(_shop,_item,_value);
    }

    function paySafe(address _shop, uint256 _item) onlyOwner onlyShop(_shop) public payable returns (bool success) {
        pay(_shop,_item);
        return true;
    }
    function payUnsafe(address _shop, uint256 _item) onlyOwner public payable returns (bool success) {
        pay(_shop,_item);
        return true;
    }
    function payCancel(address _shop, uint256 _item) onlyOwner public returns (bool success) {
        _Shop(_shop).payCancel(_item);
        return true;
    }

    function refund(address _erc20, uint256 _item, uint256 _value) public payable returns (bool success) {
        require((_erc20==address(0)?msg.value:ERC20Interface(_erc20).allowance(msg.sender,this))==_value);
        if(_erc20!=address(0))
            ERC20Interface(_erc20).transferFrom(msg.sender,this,_value);
        emit Refund(msg.sender,_item,_value);
        return true;
    }
    function prize(address _erc20, uint256 _item, uint256 _value) public payable returns (bool success) {
        require((_erc20==address(0)?msg.value:ERC20Interface(_erc20).allowance(msg.sender,this))==_value);
        if(_erc20!=address(0))
            ERC20Interface(_erc20).transferFrom(msg.sender,this,_value);
        emit Prize(msg.sender,_item,_value);
        return true;
    }

    function withdrawal(address _erc20, uint256 _value) onlyOwner public returns (bool success) {
        require((_erc20==address(0)?address(this).balance:ERC20Interface(_erc20).balanceOf(this))>=_value);
        if(_erc20==address(0))
            owner.transfer(_value);
        else
            ERC20Interface(_erc20).transfer(owner,_value);
        return true;
    }
}

contract _Shop is _Base, SafeMath{
    address erc20;
    constructor(address _who, address _erc20) public {
        owner           = _who;
        walletCenter    = msg.sender;
        erc20           = _erc20;
    }

    struct item {
        uint8                       category;   // 0 = disable, 1 = non Stock, non Expire, 2 = can Expire (after 1 week), 3 = stackable
        uint256                     price;
        uint256                     stockCount;

        mapping(address=>uint256)   customer;
    }

    uint                    index;
    mapping(uint256=>item)  items;

    function pay(uint256 _item) onlyWallet(msg.sender) public payable returns (bool success) {
        require(canBuy(msg.sender, _item));
        require((erc20==address(0)?msg.value:ERC20Interface(erc20).allowance(msg.sender,this))==items[_item].price);

        if(erc20!=address(0))
            ERC20Interface(erc20).transferFrom(msg.sender,this,items[_item].price);

        if(items[_item].category==1 || items[_item].category==2 && now > safeAdd(items[_item].customer[msg.sender], 1 weeks))
            items[_item].customer[msg.sender]   = now;
        else if(items[_item].category==2 && now < safeAdd(items[_item].customer[msg.sender], 1 weeks) )
            items[_item].customer[msg.sender]   = safeAdd(items[_item].customer[msg.sender], 1 weeks);
        else if(items[_item].category==3) {
            items[_item].customer[msg.sender]   = safeAdd(items[_item].customer[msg.sender],1);
            items[_item].stockCount             = safeSub(items[_item].stockCount,1);
        }

        emit Pay(msg.sender,_item,items[_item].customer[msg.sender]);
        return true;
    }

    function payCancel(uint256 _item) onlyWallet(msg.sender) public returns (bool success) {
        require (items[_item].category==2&&safeAdd(items[_item].customer[msg.sender],2 weeks)>now&&balanceOf(erc20)>=items[_item].price);

        items[_item].customer[msg.sender]  = safeSub(items[_item].customer[msg.sender],1 weeks);
        transfer(msg.sender, erc20, items[_item].price);
        _Wallet(msg.sender).refund(erc20,_item,items[_item].price);
        emit Refund(msg.sender,_item,items[_item].price);

        return true;
    }
    function refund(address _to, uint256 _item) onlyWallet(_to) onlyOwner public payable returns (bool success) {
        require(isBuyer(_to,_item)&&items[_item].category>0&&(items[_item].customer[_to]>0||(items[_item].category==2&&safeAdd(items[_item].customer[_to],2 weeks)>now)));
        require((erc20==address(0)?address(this).balance:ERC20Interface(erc20).balanceOf(this))>=items[_item].price);

        if(items[_item].category==1)
            items[_item].customer[_to]  = 0;
        else if(items[_item].category==2)
            items[_item].customer[_to]  = safeSub(items[_item].customer[_to],1 weeks);
        else
            items[_item].customer[_to]  = safeSub(items[_item].customer[_to],1);

        transfer(_to, erc20, items[_item].price);
        _Wallet(_to).refund(erc20,_item,items[_item].price);
        emit Refund(_to,_item,items[_item].price);

        return true;
    }

    event Item(uint256 indexed _item, uint256 _price);
    function resister(uint8 _category, uint256 _price, uint256 _stock) onlyOwner public returns (uint256 _itemId) {
        require(_category>0&&_category<4);
        require(_price>0);
        items[index]    = item(_category,_price,_stock);
        index = safeAdd(index,1);
        emit Item(index,_price);
        return safeSub(index,1);
    }
    function update(uint256 _item, uint256 _price, uint256 _stock) onlyOwner public {
        require(items[_item].category>0);
        require(_price>0);
        uint256 temp = items[_item].price;
        items[_item].price      = _price;
        items[_item].stockCount = safeAdd(items[_item].stockCount,_stock);

        if(temp!=items[_item].price)
            emit Item(index,items[_item].price);
    }

    function price(uint256 _item) public constant returns (address _erc20, uint256 _value) {
        return (erc20,items[_item].price);
    }

    function canBuy(address _who, uint256 _item) public constant returns (bool _canBuy) {
        return  (items[_item].category>0) &&
        !(items[_item].category==1&&items[_item].customer[_who]>0) &&
        (items[_item].stockCount>0);
    }

    function isBuyer(address _who, uint256 _item) public constant returns (bool _buyer) {
        return (items[_item].category==1&&items[_item].customer[_who]>0)||(items[_item].category==2&&safeAdd(items[_item].customer[_who],1 weeks)>now)||(items[_item].category==3&&items[_item].customer[_who]>0);
    }

    uint lastWithdrawal;
    function withdrawal(address _erc20, uint256 _value) onlyOwner public returns (bool success) {
        require(safeAdd(lastWithdrawal,1 weeks)<=now);
        require((_erc20==address(0)?address(this).balance:ERC20Interface(_erc20).balanceOf(this))>=_value);
        if(_erc20==address(0))
            owner.transfer(_value);
        else
            ERC20Interface(_erc20).transfer(owner,_value);
        lastWithdrawal = now;
        return true;
    }
}

contract WalletCenter {
    mapping(address=>bool) public     wallet;
    event Wallet(address indexed _owner, address indexed _wallet);
    function createWallet() public returns (address _wallet) {
        _wallet = new _Wallet(msg.sender);
        wallet[_wallet] = true;
        emit Wallet(msg.sender,_wallet);
        return _wallet;
    }
    function isWallet(address _wallet) public constant returns (bool) {
        return wallet[_wallet];
    }
    mapping(address=>bool) public     shop;
    event Shop(address indexed _owner, address indexed _shop, address indexed _erc20);
    function createShop(address _erc20) public returns (address _shop) {
        _shop   = new _Shop(msg.sender,_erc20);
        shop[_shop] = true;
        emit Shop(msg.sender,_shop,_erc20);
        return _shop;
    }
    function isShop(address _shop) public constant returns (bool) {
        return shop[_shop];
    }
}