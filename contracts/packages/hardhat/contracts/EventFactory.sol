// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts/proxy/Clones.sol"; 
import "./Event.sol";

contract EventFactory {
    address public immutable eventImplementation;
    address public immutable ticketImplementation;
    address[] public deployedEvents;              

    event EventCreated(string eventHash, address clone);
    event TicketMinted(
        uint256 indexed tokenId,
        address indexed ticketAddress,
        address indexed eventAddress,
        uint256 blockNumber,
        address minter
    );
    event TicketCreated(string ticketHash, address clone);

    struct AttendeeTicket {
        uint256 tokenId;
        address ticketAddress;
        address eventAddress;
    }

    mapping(address => mapping(address => uint256[])) public attendeeToTickets;
    mapping(uint256 => address) public ticketToAttendee;
    mapping(address => AttendeeTicket[]) public attendeeToAllTickets;

    constructor(
        address _eventImplementation,
        address _ticketImplementation 
    ) {
        eventImplementation = _eventImplementation;
        ticketImplementation = _ticketImplementation;
    }

    function createEvent(
        string memory eventHash
    ) external returns (address) {
        address eventClone = Clones.clone(eventImplementation);
        Event(eventClone).initialize(ticketImplementation, msg.sender, address(this));
        deployedEvents.push(eventClone);
        emit EventCreated(eventHash, eventClone);
        return eventClone;
    }

    function createTicket(
        address eventAddress,
        string memory name,
        uint256 maxSupply,
        uint256 mintPrice,
        string memory ticketHash
    ) external returns (address) {
        require(eventAddress != address(0), "Invalid event address");

        require(Event(eventAddress).owner() == msg.sender, "Caller is not the event owner");

        address ticketClone = Clones.clone(ticketImplementation);
        Ticket(ticketClone).initialize(name, "TNFT", maxSupply, mintPrice, msg.sender);

        Event(eventAddress).registerTicket(ticketClone);

        emit TicketCreated(ticketHash, ticketClone);
        return ticketClone;
    }

    function getDeployedEvents() external view returns (address[] memory) {
        return deployedEvents;
    }

    function recordMint(
        uint256 tokenId,
        address ticketAddress,
        address eventAddress,
        address minter
    ) external {
        emit TicketMinted(tokenId, ticketAddress, eventAddress, block.number, minter);

        ticketToAttendee[tokenId] = minter;

        attendeeToTickets[minter][eventAddress].push(tokenId);

        attendeeToAllTickets[minter].push(AttendeeTicket(tokenId, ticketAddress, eventAddress));
    }

    function getAttendeeTicketsForEvent(address attendee, address eventAddress) external view returns (uint256[] memory) {
        return attendeeToTickets[attendee][eventAddress];
    }

    function getTicketOwner(uint256 tokenId) external view returns (address) {
        return ticketToAttendee[tokenId];
    }

    function getTicketsOfAttendee(address attendee) external view returns (AttendeeTicket[] memory) {
        return attendeeToAllTickets[attendee];
    }

}
