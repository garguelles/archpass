'use client';
import {
  Address,
  Avatar,
  EthBalance,
  Identity,
  Name,
} from '@coinbase/onchainkit/identity';
import {
  ConnectWallet,
  Wallet,
  WalletDropdown,
  WalletDropdownBasename,
  WalletDropdownDisconnect,
  WalletDropdownFundLink,
  WalletDropdownLink,
} from '@coinbase/onchainkit/wallet';
import { useAccount, useDisconnect } from 'wagmi';
import {
  AlertDialog,
  AlertDialogHeader,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogTitle,
} from './ui/alert-dialog';
import { Button } from './ui/button';
import { FilePen } from 'lucide-react';

type WalletWrapperParams = {
  text?: string;
  className?: string;
  withWalletAggregator?: boolean;
};
export default function PublicWalletWrapper({
  className,
  text,
  withWalletAggregator = false,
}: WalletWrapperParams) {
  return (
    <>
      <Wallet>
        <ConnectWallet
          withWalletAggregator={withWalletAggregator}
          text={text}
          className={className}
        >
          <Avatar className="h-6 w-6" />
          <Name />
        </ConnectWallet>
        <WalletDropdown>
          <Identity className="px-4 pt-3 pb-2" hasCopyAddressOnClick={true}>
            <Avatar />
            <Name />
            <Address />
            <EthBalance />
          </Identity>
          <WalletDropdownBasename />
          <WalletDropdownLink icon="wallet" href="https://wallet.coinbase.com">
            Go to Wallet Dashboard
          </WalletDropdownLink>
          <WalletDropdownFundLink />
          <WalletDropdownDisconnect />
        </WalletDropdown>
      </Wallet>
    </>
  );
}
