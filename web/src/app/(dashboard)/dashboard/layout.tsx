'use client';
import { Card, CardContent } from '@/components/ui/card';
import { usePathname } from 'next/navigation';

import '@/app/global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';
import LoginButton from '@/components/LoginButton';
import dynamic from 'next/dynamic';
import { DashboardSidebar } from './[eventSlug]/sidebar';

const OnchainProviders = dynamic(
  () => import('@/components/OnchainProviders'),
  {
    ssr: false,
  },
);

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const pathname = usePathname();
  const showSidebar = pathname !== '/dashboard';

  return (
    <html lang="en" className="h-full">
      <body className="h-full bg-background text-foreground">
        <OnchainProviders>
          <div className="flex min-h-screen">
            {/* Sidebar */}
            {showSidebar && <DashboardSidebar />}

            {/* Main content */}
            <div className="flex-1">
              <header className="flex justify-end bg-card text-card-foreground p-4 border-b">
                <LoginButton />
              </header>

              <main className="p-4 md:p-8">
                <Card>
                  <CardContent className="p-6">{children}</CardContent>
                </Card>
              </main>
            </div>
          </div>
        </OnchainProviders>
      </body>
    </html>
  );
}
