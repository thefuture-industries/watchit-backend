use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
pub struct JsonError {
    pub error: String,
}

#[derive(Debug, Serialize)]
pub struct ErrorResponse {
    pub error: String,
    pub status: u16,
}
