use crate::models::recommendations;
use crate::models::movie;
use crate::models::error;
use dotenv::dotenv;

#[derive(Clone)]
pub struct NewStore {
  client: reqwest::Client,
  server_url: String,
}


// ------------------------------
// Инициализация default значений
// ------------------------------
impl NewStore {
  pub fn new() -> Self {
    dotenv().ok();
    let server_url = std::env::var("SERVER_URL")
      .expect("SERVER_URL environment variable not set");

    Self {
      client: reqwest::Client::new(),
      server_url,
    }
  }
}

pub trait IRecommendations {
  // Добовление избанных фильмов
  async fn add_recommendations(&self, recom: recommendations::RecommendationAddPayload) -> std::result::Result<String, error::ErrorResponse>;
  // Получение избанных фильмов
  async fn get_recommendations(&self, uuid: &str) -> std::result::Result<Vec<movie::MovieModel>, error::ErrorResponse>;
}

impl IRecommendations for NewStore {
  // ---------------------------
  // Добовление избанных фильмов
  async fn add_recommendations(&self, recom: recommendations::RecommendationAddPayload) -> std::result::Result<String, error::ErrorResponse> {
    // Отправка запроса на сервер
    let url = format!("{}recommendations", self.server_url);
    let body = serde_json::json!({
      "uuid": &recom.uuid,
      "title": &recom.title,
      "genre": &recom.genre,
    });

    let response = self.client.post(url).json(&body).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        error::ErrorResponse {
          error: "error trying to connect: tcp connect error".to_string(),
          status: 500,
        }
      } else {
        error::ErrorResponse {
          error: format!("ERROR request {}", e),
          status: 500,
        }
      }
    })?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        Ok(String::from("Recommendation is created successfully"))
      }

      status_code => {
        let error_text = response.text().await.unwrap_or_default();
        let error_json: std::result::Result<error::JsonError, _> = serde_json::from_str(&error_text);

        let error = match error_json {
          Ok(err) => err.error,
          Err(_) => error_text,
        };

        Err(error::ErrorResponse {
          error,
          status: status_code.into(),
        })
      }
    }
  }

  // --------------------------
  // Получение избанных фильмов
  async fn get_recommendations(&self, uuid: &str) -> std::result::Result<Vec<movie::MovieModel>, error::ErrorResponse> {
    let url = format!("{}recommendations/{}", self.server_url, uuid);
    let response = self.client.get(url).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        error::ErrorResponse {
          error: "error trying to connect: tcp connect error".to_string(),
          status: 500,
        }
      } else {
        error::ErrorResponse {
          error: format!("ERROR request {}", e),
          status: 500,
        }
      }
    })?;

    match response.status() {
      reqwest::StatusCode::OK => {
        let res: Vec<movie::MovieModel> = response.json().await
          .map_err(|e| error::ErrorResponse {
            error: format!("Error deserialization json: {}", e),
            status: 500,
          })?;

        Ok(res)
      }

      status_code => {
        let error_text = response.text().await.unwrap_or_default();
        let error_json: std::result::Result<error::JsonError, _> = serde_json::from_str(&error_text);

        let error = match error_json {
          Ok(err) => err.error,
          Err(_) => error_text,
        };

        Err(error::ErrorResponse {
          error,
          status: status_code.into(),
        })
      }
    }
  }
}
