import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function usePublicUserTicketListQuery(walletAddress: string) {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['publicEventItem', walletAddress],
    enabled: !!walletAddress,
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: tickets } = await api.get(`/attendeeTicket.get`, {
        params: { walletAddress: walletAddress },
      });

      return tickets;
    },
  });

  return {
    attendee: data,
    pending: isLoading || isFetching,
    refetchUserTickets: refetch,
  };
}
