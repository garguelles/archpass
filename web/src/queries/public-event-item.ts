import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function usePublicEventItemQuery(eventSlug: string) {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['publicEventItem', eventSlug],
    enabled: !!eventSlug,
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: eventItem } = await api.get(`/event.get`, {
        params: { slug: eventSlug },
      });

      return eventItem;
    },
  });

  return {
    event: data,
    pending: isLoading || isFetching,
    refetchEventItem: refetch,
  };
}
