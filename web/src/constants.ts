export const BASE_SEPOLIA_CHAIN_ID = 84532;
export const mintContractAddress = '0xA3e40bBe8E8579Cd2619Ef9C6fEA362b760dac9f';
export const mintABI = [
  {
    inputs: [
      {
        internalType: 'address',
        name: 'to',
        type: 'address',
      },
    ],
    name: 'mint',
    outputs: [],
    stateMutability: 'public',
    type: 'function',
  },
] as const;

export const eventFactoryABI = [
  {
    inputs: [
      {
        internalType: 'address',
        name: '_eventImplementation',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_ticketImplementation',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  { inputs: [], name: 'FailedDeployment', type: 'error' },
  {
    inputs: [
      { internalType: 'uint256', name: 'balance', type: 'uint256' },
      { internalType: 'uint256', name: 'needed', type: 'uint256' },
    ],
    name: 'InsufficientBalance',
    type: 'error',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'string',
        name: 'eventHash',
        type: 'string',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'clone',
        type: 'address',
      },
    ],
    name: 'EventCreated',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'string',
        name: 'ticketHash',
        type: 'string',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'clone',
        type: 'address',
      },
    ],
    name: 'TicketCreated',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'uint256',
        name: 'tokenId',
        type: 'uint256',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'ticketAddress',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'eventAddress',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'blockNumber',
        type: 'uint256',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'minter',
        type: 'address',
      },
    ],
    name: 'TicketMinted',
    type: 'event',
  },
  {
    inputs: [
      { internalType: 'address', name: '', type: 'address' },
      { internalType: 'uint256', name: '', type: 'uint256' },
    ],
    name: 'attendeeToAllTickets',
    outputs: [
      { internalType: 'uint256', name: 'tokenId', type: 'uint256' },
      { internalType: 'address', name: 'ticketAddress', type: 'address' },
      { internalType: 'address', name: 'eventAddress', type: 'address' },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: '', type: 'address' },
      { internalType: 'address', name: '', type: 'address' },
      { internalType: 'uint256', name: '', type: 'uint256' },
    ],
    name: 'attendeeToTickets',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'string', name: 'eventHash', type: 'string' }],
    name: 'createEvent',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'eventAddress', type: 'address' },
      { internalType: 'string', name: 'name', type: 'string' },
      { internalType: 'uint256', name: 'maxSupply', type: 'uint256' },
      { internalType: 'uint256', name: 'mintPrice', type: 'uint256' },
      { internalType: 'string', name: 'ticketHash', type: 'string' },
    ],
    name: 'createTicket',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    name: 'deployedEvents',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'eventImplementation',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'attendee', type: 'address' },
      { internalType: 'address', name: 'eventAddress', type: 'address' },
    ],
    name: 'getAttendeeTicketsForEvent',
    outputs: [{ internalType: 'uint256[]', name: '', type: 'uint256[]' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'getDeployedEvents',
    outputs: [{ internalType: 'address[]', name: '', type: 'address[]' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint256', name: 'tokenId', type: 'uint256' }],
    name: 'getTicketOwner',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'attendee', type: 'address' }],
    name: 'getTicketsOfAttendee',
    outputs: [
      {
        components: [
          { internalType: 'uint256', name: 'tokenId', type: 'uint256' },
          { internalType: 'address', name: 'ticketAddress', type: 'address' },
          { internalType: 'address', name: 'eventAddress', type: 'address' },
        ],
        internalType: 'struct EventFactory.AttendeeTicket[]',
        name: '',
        type: 'tuple[]',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'uint256', name: 'tokenId', type: 'uint256' },
      { internalType: 'address', name: 'ticketAddress', type: 'address' },
      { internalType: 'address', name: 'eventAddress', type: 'address' },
      { internalType: 'address', name: 'minter', type: 'address' },
    ],
    name: 'recordMint',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'ticketImplementation',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    name: 'ticketToAttendee',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
];

export const eventABI = [
  { inputs: [], name: 'FailedDeployment', type: 'error' },
  {
    inputs: [
      { internalType: 'uint256', name: 'balance', type: 'uint256' },
      { internalType: 'uint256', name: 'needed', type: 'uint256' },
    ],
    name: 'InsufficientBalance',
    type: 'error',
  },
  { inputs: [], name: 'InvalidInitialization', type: 'error' },
  { inputs: [], name: 'NotInitializing', type: 'error' },
  {
    inputs: [{ internalType: 'address', name: 'owner', type: 'address' }],
    name: 'OwnableInvalidOwner',
    type: 'error',
  },
  {
    inputs: [{ internalType: 'address', name: 'account', type: 'address' }],
    name: 'OwnableUnauthorizedAccount',
    type: 'error',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint64',
        name: 'version',
        type: 'uint64',
      },
    ],
    name: 'Initialized',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: true,
        internalType: 'address',
        name: 'previousOwner',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'newOwner',
        type: 'address',
      },
    ],
    name: 'OwnershipTransferred',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'string',
        name: 'ticketHash',
        type: 'string',
      },
      {
        indexed: false,
        internalType: 'address',
        name: 'clone',
        type: 'address',
      },
    ],
    name: 'TicketCreated',
    type: 'event',
  },
  {
    inputs: [
      { internalType: 'string', name: 'name', type: 'string' },
      { internalType: 'uint256', name: 'maxSupply', type: 'uint256' },
      { internalType: 'uint256', name: 'mintPrice', type: 'uint256' },
      { internalType: 'string', name: 'ticketHash', type: 'string' },
    ],
    name: 'createTicket',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'ticketAddress', type: 'address' },
    ],
    name: 'doesTicketExist',
    outputs: [{ internalType: 'bool', name: '', type: 'bool' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: '', type: 'address' }],
    name: 'eventMinters',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    name: 'eventTickets',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'factoryAddress',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: '_ticketFactory', type: 'address' },
      { internalType: 'address', name: 'owner', type: 'address' },
      { internalType: 'address', name: '_factoryAddress', type: 'address' },
    ],
    name: 'initialize',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'ticketAddress', type: 'address' },
      { internalType: 'string', name: 'tokenURI', type: 'string' },
    ],
    name: 'mintNFT',
    outputs: [{ internalType: 'uint256', name: '', type: 'uint256' }],
    stateMutability: 'payable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'owner',
    outputs: [{ internalType: 'address', name: '', type: 'address' }],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      { internalType: 'address', name: 'ticketAddress', type: 'address' },
    ],
    name: 'registerTicket',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'renounceOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [{ internalType: 'address', name: 'newOwner', type: 'address' }],
    name: 'transferOwnership',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'withdraw',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
] as const;
