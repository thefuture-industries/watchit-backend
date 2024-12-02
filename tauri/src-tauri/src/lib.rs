mod store;
mod models;

use tokio::runtime::Runtime;
use models::{movie::{MovieModel},youtube::{YoutubeModel},user::{UserModel}};
use store::movie;
use store::youtube;
use store::user;
use crate::store::movie::IMovie;
use crate::store::youtube::IYoutube;
use crate::store::user::IUser;

// Создлание команды для использования из React
// Получение массива популярных фильмов
#[tauri::command]
fn get_popular_movies(total_page: &str) -> std::result::Result<Vec<MovieModel>, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime {}", e)),
  };

  let movies = rt.block_on(store.popular_movies(total_page));
  match movies {
    Ok(m) => Ok(m),
    Err(err) => Err(format!("Error {}", err))
  }
}

// -----------------------------------
// Получение популярного видео youtube
// -----------------------------------
#[tauri::command]
fn get_popular_video() -> std::result::Result<Vec<YoutubeModel>, std::string::String> {
  let mut store = youtube::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let videos = rt.block_on(store.popular_youtube());
  match videos {
    Ok(v) => Ok(v),
    Err(err) => Err(format!("Error {}", err)),
  }
}

// -----------------------------------
// Получение видео youtube по фильтрам
// -----------------------------------
#[tauri::command]
fn get_youtube_videos(category: &str, search: &str, channel: &str) -> std::result::Result<Vec<YoutubeModel>, std::string::String> {
  let mut store = youtube::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let videos = rt.block_on(store.get_youtube_videos(category, search, channel));
  match videos {
    Ok(v) => Ok(v),
    Err(err) => Err(format!("Error {}", err)),
  }
}

// -------------------------
// Получение изоюражений фильмов
// -------------------------
#[tauri::command]
fn image_movie(img: &str) -> std::result::Result<Vec<u8>, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let image = rt.block_on(store.image_movie(img));
  match image {
    Ok(v) => Ok(v),
    Err(err) => Err(format!("Error {}", err)),
  }
}

// ----------------------
// Поиск фильмов по title
// ----------------------
#[tauri::command]
fn search_movies(s: &str) -> std::result::Result<Vec<MovieModel>, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let search_movies = rt.block_on(store.search_movies(s));
  match search_movies {
    Ok(m) => Ok(m),
    Err(err) => Err(format!("Error {}", err))
  }
}

// -------------------------
// Получение деталий фильмов
// -------------------------
#[tauri::command]
fn movie_details(id: i32) -> std::result::Result<MovieModel, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let movie_detail = rt.block_on(store.movie_details(id));
  match movie_detail {
    Ok(md) => Ok(md),
    Err(err) => Err(format!("Error {}", err))
  }
}

// -------------------------
// Получение похожих фильмов
// -------------------------
#[tauri::command]
fn similar_movies(genre_id: Vec<i32>, title: &str, overview: &str) -> std::result::Result<Vec<MovieModel>, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let movie_similar = rt.block_on(store.similar_movies(genre_id, title, overview));
  match movie_similar {
    Ok(ms) => Ok(ms),
    Err(err) => Err(format!("Error {}", err))
  }
}

// -------------------------
// Получение похожих фильмов
// -------------------------
#[tauri::command]
fn get_movies(search: &str, genre: &str, date: &str) -> std::result::Result<Vec<MovieModel>, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let movies = rt.block_on(store.get_movies(search, genre, date));
  match movies {
    Ok(m) => Ok(m),
    Err(err) => Err(format!("Error {}", err))
  }
}

// ---------------------------
// Получение фильмов по сюжету
// ---------------------------
#[tauri::command]
fn get_plot_movies(ip_address: &str, text: &str, lege: &str) -> std::result::Result<Vec<MovieModel>, std::string::String> {
  let mut store = movie::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e)),
  };

  let movies = rt.block_on(store.get_plot_movies(ip_address, text, lege));
  match movies {
    Ok(m) => Ok(m),
    Err(err) => Err(format!("Error {}", err))
  }
}

// --------------------------------------
// Проверка на существования пользователя
// --------------------------------------
#[tauri::command]
fn check_user(ip_address: &str, created_at: &str) -> std::result::Result<String, String> {
  let mut store = user::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e))
  };

  let response = rt.block_on(store.check_user(ip_address, created_at));
  match response {
    Ok(res) => Ok(res),
    Err(err) => Err(format!("Error {}", err))
  }
}

// ---------------------
// Создание пользователя
// ---------------------
#[tauri::command]
fn add_user(secret_word: &str, ip_address: &str, latitude: &str, longitude: &str, country: &str, region_name: &str, zip: &str) -> std::result::Result<String, String> {
  let mut store = user::NewStore::new();

  let rt = match Runtime::new() {
    Ok(rt) => rt,
    Err(e) => return Err(format!("Error create Runtime: {}", e))
  };

  let response = rt.block_on(store.add_user(secret_word, ip_address, latitude, longitude, country, region_name, zip));
  match response {
    Ok(res) => Ok(res),
    Err(err) => Err(format!("Error {}", err))
  }
}

// -------------------
// Выход из приложения
// -------------------
#[tauri::command]
fn exist_app(app_handle: tauri::AppHandle) {
  app_handle.exit(0);
}
// -----------------------
// Перезагрузка приложения
// -----------------------
#[tauri::command]
fn restart_app(app_handle: tauri::AppHandle) {
  app_handle.restart();
}

// Старт приложения
// Подключение команд
#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_shell::init())
        .invoke_handler(tauri::generate_handler![
                                                // MOVIES
                                                get_popular_movies,
                                                image_movie,
                                                search_movies,
                                                movie_details,
                                                similar_movies,
                                                get_movies,
                                                get_plot_movies,
                                                // YOUTUBE
                                                get_popular_video,
                                                get_youtube_videos,
                                                // USER
                                                add_user,
                                                check_user,
                                                // APP
                                                exist_app,
                                                restart_app])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
