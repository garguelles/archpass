//SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract Ticket is Initializable, ERC721URIStorageUpgradeable, OwnableUpgradeable {
    uint256 private _tokenIdCounter;

    uint256 public maxSupply;
    uint256 public mintPrice;
    bool public allowMint = true;

    mapping(address => uint256) public userMintCount; 

    function initialize(
        string memory name,
        string memory symbol,
        uint256 _maxSupply,
        uint256 _mintPrice,
        address owner 
    ) public initializer {
        __ERC721_init(name, symbol);
        __Ownable_init(owner);

        maxSupply = _maxSupply;
        mintPrice = _mintPrice;
        allowMint = true;  
        _tokenIdCounter = 0;
    }

    function mintNFT(address to, string memory tokenURI) public returns (uint256) {
        require(_tokenIdCounter < maxSupply, "Max supply reached");
        require(userMintCount[to] == 0, "Attendee has already minted a ticket");
        require(allowMint, "Minting is not allowed");

        _tokenIdCounter++;
        uint256 newItemId = _tokenIdCounter;
        
        _mint(to, newItemId);
        _setTokenURI(newItemId, tokenURI);

        userMintCount[to]++;

        return newItemId;
    }

    function setMintPrice(uint256 _mintPrice) public onlyOwner {
        mintPrice = _mintPrice;
    }

    function setAllowMint(bool _allowMint) public onlyOwner {
        allowMint = _allowMint;
    }

    function setMaxSupply(uint256 _maxSupply) public onlyOwner {
        maxSupply = _maxSupply;
    }

    function totalSupply() public view returns (uint256) {
        return _tokenIdCounter;
    }
}
