'use client';
import { useAuth } from '@/providers/auth-provider';
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

type WalletWrapperParams = {
  text?: string;
  className?: string;
  withWalletAggregator?: boolean;
};
export default function WalletWrapper({
  className,
  text,
  withWalletAggregator = false,
}: WalletWrapperParams) {
  const { address, isConnected } = useAccount();
  const { disconnect } = useDisconnect();
  const { login, accessToken } = useAuth();

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

      <AlertDialog open={isConnected && !accessToken}>
        <AlertDialogContent
          onEscapeKeyDown={cancelLogin}
          className="flex flex-col"
        >
          <AlertDialogHeader>
            <AlertDialogTitle>Verify your account</AlertDialogTitle>
            <AlertDialogDescription>
              To finish connecting, you must sign a message in your wallet to
              verify that you are the owner of this account.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <Button
            onClick={() => login(address as `0xString`)}
            className="w-full"
          >
            <FilePen className="mr-2 h-4 w-4" /> Sign message
          </Button>
        </AlertDialogContent>
      </AlertDialog>
    </>
  );
}
