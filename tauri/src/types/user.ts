// User: Моделька пользователя в системе БД
export type UserModel = {
  id: string;
  uuid: string;
  secret_word: string;
  username: string;
  username_upper: string;
  email: string;
  ip_address: string;
  latitude: string;
  longitude: string;
  country: string;
  region_name: string;
  zip: string;
  created_at: string;
};

export type UserAddPayload = {
  secret_word: string;
  ip_address: string;
  latitude: string;
  longitude: string;
  country: string;
  region_name: string;
  zip: string;
};

export type UserUpdatePayload = {
  uuid?: string;
  username?: string;
  email?: string;
  secret_word?: string;
};
