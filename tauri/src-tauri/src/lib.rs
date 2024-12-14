mod models;
mod store;

use crate::store::favourites::IFavourites;
use crate::store::movie::IMovie;
use crate::store::recommendations::IRecommendations;
use crate::store::user::IUser;
use crate::store::youtube::IYoutube;

use models::{
    favourites::{FavouriteAddPayload, Favourites},
    movie::MovieModel,
    recommendations::{RecommendationAddPayload, Recommendations},
    user::{IsUser, UserAddPayload},
    youtube::YoutubeModel,
};

use store::favourites;
use store::movie;
use store::recommendations;
use store::user;
use store::youtube;

// Создлание команды для использования из React
// Получение массива популярных фильмов
#[tauri::command]
fn get_popular_movies(total_page: &str) -> std::result::Result<Vec<MovieModel>, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _movies = match rt.block_on(store.popular_movies(total_page)) {
        Ok(m) => return Ok(m),
        Err(e) => return Err(format!("{}", e)),
    };
}

// -----------------------------------
// Получение популярного видео youtube
// -----------------------------------
#[tauri::command]
fn get_popular_video() -> std::result::Result<Vec<YoutubeModel>, String> {
    let mut store = youtube::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _videos = match rt.block_on(store.popular_youtube()) {
        Ok(v) => return Ok(v),
        Err(e) => return Err(format!("{}", e)),
    };
}

// -----------------------------------
// Получение видео youtube по фильтрам
// -----------------------------------
#[tauri::command]
fn get_youtube_videos(
    category: &str,
    search: &str,
    channel: &str,
) -> std::result::Result<Vec<YoutubeModel>, String> {
    let store = youtube::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _videos = match rt.block_on(store.get_youtube_videos(category, search, channel)) {
        Ok(v) => return Ok(v),
        Err(e) => return Err(format!("{}", e)),
    };
}

// -------------------------
// Получение изоюражений фильмов
// -------------------------
#[tauri::command]
fn image_movie(img: &str) -> std::result::Result<Vec<u8>, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _image = match rt.block_on(store.image_movie(img)) {
        Ok(v) => return Ok(v),
        Err(e) => return Err(format!("{}", e)),
    };
}

// ----------------------
// Поиск фильмов по title
// ----------------------
#[tauri::command]
fn search_movies(s: &str) -> std::result::Result<Vec<MovieModel>, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _search_movies = match rt.block_on(store.search_movies(s)) {
        Ok(m) => return Ok(m),
        Err(e) => return Err(format!("{}", e)),
    };
}

// -------------------------
// Получение деталий фильмов
// -------------------------
#[tauri::command]
fn movie_details(id: i32) -> std::result::Result<MovieModel, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _movie_detail = match rt.block_on(store.movie_details(id)) {
        Ok(md) => return Ok(md),
        Err(e) => return Err(format!("{}", e)),
    };
}

// -------------------------
// Получение похожих фильмов
// -------------------------
#[tauri::command]
fn similar_movies(
    genre_id: Vec<i32>,
    title: &str,
    overview: &str,
) -> std::result::Result<Vec<MovieModel>, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _movie_similar = match rt.block_on(store.similar_movies(genre_id, title, overview)) {
        Ok(ms) => return Ok(ms),
        Err(e) => return Err(format!("{}", e)),
    };
}

// -------------------------
// Получение похожих фильмов
// -------------------------
#[tauri::command]
fn get_movies(
    search: &str,
    genre: &str,
    date: &str,
) -> std::result::Result<Vec<MovieModel>, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _movies = match rt.block_on(store.get_movies(search, genre, date)) {
        Ok(m) => return Ok(m),
        Err(e) => return Err(format!("{}", e)),
    };
}

// ---------------------------
// Получение фильмов по сюжету
// ---------------------------
#[tauri::command]
fn get_plot_movies(
    uuid: &str,
    text: &str,
    lege: &str,
) -> std::result::Result<Vec<MovieModel>, String> {
    let store = movie::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _movies = match rt.block_on(store.get_plot_movies(uuid, text, lege)) {
        Ok(m) => return Ok(m),
        Err(e) => return Err(format!("{}", e)),
    };
}

// --------------------------------------
// Проверка на существования пользователя
// --------------------------------------
#[tauri::command]
fn check_user(ip_address: &str, created_at: &str) -> std::result::Result<String, String> {
    let store = user::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _response = match rt.block_on(store.check_user(ip_address, created_at)) {
        Ok(res) => return Ok(res),
        Err(e) => return Err(format!("{}", e)),
    };
}

// ---------------------
// Создание пользователя
// ---------------------
#[tauri::command]
fn add_user(user: UserAddPayload) -> std::result::Result<IsUser, String> {
    let store = user::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _response = match rt.block_on(store.add_user(user)) {
        Ok(res) => return Ok(res),
        Err(e) => return Err(format!("{}", e)),
    };
}

// ---------------------------
// Добовление избанных фильмов
// ---------------------------
#[tauri::command]
fn add_favourites(favourite: FavouriteAddPayload) -> std::result::Result<String, String> {
    let store = favourites::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _response = match rt.block_on(store.add_favourites(favourite)) {
        Ok(r) => return Ok(r),
        Err(e) => return Err(format!("{}", e)),
    };
}

// ---------------------------
// Получение избанных фильмов
// ---------------------------
#[tauri::command]
fn get_favourites(uuid: &str) -> std::result::Result<Vec<Favourites>, String> {
    let store = favourites::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _response = match rt.block_on(store.get_favourites(uuid)) {
        Ok(r) => return Ok(r),
        Err(e) => return Err(format!("{}", e)),
    };
}

#[tauri::command]
fn add_recommendations(recom: RecommendationAddPayload) -> std::result::Result<(), String> {
    let store = recommendations::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _response = match rt.block_on(store.add_recommendations(recom)) {
        Ok(r) => return Ok(r),
        Err(e) => return Err(format!("{}", e)),
    };
}

#[tauri::command]
fn get_recommendations(uuid: &str) -> std::result::Result<Vec<MovieModel>, String> {
    let store = recommendations::NewStore::new();
    let rt = tokio::runtime::Runtime::new().map_err(|e| format!("{}", e))?;

    let _response = match rt.block_on(store.get_recommendations(uuid)) {
        Ok(r) => return Ok(r),
        Err(e) => return Err(format!("{}", e)),
    };
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
            // FAVOURITES
            add_favourites,
            get_favourites,
            // RECOMMENDATIONS
            add_recommendations,
            get_recommendations,
            // APP
            exist_app,
            restart_app
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
