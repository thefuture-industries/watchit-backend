// Импортируем модель фильмов
// из mod файла models/movie.rs
use crate::models::movie;
use dotenv::dotenv;

// --------------------------------------
// Структура данных для работы с фильмами
// --------------------------------------
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
pub trait IMovie {
  // Отправка запроса на сервер
  async fn send_request(&self, url: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String>;
  // Получение с сервера популярных фильмов
  async fn popular_movies(&self, total_page: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String>;
  // Получение фильмов по фильтрам
  async fn get_movies(&self, search: &str, genre: &str, date: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String>;
  // Получение изображение фильма
  async fn image_movie(&self, img: &str) -> std::result::Result<Vec<u8>, std::string::String>;
  // Поиск фильмов по title и description
  async fn search_movies(&self, s: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String>;
  // Получение детали фильма по ID
  async fn movie_details(&self, id: i32) -> std::result::Result<movie::MovieModel, std::string::String>;
  // Получение похожих фильмов
  async fn similar_movies(&self, genre_id: Vec<i32>, title: &str, overview: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
  // Получение фильмов по сюжету
  async fn get_plot_movies(&self, ip_address: &str, text: &str, lege: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
}

// ---------------------------
// Инициализация этой оболочки
// ---------------------------
impl IMovie for NewStore { // for NewStore
  // --------------------------
  // Отправка запроса на сервер
  async fn send_request(&self, url: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String> {
    // Создания клиента
    let client = reqwest::Client::new();

    // Отправка запроса на сервер
    let response = client.get(url).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: Vec<movie::MovieModel> = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // Получение с сервера популярных фильмов
  // return: 1) массив с фильмами 2) ошибка
  async fn popular_movies(&self, total_page: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String> {
    dotenv().ok();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}movies/popular?page={}", server, total_page);
    self.send_request(&url).await
  }

  // Получение фильмов по фильтрам
  // -----------------------------
  async fn get_movies(&self, search: &str, genre: &str, date: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String> {
    dotenv().ok();

    // Получение url на сервер
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    let url = format!("{}movies?s={}&genre_id={}&date={}", server, search, genre, date);
    self.send_request(&url).await
  }

  // Получение изображение фильма
  // ----------------------------
  async fn image_movie(&self, img: &str) -> std::result::Result<Vec<u8>, std::string::String> {
    dotenv().ok();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}image/w500/{}", server, img);
    // Создания клиента
    let client = reqwest::Client::new();

    // Отправка запроса на сервер
    let response = client.get(url).send().await.map_err(|e| format!("Error request: {}", e))?;
    // let response_clone = response.clone();

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let image_bytes_result = response.bytes().await;
        match image_bytes_result {
          Ok(image_bytes) => Ok(image_bytes.to_vec()),
          Err(e) => Err(format!("Error {}", e)),
        }

        // Ok(image_bytes.to_vec())
      }

      status_code => {
        Err(format!("HTTP error: {} {}", status_code, response.status()))
      }
    }
  }

  // ----------------------
  // Поиск фильмов по title
  async fn search_movies(&self, s: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String> {
    dotenv().ok();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}movies?s={}", server, s);
    self.send_request(&url).await
  }

  // ------------------------
  // Получение деталий фильма
  async fn movie_details(&self, id: i32) -> std::result::Result<movie::MovieModel, std::string::String> {
    dotenv().ok();
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Создания клиента
    let client = reqwest::Client::new();

    // Отправка запроса на сервер
    let url = format!("{}movie/{}", server, id);
    let response = client.get(url).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: movie::MovieModel = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // -------------------------
  // Получение похожих фильмов
  // -------------------------
  async fn similar_movies(&self, genre_id: Vec<i32>, title: &str, overview: &str) -> std::result::Result<Vec<movie::MovieModel>, std::string::String> {
    dotenv().ok();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    let genre_id_str = genre_id.iter()
      .map(|&x| x.to_string())
      .collect::<Vec<String>>()
      .join(",");

    // Отправка запроса на сервер
    let url = format!("{}movies/similar?genre_id={}&title={}&overview={}", server, genre_id_str, title, overview);

    // Создания клиента
    let client = reqwest::Client::new();

    // Отправка запроса на сервер
    let response = client.get(url).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: Vec<movie::MovieModel> = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // ---------------------------
  // Получение фильмов по сюжету
  // ---------------------------
  async fn get_plot_movies(&self, ip_address: &str, text: &str, lege: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    dotenv().ok();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}text/movies", server);
    let body = serde_json::json!({
      "ip_address": ip_address,
      "text": text,
      "lege": lege
    });

    // Создания клиента
    let client = reqwest::Client::new();

    // Отправка запроса на сервер
    let response = client.post(url).json(&body).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: Vec<movie::MovieModel> = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }
}
