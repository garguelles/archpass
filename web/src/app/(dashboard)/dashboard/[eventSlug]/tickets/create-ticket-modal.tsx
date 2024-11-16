import { useCallback, useState } from 'react';
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
import { Label } from '@/components/ui/label';
import { Plus } from 'lucide-react';
import { AP_EVENT_FACTORY_CONTRACT_ADDRESS } from '@/config';
import { BASE_SEPOLIA_CHAIN_ID, eventFactoryABI } from '@/constants';
import { type ContractFunctionParameters, parseEther } from 'viem';
import {
  Transaction,
  TransactionButton,
  TransactionError,
  TransactionResponse,
  TransactionStatus,
  TransactionStatusAction,
  TransactionStatusLabel,
} from '@coinbase/onchainkit/transaction';
import { TEvent } from '@/types';
import { useCreateTicketMutation } from '@/queries/create-ticket';

type TicketFormData = {
  name: string;
  price: number;
  maxSupply: number;
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

  const [open, setOpen] = useState(false);
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
    getValues,
  } = useForm<TicketFormData>();

  const createArgs = useCallback(() => {
    const formValues = getValues();
    return [
      eventContractAddress,
      formValues.name,
      formValues.maxSupply,
      parseEther(formValues.mintPrice),
      '0xtesthash',
    ];
  }, [getValues()]);

  const contracts = [
    {
      address: AP_EVENT_FACTORY_CONTRACT_ADDRESS,
      abi: eventFactoryABI,
      functionName: 'createTicket',
      args: [...createArgs()],
    },
  ] as unknown as ContractFunctionParameters[];

  const handleError = useCallback((err: TransactionError) => {
    console.error('Transaction error:', err);
  }, []);

  const handleSuccess = useCallback(
    (response: TransactionResponse) => {
      const formValues = getValues();
      const ticketAddress = response.transactionReceipts?.[0].logs?.[0].address;
      const payload = {
        name: formValues.name,
        eventId: event.id,
        quantity: formValues.maxSupply,
        mintPrice: formValues.mintPrice,
        contractAddress: ticketAddress,
      };

      mutateAsync(payload).then(() => {
        setOpen(false);
        refetchTicketList();
      });
    },
    [getValues, mutateAsync],
  );

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
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
              id="price"
              type="number"
              step="0.01"
              {...register('price', {
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