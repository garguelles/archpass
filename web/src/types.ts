export type TEvent = {
  id: number;
  name: string;
  location: string;
  imageUrl: string;
  date: string;
  contractAddress: string;
  eventSlug: string;
  createdAt: Date;
  modifiedAt: Date;
};

export type TTicket = {
  id: number;
  name: string;
  mintPrice: string;
  quantity: number;
};
