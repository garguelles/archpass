import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function useEventItemQuery(eventSlug: string) {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['userProjects'],
    enabled: !!eventSlug,
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: eventItem } = await api.get(`/admin/event.get`, {
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
