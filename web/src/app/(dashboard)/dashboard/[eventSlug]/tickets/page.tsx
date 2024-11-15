'use client';

import { useState } from 'react';
import { useParams } from 'next/navigation';
import { Button } from '@/components/ui/button';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { Edit, Trash } from 'lucide-react';
import { CreateTicketModal } from './create-ticket-modal';

// Mock data for tickets
const ticketsData = [
  { id: 1, name: 'General Admission', price: 50, quantity: 500, sold: 300 },
  { id: 2, name: 'VIP', price: 150, quantity: 100, sold: 75 },
  { id: 3, name: 'Early Bird', price: 35, quantity: 200, sold: 200 },
];

type Ticket = {
  id: number;
  name: string;
  price: number;
  quantity: number;
  sold: number;
};

export default function TicketsPage() {
  const [tickets, setTickets] = useState<Ticket[]>(ticketsData);
  const params = useParams();
  const eventSlug = params.eventSlug as string;

  const handleTicketCreated = (newTicket: Omit<Ticket, 'id' | 'sold'>) => {
    const ticketWithId: Ticket = {
      ...newTicket,
      id: tickets.length + 1,
      sold: 0,
    };
    setTickets([...tickets, ticketWithId]);
  };

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Tickets for {eventSlug}</h1>
        <CreateTicketModal onTicketCreated={handleTicketCreated} />
      </div>

      <div className="bg-background rounded-lg border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Price</TableHead>
              <TableHead>Quantity</TableHead>
              <TableHead>Sold</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {tickets.map((ticket) => (
              <TableRow key={ticket.id}>
                <TableCell>{ticket.name}</TableCell>
                <TableCell>${ticket.price.toFixed(2)}</TableCell>
                <TableCell>{ticket.quantity}</TableCell>
                <TableCell>{ticket.sold}</TableCell>
                <TableCell>
                  <div className="flex space-x-2">
                    <Button variant="outline" size="icon">
                      <Edit className="h-4 w-4" />
                      <span className="sr-only">Edit</span>
                    </Button>
                    <Button variant="outline" size="icon">
                      <Trash className="h-4 w-4" />
                      <span className="sr-only">Delete</span>
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      <p className="text-sm text-muted-foreground">
        Total tickets:{' '}
        {tickets.reduce((sum, ticket) => sum + ticket.quantity, 0)}
      </p>
    </div>
  );
}
