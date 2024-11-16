// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts/proxy/Clones.sol"; 
import "./Event.sol";

contract EventFactory {
    address public immutable eventImplementation;
    address public immutable ticketImplementation;
    address[] public deployedEvents;              

    event EventCreated(address clone);
    event TicketMinted(
        uint256 indexed tokenId,
        address indexed ticketAddress,
        address indexed eventAddress,
        uint256 blockNumber,
        address minter
    );

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

    function createEvent() external returns (address) {
        address eventClone = Clones.clone(eventImplementation);
        Event(eventClone).initialize(ticketImplementation, msg.sender, address(this));
        deployedEvents.push(eventClone);
        emit EventCreated(eventClone);
        return eventClone;
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
