pragma solidity ^0.4.24;

contract SendTokens {
    function receiverSend(address receiver, uint256 amount) public payable {
        receiver.send(amount);
    }

    function receiverTransfer(address receiver, uint256 amount) public payable {
        receiver.transfer(amount);
    }

    function receiverCall(address receiver, uint256 amount) public payable {
        receiver.call.value(amount);
    }
}