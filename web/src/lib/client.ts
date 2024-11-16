import { NEXT_PUBLIC_API_BASE_URL } from '@/config';
import axios from 'axios';

export const createAuthenticatedClient = () => {
  const accessToken = window.localStorage.getItem('accessToken');
  const authenticatedApi = axios.create({
    baseURL: `${NEXT_PUBLIC_API_BASE_URL}`,
    timeout: 60000,
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
      Authorization: `Bearer ${accessToken}`,
    },
  });

  return authenticatedApi;
};
