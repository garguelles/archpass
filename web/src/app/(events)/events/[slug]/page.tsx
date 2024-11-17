'use client';

import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Label } from '@/components/ui/label';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import { BASE_SEPOLIA_CHAIN_ID, eventABI } from '@/constants';
import { useUpload } from '@/hooks/useUpload';
import { getSlicedAddress } from '@/lib/utils';
import { useCreateTicketImageMutation } from '@/queries/create-ticket-image';
import { usePublicEventItemQuery } from '@/queries/public-event-item';
import type { TTicket } from '@/types';
import type {
  TransactionError,
  TransactionResponse,
} from '@coinbase/onchainkit/transaction';
import { CalendarIcon, MapPinIcon, StarIcon, Ticket } from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';
import { useCallback, useEffect, useState } from 'react';
import { type Address, parseEther } from 'viem';
import { useAccount, useWriteContract } from 'wagmi';

// This would typically come from a database or API
const eventData = {
  id: 'eth2024',
  name: 'ETH Global 2024',
  date: 'November 15-17, 2024',
  location: 'San Francisco, CA',
  description:
    "Join the world's largest Ethereum hackathon and conference. Connect with leading developers, entrepreneurs, and enthusiasts in the Ethereum ecosystem. Experience three days of intense coding, workshops, and networking opportunities.",
  imageUrl: 'https://place-hold.it/800x400',
  tickets: [
    { id: 'general', name: 'General Admission', price: '0.1 ETH' },
    { id: 'vip', name: 'VIP Access', price: '0.3 ETH' },
  ],
  organizer: {
    name: 'ETH Global',
    description:
      "ETH Global is the world's largest Ethereum hackathon and conference organizer, bringing together developers, entrepreneurs, and enthusiasts from around the globe.",
    avatarUrl: 'https://place-hold.it/100x100',
    reviewCount: 128,
  },
};

export default function EventPage({ params }: { params: { slug: string } }) {
  const { address } = useAccount();
  const { event } = usePublicEventItemQuery(params.slug);
  const { mutateAsync, isPending: isApiLoading } =
    useCreateTicketImageMutation();
  const { upload, isLoading } = useUpload();
  const [selectedTicketContract, setSelectedTicketContract] =
    useState<string>('');
  const {
    data: hash,
    writeContract,
    error: contractError,
    isPending,
  } = useWriteContract();

  const mintTicket = async () => {
    if (!selectedTicketContract) {
      return;
    }

    const ticket = event?.tickets.find(
      (t: TTicket) => t.contractAddress === selectedTicketContract,
    );

    if (!ticket) {
      return;
    }

    const imageBlob = await mutateAsync({
      eventName: event.name,
      eventLocation: event.location,
      eventDate: event.date,
      attendeeName: `${getSlicedAddress(address ?? '0x0')}`,
      ticketName: ticket?.name?.toUpperCase(),
    });

    const uploadResponse = await upload({
      image: imageBlob as File,
      metadata: {
        name: ticket?.name || '',
        description: ticket?.description || '',
        attributes: [
          {
            trait_type: 'eventName',
            value: event.name,
          },
          {
            trait_type: 'eventLocation',
            value: event.location,
          },
          {
            trait_type: 'eventDate',
            value: event.date,
          },
        ],
      },
    });

    if (!uploadResponse?.tokenURI) {
      return;
    }

    writeContract({
      address: event?.contractAddress as Address,
      abi: eventABI,
      functionName: 'mintNFT',
      args: [selectedTicketContract as Address, uploadResponse?.tokenURI],
      chainId: BASE_SEPOLIA_CHAIN_ID,
      value: parseEther(ticket.mintPrice),
    });
  };

  useEffect(() => {
    console.log('THE CONTRACT', selectedTicketContract);
  }, []);

  let mintPrice = event?.tickets.find(
    (t: TTicket) => t.contractAddress === selectedTicketContract,
  )?.mintPrice;
  mintPrice = mintPrice ? parseEther(mintPrice) : 0n;

  const contracts = [
    {
      address: event?.contractAddress as Address,
      abi: eventABI,
      functionName: 'mintNFT',
      args: [selectedTicketContract, 'test'],
      value: mintPrice,
    },
  ];

  // const t = async () => {
  //   const ticket = event?.tickets.find(
  //     (t: TTicket) => t.contractAddress === selectedTicketContract,
  //   );
  //
  //   const imageBlob = await mutateAsync({
  //     eventName: event.name,
  //     eventLocation: event.location,
  //     eventDate: event.date,
  //     attendeeName: `[${getSlicedAddress(address)}]`,
  //     ticketName: ticket?.name?.toUpperCase(),
  //   });
  //
  //   const uploadResponse = await upload({
  //     image: imageBlob as File,
  //     metadata: {
  //       name: ticket?.name || '',
  //       description: ticket?.description || '',
  //       attributes: [{
  //         trait_type: "eventName",
  //         value: event.name,
  //       }, {
  //         trait_type: "eventLocation",
  //         value: event.location,
  //       }, {
  //         trait_type: "eventDate",
  //         value: event.date,
  //       }],
  //     },
  //   });
  // }

  const handleError = useCallback((err: TransactionError) => {
    console.error('Transaction error:', err);
  }, []);

  const handleSuccess = useCallback(async (response: TransactionResponse) => {
    console.log(response);
  }, []);

  return (
    <div className="min-h-screen bg-background text-foreground">
      <header className="relative h-[400px] w-full">
        <Image
          src={eventData.imageUrl}
          alt={event?.name}
          fill={true}
          style={{ objectFit: 'cover' }}
          priority={true}
        />
        <div className="absolute inset-0 bg-black/60 flex items-end">
          <div className="container mx-auto px-4 py-6">
            <h1 className="text-4xl font-bold text-white mb-2">
              {event?.name}
            </h1>
            <div className="flex items-center text-white/80 space-x-4">
              <span className="flex items-center">
                <CalendarIcon className="mr-2 h-5 w-5" />
                {event?.date}
              </span>
              <span className="flex items-center">
                <MapPinIcon className="mr-2 h-5 w-5" />
                {eventData.location}
              </span>
            </div>
          </div>
        </div>
      </header>

      <main className="container mx-auto px-4 py-12">
        <div className="grid md:grid-cols-3 gap-8">
          <div className="md:col-span-2 space-y-8">
            <Card>
              <CardHeader>
                <CardTitle>Organizer</CardTitle>
              </CardHeader>
              <CardContent className="flex items-center space-x-4">
                <Avatar className="h-16 w-16">
                  <AvatarImage
                    src={event?.organizer?.avatarUrl}
                    alt={event?.organizer?.name}
                  />
                  <AvatarFallback>
                    {eventData.organizer.name.slice(0, 2).toUpperCase()}
                  </AvatarFallback>
                </Avatar>
                <div className="space-y-1">
                  <h3 className="text-lg font-semibold">
                    {event?.organizer?.name || event?.organizer?.walletAddress}
                  </h3>
                  <p className="text-sm text-muted-foreground">
                    {event?.organizer?.description}
                  </p>
                  <Link
                    href="#reviews"
                    className="text-sm text-primary flex items-center hover:underline"
                  >
                    <StarIcon className="mr-1 h-4 w-4" />
                    {eventData.organizer.reviewCount} attestations
                  </Link>
                </div>
              </CardContent>
            </Card>

            <Card>
              <CardHeader>
                <CardTitle>About This Event</CardTitle>
              </CardHeader>
              <CardContent>
                <p>{event?.description}</p>
              </CardContent>
            </Card>
          </div>

          <Card>
            <CardHeader>
              <CardTitle>Ticket Information</CardTitle>
              <CardDescription>
                Secure your spot at {eventData.name}
              </CardDescription>
            </CardHeader>
            <CardContent>
              <RadioGroup
                onValueChange={(value: string) => {
                  value && setSelectedTicketContract(value);
                }}
                defaultValue="general"
                className="space-y-4"
              >
                {event?.tickets.map((ticket: TTicket) => (
                  <div key={ticket.id} className="flex items-center space-x-2">
                    <RadioGroupItem
                      value={ticket?.contractAddress}
                      id={ticket.contractAddress}
                    />
                    <Label
                      htmlFor={ticket.contractAddress}
                      className="flex justify-between w-full"
                    >
                      <span>{ticket?.name}</span>
                      <span>{ticket?.mintPrice} ETH</span>
                    </Label>
                  </div>
                ))}
              </RadioGroup>
            </CardContent>
            <CardFooter>
              <Button
                onClick={() => mintTicket()}
                className="w-full"
                disabled={
                  isApiLoading ||
                  isLoading ||
                  isPending ||
                  !selectedTicketContract
                }
              >
                <Ticket className="mr-2 h-4 w-4" /> Mint Ticket
              </Button>
              {/*<Transaction*/}
              {/*  contracts={contracts}*/}
              {/*  chainId={BASE_SEPOLIA_CHAIN_ID}*/}
              {/*  onError={handleError}*/}
              {/*  onSuccess={handleSuccess}*/}
              {/*  onStatus={async (status) => {*/}
              {/*    console.log(status)*/}
              {/*    if (status.statusName === "init") {*/}

              {/*    }*/}
              {/*  }}*/}
              {/*>*/}
              {/*  <TransactionButton*/}
              {/*    text="Mint ticket"*/}
              {/*    className="mt-0 mr-auto ml-auto max-w-full text-[white]"*/}
              {/*  />*/}
              {/*  <TransactionStatus>*/}
              {/*    <TransactionStatusLabel />*/}
              {/*    <TransactionStatusAction />*/}
              {/*  </TransactionStatus>*/}
              {/*</Transaction>*/}
            </CardFooter>
          </Card>
        </div>
      </main>
    </div>
  );
}
