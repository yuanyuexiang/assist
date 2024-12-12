// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

contract StatusContract {
    // 定义最多账户
    uint256 constant MAX_CHARACTER_AMOUNT = 140;
    // 用户状态账本
    mapping(address => string) public statuses;
    // 定义事件
    event StatusUpdated(address indexed user, string newStatus, uint64 timestamp);

    // 设置状态
    function setStatus(string memory _status) public {
        require(bytes(_status).length <= MAX_CHARACTER_AMOUNT, "Status is over the max character amount.");
        statuses[msg.sender] = _status;
        emit StatusUpdated(msg.sender, _status, uint64(block.timestamp));
    }

    // 获取状态
    function getStatus(address _user) public view returns (string memory) {
        string memory userStatus = statuses[_user];
        if (bytes(userStatus).length == 0) {
            return "No user status";
        }
        return userStatus;
    }
}