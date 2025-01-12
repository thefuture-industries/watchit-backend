class AppConfig {
  public SERVER_URL_WEB: string;
  public SERVER_URL_MOBILE: string;

  constructor() {
    // @type {string}
    this.SERVER_URL_WEB = 'http://localhost:8080/api/v1';
    this.SERVER_URL_MOBILE = `http://192.168.0.5:8080/api/v1`;
  }
}

class ConfigService {
  // returns {AppConfig}
  public ReturnConfig(): AppConfig {
    return new AppConfig();
  }
}

export default new ConfigService();
