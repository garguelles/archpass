'use client';

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { Eye } from 'lucide-react';
import { useParams } from 'next/navigation';
import { useState } from 'react';

// Mock data for attendees
const attendeesData = [
  {
    id: 1,
    name: 'Alice Johnson',
    email: 'alice@example.com',
    ticketType: 'VIP',
    purchaseDate: '2023-11-15',
  },
  {
    id: 2,
    name: 'Bob Smith',
    email: 'bob@example.com',
    ticketType: 'General Admission',
    purchaseDate: '2023-11-16',
  },
  {
    id: 3,
    name: 'Charlie Brown',
    email: 'charlie@example.com',
    ticketType: 'Early Bird',
    purchaseDate: '2023-11-14',
  },
  {
    id: 4,
    name: 'Diana Prince',
    email: 'diana@example.com',
    ticketType: 'VIP',
    purchaseDate: '2023-11-17',
  },
  {
    id: 5,
    name: 'Ethan Hunt',
    email: 'ethan@example.com',
    ticketType: 'General Admission',
    purchaseDate: '2023-11-18',
  },
];

type Attendee = {
  id: number;
  name: string;
  email: string;
  ticketType: string;
  purchaseDate: string;
};

export default function AttendeesPage() {
  const [attendees] = useState<Attendee[]>(attendeesData);
  const [selectedAttendee, setSelectedAttendee] = useState<Attendee | null>(
    null,
  );
  const params = useParams();
  const eventSlug = params.eventSlug as string;

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Attendees for {eventSlug}</h1>
      </div>

      <div className="bg-background rounded-lg border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Name</TableHead>
              <TableHead>Email</TableHead>
              <TableHead>Ticket Type</TableHead>
              <TableHead>Purchase Date</TableHead>
              <TableHead>Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {attendees.map((attendee) => (
              <TableRow key={attendee.id}>
                <TableCell className="font-medium">{attendee.name}</TableCell>
                <TableCell>{attendee.email}</TableCell>
                <TableCell>{attendee.ticketType}</TableCell>
                <TableCell>{attendee.purchaseDate}</TableCell>
                <TableCell>
                  <Dialog>
                    <DialogTrigger asChild={true}>
                      <Button
                        variant="outline"
                        size="icon"
                        onClick={() => setSelectedAttendee(attendee)}
                      >
                        <Eye className="h-4 w-4" />
                        <span className="sr-only">View</span>
                      </Button>
                    </DialogTrigger>
                    <DialogContent>
                      <DialogHeader>
                        <DialogTitle>Attendee Details</DialogTitle>
                      </DialogHeader>
                      {selectedAttendee && (
                        <div className="flex flex-col items-center space-y-4">
                          <Avatar className="h-24 w-24">
                            <AvatarImage
                              src={`https://api.dicebear.com/6.x/initials/svg?seed=${selectedAttendee.name}`}
                            />
                            <AvatarFallback>
                              {selectedAttendee.name
                                .split(' ')
                                .map((n) => n[0])
                                .join('')}
                            </AvatarFallback>
                          </Avatar>
                          <h2 className="text-2xl font-bold">
                            {selectedAttendee.name}
                          </h2>
                          <p className="text-muted-foreground">
                            {selectedAttendee.email}
                          </p>
                          <div className="grid grid-cols-2 gap-4 text-sm">
                            <div>
                              <p className="font-semibold">Ticket Type</p>
                              <p>{selectedAttendee.ticketType}</p>
                            </div>
                            <div>
                              <p className="font-semibold">Purchase Date</p>
                              <p>{selectedAttendee.purchaseDate}</p>
                            </div>
                          </div>
                        </div>
                      )}
                    </DialogContent>
                  </Dialog>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      <p className="text-sm text-muted-foreground">
        Total attendees: {attendees.length}
      </p>
    </div>
  );
}
