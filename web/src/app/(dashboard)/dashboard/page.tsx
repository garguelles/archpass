import Link from 'next/link';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { CalendarDays, MapPin, Users } from 'lucide-react';
import { CreateEventModal } from './create-event-modal';

// Mock data for events (in a real app, this would come from an API or database)
const events = [
  {
    id: 1,
    name: 'ETH Global 2024',
    slug: 'eth-global-24',
    date: 'Nov 15-17, 2024',
    location: 'San Francisco, CA',
    attendees: 1000,
  },
  {
    id: 2,
    name: 'Devcon 2024',
    slug: 'devcon-24',
    date: 'Oct 1-4, 2024',
    location: 'Bangkok, Thailand',
    attendees: 3500,
  },
  {
    id: 3,
    name: 'NFT NYC',
    slug: 'nft-nyc',
    date: 'Apr 5-7, 2024',
    location: 'New York, NY',
    attendees: 1500,
  },
];

export default function DashboardPage() {
  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Event Dashboard</h1>
        <CreateEventModal />
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {events.map((event) => (
          <Link
            href={`/dashboard/${event.slug}`}
            key={event.id}
            className="transition-transform hover:scale-105"
          >
            <Card className="h-full">
              <CardHeader>
                <CardTitle>{event.name}</CardTitle>
                <CardDescription>
                  <div className="flex items-center">
                    <CalendarDays className="mr-2 h-4 w-4" />
                    {event.date}
                  </div>
                </CardDescription>
              </CardHeader>
              <CardContent>
                <div className="flex items-center mb-2">
                  <MapPin className="mr-2 h-4 w-4" />
                  {event.location}
                </div>
                <div className="flex items-center">
                  <Users className="mr-2 h-4 w-4" />
                  {event.attendees} attendees
                </div>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  );
}
