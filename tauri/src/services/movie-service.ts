import { invoke } from "@tauri-apps/api/core";
import { MovieModel } from "~/types/movie";
import userService from "./user-service";

// Класс для работы с фильмами
class MovieService {
  /**
   * Проверка на хеш фильмов
   */
  public getMovieArrayFromSessionStorage(): MovieModel[] {
    const storedDataMovies = sessionStorage.getItem("sess_movies");

    if (storedDataMovies) {
      try {
        return JSON.parse(storedDataMovies) as MovieModel[];
      } catch (error) {
        console.error("Error parsing movie array from sessionStorage:", error);
        return [];
      }
    } else {
      return [];
    }
  }

  // /**
  //  * Получение популярных фильмов
  //  */
  // public async get_popular(): Promise<MovieModel[]> {
  //   // Проверка на хеш в sessionStorage
  //   let sessionMovies = this.getMovieArrayFromSessionStorage();

  //   if (sessionMovies.length > 0) {
  //     return sessionMovies;
  //   }

  //   let movie: MovieModel[] = await invoke("get_popular_movies", {
  //     totalPage: "",
  //   });
  //   this.set_movies(movie);

  //   sessionStorage.setItem("m_array", JSON.stringify(movie));
  //   return this.movies;
  // }

  /**
   * Получение фильмов по фильтрам
   */
  public async get(
    search: string,
    genre: string,
    date: string
  ): Promise<MovieModel[]> {
    try {
      let movies: MovieModel[] = await invoke("get_movies", {
        search: search,
        genre: genre,
        date: date,
      });

      return movies;
    } catch (err) {
      return [];
    }
  }

  /**
   * Получение изображение фильмов
   */
  public async image_movie(img: string) {
    let get_img = await invoke("image_movie", {
      img: img,
    });

    return get_img;
  }

  /**
   * Поиск фильма по titla или overview (...)
   */
  public async search(search: string): Promise<MovieModel[]> {
    let movies: MovieModel[] = await invoke("search_movies", {
      s: encodeURIComponent(search),
    });

    return movies;
  }

  /**
   * Получние деталей фильма по ID
   */
  public async details(id: number): Promise<MovieModel> {
    let movie_details: MovieModel = await invoke("movie_details", {
      id: id,
    });

    return movie_details;
  }

  /**
   * Получение похожих фильмов
   */
  public async similar(
    genre_id: Array<number>,
    title: string,
    overview: string
  ): Promise<MovieModel[]> {
    let similar_movies: MovieModel[] = await invoke("similar_movies", {
      genreId: genre_id,
      title: encodeURIComponent(title),
      overview: encodeURIComponent(overview),
    });

    return similar_movies;
  }

  /**
   * Поиск фильмов по сюжету
   */
  public async plot(text: string, lege: string): Promise<MovieModel[]> {
    let movies: MovieModel[] = await invoke("get_plot_movies", {
      uuid: userService.get_uuid(),
      text: encodeURIComponent(text),
      lege: lege,
    });

    return movies;
  }
}

export default new MovieService();
