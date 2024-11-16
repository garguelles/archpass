'use client';

import {
  type ReactNode,
  createContext,
  useCallback,
  useContext,
  useEffect,
  useMemo,
  useState,
} from 'react';
import { SiweMessage } from 'siwe';
import { useAccountEffect } from 'wagmi';

import { useSignMessage } from 'wagmi';

import { NEXT_PUBLIC_API_BASE_URL, DEFAULT_CHAIN_ID } from '../config';
import { useRouter } from 'next/navigation';

export type TAuthContext = {
  login: (address: string) => void;
  accessToken: string | null;
};

export type TVerifyProps = {
  signature: string;
  message: string;
  nonce: string;
};

const authContextDefaultValue = {
  accessToken: null,
  login: (_: string) => {},
};

export const AuthContext = createContext<TAuthContext>(authContextDefaultValue);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const router = useRouter();

  const [accessToken, setAccessToken] = useState<string | null>(null);
  useAccountEffect({
    onDisconnect() {
      window.localStorage.removeItem('accessToken');
      setAccessToken(null);
      router.push('/');
    },
  });

  const { signMessageAsync } = useSignMessage();

  useEffect(() => {
    // initialize token onload
    const token = window.localStorage.getItem('accessToken');
    if (token) {
      setAccessToken(token);
    }
  }, []);

  const getNonce = useCallback(async () => {
    const response = await fetch(`${NEXT_PUBLIC_API_BASE_URL}/nonce`);
    if (response.ok) {
      const jsonResponse = await response.json();
      return jsonResponse.nonce;
    }

    return '';
  }, []);

  const verify = useCallback(
    async ({ signature, message, nonce }: TVerifyProps) => {
      const response = await fetch(`${NEXT_PUBLIC_API_BASE_URL}/verify`, {
        method: 'POST',
        body: JSON.stringify({ signature, message, nonce }),
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (response.ok) {
        const jsonResponse = await response.json();
        return jsonResponse.accessToken;
      }
    },
    [],
  );

  const createSiweMessage = useCallback(
    async (address: string, statement: string) => {
      const nonce = await getNonce();
      const message = new SiweMessage({
        domain: window.location.host,
        address,
        statement,
        uri: window.location.origin,
        version: '1',
        nonce,
        chainId: Number(DEFAULT_CHAIN_ID!),
      });
      return { message: message.prepareMessage(), nonce };
    },
    [getNonce],
  );

  const login = useCallback(
    async (address: string) => {
      const { message, nonce } = await createSiweMessage(
        address ?? '',
        'Sign in with ethereum',
      );
      const signature = await signMessageAsync({ message });
      const token = await verify({ signature, message, nonce });
      if (!token) {
        console.log('ERROR');
      }
      setAccessToken(token);
      window.localStorage.setItem('accessToken', token);
      router.push('/dashboard');
    },
    [signMessageAsync, verify, createSiweMessage],
  );

  const value = useMemo(
    () => ({
      accessToken,
      login,
    }),
    [login, accessToken],
  );
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => useContext(AuthContext);
