use crate::models::recommendations;
use crate::models::movie;
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
  async fn add_recommendations(&self, recom: recommendations::RecommendationAddPayload) -> std::result::Result<String, String>;
  // Получение избанных фильмов
  async fn get_recommendations(&self, uuid: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
}

impl IRecommendations for NewStore {
  // ---------------------------
  // Добовление избанных фильмов
  async fn add_recommendations(&self, recom: recommendations::RecommendationAddPayload) -> std::result::Result<String, String> {
    // Отправка запроса на сервер
    let url = format!("{}recommendations", self.server_url);
    let body = serde_json::json!({
      "uuid": &recom.uuid,
      "title": &recom.title,
      "genre": &recom.genre,
    });

    let response = self.client.post(url).json(&body).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        "error trying to connect: tcp connect error".to_string()
      } else {
        format!("ERROR request {}", e)
      }
    })?;

    match response.status() {
      reqwest::StatusCode::OK => {
        Ok(String::from("Recommendation is created successfully"))
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }

  // --------------------------
  // Получение избанных фильмов
  async fn get_recommendations(&self, uuid: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    let url = format!("{}recommendations/{}", self.server_url, uuid);
    let response = self.client.get(url).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        "error trying to connect: tcp connect error".to_string()
      } else {
        format!("ERROR request {}", e)
      }
    })?;

    match response.status() {
      reqwest::StatusCode::OK => {
        let res: Vec<movie::MovieModel> = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;

        Ok(res)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }
}
