'use client';

import { useState } from 'react';
import { useParams } from 'next/navigation';
import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table';
import { useForm } from 'react-hook-form';
import { Plus, Edit, Trash } from 'lucide-react';

// Mock data for tickets
const ticketsData = [
  { id: 1, name: 'General Admission', price: 50, quantity: 500, sold: 300 },
  { id: 2, name: 'VIP', price: 150, quantity: 100, sold: 75 },
  { id: 3, name: 'Early Bird', price: 35, quantity: 200, sold: 200 },
];

type TicketFormData = {
  name: string;
  price: number;
  quantity: number;
};

export default function TicketsPage() {
  const [tickets, setTickets] = useState(ticketsData);
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const params = useParams();
  const eventSlug = params.eventSlug as string;

  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<TicketFormData>();

  const onSubmit = (data: TicketFormData) => {
    const newTicket = {
      id: tickets.length + 1,
      name: data.name,
      price: data.price,
      quantity: data.quantity,
      sold: 0,
    };
    setTickets([...tickets, newTicket]);
    setIsDialogOpen(false);
    reset();
  };

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Tickets for {eventSlug}</h1>
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" /> Create Ticket
            </Button>
          </DialogTrigger>
          <DialogContent>
            <form onSubmit={handleSubmit(onSubmit)}>
              <DialogHeader>
                <DialogTitle>Create New Ticket</DialogTitle>
                <DialogDescription>
                  Add a new ticket type for your event.
                </DialogDescription>
              </DialogHeader>
              <div className="grid gap-4 py-4">
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="name" className="text-right">
                    Name
                  </Label>
                  <Input
                    id="name"
                    className="col-span-3"
                    {...register('name', {
                      required: 'Ticket name is required',
                    })}
                  />
                  {errors.name && (
                    <p className="text-sm text-red-500 col-start-2 col-span-3">
                      {errors.name.message}
                    </p>
                  )}
                </div>
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="price" className="text-right">
                    Price
                  </Label>
                  <Input
                    id="price"
                    type="number"
                    className="col-span-3"
                    {...register('price', {
                      required: 'Price is required',
                      min: 0,
                    })}
                  />
                  {errors.price && (
                    <p className="text-sm text-red-500 col-start-2 col-span-3">
                      {errors.price.message}
                    </p>
                  )}
                </div>
                <div className="grid grid-cols-4 items-center gap-4">
                  <Label htmlFor="quantity" className="text-right">
                    Quantity
                  </Label>
                  <Input
                    id="quantity"
                    type="number"
                    className="col-span-3"
                    {...register('quantity', {
                      required: 'Quantity is required',
                      min: 1,
                    })}
                  />
                  {errors.quantity && (
                    <p className="text-sm text-red-500 col-start-2 col-span-3">
                      {errors.quantity.message}
                    </p>
                  )}
                </div>
              </div>
              <DialogFooter>
                <Button type="submit">Create Ticket</Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
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
                <TableCell>${ticket.price}</TableCell>
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
