
pragma solidity ^0.4.22;

///@title vote
contract ballot {
    struct Voter {
        uint    weight;
        bool    voted;
        address delegate;
        uint    vote;
    }

    struct Proposal {
        bytes32 name;
        uint    voteCount;
    }

    address public chairPerson;

    mapping(address => Voter) public voters;

    Proposal[] public proposals;

    function ballot(){

    }
}
