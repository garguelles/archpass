import Image from 'next/image';
import Link from 'next/link';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import { Twitter, Github, Linkedin } from 'lucide-react';

// Mock data (in a real app, this would come from an API or database)
const userData = {
  displayName: 'Vitalik Buterin',
  walletAddress: '0x1234...5678',
  bio: 'Ethereum co-founder. Blockchain enthusiast. Decentralization advocate.',
  avatarUrl: 'https://place-hold.it/100x100',
  socials: {
    twitter: 'https://twitter.com/VitalikButerin',
    github: 'https://github.com/vbuterin',
    linkedin: 'https://www.linkedin.com/in/vitalik-buterin-267a7450/',
  },
  tickets: [
    {
      id: 1,
      name: 'ETH Global 2024',
      imageUrl: 'https://place-hold.it/300x300',
    },
    {
      id: 2,
      name: 'Devcon 2024',
      imageUrl: 'https://place-hold.it/300x300',
    },
    {
      id: 3,
      name: 'NFT NYC',
      imageUrl: 'https://place-hold.it/300x300',
    },
  ],
};

export default function UserPage({
  params,
}: { params: { walletAddress: string } }) {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <main className="container mx-auto px-4 py-12">
        <Card className="mb-8">
          <CardContent className="flex flex-col sm:flex-row items-center sm:items-start gap-4 pt-6">
            <Avatar className="w-24 h-24">
              <AvatarImage
                src={userData.avatarUrl}
                alt={userData.displayName}
              />
              <AvatarFallback>
                {userData.displayName.slice(0, 2).toUpperCase()}
              </AvatarFallback>
            </Avatar>
            <div className="text-center sm:text-left">
              <h1 className="text-2xl font-bold mb-2">
                {userData.displayName}
              </h1>
              <p className="text-muted-foreground mb-2">
                {userData.walletAddress}
              </p>
              <p className="mb-4">{userData.bio}</p>
              <div className="flex justify-center sm:justify-start space-x-4">
                {userData.socials.twitter && (
                  <Button variant="ghost" size="icon" asChild>
                    <Link href={userData.socials.twitter}>
                      <Twitter className="h-5 w-5" />
                      <span className="sr-only">Twitter</span>
                    </Link>
                  </Button>
                )}
                {userData.socials.github && (
                  <Button variant="ghost" size="icon" asChild>
                    <Link href={userData.socials.github}>
                      <Github className="h-5 w-5" />
                      <span className="sr-only">GitHub</span>
                    </Link>
                  </Button>
                )}
                {userData.socials.linkedin && (
                  <Button variant="ghost" size="icon" asChild>
                    <Link href={userData.socials.linkedin}>
                      <Linkedin className="h-5 w-5" />
                      <span className="sr-only">LinkedIn</span>
                    </Link>
                  </Button>
                )}
              </div>
            </div>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Tickets (NFTs)</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
              {userData.tickets.map((ticket) => (
                <Card key={ticket.id}>
                  <CardContent className="p-4">
                    <Image
                      src={ticket.imageUrl}
                      alt={ticket.name}
                      width={300}
                      height={300}
                      className="rounded-lg mb-2"
                    />
                    <h3 className="font-semibold text-center">{ticket.name}</h3>
                  </CardContent>
                </Card>
              ))}
            </div>
          </CardContent>
        </Card>
      </main>
    </div>
  );
}
