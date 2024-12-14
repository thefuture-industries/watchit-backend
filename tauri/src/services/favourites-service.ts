import { FavouriteAddPayload, FavouriteModel } from "~/types/favourites";
import userService from "./user-service";
import { invoke } from "@tauri-apps/api/core";

class FavouritesService {
  /**
   * Получение избранных фильмов
   */
  public async get(): Promise<FavouriteModel[]> {
    // Получение uuid пользователя
    const uuid = userService.get_uuid();

    // Получение избранных фильмов с сервера
    const favourites: FavouriteModel[] = await invoke("get_favourites", {
      uuid,
    });
    return favourites;
  }

  /**
   * Добавление фильма в избранное
   */
  public async add(favourites: FavouriteAddPayload): Promise<void> {
    const payload: FavouriteAddPayload = {
      uuid: userService.get_uuid(),
      movie_id: favourites.movie_id,
      movie_poster: favourites.movie_poster,
    };

    // Добавление фильма в избранное на сервере
    await invoke("add_favourites", {
      favourite: payload,
    });
  }

  /**
   * Удаление фильма из избранного
   */
  public async delete(): Promise<void> {
    return;
  }
}

export default new FavouritesService();
