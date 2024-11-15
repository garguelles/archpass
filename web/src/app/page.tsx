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
    <div className="min-h-screen bg-background text-foreground">
      <header className="container mx-auto py-6 px-4">
        <nav className="flex justify-between items-center">
          <h1 className="text-2xl font-bold">ArchPass</h1>
          <div className="space-x-4">
            <Link
              href="/events"
              className="hover:text-muted-foreground transition-colors"
            >
              Events
            </Link>
            <Link
              href="/about"
              className="hover:text-muted-foreground transition-colors"
            >
              About
            </Link>
            <Link
              href="/contact"
              className="hover:text-muted-foreground transition-colors"
            >
              Contact
            </Link>
          </div>
        </nav>
      </header>

      <main className="container mx-auto px-4 py-12">
        <section className="text-center mb-16">
          <h2 className="text-4xl font-bold mb-4">
            Decentralized Ticketing for the Web3 Era
          </h2>
          <p className="text-xl text-muted-foreground mb-8">
            Secure, transparent, and efficient event ticketing powered by
            blockchain technology.
          </p>
          <Button variant="outline" size="lg">
            Get Started
          </Button>
        </section>

        <section className="grid md:grid-cols-3 gap-8">
          <Card className="bg-card/50 border-muted">
            <CardHeader>
              <CardTitle>Secure Transactions</CardTitle>
              <CardDescription>
                Blockchain-backed ticketing ensures tamper-proof and verifiable
                transactions.
              </CardDescription>
            </CardHeader>
          </Card>
          <Card className="bg-card/50 border-muted">
            <CardHeader>
              <CardTitle>Transparent Pricing</CardTitle>
              <CardDescription>
                Fair pricing with no hidden fees. All transactions are visible
                on the blockchain.
              </CardDescription>
            </CardHeader>
          </Card>
          <Card className="bg-card/50 border-muted">
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
    </div>
  );
}
