import { Platform } from 'react-native';
import configService from './config.service';
import axios from 'axios';

class MovieService {
	/**
	 * popular
	 */
	public async popular() {
		const platform = Platform.OS;
		let url = ``;

		if (platform === 'ios' || platform === 'android') {
			url = `${configService.ReturnConfig().SERVER_URL_MOBILE}/movies/popular?page=3`;
		} else {
			url = `${configService.ReturnConfig().SERVER_URL_WEB}/movies/popular?page=3`;
		}

		const response = await axios.get(url);
		return response.data;
	}
}

export default new MovieService();
