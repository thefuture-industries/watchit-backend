// User: Моделька пользователя в системе БД
// Id, UserName, UserNameUpper, Email, EmailUpper, IPAddress, Lat, Lon, Country, RegionName, Zip, CreatedAt
export type UserModel = {
  // ID - уникальный идентификатор пользователя в системе.
  id: string | null;

  // UserName - имя пользователя.
  username: string | null;

  // UserNameUpper - имя пользователя в верхнем регистре.
  // Используется для ускорения поиска и сравнения имен пользователей,
  // так как сравнение строк в верхнем регистре быстрее, чем в нижнем.
  username_upper: string | null;

  // Email - адрес электронной почты пользователя.
  email: string | null;

  // EmailUpper - адрес электронной почты пользователя в верхнем регистре.
  // Используется для ускорения поиска и сравнения адресов электронной почты,
  // так как сравнение строк в верхнем регистре быстрее, чем в нижнем.
  email_upper: string | null;

  // IPAddress - ip address пользователей.
  ip_address: string;

  // Lat - Ширина пользователь.
  latitude: string;

  // Lon - Долгота пользователь.
  longitude: string;

  // Country - Страна пользователя.
  country: string;

  // RegionName - Город пользователя.
  region_name: string;

  // Zip - Индекс пользователя.
  zip: string;

  // CreatedAt - дата и время создания записи пользователя в системе.
  // Используется для отслеживания истории изменений данных пользователя.
  created_at: string | null;
};
