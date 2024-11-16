'use client';
import PublicWalletWrapper from './PublicWalletWrapper';
import WalletWrapper from './WalletWrapper';

export default function SignupButton({
  requireSIWE = true,
}: { requireSIWE?: boolean }) {
  console.log('REQUIRE_SIGNIN', requireSIWE);
  return requireSIWE ? (
    <WalletWrapper
      className="ockConnectWallet_Container min-w-[90px] shrink bg-slate-200 text-[#030712] hover:bg-slate-300"
      text="Sign up"
    />
  ) : (
    <PublicWalletWrapper
      className="ockConnectWallet_Container min-w-[90px] shrink bg-slate-200 text-[#030712] hover:bg-slate-300"
      text="Sign up"
    />
  );
}
