//SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
//import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";



/**
 * @title LLM as NFT
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
contract LLM is ERC721URIStorage{
    uint256 token_ids_count;
    string Author;
    string ModelName;
    uint256 Price;
    address author_wallet;

    constructor(string memory name_, string memory smbl_, string memory author_, string memory model_name_, uint256 price, address wallet) ERC721(name_, smbl_) ERC721URIStorage() {
        Author = author_;   // 'TheDrummer' -- hf user
        ModelName = model_name_; // 'Tiger-Gemma-9B-v1-GGUF' -- repo of model
        Price = price;          // price per token to go as royalty
        author_wallet = wallet; // author wallet
    }


    // base url. may be used as hugginface link
    function _baseURI() internal view virtual override returns (string memory) {
        return "https://huggingface.co/";
    }




    /**
     *  
     */
    function CreateItem(string memory file_id) public returns (uint256 token_id) {
        token_ids_count +=1;
        token_id = token_ids_count;
        _safeMint(msg.sender,token_id);
        _setTokenURI(token_id,file_id);

    }


}