import { invoke } from "@tauri-apps/api/core";
import { YoutubeModel } from "~/types/youtube";

class YoutubeService {
  private videos: YoutubeModel[] = [];

  /**
   * // Получение популярного видео с ютуба
   */
  public async get_video(): Promise<YoutubeModel[]> {
    const cacheVideo = sessionStorage.getItem("pop__v1");

    if (cacheVideo) {
      try {
        return JSON.parse(cacheVideo) as YoutubeModel[];
      } catch (err) {
        return [];
      }
    }

    try {
      const video = await invoke("get_popular_video");
      sessionStorage.setItem("pop__v1", JSON.stringify(video));

      return video as YoutubeModel[];
    } catch (err) {
      return [];
    }
  }

  /**
   * Получение массив видео с ютуба
   */
  public async get_videos(): Promise<YoutubeModel[]> {
    if (this.videos.length == 0) {
      let video: YoutubeModel[] = await invoke("");
      this.set_videos(video);
    }

    return this.videos;
  }

  /**
   * name
   */
  public async get_youtube_videos(
    category: string,
    search: string,
    channel: string
  ): Promise<YoutubeModel[]> {
    let videos: YoutubeModel[] = await invoke("get_youtube_videos", {
      category: category,
      search: search,
      channel: channel,
    });

    return videos;
  }

  // Установка в массив видео
  private set_videos(videos: YoutubeModel[]): void {
    this.videos = [...this.videos, ...videos];
  }
}

export default new YoutubeService();
