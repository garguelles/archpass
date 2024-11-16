import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function useEventItemQuery(eventId: number) {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['userProjects'],
    enabled: !!eventId,
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: eventItem } = await api.get(`/admin/event.get`, {
        params: { id: eventId },
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
