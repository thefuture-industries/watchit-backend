use serde::{Deserialize, Serialize};

// Recommendations: Моделька рекомендаций фильмов в системе БД
// Id, UUID, Title, Genre
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Recommendations {
    pub id: i32,
    pub uuid: String,
    pub title: i32,
    pub genre: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct RecommendationAddPayload {
    pub uuid: String,
    pub title: String,
    pub genre: String,
}
