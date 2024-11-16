// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/proxy/Clones.sol"; 
import "./IEventFactory.sol";
import "./Ticket.sol";

contract Event is Initializable, OwnableUpgradeable {
    address[] public eventTickets;
    address ticketImplementation;
    address public factoryAddress; 
    mapping(address => address) public eventMinters;

    // address -> ticket
    event TicketCreated(string ticketHash, address clone);

    function doesTicketExist(address ticketAddress) public view returns (bool) {
        for (uint i = 0; i < eventTickets.length; i++) {
            if (eventTickets[i] == ticketAddress) {
                return true;
            }
        }

        return false;
    }

    function initialize(
        address _ticketFactory,
        address owner,
        address _factoryAddress
    ) public initializer {
        __Ownable_init(owner);
        ticketImplementation = _ticketFactory;
        factoryAddress = _factoryAddress;
    }

    function createTicket(
        string memory name,
        uint256 maxSupply,
        uint256 mintPrice,
        string memory ticketHash 
    ) public onlyOwner returns (address) {
        address clone = Clones.clone(ticketImplementation);
        Ticket(clone).initialize(name, "TNFT", maxSupply, mintPrice, owner());
        eventTickets.push(clone);
        emit TicketCreated(ticketHash, clone);
        return clone;
    }
    
    function mintNFT(
        address ticketAddress,
        string memory tokenURI
    ) public payable returns (uint256) {
        require(doesTicketExist(ticketAddress), "Ticket doesn't belong to Event");
        Ticket ticket = Ticket(ticketAddress);
        require(msg.value >= ticket.mintPrice(), "Insufficient ETH sent");
        
        // Refund excess payment
        if (msg.value > ticket.mintPrice()) {
            payable(msg.sender).transfer(msg.value - ticket.mintPrice());
        }
        
        uint256 tokenId = ticket.mintNFT(msg.sender, tokenURI);
        eventMinters[msg.sender] = ticketAddress;

        IEventFactory(factoryAddress).recordMint(tokenId, ticketAddress, address(this), msg.sender);

        return tokenId;
    }

    function withdraw() public onlyOwner {
        uint256 balance = address(this).balance;
        payable(owner()).transfer(balance);
    }
}
