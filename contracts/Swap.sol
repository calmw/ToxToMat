// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./RoleAccess.sol";
import "./interface/IERC20.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";


contract Swap is RoleAccess, ReentrancyGuardUpgradeable {
    using Address for address payable;

    event SwapRecord(
        address user,
        uint256 toxNum,
        uint256 matNum,
        uint256 ctime
    );

    // TOX合约地址
    address public toxAddress;
    // 每次兑换，最多使用的TOX数量
    uint256 public maxToxPerTime; // 每次兑换，最多使用的TOX数量
    // 每次兑换，最少使用的TOX数量
    uint256 public minToxPerTime; // 每次兑换，最少使用的TOX数量
    // 兑换汇率
    uint256 public rate; // 兑换汇率
    // 用户，最最多可使用的TOX数量
    uint256 public swapTotalMax; // 每个用户兑换，最最多可使用的TOX数量
    // 用户兑换累计使用的TOX数量
    mapping(address => uint256)  public userSwapTotal; // 用户兑换累计使用的TOX数量

    function initialize() public initializer {
        _addAdmin(msg.sender);
    }

    // 管理员设置每次兑换，最多使用的TOX数量
    function adminSetToxAddress(address toxAddress_) public onlyAdmin {
        toxAddress = toxAddress_;
    }

    // 管理员设置每次兑换，最多使用的TOX数量
    function adminSetMaxToxPerTime(uint256 maxToxPerTime_) public onlyAdmin {
        maxToxPerTime = maxToxPerTime_;
    }

    // 管理员设置每次兑换，最少使用的TOX数量
    function adminSetMinToxPerTime(uint256 minToxPerTime_) public onlyAdmin {
        minToxPerTime = minToxPerTime_;
    }

    // 管理员设置兑换汇率
    function adminSetRate(uint256 rate_) public onlyAdmin {
        rate = rate_;
    }

    // 管理员设置用户，最最多可使用的TOX数量
    function adminSetSwapTotalMax(uint256 swapTotalMax_) public onlyAdmin {
        swapTotalMax = swapTotalMax_;
    }

    // 兑换
    function swap(uint256 toxNum_) external nonReentrant {
        require(toxNum_ >= minToxPerTime, "TOX number is too less");
        require(toxNum_ <= maxToxPerTime, "TOX number is too much");
        require(userSwapTotal[msg.sender] + toxNum_ <= swapTotalMax, "the quantity you redeemed exceeds the limit");

        IERC20Upgradeable TOXContract = IERC20Upgradeable(toxAddress);
        TOXContract.transferFrom(msg.sender, address(this), toxNum_);
        userSwapTotal[msg.sender] += toxNum_;
        uint256 matNum = toxNum_ / rate;
        address payable sender = payable(msg.sender);
        sender.sendValue(matNum);

        emit SwapRecord(
            msg.sender,
            toxNum_,
            matNum,
            block.timestamp
        );
    }

    // 获取当前合约余额
    function getMatBalance() external view onlyAdmin returns (uint256){
        return address(this).balance;
    }

    // 提现
    function withdrawMat(address user_, uint256 amount_) external onlyAdmin {
        address payable user = payable(user_);
        user.sendValue(amount_);
    }

    receive() external payable {}
}