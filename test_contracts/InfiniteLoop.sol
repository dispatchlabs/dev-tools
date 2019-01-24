pragma solidity ^0.4.24;

contract InfiniteLoop {
  function infiniteLoop() public {
    int input1 = 0;
    while(input1 == 0) {
      logEvent();
    }
  }
}