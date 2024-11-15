'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@/components/ui/label';

export function CreateEventModal() {
  const [open, setOpen] = useState(false);
  const [eventName, setEventName] = useState('');
  const [description, setDescription] = useState('');
  const [location, setLocation] = useState('');
  const [headerImage, setHeaderImage] = useState('');
  const router = useRouter();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // In a real app, you would send this data to your backend
    console.log({ eventName, description, location, headerImage });

    // Create a slug from the event name
    const slug = eventName.toLowerCase().replace(/\s+/g, '-');

    // Close the modal
    setOpen(false);

    // Redirect to the new event page
    router.push(`/dashboard/${slug}`);
  };

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Create Event</Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create New Event</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <Label htmlFor="eventName">Event Name</Label>
            <Input
              id="eventName"
              value={eventName}
              onChange={(e) => setEventName(e.target.value)}
              required
            />
          </div>
          <div>
            <Label htmlFor="description">Description</Label>
            <Textarea
              id="description"
              value={description}
              onChange={(e) => setDescription(e.target.value)}
              required
            />
          </div>
          <div>
            <Label htmlFor="location">Location</Label>
            <Input
              id="location"
              value={location}
              onChange={(e) => setLocation(e.target.value)}
              required
            />
          </div>
          <div>
            <Label htmlFor="headerImage">Header Image URL</Label>
            <Input
              id="headerImage"
              type="url"
              value={headerImage}
              onChange={(e) => setHeaderImage(e.target.value)}
              required
            />
          </div>
          <Button type="submit">Create Event</Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
