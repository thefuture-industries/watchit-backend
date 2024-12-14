import { invoke } from "@tauri-apps/api/core";
import { UserAddPayload } from "~/types/user";

class UserService {
  /**
   * Создания пользователя
   */
  public async add_user(user: UserAddPayload): Promise<string> {
    try {
      let response: string = await invoke("add_user", {
        user: user,
      });
      return response;
    } catch (err: any) {
      return err;
    }
  }

  /**
   * Проверка на существования пользователя
   */
  public async check_user(
    ip_address: string,
    create_at: string
  ): Promise<string> {
    try {
      let response: string = await invoke("check_user", {
        ipAddress: ip_address,
        createdAt: create_at,
      });

      return response;
    } catch (err: any) {
      return err;
    }
  }

  /**
   * Получение из сессии uuid пользователя
   */
  public get_uuid(): string {
    const uuid = sessionStorage.getItem("_sess");

    if (!uuid) {
      return "";
    }

    return JSON.parse(uuid).uuid;
  }
}

export default new UserService();
