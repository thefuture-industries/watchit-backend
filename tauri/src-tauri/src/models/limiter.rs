use serde::{Deserialize, Serialize};

// Limiter: Моделька рекомендаций фильмов в системе БД
// Id, UUID, text_limit, youtube_limit, update_at
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Limiter {
    pub id: i32,

    pub uuid: String,

    pub text_limit: i32,

    pub youtube_limit: i32,

    pub update_at: String,
}
