export type RecommendationsModel = {
  id: number;
  uuid: string;
  title: number;
  genre: string;
};

export type RecommendationAddPayload = {
  uuid?: string;
  title: string;
  genre: string;
};
