import { Button } from '@/components/ui/button';
import { useEventItemQuery } from '@/queries/event-item';
import { LayoutDashboard, Settings, Tickets, Users } from 'lucide-react';
import Link from 'next/link';
import { useParams } from 'next/navigation';

export function DashboardSidebar() {
  const params = useParams();
  const eventId = params.eventSlug as string;
  const { event } = useEventItemQuery(eventId);

  return (
    <aside className="hidden md:flex flex-col w-64 bg-card text-card-foreground border-r min-h-screen">
      <div className="p-4">
        <div className="text-2xl font-bold mb-6">ArchPass</div>
        <nav className="space-y-2">
          <Button variant="ghost" className="w-full justify-start" asChild>
            <Link href="/dashboard">
              <LayoutDashboard className="mr-2 h-4 w-4" />
              Dashboard
            </Link>
          </Button>
          <Button variant="ghost" className="w-full justify-start" asChild>
            <Link href={`/dashboard/${event?.eventSlug}/tickets`}>
              <Tickets className="mr-2 h-4 w-4" />
              Tickets
            </Link>
          </Button>

          <Button variant="ghost" className="w-full justify-start" asChild>
            <Link href={`/dashboard/${event?.eventSlug}/attendees`}>
              <Users className="mr-2 h-4 w-4" />
              Attendees
            </Link>
          </Button>
          <Button variant="ghost" className="w-full justify-start" asChild>
            <Link href={`/dashboard/${event?.eventSlug}/settings`}>
              <Settings className="mr-2 h-4 w-4" />
              Settings
            </Link>
          </Button>
        </nav>
      </div>
    </aside>
  );
}
