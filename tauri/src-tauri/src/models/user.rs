use serde::{Deserialize, Serialize};

// User: Моделька пользователя в системе БД
// Id, UserName, UserNameUpper, Email, EmailUpper, IPAddress, Lat, Lon, Country, RegionName, Zip, CreatedAt
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct UserModel {
    // ID - уникальный идентификатор пользователя в системе.
    pub id: u64,

    // UUID - уникальный идентификатор пользователя в системе.
    pub uuid: String,

    // SecretWord - уникальное слова  пользователя в системе.
    pub secret_word: String,

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

    // Country - Страна пользователя.
    pub country: String,

    // RegionName - Город пользователя.
    #[serde(rename = "regionName")]
    pub region_name: String,

    // Zip - Индекс пользователя.
    pub zip: String,

    // CreatedAt - дата и время создания записи пользователя в системе.
    // Используется для отслеживания истории изменений данных пользователя.
    #[serde(rename = "createdAt")]
    pub created_at: String,
}

// IsUser: Моделька пользователя возврата сервера
// uuid, username, email
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct IsUser {
    // UUID - уникальный идентификатор пользователя в системе.
    pub uuid: String,

    // UserName - имя пользователя.
    pub username: String,

    // Email - адрес электронной почты пользователя.
    #[serde(rename = "email", default)]
    pub email: Option<String>,
}

// Тип DTO для добавления данных пользователя
#[derive(Serialize, Deserialize, Debug)]
pub struct UserAddPayload {
    pub secret_word: String,
    pub ip_address: String,
    pub country: String,
    pub region_name: String,
    pub zip: String,
}

// Тип DTO для обновления данных пользователя
#[derive(Serialize, Deserialize, Debug)]
pub struct UserUpdatePayload {
    pub uuid: String,
    #[serde(rename = "username", default)]
    pub username: Option<String>,
    #[serde(rename = "email", default)]
    pub email: Option<String>,
    #[serde(rename = "secret_word", default)]
    pub secret_word: Option<String>,
    #[serde(rename = "secret_word_old", default)]
    pub secret_word_old: Option<String>,
}
