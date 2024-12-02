use crate::models::user;
use dotenv::dotenv;

#[derive(Clone)]
pub struct NewStore {}


impl NewStore {
  pub fn new() -> Self {
    NewStore {}
  }
}

// -------------------------
// Создания оболочки методов
// -------------------------
pub trait IUser {
  // Создание пользователя
  async fn add_user(&self, secret_word: &str, ip_address: &str, latitude: &str, longitude: &str, country: &str, region_name: &str, zip: &str) -> std::result::Result<String, String>;
  // Проверка на существования пользователя
  async fn check_user(&self, ip_address: &str, created_at: &str) -> std::result::Result<String, String>;
}

impl IUser for NewStore {
  // ---------------------
  // Создание пользователя
  async fn add_user(&self, secret_word: &str, ip_address: &str, latitude: &str, longitude: &str, country: &str, region_name: &str, zip: &str) -> std::result::Result<String, String> {
    dotenv().ok();

    // Создания клиента
    let client = reqwest::Client::new();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}user/add", server);
    let body = serde_json::json!({
      "secret_word": secret_word,
      "ip_address": ip_address,
      "latitude": latitude,
      "longitude": longitude,
      "country": country,
      "regionName": region_name,
      "zip": zip
    });
    let response = client.post(url).json(&body).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::CREATED => {
        let res: std::string::String = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(res)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }

  // --------------------------------------
  // Проверка на существования пользователя
  async fn check_user(&self, ip_address: &str, created_at: &str) -> std::result::Result<String, String> {
    dotenv().ok();

    // Создания клиента
    let client = reqwest::Client::new();

    // Получение url сервера из .env
    let server = std::env::var("SERVER_URL").map_err(|e| format!("Error get env: {}", e))?;

    // Отправка запроса на сервер
    let url = format!("{}user/check", server);
    let body = serde_json::json!({
      "ip_address": ip_address,
      "created_at": created_at,
    });
    let response = client.post(url).json(&body).send().await.map_err(|e| format!("Error request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let res: std::string::String = response.json().await.map_err(|e| format!("Error deserialization json: {}", e))?;
        Ok(res)
      }

      status_code => Err(format!("Error HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }
}
