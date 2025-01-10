//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


// in-direct imports

import "./SampleNFT.sol";
import "./SingletonNFT.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


// direct imports
//import "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract FactoryNFT is Ownable {


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
     *  @param file_ids unique ids of file's, obtained from telegram
     *  *IMPORTANT* -- as this function creates new contract -- it may be very expensive buy gas cost
     *  if user just want to create single nft token, it should use SingletonNFT.CreateItem(file_id) instead
     */
    function CreateCollection(string memory name_, string memory smbl_, string[] memory file_ids) public returns (address) {
        address collection_address = address(new SampleNFT(name_,smbl_,file_ids,msg.sender));
        emit CollectionCreated(msg.sender,collection_address);
        return collection_address;
    }


}