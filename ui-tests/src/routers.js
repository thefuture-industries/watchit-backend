import { fileURLToPath } from 'url';
import { dirname } from 'path';
import path from 'path';

class Routers {
	constructor() {
		this.__filename = fileURLToPath(import.meta.url);
		this.__dirname = dirname(this.__filename);
	}

	home(req, res) {
		res.sendFile(
			path.join(this.__dirname, '../www/template', 'index.html')
		);
	}
}

const routers = new Routers();
export default (router) => {
	router.get('/', routers.home.bind(routers));
};
