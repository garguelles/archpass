'use client';
import PublicWalletWrapper from './PublicWalletWrapper';
import WalletWrapper from './WalletWrapper';

export default function LoginButton({
  requireSIWE = true,
}: { requireSIWE?: boolean }) {
  console.log('REQUIRE', requireSIWE);
  return requireSIWE ? (
    <WalletWrapper
      className="min-w-[90px]"
      text="Log in"
      withWalletAggregator={true}
    />
  ) : (
    <PublicWalletWrapper
      className="min-w-[90px]"
      text="Log in"
      withWalletAggregator={true}
    />
  );
}
