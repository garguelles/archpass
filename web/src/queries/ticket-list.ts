import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function useTicketListQuery(eventId: number) {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['tickets'],
    enabled: !!eventId,
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: ticketList } = await api.get(`/admin/ticket.list`, {
        params: { eventId },
      });

      return ticketList;
    },
  });

  return {
    tickets: data,
    pending: isLoading || isFetching,
    refetchTicketList: refetch,
  };
}
