import type { Metadata } from 'next';
import { NEXT_PUBLIC_URL } from '../../config';

import '../global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';
import LoginButton from '@/components/LoginButton';
import dynamic from 'next/dynamic';
import Link from 'next/link';

const OnchainProviders = dynamic(
  () => import('@/components/OnchainProviders'),
  {
    ssr: false,
  },
);

export const viewport = {
  width: 'device-width',
  initialScale: 1.0,
};

export const metadata: Metadata = {
  title: 'ArchPass',
  description: 'On-chain events',
  openGraph: {
    title: 'ArchPass',
    description: 'Built with OnchainKit',
    images: [`${NEXT_PUBLIC_URL}/vibes/vibes-19.png`],
  },
};

export default function RootLayout({
  children,
}: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className="bg-background items-center justify-center">
        <OnchainProviders>
          <div className="min-h-screen bg-background text-foreground">
            <header className="container mx-auto py-6 px-4">
              <nav className="flex justify-between items-center">
                <h1 className="text-2xl font-bold">ArchPass</h1>
                <div className="flex items-center space-x-4">
                  <div className="hidden sm:flex space-x-4">
                    <Link
                      href="/events"
                      className="hover:text-primary transition-colors"
                    >
                      Events
                    </Link>
                    <Link
                      href="/about"
                      className="hover:text-primary transition-colors"
                    >
                      About
                    </Link>
                    <Link
                      href="/contact"
                      className="hover:text-primary transition-colors"
                    >
                      Contact
                    </Link>
                  </div>
                  <LoginButton />
                  {/*<Button variant="outline" asChild>
                  <Link href="/login">Login</Link>
                </Button>
                  <Button asChild>
                    <Link href="/signup">Sign up</Link>
                  </Button>*/}
                </div>
              </nav>
            </header>
            {children}
          </div>
        </OnchainProviders>
      </body>
    </html>
  );
}
