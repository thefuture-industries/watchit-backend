// Импортируем модель фильмов
// из mod файла models/movie.rs
use crate::models::movie;
use dotenv::dotenv;

// --------------------------------------
// Структура данных для работы с фильмами
// --------------------------------------
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
pub trait IMovie {
  // Отправка запроса на сервер
  async fn send_request(&self, url: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
  // Получение с сервера популярных фильмов
  async fn popular_movies(&self, total_page: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
  // Получение фильмов по фильтрам
  async fn get_movies(&self, search: &str, genre: &str, date: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
  // Получение изображение фильма
  async fn image_movie(&self, img: &str) -> std::result::Result<Vec<u8>, String>;
  // Поиск фильмов по title и description
  async fn search_movies(&self, s: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
  // Получение детали фильма по ID
  async fn movie_details(&self, id: i32) -> std::result::Result<movie::MovieModel, String>;
  // Получение похожих фильмов
  async fn similar_movies(&self, genre_id: Vec<i32>, title: &str, overview: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
  // Получение фильмов по сюжету
  async fn get_plot_movies(&self, uuid: &str, text: &str, lege: &str) -> std::result::Result<Vec<movie::MovieModel>, String>;
}

// ---------------------------
// Инициализация этой оболочки
// ---------------------------
impl IMovie for NewStore {
  // --------------------------
  // Отправка запроса на сервер
  async fn send_request(&self, url: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    // Отправка запроса на сервер
    let response = self.client.get(url).send().await.map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: Vec<movie::MovieModel> = response.json().await.map_err(|e| format!("ERROR deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // Получение с сервера популярных фильмов
  // return: 1) массив с фильмами 2) ошибка
  async fn popular_movies(&self, total_page: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    // Отправка запроса на сервер
    let url = format!("{}movies/popular?page={}", self.server_url, total_page);
    self.send_request(&url).await
  }

  // Получение фильмов по фильтрам
  // -----------------------------
  async fn get_movies(&self, search: &str, genre: &str, date: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    let url = format!("{}movies?s={}&genre_id={}&date={}", self.server_url, search, genre, date);
    self.send_request(&url).await
  }

  // Получение изображение фильма
  // ----------------------------
  async fn image_movie(&self, img: &str) -> std::result::Result<Vec<u8>, String> {
    // Отправка запроса на сервер
    let url = format!("{}image/w500/{}", self.server_url, img);

    // Отправка запроса на сервер
    let response = self.client.get(&url).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let image_bytes_result = response.bytes().await;
        match image_bytes_result {
          Ok(image_bytes) => Ok(image_bytes.to_vec()),
          Err(e) => Err(format!("ERROR {}", e)),
        }
      }

      status_code => {
        Err(format!("HTTP ERROR: {} {}", status_code, response.status()))
      }
    }
  }

  // ----------------------
  // Поиск фильмов по title
  async fn search_movies(&self, s: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    // Отправка запроса на сервер
    let url = format!("{}movies?s={}", self.server_url, s);
    self.send_request(&url).await
  }

  // ------------------------
  // Получение деталий фильма
  async fn movie_details(&self, id: i32) -> std::result::Result<movie::MovieModel, String> {
    // Отправка запроса на сервер
    let url = format!("{}movie/{}", self.server_url, id);
    let response = self.client.get(url).send().await.map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: movie::MovieModel = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // -------------------------
  // Получение похожих фильмов
  // -------------------------
  async fn similar_movies(&self, genre_id: Vec<i32>, title: &str, overview: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    let genre_id_str = genre_id.iter()
      .map(|&x| x.to_string())
      .collect::<Vec<String>>()
      .join(",");

    // Отправка запроса на сервер
    let url = format!("{}movies/similar?genre_id={}&title={}&overview={}", self.server_url, genre_id_str, title, overview);

    // Отправка запроса на сервер
    let response = self.client.get(url).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: Vec<movie::MovieModel> = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }

  // ---------------------------
  // Получение фильмов по сюжету
  // ---------------------------
  async fn get_plot_movies(&self, uuid: &str, text: &str, lege: &str) -> std::result::Result<Vec<movie::MovieModel>, String> {
    // Отправка запроса на сервер
    let url = format!("{}text/movies", self.server_url);
    let body = serde_json::json!({
      "uuid": uuid,
      "text": text,
      "lege": lege
    });

    // Отправка запроса на сервер
    let response = self.client.post(url).json(&body).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let movies: Vec<movie::MovieModel> = response.json().await.map_err(|e| format!("ERROR deserialization json: {}", e))?;
        Ok(movies)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default())),
    }
  }
}
