use crate::models::youtube;
use crate::models::error;
use dotenv::dotenv;

// ----------------------------------------
// Структура данных для работы с ютуб видео
// ----------------------------------------
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

// -------------------------
// Создания оболочки методов
// -------------------------
pub trait IYoutube {
  // Получение популярного видео youtube
  async fn popular_youtube(&mut self) -> std::result::Result<Vec<youtube::YoutubeModel>, error::ErrorResponse>;
  // Получение видео youtube по фильтрам
  async fn get_youtube_videos(&self, category: &str, search: &str, channel: &str) -> std::result::Result<Vec<youtube::YoutubeModel>, error::ErrorResponse>;
}

// ---------------------------
// Инициализация этой оболочки
// ---------------------------
impl IYoutube for NewStore {
  // Получение популярного видео youtube
  // -----------------------------------
  async fn popular_youtube(&mut self) -> std::result::Result<Vec<youtube::YoutubeModel>, error::ErrorResponse> {
    // Отправка запроса на сервер
    let url = format!("{}/youtube/video/popular", self.server_url);
    let response = self.client.get(&url).send().await
      .map_err(|e| {
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
        let videos: Vec<youtube::YoutubeModel> = response.json().await
          .map_err(|e| {
            error::ErrorResponse {
              error: format!("Error deserialization json: {}", e),
              status: 500,
            }
          })?;

        Ok(videos)
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

  // Получение видео youtube по фильтрам
  // -----------------------------------
  async fn get_youtube_videos(&self, category: &str, search: &str, channel: &str) -> std::result::Result<Vec<youtube::YoutubeModel>, error::ErrorResponse> {
    // Отправка запроса на сервер
    let url = format!("{}/youtube/video?categoryId={}&s={}&ch={}&y=", self.server_url, category, search, channel);
    let response = self.client.get(&url).send().await
      .map_err(|e| {
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
        let videos: Vec<youtube::YoutubeModel> = response.json().await
          .map_err(|e| {
            error::ErrorResponse {
              error: format!("Error deserialization json: {}", e),
              status: 500,
            }
          })?;

        Ok(videos)
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
