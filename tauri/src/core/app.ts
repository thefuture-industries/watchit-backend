import { invoke } from "@tauri-apps/api/core";

class App {
  public async exit(): Promise<void> {
    await invoke("exist_app");
  }

  public async restart(): Promise<void> {
    await invoke("restart_app");
  }
}

export default new App();
