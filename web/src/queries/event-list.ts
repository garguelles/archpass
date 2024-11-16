import { createAuthenticatedClient } from '@/lib/client';
import { useQuery } from '@tanstack/react-query';

export function useEventListQuery() {
  const { data, isLoading, isFetching, refetch } = useQuery({
    queryKey: ['userProjects'],
    queryFn: async () => {
      const api = createAuthenticatedClient();
      const { data: eventList } = await api.get(`/admin/event.list`);

      return eventList;
    },
  });

  return {
    events: data,
    pending: isLoading || isFetching,
    refetchEventList: refetch,
  };
}
