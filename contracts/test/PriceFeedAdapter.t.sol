pragma solidity 0.8.12;

import "forge-std/Test.sol";
import "forge-std/console2.sol";

import { PriceFeedAdapter } from "../src/PriceFeedAdapter.sol";

interface IDIAOracleV2{
    function getValue(string memory) external returns (uint128, uint128);
}

contract PriceAdapterTest {
    PriceFeedAdapter internal priceFeedAdapter;

    event TestValue(int value);
    event TestValue2(uint128 value);

    function setUp() public {
        priceFeedAdapter = new PriceFeedAdapter();

        // BTC/USD on sepolia
        priceFeedAdapter.addFeed("BTC/USD", address(0x1b44F3514812d835EB1BDB0acB33d3fA3351Ee43));
        priceFeedAdapter.setDiaOracleAddress(address(0xCD5F78206ca1FF96Ff4c043C61a2299B2Febf3cB));
    }

    function test_fetch_btc_usd_price() public {
        emit TestValue(priceFeedAdapter.getLatestPrice("BTC/USD"));
       priceFeedAdapter.getPriceDia("BTC/USD"); // https://www.unixtimestamp.com/
    }
}
