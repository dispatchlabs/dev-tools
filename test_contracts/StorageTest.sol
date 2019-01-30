// Tests contract -> contract opcodes
pragma solidity ^0.4.24;


contract StorageTest {
  mapping (string => string) private _data;

  function get(string key) public view returns(string) {
    return _data[key];
  }

  function set(string key, string value) public returns(bool) {
    _data[key] = value;

    return true;
  }
}
