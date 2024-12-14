use crate::models::youtube;
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
  async fn popular_youtube(&mut self) -> std::result::Result<Vec<youtube::YoutubeModel>, std::string::String>;
  // Получение видео youtube по фильтрам
  async fn get_youtube_videos(&self, category: &str, search: &str, channel: &str) -> std::result::Result<Vec<youtube::YoutubeModel>, std::string::String>;
}

// ---------------------------
// Инициализация этой оболочки
// ---------------------------
impl IYoutube for NewStore {
  // Получение популярного видео youtube
  // -----------------------------------
  async fn popular_youtube(&mut self) -> std::result::Result<Vec<youtube::YoutubeModel>, std::string::String> {
    // Отправка запроса на сервер
    let url = format!("{}/youtube/video/popular", self.server_url);
    let response = self.client.get(&url).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let videos: Vec<youtube::YoutubeModel> = response.json().await
          .map_err(|e| format!("Error deserialization json: {}", e))?;

        Ok(videos)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // Получение видео youtube по фильтрам
  // -----------------------------------
  async fn get_youtube_videos(&self, category: &str, search: &str, channel: &str) -> std::result::Result<Vec<youtube::YoutubeModel>, std::string::String> {
    // Отправка запроса на сервер
    let url = format!("{}/youtube/video?categoryId={}&s={}&ch={}&y=", self.server_url, category, search, channel);
    let response = self.client.get(&url).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let videos: Vec<youtube::YoutubeModel> = response.json().await
          .map_err(|e| format!("Error deserialization json: {}", e))?;

        Ok(videos)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }
}
