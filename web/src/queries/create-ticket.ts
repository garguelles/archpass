import { NEXT_PUBLIC_API_BASE_URL } from '@/config';
import { createAuthenticatedClient } from '@/lib/client';
import { useMutation } from '@tanstack/react-query';

export function useCreateTicketMutation() {
  return useMutation({
    mutationFn: async (ticket: Record<string, any>) => {
      const api = createAuthenticatedClient();
      const { data } = await api.post(
        `${NEXT_PUBLIC_API_BASE_URL}/admin/ticket.create`,
        ticket,
      );

      return data;
    },
  });
}
