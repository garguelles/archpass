import { create } from '@web3-storage/w3up-client';
import { extract } from '@web3-storage/w3up-client/delegation';
import { useState } from 'react';

type NFTMetadataAttribute = {
  trait_type: string;
  value: string | number;
};

type NFTMetadata = {
  name: string;
  description: string;
  image?: string;
  attributes?: NFTMetadataAttribute[];
};

type UploadProps = {
  image: File;
  metadata?: NFTMetadata;
};

export const useUpload = () => {
  const [error, setError] = useState<string | null>(null);
  const [isError, setIsError] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState(false);

  const reset = () => setError(null);

  const getProof = async (did: string) => {
    const response = await fetch(`/api/ucan-proof?did=${did}`);
    if (!response.ok) {
      throw new Error('Error fetching proof');
    }
    return await response.arrayBuffer();
  };

  const upload = async ({ metadata, image }: UploadProps) => {
    setIsLoading(true);
    setError(null);
    setIsError(false);

    try {
      const client = await create();
      const buffer = await getProof(client.agent.did());

      const delegation = await extract(new Uint8Array(buffer));
      if (!delegation.ok) {
        throw new Error('Problem with extracting proof.');
      }

      const space = await client.addSpace(delegation.ok);
      await client.setCurrentSpace(space.did());

      const imageCid = await client.uploadFile(image);
      const imageUri = `ipfs://${imageCid.toString()}`;
      const response: { imageURI: string; tokenURI?: string | undefined } = {
        imageURI: imageUri,
        tokenURI: undefined,
      };

      if (metadata) {
        metadata.image = imageUri;
        const blob = new Blob([JSON.stringify(metadata)], {
          type: 'application/json',
        });

        const metadataCid = await client.uploadFile(
          new File([blob], 'metadata.json'), // Updated to a more descriptive filename
        );
        response.tokenURI = `ipfs://${metadataCid.toString()}`;
      }

      return response;
    } catch (err: unknown) {
      setIsError(true);
      setError(
        err instanceof Error
          ? err.message
          : 'An error occurred during the upload process',
      );
      return null;
    } finally {
      setIsLoading(false);
    }
  };

  return { upload, error, isLoading, reset, isError };
};
