import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

/**
 * Converts an IPFS URL (ipfs://<CID>) to a Web3 Storage HTTPS URL.
 * @param {string} ipfsUrl - The IPFS URL to convert.
 * @returns {string} - The corresponding HTTPS URL.
 */
export function convertIpfsToHttps(ipfsUrl: string): string {
  if (!ipfsUrl.startsWith('ipfs://')) {
    throw new Error('Invalid IPFS URL format');
  }

  const cid = ipfsUrl.replace('ipfs://', '');

  return `https://${cid}.ipfs.w3s.link/`;
}