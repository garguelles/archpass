import {NEXT_PUBLIC_API_BASE_URL, NEXT_PUBLIC_IMAGE_API_BASE_URL} from '@/config';
import { createAuthenticatedClient } from '@/lib/client';
import { useMutation } from '@tanstack/react-query';

export function useCreateTicketImageMutation() {
  return useMutation({
    mutationFn: async (ticket: Record<string, any>) => {
      const api = createAuthenticatedClient();
      const { data } = await api.post(
        `${NEXT_PUBLIC_IMAGE_API_BASE_URL}/generate-ticket`,
        ticket,
        {
          responseType: "blob",
        }
      );

      return data;
    },
  });
}
