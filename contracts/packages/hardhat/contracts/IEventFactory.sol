// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

interface IEventFactory {
    /**
     * @dev Records a minting event.
     * @param tokenId The ID of the minted token.
     * @param ticketAddress The address of the ticket contract.
     * @param eventAddress The address of the event contract.
     * @param minter The address of the user who minted the token.
     */
    function recordMint(
        uint256 tokenId,
        address ticketAddress,
        address eventAddress,
        address minter
    ) external;
}
