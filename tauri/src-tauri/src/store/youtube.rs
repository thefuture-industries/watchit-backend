use crate::models::youtube;
use dotenv::dotenv;

// ----------------------------------------
// Структура данных для работы с ютуб видео
// ----------------------------------------
#[derive(Clone)]
pub struct NewStore {}

// ------------------------------
// Инициализация default значений
// ------------------------------
impl NewStore {
  pub fn new() -> Self {
    NewStore {}
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
    dotenv().ok();

    // Создания клиента
    let client = reqwest::Client::new();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}/youtube/video/popular", server);
    let response = client.get(&url).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let videos: Vec<youtube::YoutubeModel> = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(videos)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // Получение видео youtube по фильтрам
  // -----------------------------------
  async fn get_youtube_videos(&self, category: &str, search: &str, channel: &str) -> std::result::Result<Vec<youtube::YoutubeModel>, std::string::String> {
    dotenv().ok();

    // Создания клиента
    let client = reqwest::Client::new();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}/youtube/video?categoryId={}&s={}&ch={}&y=", server, category, search, channel);
    let response = client.get(&url).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let videos: Vec<youtube::YoutubeModel> = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(videos)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }
}
