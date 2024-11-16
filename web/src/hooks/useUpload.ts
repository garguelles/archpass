import { create } from '@web3-storage/w3up-client';
import { extract } from '@web3-storage/w3up-client/delegation';

export const useUpload = () => {
  const getProof = async (did: string) => {
    const response = await fetch(`api/ucan-proof?did=${did}`);
    if (!response.ok) {
      throw new Error('Error');
    }

    return await response.arrayBuffer();
  };

  const upload = async () => {
    const client = await create();
    const buffer = await getProof(client.agent.did());
    const delegation = await extract(new Uint8Array(buffer));
    if (!delegation.ok) {
      console.error('Problem with extracting proof.');
      return;
    }

    const space = await client.addSpace(delegation.ok);
    await client.setCurrentSpace(space.did());

    // NOTE: Sample implementation of uploading data
    // const metadata = {
    // 	"name": "Tier 1",
    // 	"description": "From delegating users",
    // 	"image": null,
    // 	"attributes": []
    // }
    // const blob = new Blob([JSON.stringify(metadata)], { type: 'application/json' });
    // const jsonCid = await client.uploadFile(new File([blob], 'ticket.png'))
    // console.log(`https://${jsonCid.toString()}.ipfs.w3s.link`);
  };

  return { upload };
};