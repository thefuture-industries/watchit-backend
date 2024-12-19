import { RecommendationAddPayload } from "~/types/recommendations";
import userService from "./user-service";
import { invoke } from "@tauri-apps/api/core";
import { MovieModel } from "~/types/movie";
import movieService from "./movie-service";

class RecommendationsService {
  /**
   * Получение рекомендаций
   */
  public async get(): Promise<MovieModel[]> {
    const uuid = userService.get_uuid();
    const cacheMovies = movieService.getMovieArrayFromSessionStorage();

    if (cacheMovies.length > 0) {
      return cacheMovies;
    }

    const recommendations: MovieModel[] = await invoke("get_recommendations", {
      uuid: uuid,
    });

    sessionStorage.setItem("sess_movies", JSON.stringify(recommendations));
    return recommendations;
  }

  /**
   * Добавление рекомендаций
   */
  public async add(recommendations: RecommendationAddPayload): Promise<void> {
    const payload: RecommendationAddPayload = {
      uuid: userService.get_uuid(),
      title: recommendations.title,
      genre: recommendations.genre,
    };

    await invoke("add_recommendations", {
      recom: payload,
    });
  }
}

export default new RecommendationsService();
