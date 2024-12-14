use serde::{Deserialize, Serialize};

// Favourites: Моделька избранных фильмов в системе БД
// Id, UUID, MovieID, MoviePoster, CreatedAt
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Favourites {
    pub id: i32,

    pub uuid: String,

    #[serde(rename = "movieId")]
    pub movie_id: i32,

    #[serde(rename = "moviePoster")]
    pub movie_poster: String,

    #[serde(rename = "createdAt")]
    pub created_at: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct FavouriteAddPayload {
    pub uuid: String,
    pub movie_id: i32,
    pub movie_poster: String,
}
