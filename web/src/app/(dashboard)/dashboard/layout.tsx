'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { LayoutDashboard, Tickets, Users, Settings } from 'lucide-react';
import dynamic from 'next/dynamic';

import '@/app/global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';

const OnchainProviders = dynamic(
  () => import('src/components/OnchainProviders'),
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
        <div className="flex min-h-screen">
          {/* Sidebar */}
          {showSidebar && (
            <aside className="hidden md:flex flex-col w-64 bg-card text-card-foreground border-r min-h-screen">
              <div className="p-4">
                <div className="text-2xl font-bold mb-6">ArchPass</div>
                <nav className="space-y-2">
                  <Button
                    variant="ghost"
                    className="w-full justify-start"
                    asChild
                  >
                    <Link href="/dashboard">
                      <LayoutDashboard className="mr-2 h-4 w-4" />
                      Dashboard
                    </Link>
                  </Button>
                  <Button
                    variant="ghost"
                    className="w-full justify-start"
                    asChild
                  >
                    <Link href="/dashboard/slug/tickets">
                      <Tickets className="mr-2 h-4 w-4" />
                      Tickets
                    </Link>
                  </Button>

                  <Button
                    variant="ghost"
                    className="w-full justify-start"
                    asChild
                  >
                    <Link href="/dashboard/slug/attendees">
                      <Users className="mr-2 h-4 w-4" />
                      Attendees
                    </Link>
                  </Button>
                  <Button
                    variant="ghost"
                    className="w-full justify-start"
                    asChild
                  >
                    <Link href="/dashboard/settings">
                      <Settings className="mr-2 h-4 w-4" />
                      Settings
                    </Link>
                  </Button>
                </nav>
              </div>
            </aside>
          )}

          {/* Main content */}
          <div className="flex-1">
            <header className="bg-card text-card-foreground p-4 border-b">
              <div className="text-2xl font-bold">ArchPass</div>
            </header>
            <main className="p-4 md:p-8">
              <Card>
                <CardContent className="p-6">
                  <OnchainProviders>{children}</OnchainProviders>
                </CardContent>
              </Card>
            </main>
          </div>
        </div>
      </body>
    </html>
  );
}
