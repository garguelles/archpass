import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function useEventTicketQuery(ticketSlug: string) {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['eventItem', ticketSlug],
    enabled: !!ticketSlug,
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: eventItem } = await api.get(`/eventTicket.get`, {
        params: { ticketSlug: ticketSlug},
      });

      return eventItem;
    },
  });

  return {
    ticket: data,
    pending: isLoading || isFetching,
  };
}
