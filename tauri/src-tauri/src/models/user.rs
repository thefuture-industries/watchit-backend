use serde::{Deserialize,Serialize};

// User: Моделька пользователя в системе БД
// Id, UserName, UserNameUpper, Email, EmailUpper, IPAddress, Lat, Lon, Country, RegionName, Zip, CreatedAt
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct UserModel {
  // ID - уникальный идентификатор пользователя в системе.
  pub id: u64,

  // UUID - уникальный идентификатор пользователя в системе.
  pub uuid: String,

  // SecretWord - уникальное слова  пользователя в системе.
  pub secret_word: u64,

  // UserName - имя пользователя.
  pub username: String,

  // UserNameUpper - имя пользователя в верхнем регистре.
	// Используется для ускорения поиска и сравнения имен пользователей,
	// так как сравнение строк в верхнем регистре быстрее, чем в нижнем.
  pub username_upper: String,

  // Email - адрес электронной почты пользователя.
  #[serde(rename = "email", default)]
  pub email: Option<String>,

  // EmailUpper - адрес электронной почты пользователя в верхнем регистре.
	// Используется для ускорения поиска и сравнения адресов электронной почты,
	// так как сравнение строк в верхнем регистре быстрее, чем в нижнем.
  #[serde(rename = "email_upper", default)]
  pub email_upper: Option<String>,

  // IPAddress - ip address пользователей.
  pub ip_address: String,

  // Lat - Ширина пользователь.
  pub latitude: String,

  // Lon - Долгота пользователь.
  pub longitude: String,

  // Country - Страна пользователя.
  pub country: String,

  // RegionName - Город пользователя.
  #[serde(rename="regionName")]
  pub region_name: String,

  // Zip - Индекс пользователя.
  pub zip: String,

	// CreatedAt - дата и время создания записи пользователя в системе.
	// Используется для отслеживания истории изменений данных пользователя.
  #[serde(rename="createdAt")]
  pub created_at: String,
}
