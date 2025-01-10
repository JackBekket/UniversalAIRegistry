//SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// ../node_modules/
// in-direct imports

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
//import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";



// direct imports
//import "../node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";



// collection, singleton NFT(?)
contract SampleNFT is ERC721URIStorage{

    //using Counters for Counters.Counter;
    uint256 token_ids_count;

    constructor(string memory name_, string memory smbl_) ERC721(name_, smbl_) ERC721URIStorage() {

    }


    /**
     *  @dev create nft (single record in this contract) with provided file_id and returns token_id
     *  @param file_id unique file_id string obtained from telegram(?)
     */
    function CreateItem(string memory file_id) public returns (uint256 token_id) {
        token_ids_count +=1;
        token_id = token_ids_count;
        _safeMint(msg.sender,token_id);
        _setTokenURI(token_id,file_id);

    }


}