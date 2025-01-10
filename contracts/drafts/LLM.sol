//SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
//import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";



/**
 * @title 
 * @author zer0_eX (JackBekket) 
 * @notice
 * 
 * Input data:
 *   model: huggingface://TheDrummer/Tiger-Gemma-9B-v1-GGUF/Tiger-Gemma-9B-v1a-Q4_K_M.gguf
 *  # https://huggingface.co/TheDrummer/Tiger-Gemma-9B-v1
 * 
 *  base url should be: https://huggingface.co/
 *  token_id should be masked as 'TheDrummer/Tiger-Gemma-9B-v1-GGUF/Tiger-Gemma-9B-v1a-'   --- there should be a variable for that?
 *  token_ids should be differ via quantization methods, so each token id is something like Q4_K_M.gguf
 *  there should be a price for a whole model, without diffirienting from a quantization type... ?
 */




// collection, singleton NFT(?)
contract SampleNFT is ERC721URIStorage{
    uint256 token_ids_count;

    constructor(string memory name_, string memory smbl_) ERC721(name_, smbl_) ERC721URIStorage() {

    }


    // base url. may be used as hugginface link
    function _baseURI() internal view virtual override returns (string memory) {
        return "https://huggingface.co/";
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