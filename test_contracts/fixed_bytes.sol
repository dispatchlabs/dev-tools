pragma solidity ^0.4.24;

contract FixedBytes {
    bytes32 bytes32Amount;
    uint uintAmount;
    bool booleanValue;
    address addressValue;
    string stringValue;
    
    function returnBytes32Amount(bytes32 _amount) public returns(bytes32) {
        bytes32Amount = _amount; 
        return bytes32Amount;
    }
}