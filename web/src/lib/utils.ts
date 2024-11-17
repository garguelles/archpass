import { type ClassValue, clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';
import type { Address } from 'viem';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

/**
 * Converts an IPFS URL (ipfs://<CID>) to a Web3 Storage HTTPS URL.
 * @param {string} ipfsUrl - The IPFS URL to convert.
 * @returns {string} - The corresponding HTTPS URL.
 */
export function convertIpfsToHttps(ipfsUrl: string): string {
  if (!ipfsUrl) return ''
  
  if (!ipfsUrl.startsWith('ipfs://')) {
    throw new Error('Invalid IPFS URL format');
  }

  const cid = ipfsUrl.replace('ipfs://', '');

  return `https://${cid}.ipfs.w3s.link/`;
}

/**
 * Returns the first 6 and last 4 characters of an address.
 */
export const getSlicedAddress = (address: Address) => {
  return `${address.slice(0, 6)}...${address.slice(-4)}`;
};
