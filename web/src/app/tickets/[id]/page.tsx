import Image from 'next/image';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { CalendarIcon, MapPinIcon, UserIcon } from 'lucide-react';

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
}: { params: { ticketId: string } }) {
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
              src={ticketData.imageUrl}
              alt={ticketData.eventName}
              fill
              style={{ objectFit: 'cover' }}
              className="rounded-md"
            />
          </div>
          <div className="space-y-4">
            <div className="flex items-center space-x-2">
              <UserIcon className="h-5 w-5 text-muted-foreground" />
              <span className="font-semibold">{ticketData.name}</span>
            </div>
            <div className="flex items-center space-x-2">
              <CalendarIcon className="h-5 w-5 text-muted-foreground" />
              <span>{ticketData.date}</span>
            </div>
            <div className="flex items-center space-x-2">
              <MapPinIcon className="h-5 w-5 text-muted-foreground" />
              <span>{ticketData.location}</span>
            </div>
          </div>
          <div className="mt-6 text-center">
            <h2 className="text-2xl font-bold">{ticketData.eventName}</h2>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
