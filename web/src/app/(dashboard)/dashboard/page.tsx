'use client';

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
import { useEventListQuery } from '@/queries/event-list';
import { TEvent } from '@/types';

export default function DashboardPage() {
  const { events, refetchEventList } = useEventListQuery();

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Event Dashboard</h1>
        <CreateEventModal refetchEventList={refetchEventList} />
      </div>

      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
        {events?.map((event: TEvent) => (
          <Link
            href={`/dashboard/${event.eventSlug}`}
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
                  890 attendees
                </div>
              </CardContent>
            </Card>
          </Link>
        ))}
      </div>
    </div>
  );
}
