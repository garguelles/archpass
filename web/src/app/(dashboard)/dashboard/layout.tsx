import Link from 'next/link';
import { Button } from '@/components/ui/button';
import { Card, CardContent } from '@/components/ui/card';
import { LayoutDashboard, CalendarDays, Users, Settings } from 'lucide-react';
import OnchainProviders from '@/components/OnchainProviders';

import '@/app/global.css';
import '@coinbase/onchainkit/styles.css';
import '@rainbow-me/rainbowkit/styles.css';

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="h-full">
      <body className="h-full bg-background text-foreground">
        <div className="flex min-h-screen">
          {/* Sidebar */}
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
                  <Link href="/dashboard/events">
                    <CalendarDays className="mr-2 h-4 w-4" />
                    Events
                  </Link>
                </Button>
                <Button
                  variant="ghost"
                  className="w-full justify-start"
                  asChild
                >
                  <Link href="/dashboard/attendees">
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

          {/* Mobile header */}
          <header className="md:hidden w-full bg-card text-card-foreground p-4 border-b">
            <div className="text-2xl font-bold">ArchPass</div>
          </header>

          {/* Main content */}
          <div className="flex-1">
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
