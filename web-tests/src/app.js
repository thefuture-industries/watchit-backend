import { config } from './config.server.js';
import { fileURLToPath } from 'url';
import logging from './logging.js';
import routers from './routers.js';
import { dirname } from 'path';
import express from 'express';
import http from 'http';
import path from 'path';
import util from 'util';

// Инициализация Express приложения
const app = express();

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

/**
 * Дефолт настройки сервера
 */
app.use(express.static(path.join(__dirname, '..', 'www')));
app.use(express.json());

/**
 * Конфигурация маршрутов API
 * Все маршруты имеют префикс '/'
 * Версия API v1
 */
routers(app);

/**
 * Инициализация сервера
 * Создает и запускает HTTP сервер с настроенными хостом и портом
 * Конфигурация загружается из config.server
 */
const server = http.createServer(app);
server.listen(config.server.port, config.server.host, () => {
	console.log(`        _   _            _   _
  _   _(_) | |_ ___  ___| |_(_)_ __   __ _
 | | | | | | __/ _ \\/ __| __| |  _ \\ / _  |
 | |_| | | | ||  __/\\__ \\ |_| | | | | (_| |
  \\__,_|_| \\___\\___||___/\\__|_|_| |_|\__,  |
                                     |___/ `);

	logging.appLogger.info(
		util.format(
			'Server started: %s:%d',
			config.server.host,
			config.server.port
		)
	);
});
