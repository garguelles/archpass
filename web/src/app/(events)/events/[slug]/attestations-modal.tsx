import { useState } from 'react';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Label } from '@/components/ui/label';
import { RadioGroup, RadioGroupItem } from '@/components/ui/radio-group';
import { ScrollArea } from '@/components/ui/scroll-area';

type Attestation = {
  id: string;
  attester: string;
  reaction: '👍' | '👎';
  review: string;
};

type AttestationsModalProps = {
  isOpen: boolean;
  onClose: () => void;
  attestations: Attestation[];
  eventId: string;
};

export function AttestationsModal({
  isOpen,
  onClose,
  attestations,
  eventId,
}: AttestationsModalProps) {
  const [newAttestation, setNewAttestation] = useState<Omit<Attestation, 'id'>>(
    {
      attester: '',
      reaction: '👍',
      review: '',
    },
  );

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // Here you would typically send the new attestation to your backend
    console.log('New attestation:', newAttestation);
    // Reset the form
    setNewAttestation({ attester: '', reaction: '👍', review: '' });
  };

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Attestations for Event {eventId}</DialogTitle>
        </DialogHeader>
        <ScrollArea className="h-[300px] pr-4">
          {attestations.map((attestation) => (
            <div key={attestation.id} className="mb-4 p-4 border rounded">
              <div className="flex justify-between items-center mb-2">
                <span className="font-semibold">{attestation.attester}</span>
                <span>{attestation.reaction}</span>
              </div>
              <p className="text-sm">{attestation.review}</p>
            </div>
          ))}
        </ScrollArea>
        <form onSubmit={handleSubmit} className="space-y-4 mt-4">
          <div>
            <Label htmlFor="attester">Wallet Address or ENS</Label>
            <Input
              id="attester"
              value={newAttestation.attester}
              onChange={(e) =>
                setNewAttestation({
                  ...newAttestation,
                  attester: e.target.value,
                })
              }
              required
            />
          </div>
          <div>
            <Label>Reaction</Label>
            <RadioGroup
              value={newAttestation.reaction}
              onValueChange={(value: '👍' | '👎') =>
                setNewAttestation({ ...newAttestation, reaction: value })
              }
              className="flex space-x-4"
            >
              <div className="flex items-center space-x-2">
                <RadioGroupItem value="👍" id="thumbsUp" />
                <Label htmlFor="thumbsUp">👍</Label>
              </div>
              <div className="flex items-center space-x-2">
                <RadioGroupItem value="👎" id="thumbsDown" />
                <Label htmlFor="thumbsDown">👎</Label>
              </div>
            </RadioGroup>
          </div>
          <div>
            <Label htmlFor="review">Review</Label>
            <Textarea
              id="review"
              value={newAttestation.review}
              onChange={(e) =>
                setNewAttestation({ ...newAttestation, review: e.target.value })
              }
              required
            />
          </div>
          <Button type="submit">Submit Attestation</Button>
        </form>
      </DialogContent>
    </Dialog>
  );
}
