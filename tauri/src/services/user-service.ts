import { invoke } from "@tauri-apps/api/core";
import { UserModel } from "~/types/user";

class UserService {
  /**
   * Создания пользователя
   */
  public async add_user(user: UserModel): Promise<string> {
    try {
      let response: string = await invoke("add_user", {
        ipAddress: user.ip_address,
        latitude: user.latitude,
        longitude: user.longitude,
        country: user.country,
        regionName: user.region_name,
        zip: user.zip,
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
}

export default new UserService();
