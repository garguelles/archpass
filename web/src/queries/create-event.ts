import { NEXT_PUBLIC_API_BASE_URL } from '@/config';
import { createAuthenticatedClient } from '@/lib/client';
import { useMutation } from '@tanstack/react-query';

export function useCreateEventMutation() {
  return useMutation({
    mutationFn: (event: Record<string, any>) => {
      const api = createAuthenticatedClient();
      return api.post(`${NEXT_PUBLIC_API_BASE_URL}/admin/event.create`, event);
    },
  });
}
