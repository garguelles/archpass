'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';
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

type FormData = {
  eventName: string;
  description: string;
  location: string;
  eventDate: string;
  headerImage: string;
};

export function CreateEventModal() {
  const [open, setOpen] = useState(false);
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>();

  const onSubmit = (data: FormData) => {
    // In a real app, you would send this data to your backend
    console.log(data);

    // Create a slug from the event name
    const slug = data.eventName.toLowerCase().replace(/\s+/g, '-');

    // Close the modal
    setOpen(false);

    // Redirect to the new event page
    // router.push(`/dashboard/${slug}`);
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
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <Label htmlFor="eventName">Event Name</Label>
            <Input
              id="eventName"
              {...register('eventName', { required: 'Event name is required' })}
            />
            {errors.eventName && (
              <p className="text-sm text-red-500">{errors.eventName.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="description">Description</Label>
            <Textarea
              id="description"
              {...register('description', {
                required: 'Description is required',
              })}
            />
            {errors.description && (
              <p className="text-sm text-red-500">
                {errors.description.message}
              </p>
            )}
          </div>
          <div>
            <Label htmlFor="location">Location</Label>
            <Input
              id="location"
              {...register('location', { required: 'Location is required' })}
            />
            {errors.location && (
              <p className="text-sm text-red-500">{errors.location.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="eventDate">Event Date</Label>
            <Input
              id="eventDate"
              {...register('eventDate', { required: 'Event date is required' })}
              placeholder="Aug 15-19, 2024"
            />
            {errors.eventDate && (
              <p className="text-sm text-red-500">{errors.eventDate.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="headerImage">Header Image URL</Label>
            <Input
              id="headerImage"
              type="url"
              {...register('headerImage', {
                required: 'Header image URL is required',
                pattern: {
                  value:
                    /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/,
                  message: 'Invalid URL format',
                },
              })}
            />
            {errors.headerImage && (
              <p className="text-sm text-red-500">
                {errors.headerImage.message}
              </p>
            )}
          </div>
          <Button type="submit">Create Event</Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
