import { NEXT_PUBLIC_API_BASE_URL } from '@/config';
import { createAuthenticatedClient } from '@/lib/client';
import { useMutation } from '@tanstack/react-query';

export function useCreateEventMutation() {
  return useMutation({
    mutationFn: async (event: Record<string, any>) => {
      const api = createAuthenticatedClient();
      const { data } = await api.post(
        `${NEXT_PUBLIC_API_BASE_URL}/admin/event.create`,
        event,
      );

      return data;
    },
  });
}
