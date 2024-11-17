'use client';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { CalendarIcon, MapPinIcon, UserIcon } from 'lucide-react';
import Image from 'next/image';
import { useEventTicketQuery } from '@/queries/event-ticket';
import { convertIpfsToHttps } from '@/lib/utils';

// Mock data (in a real app, this would come from an API or database)
const ticketData = {
  id: '123',
  name: 'John Doe',
  eventName: 'ETH Global 2024',
  date: 'November 15-17, 2024',
  location: 'San Francisco, CA',
  imageUrl: 'https://place-hold.it/600x400',
};

export default function TicketPage({
  params,
}: { params: { id: string } }) {
  console.log(params)
  const { ticket } = useEventTicketQuery(params?.id);
  console.log(ticket);
  return (
    <div className="min-h-screen bg-background text-foreground flex items-center justify-center p-4">
      <Card className="w-full max-w-md overflow-hidden">
        <CardHeader className="bg-primary text-primary-foreground p-6">
          <CardTitle className="text-2xl font-bold text-center">
            Ticket
          </CardTitle>
        </CardHeader>
        <CardContent className="p-6">
          <div className="relative w-full aspect-video mb-6">
            <Image
              src={convertIpfsToHttps(ticket?.imageUrl ?? '')}
              alt={ticket?.eventName}
              width={400}
              height={400}
              fill={false}
              style={{ objectFit: 'cover' }}
              className="rounded-md"
            />
          </div>
          <div className="space-y-4">
            <div className="flex items-center space-x-2">
              <CalendarIcon className="h-5 w-5 text-muted-foreground" />
              <span>{ticket?.date}</span>
            </div>
            <div className="flex items-center space-x-2">
              <MapPinIcon className="h-5 w-5 text-muted-foreground" />
              <span>{ticket?.location}</span>
            </div>
          </div>
          <div className="mt-6 text-center">
            <h2 className="text-2xl font-bold">{ticket?.eventName}</h2>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
