// use NODE_ENV to not have to change config based on where it's deployed
export const NEXT_PUBLIC_URL =
  process.env.NODE_ENV === 'development'
    ? 'http://localhost:3000'
    : 'https://onchain-app-template.vercel.app';
// Add your API KEY from the Coinbase Developer Portal
export const NEXT_PUBLIC_CDP_API_KEY = process.env.NEXT_PUBLIC_CDP_API_KEY;
export const NEXT_PUBLIC_WC_PROJECT_ID = process.env.NEXT_PUBLIC_WC_PROJECT_ID;
export const NEXT_PUBLIC_API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL;
export const NEXT_PUBLIC_IMAGE_API_BASE_URL = process.env.NEXT_PUBLIC_IMAGE_API_BASE_URL;
export const DEFAULT_CHAIN_ID = process.env.NEXT_PUBLIC_DEFAULT_CHAIN_ID;
export const RPC_URL = process.env.NEXT_PUBLIC_RPC_URL;
export const AP_EVENT_FACTORY_CONTRACT_ADDRESS =
  process.env.NEXT_PUBLIC_AP_EVENT_FACTORY_CONTRACT_ADDRESS;
export const AP_EVENT_CONTRACT_ADDRESS =
  process.env.NEXT_PUBLIC_EVENT_CONTRACT_ADDRESS;
export const AP_TICKET_CONTRACT_ADDRESS =
  process.env.NEXT_PUBLIC_AP_TICKET_CONTRACT_ADDRESS;
