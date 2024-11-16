'use client';

import LoginButton from '@/components/LoginButton';
import SignupButton from '@/components/SignupButton';
import { useAccount } from 'wagmi';

export function EventLayoutHeader() {
  const { address } = useAccount();
  return (
    <header className="container mx-auto py-6 px-4">
      <nav className="flex justify-between items-center">
        <h1 className="text-2xl font-bold">ArchPass</h1>
        <div className="flex items-center space-x-4">
          {!address && <LoginButton requireSIWE={false} />}

          <SignupButton requireSIWE={false} />
        </div>
      </nav>
    </header>
  );
}
