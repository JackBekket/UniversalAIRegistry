//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;


// in-direct imports

import './LLM.sol';
import "@openzeppelin/contracts/access/Ownable.sol";


abstract contract FactoryNFT is Ownable {


event CollectionCreated(address indexed creator, address collection);

// from creator to address
mapping (address => address) public Collections ;

address payable public treasure_fund;
//address singleton_nft;

constructor() {
    treasure_fund = payable(msg.sender);
    //singleton_nft = singleton_collection;
}



    /**
     *  @dev Deploy new ERC721URIStorage contract and return it's address
     *  @param name_ name of collection
     *  @param smbl_ symbol of collection
     *  
     */
    function CreateCollection(string memory name_, string memory smbl_, string memory author_,string memory model_name_, uint256 price, address wallet) public returns (address) {
        address collection_address = address(new LLM(name_,smbl_,author_,model_name_,price,msg.sender));
        emit CollectionCreated(msg.sender,collection_address);
        return collection_address;
    }


}