import { invoke } from "@tauri-apps/api/core";
import { MovieModel } from "~/types/movie";

// Класс для работы с фильмами
class MovieService {
  /**
   * Массив для хранения фильмов
   */
  private movies: MovieModel[] = [];
  // private s_movies: MovieModel[] = [];

  private total_page_popular: number = 1;
  // private total_page_movie_search: number = 0;

  /**
   * Установка фильмов в массив
   */
  private set_movies(movie: MovieModel[]): void {
    this.movies = [...this.movies, ...movie];
  }

  /**
   * Проверка на хеш фильмов
   */
  private getMovieArrayFromSessionStorage(): MovieModel[] {
    const storedData = sessionStorage.getItem("m_array");

    if (storedData) {
      try {
        return JSON.parse(storedData);
      } catch (error) {
        console.error("Error parsing movie array from sessionStorage:", error);
        return [];
      }
    } else {
      return [];
    }
  }

  /**
   * Получение популярных фильмов
   */
  public async get_popular_movies(): Promise<MovieModel[]> {
    // Проверка на хеш в sessionStorage
    let sessionMovies = this.getMovieArrayFromSessionStorage();

    if (sessionMovies.length > 0) {
      return sessionMovies;
    }

    let movie: MovieModel[] = await invoke("get_popular_movies", {
      totalPage: "",
    });
    this.set_movies(movie);

    sessionStorage.setItem("m_array", JSON.stringify(movie));
    return this.movies;
  }

  /**
   * Получение фильмов по фильтрам
   */
  public async get_movies(
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
   * increment_page_movie_search
   */
  // public async increment_page_movie_search() {
  //   this.total_page_movie_search += 1;

  //   let movies: MovieModel[] = await invoke("get_popular_movies", {
  //     totalPage: this.total_page_popular.toString(),
  //   });

  //   return movies;
  // }

  /**
   * increment_page_popular
   */
  public async increment_page_popular() {
    let sessionMovies: MovieModel[] = this.getMovieArrayFromSessionStorage();
    this.total_page_popular += 1;

    let movie: MovieModel[] = await invoke("get_popular_movies", {
      totalPage: this.total_page_popular.toString(),
    });

    // Запись в хеш
    movie.forEach((singleMovie) => {
      sessionMovies.push(singleMovie);
    });

    // Запись в хеш
    sessionStorage.setItem("m_array", JSON.stringify(sessionMovies));
    this.set_movies(movie);
  }

  /**
   * Показ больше фильмов total_page++(rust);
   */
  public async show_more(): Promise<void> {
    await invoke("increment_page");
    let movie: MovieModel[] = await this.get_popular_movies();
    this.set_movies(movie);
  }

  /**
   * Поиск фильма по titla или overview (...)
   */
  public async search_movies(search: string): Promise<MovieModel[]> {
    sessionStorage.removeItem("m_array");

    let movies: MovieModel[] = await invoke("search_movies", {
      s: encodeURIComponent(search),
    });

    // this.s_movies = [...movies];
    return movies;
  }

  /**
   * Получние деталей фильма по ID
   */
  public async details_movie(id: number): Promise<MovieModel> {
    sessionStorage.removeItem("m_array");

    let movie_details: MovieModel = await invoke("movie_details", {
      id: id,
    });

    return movie_details;
  }

  /**
   * Получение похожих фильмов
   */
  public async similar_movies(
    genre_id: Array<number>,
    title: string,
    overview: string
  ): Promise<MovieModel[]> {
    sessionStorage.removeItem("m_array");

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
  public async plot_movies(text: string, lege: string): Promise<MovieModel[]> {
    let movies: MovieModel[] = await invoke("get_plot_movies", {
      ipAddress: sessionStorage.getItem("_sess"),
      text: encodeURIComponent(text),
      lege: lege,
    });

    return movies;
  }
}

export default new MovieService();
