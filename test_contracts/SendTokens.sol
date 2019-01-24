pragma solidity ^0.4.24;

contract SendTokens {
  // Should require 2300 Hz
  function sendAmountTo(address receiver, uint256 amount) public payable returns (bool) {
    return receiver.send(amount);
  }
}
