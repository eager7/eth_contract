pragma solidity ^0.5.0;

contract SimpleStorage {
    uint256 storedData;

    event Log(bytes32 indexed log, uint256 indexed x);
    function set(uint256 x) public {
        storedData = x;
        emit Log("set data", x);
    }

    function get()public view returns (uint) {
        return storedData;
    }
}