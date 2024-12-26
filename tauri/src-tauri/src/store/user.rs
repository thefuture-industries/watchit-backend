use crate::models::user;
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

// -------------------------
// Создания оболочки методов
// -------------------------
pub trait IUser {
  // Создание пользователя
  async fn add_user(&self, user: user::UserAddPayload) -> std::result::Result<user::IsUser, error::ErrorResponse>;
  // Обновление данных пользователя
  async fn update_user(&self, user: user::UserUpdatePayload) -> std::result::Result<user::IsUser, error::ErrorResponse>;
}

impl IUser for NewStore {
  // ---------------------
  // Создание пользователя
  async fn add_user(&self, user: user::UserAddPayload) -> std::result::Result<user::IsUser, error::ErrorResponse> {
    // Отправка запроса на сервер
    let url = format!("{}user/add", self.server_url);
    let body = serde_json::json!({
      "secret_word": &user.secret_word,
      "ip_address": &user.ip_address,
      "country": &user.country,
      "regionName": &user.region_name,
      "zip": &user.zip
    });

    let response = self.client.post(url).json(&body).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        error::ErrorResponse {
          error: "ERROR trying to connect: tcp connect error".to_string(),
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
      reqwest::StatusCode::CREATED => {
        let res: user::IsUser = response.json().await
          .map_err(|e| error::ErrorResponse {
            error: format!("ERROR deserialization json: {}", e),
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

  // ------------------------------
  // Обновление данных пользователя
  async fn update_user(&self, user: user::UserUpdatePayload) -> std::result::Result<user::IsUser, error::ErrorResponse> {
    // Отправка запроса на сервер
    let url = format!("{}user/update", self.server_url);
    let body = serde_json::json!({
      "uuid": &user.uuid,
      "username": &user.username,
      "email": &user.email,
      "secret_word": &user.secret_word,
      "secret_word_old": &user.secret_word_old,
    });

    let response = self.client.put(url).json(&body).send().await.map_err(|e| {
      if e.to_string().contains("connect error") {
        error::ErrorResponse {
          error: "error trying to connect: tcp connect error".to_string(),
          status: 500,
        }
      } else {
        error::ErrorResponse {
          error: format!("error request {}", e),
          status: 500,
        }
      }
    })?;

    // Обработка ошибки
    match response.status() {
      reqwest::StatusCode::OK => {
        let res: user::IsUser = response.json().await
          .map_err(|e| error::ErrorResponse {
            error: format!("ERROR deserialization json: {}", e),
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
