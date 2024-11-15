'use client';

import { useState } from 'react';
import { useParams } from 'next/navigation';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Progress } from '@/components/ui/progress';
import {
  CalendarDays,
  DollarSign,
  MapPin,
  Users,
  Ticket,
  ArrowUpRight,
} from 'lucide-react';
import { EditEventModal } from './edit-event-modal';

// Mock data (in a real app, this would come from an API or database)
const initialEventData = {
  name: 'ETH Global 2024',
  description: 'The largest Ethereum hackathon and conference',
  date: 'Aug 15-19, 2024',
  location: 'San Francisco, CA',
  headerImage: 'https://example.com/event-image.jpg',
  attendees: 1000,
  totalRevenue: 50000,
  ticketsSold: 800,
  totalTickets: 1000,
  recentTransactions: [
    { id: 1, name: 'John Doe', amount: 250, date: '2023-11-15' },
    { id: 2, name: 'Jane Smith', amount: 250, date: '2023-11-14' },
    { id: 3, name: 'Bob Johnson', amount: 250, date: '2023-11-13' },
  ],
};

export default function EventDashboard() {
  const [eventData, setEventData] = useState(initialEventData);
  const params = useParams();
  const eventSlug = params.eventSlug as string;

  const handleEventUpdate = (updatedEvent: Partial<typeof eventData>) => {
    setEventData((prevData) => ({ ...prevData, ...updatedEvent }));
  };

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold">{eventData.name} Dashboard</h1>
        <EditEventModal
          event={{
            name: eventData.name,
            description: eventData.description,
            location: eventData.location,
            eventDate: eventData.date,
            headerImage: eventData.headerImage,
          }}
          onEventUpdated={handleEventUpdate}
        />
      </div>

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">
              Total Attendees
            </CardTitle>
            <Users className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{eventData.attendees}</div>
            <p className="text-xs text-muted-foreground">
              {((eventData.attendees / eventData.totalTickets) * 100).toFixed(
                1,
              )}
              % of capacity
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Revenue</CardTitle>
            <DollarSign className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              ${eventData.totalRevenue.toLocaleString()}
            </div>
            <Button className="mt-2" size="sm">
              Withdraw
            </Button>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Ticket Sales</CardTitle>
            <Ticket className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">
              {eventData.ticketsSold} / {eventData.totalTickets}
            </div>
            <Progress
              value={(eventData.ticketsSold / eventData.totalTickets) * 100}
              className="mt-2"
            />
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Event Details</CardTitle>
            <CalendarDays className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <p className="text-xs text-muted-foreground">{eventData.date}</p>
            <div className="flex items-center mt-1">
              <MapPin className="h-4 w-4 text-muted-foreground mr-1" />
              <p className="text-xs text-muted-foreground">
                {eventData.location}
              </p>
            </div>
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Recent Transactions</CardTitle>
          <CardDescription>
            Latest ticket purchases for your event
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-2">
            {eventData.recentTransactions.map((transaction) => (
              <div
                key={transaction.id}
                className="flex items-center justify-between border-b py-2"
              >
                <div>
                  <p className="font-medium">{transaction.name}</p>
                  <p className="text-sm text-muted-foreground">
                    {transaction.date}
                  </p>
                </div>
                <div className="flex items-center">
                  <span className="font-bold mr-2">${transaction.amount}</span>
                  <ArrowUpRight className="h-4 w-4 text-green-500" />
                </div>
              </div>
            ))}
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
