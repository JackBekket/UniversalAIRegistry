//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// in-direct imports
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";


//direct imports for generating golang files
//import "../node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
//import "../node_modules/@openzeppelin/contracts/utils/Counters.sol";



/**
 *  Sample NFT which use as a dummy from factory
 *  Each file is a record inside this single contract
 *  Constructor accept file_ids string[] array and create tokens from it
 */
contract SampleNFT is ERC721URIStorage{

    using Counters for Counters.Counter;
    Counters.Counter token_ids_count;
    event ItemCreated (string file_id, uint256 token_id);

    constructor(string memory name_, string memory smbl_,string[] memory file_ids, address minter_) ERC721(name_, smbl_) ERC721URIStorage() {
        for (uint i=0; i<file_ids.length;i++) {
            CreateItem(file_ids[i],minter_);
        }
    }


    /**
     *  @dev create nft (single record in this contract) with provided file_id and returns token_id
     *  @param file_id unique file_id string obtained from telegram(?)
     *  @param minter_ address of token minter, I've changed it from msg.sender to this,cause constructor may be called from factory contract
     */
    function CreateItem(string memory file_id, address minter_) public returns (uint256 token_id) {
        token_ids_count.increment();
        token_id = token_ids_count.current();
        _safeMint(minter_,token_id);
        _setTokenURI(token_id,file_id);
        emit ItemCreated(file_id,token_id);
    }


}