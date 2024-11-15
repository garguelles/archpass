import Link from 'next/link';
import { Button } from '@/components/ui/button';
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
} from '@/components/ui/card';

export default function Home() {
  return (
    <>
      <main className="container mx-auto px-4 py-12 text-foreground">
        <section className="text-center mb-16">
          <h2 className="text-4xl font-bold mb-4">
            Decentralized Ticketing for the Web3 Era
          </h2>
          <p className="text-xl text-muted-foreground mb-8">
            Secure, transparent, and efficient event ticketing powered by
            blockchain technology.
          </p>
          <Button size="lg" asChild>
            <Link href="/events">Explore Events</Link>
          </Button>
        </section>

        <section className="grid md:grid-cols-3 gap-8">
          <Card className="bg-card text-card-foreground">
            <CardHeader>
              <CardTitle>Secure Transactions</CardTitle>
              <CardDescription>
                Blockchain-backed ticketing ensures tamper-proof and verifiable
                transactions.
              </CardDescription>
            </CardHeader>
          </Card>
          <Card className="bg-card text-card-foreground">
            <CardHeader>
              <CardTitle>Transparent Pricing</CardTitle>
              <CardDescription>
                Fair pricing with no hidden fees. All transactions are visible
                on the blockchain.
              </CardDescription>
            </CardHeader>
          </Card>
          <Card className="bg-card text-card-foreground">
            <CardHeader>
              <CardTitle>Easy Integration</CardTitle>
              <CardDescription>
                Seamlessly integrate with existing event management systems.
              </CardDescription>
            </CardHeader>
          </Card>
        </section>
      </main>

      <footer className="container mx-auto py-6 px-4 text-center text-muted-foreground">
        <p>&copy; 2024 ArchPass. All rights reserved.</p>
      </footer>
    </>
  );
}
