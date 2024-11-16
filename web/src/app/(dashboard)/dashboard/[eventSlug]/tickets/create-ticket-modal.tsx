import { Button } from '@/components/ui/button';
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { AP_EVENT_FACTORY_CONTRACT_ADDRESS } from '@/config';
import { BASE_SEPOLIA_CHAIN_ID, eventFactoryABI } from '@/constants';
import { useUpload } from '@/hooks/useUpload';
import { useCreateTicketMutation } from '@/queries/create-ticket';
import { useCreateTicketImageMutation } from '@/queries/create-ticket-image';
import type { TEvent } from '@/types';
import {
  Transaction,
  TransactionButton,
  type TransactionError,
  type TransactionResponse,
  TransactionStatus,
  TransactionStatusAction,
  TransactionStatusLabel,
} from '@coinbase/onchainkit/transaction';
import { Plus } from 'lucide-react';
import { useCallback, useState } from 'react';
import { useForm } from 'react-hook-form';
import { type ContractFunctionParameters, parseEther } from 'viem';

type TicketFormData = {
  name: string;
  price: number;
  quantity: number;
  mintPrice: string;
};

type CreateTicketModalProps = {
  eventContractAddress: string;
  event: TEvent;
  refetchTicketList: () => void;
};

export function CreateTicketModal({
  eventContractAddress,
  event,
  refetchTicketList,
}: CreateTicketModalProps) {
  const { mutateAsync } = useCreateTicketMutation();
  const { mutateAsync: createImage } = useCreateTicketImageMutation();
  const { upload } = useUpload();

  const [open, setOpen] = useState(false);
  const [contracts, setContracts] = useState<ContractFunctionParameters[]>([]);
  const {
    register,
    formState: { errors },
    reset,
    getValues,
    watch,
  } = useForm<TicketFormData>();

  // Watch form values to update contracts state
  watch((formData) => {
    if (!formData) return;

    try {
      const mintPriceWei = formData.mintPrice
        ? parseEther(formData.mintPrice)
        : 0n;

      const newContracts = [
        {
          address: AP_EVENT_FACTORY_CONTRACT_ADDRESS,
          abi: eventFactoryABI,
          functionName: 'createTicket',
          args: [
            eventContractAddress,
            formData.name || '',
            formData.quantity || 0,
            mintPriceWei,
            '0xtesthash',
          ],
        },
      ] as ContractFunctionParameters[];

      setContracts(newContracts);
    } catch (error) {
      console.error('Error updating contract parameters:', error);
    }
  });

  const handleError = useCallback((err: TransactionError) => {
    console.error('Transaction error:', err);
  }, []);

  const handleSuccess = useCallback(
    async (response: TransactionResponse) => {
      const formValues = getValues();
      const ticketAddress = response.transactionReceipts?.[0].logs?.[0].address;
      const payload = {
        name: formValues.name,
        eventId: event.id,
        quantity: Number(formValues.quantity),
        mintPrice: formValues.mintPrice,
        contractAddress: ticketAddress,
        imageUrl: '',
      };

      const imageBlob = await createImage({
        eventName: event.name,
        eventLocation: event.location,
        eventDate: event.date,
        attendeeName: '[YOUR NAME]',
        ticketName: payload.name.toUpperCase(),
      });

      const uploadResponse = await upload({
        image: imageBlob as File,
      });
      payload.imageUrl = uploadResponse?.imageURI as string;

      mutateAsync(payload).then(() => {
        setOpen(false);
        reset();
        refetchTicketList();
      });
    },
    [getValues, mutateAsync],
  );

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild={true}>
        <Button>
          <Plus className="mr-2 h-4 w-4" /> Create Ticket
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Create New Ticket</DialogTitle>
        </DialogHeader>
        <form className="space-y-4">
          <div>
            <Label htmlFor="name">Ticket Name</Label>
            <Input
              id="name"
              {...register('name', { required: 'Ticket name is required' })}
            />
            {errors.name && (
              <p className="text-sm text-red-500">{errors.name.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="price">Price</Label>
            <Input
              id="mintPrice"
              type="number"
              step="0.01"
              {...register('mintPrice', {
                required: 'Price is required',
                min: { value: 0, message: 'Price must be 0 or greater' },
              })}
            />
            {errors.price && (
              <p className="text-sm text-red-500">{errors.price.message}</p>
            )}
          </div>
          <div>
            <Label htmlFor="quantity">Quantity</Label>
            <Input
              id="quantity"
              type="number"
              {...register('quantity', {
                required: 'Quantity is required',
                min: { value: 1, message: 'Quantity must be at least 1' },
              })}
            />
            {errors.quantity && (
              <p className="text-sm text-red-500">{errors.quantity.message}</p>
            )}
          </div>
          <Transaction
            contracts={contracts}
            chainId={BASE_SEPOLIA_CHAIN_ID}
            onError={handleError}
            onSuccess={handleSuccess}
          >
            <TransactionButton
              text="Create ticket"
              className="mt-0 mr-auto ml-auto max-w-full text-[white]"
            />
            <TransactionStatus>
              <TransactionStatusLabel />
              <TransactionStatusAction />
            </TransactionStatus>
          </Transaction>
        </form>
      </DialogContent>
    </Dialog>
  );
}
