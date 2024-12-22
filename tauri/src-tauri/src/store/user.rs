use crate::models::user;
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

// -------------------------
// Создания оболочки методов
// -------------------------
pub trait IUser {
  // Создание пользователя
  async fn add_user(&self, user: user::UserAddPayload) -> std::result::Result<user::IsUser, String>;
  // Проверка на существования пользователя
  async fn check_user(&self, ip_address: &str, created_at: &str) -> std::result::Result<String, String>;
  // Обновление данных пользователя
  async fn update_user(&self, user: user::UserUpdatePayload) -> std::result::Result<user::IsUser, String>;
}

impl IUser for NewStore {
  // ---------------------
  // Создание пользователя
  async fn add_user(&self, user: user::UserAddPayload) -> std::result::Result<user::IsUser, String> {
    // Отправка запроса на сервер
    let url = format!("{}user/add", self.server_url);
    let body = serde_json::json!({
      "secret_word": &user.secret_word,
      "ip_address": &user.ip_address,
      "latitude": &user.latitude,
      "longitude": &user.longitude,
      "country": &user.country,
      "regionName": &user.region_name,
      "zip": &user.zip
    });

    let response = self.client.post(url).json(&body).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::CREATED => {
        let res: user::IsUser = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;

        Ok(res)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }

  // --------------------------------------
  // Проверка на существования пользователя
  async fn check_user(&self, ip_address: &str, created_at: &str) -> std::result::Result<String, String> {
    // Отправка запроса на сервер
    let url = format!("{}user/check", self.server_url);
    let body = serde_json::json!({
      "ip_address": ip_address,
      "created_at": created_at,
    });

    let response = self.client.post(url).json(&body).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let res: std::string::String = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;

        Ok(res)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }

  // ------------------------------
  // Обновление данных пользователя
  async fn update_user(&self, user: user::UserUpdatePayload) -> std::result::Result<user::IsUser, String> {
    // Отправка запроса на сервер
    let url = format!("{}user/update", self.server_url);
    let body = serde_json::json!({
      "uuid": &user.uuid,
      "username": &user.username,
      "email": &user.email,
      "secret_word": &user.secret_word,
    });

    let response = self.client.put(url).json(&body).send().await
      .map_err(|e| format!("ERROR request: {}", e))?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let res: user::IsUser = response.json().await
          .map_err(|e| format!("ERROR deserialization json: {}", e))?;

        Ok(res)
      }

      status_code => Err(format!("ERROR HTTP: {} {}", status_code, response.text().await.unwrap_or_default()))
    }
  }
}
