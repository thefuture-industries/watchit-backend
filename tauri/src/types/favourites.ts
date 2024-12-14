export type FavouriteModel = {
  id: number;
  uuid: string;
  movieId: number;
  moviePoster: string;
  CreatedAt: string;
};

export type FavouriteAddPayload = {
  uuid?: string;
  movie_id: number;
  movie_poster: string;
};
