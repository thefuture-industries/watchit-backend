use crate::models::favourites;
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

pub trait IFavourites {
  // Добовление избанных фильмов
  async fn add_favourites(&self, favourite: favourites::FavouriteAddPayload) -> std::result::Result<String, String>;
  // Получение избанных фильмов
  async fn get_favourites(&self, uuid: &str) -> std::result::Result<Vec<favourites::Favourites>, String>;
  // Удаление избанных фильмов
  async fn delete_favourites(&self, uuid: &str, movie_id: i32) -> std::result::Result<String, String>;
}

impl IFavourites for NewStore {
  // ---------------------------
  // Добовление избанных фильмов
  async fn add_favourites(&self, favourite: favourites::FavouriteAddPayload) -> std::result::Result<String, String> {
    // Отправка запроса на сервер
    let url = format!("{}favourites", self.server_url);
    let body = serde_json::json!({
      "uuid": &favourite.uuid,
      "movieId": &favourite.movie_id,
      "moviePoster": &favourite.movie_poster,
    });

    let response = self.client.post(url).json(&body).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        "error trying to connect: tcp connect error".to_string()
      } else {
        format!("ERROR request {}", e)
      }
    })?;

    match response.status() {
      reqwest::StatusCode::OK => Ok(String::from("Favourite added successfully")),

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }

  // --------------------------
  // Получение избанных фильмов
  async fn get_favourites(&self, uuid: &str) -> std::result::Result<Vec<favourites::Favourites>, String> {
    let url = format!("{}favourites/{}", self.server_url, uuid);
    let response = self.client.get(url).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        "error trying to connect: tcp connect error".to_string()
      } else {
        format!("ERROR request {}", e)
      }
    })?;

    match response.status() {
      reqwest::StatusCode::OK => {
        let res: Vec<favourites::Favourites> = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;

        Ok(res)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }

  // --------------------------
  // Удаление избанных фильмов
  async fn delete_favourites(&self, uuid: &str, movie_id: i32) -> std::result::Result<String, String> {
    let url = format!("{}favourites", self.server_url);
    let body = serde_json::json!({
      "uuid": uuid,
      "movieId": movie_id,
    });

    let response = self.client.delete(url).json(&body).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        "error trying to connect: tcp connect error".to_string()
      } else {
        format!("ERROR request {}", e)
      }
    })?;

    match response.status() {
      reqwest::StatusCode::OK => Ok(String::from("Favourite deleted successfully")),
      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }
}
