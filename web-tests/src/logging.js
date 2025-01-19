import winston, { createLogger, format, transports } from 'winston';
const transport = { transports };
import 'winston-daily-rotate-file';
import { config } from './config.server.js';
import util from 'util';

const DEVICE_LABEL = 'DEV';
const CONFIGURATOR_LABEL = 'CONF';
const SESSION_LABEL = 'SESS';

const LOG_ROOT_DIR = config.logging.loggingRootDir;
const LOG_APP_DIR = 'app';
const LOG_EXCEPTION_DIR = 'exception';
const LOG_SESSION_DIR = 'session';
const LOG_TCP_DIR = 'tcp';

const { combine, timestamp, label, printf, splat } = format;

/**
 * Форматирует сообщение о событии сокета для логирования.
 */
function formatSockEventMsg(socket, credentials, eventName, message, ...args) {
	const credentialString = credentials == null ? 'null' : credentials;
	const socketAddr =
		socket == null
			? 'null'
			: util.format('%s:%d', socket.remoteAddress, socket.remotePort);
	const messageString = message ? util.format(message, ...args) : '';

	return util.format(
		"sock: %s. cred: %s. event: '%s'. %s",
		socketAddr,
		credentialString,
		eventName,
		messageString
	);
}

/**
 * Форматирует содержимое буфера в шестнадцатеричную строку.
 * @param {Buffer} dataBuffer - Буфер данных.
 * @param {number} startPosition - Начальная позиция для чтения данных из буфера.
 * @param {number} length - Длина данных для чтения из буфера.
 * @returns {string} - Строка, содержащая шестнадцатеричное представление буфера.
 */
function formatBufferMsg(dataBuffer, startPosition, length) {
	let logString = '';
	for (let i = 0; i < length; i++) {
		logString +=
			('00' + dataBuffer[startPosition + i].toString(16)).substr(-2) +
			' ';
	}
	return logString;
}

/**
 * Создает транспортер для логирования с ротацией файлов по дням.
 * @param {string} logDir - Имя директории для хранения логов (относительно LOG_ROOT_DIR).
 * @returns {transport} - Транспортер для логирования с ротацией файлов.
 */
function createDaylyRotaitTransport(logDir) {
	return new transports.DailyRotateFile({
		datePattern: 'YYYY-MM-DD',
		filename: logDir + '-%DATE%.log',
		dirname: `${LOG_ROOT_DIR}/${logDir}`,
		maxSize: '20m',
		maxFiles: '30d',
	});
}

/**
 * Создает логгер приложения с настроенным форматом и транспортом.
 * @constant
 */
const appLogger = createLogger({
	level: config.logging.loggingLevel,
	format: combine(
		timestamp({ format: 'YYYY.MM.DD HH:mm:ss.SSS' }),
		splat(),
		printf(({ level, message, timestamp }) => {
			return `${timestamp} | ${level}: ${message}`;
		})
	),
	transports: [
		new transports.Console(),
		createDaylyRotaitTransport(LOG_APP_DIR),
	],
});

/**
 * Создает логгер для неперехваченных ошибок с настроенным форматом и транспортом.
 * @constant
 */
const uncaughErrorLogger = createLogger({
	level: 'error',
	format: combine(
		timestamp({ format: 'YYYY.MM.DD HH:mm:ss.SSS' }),
		splat(),
		printf(({ level, message, timestamp }) => {
			return `${timestamp} | ${level}: ${message}`;
		})
	),
	exceptionHandlers: [
		new transports.Console(),
		createDaylyRotaitTransport(LOG_EXCEPTION_DIR),
	],
});

/**
 • Создает специализированный логгер с заданными настройками.
 • @returns {winston.Logger} - Созданный логгер.
 */
function createSpecificLogger(logDirName, logLabel) {
	return createLogger({
		level: config.logging.loggingLevel,
		format: combine(
			label({ label: logLabel }),
			timestamp({ format: 'YYYY.MM.DD HH:mm:ss.SSS' }),
			splat(),
			printf(({ level, message, label, timestamp }) => {
				return `${timestamp} | ${level}. [${label}]: ${message}`;
			})
		),
		transports: [createDaylyRotaitTransport(logDirName)],
	});
}

/**
 • Логгер для TCP устройств.
 • @constant
 */
const tcpDeviceLogger = createSpecificLogger(LOG_TCP_DIR, DEVICE_LABEL);

/**
 • Логгер для TCP конфигуратора.
 • @constant
 */
const tcpConfiguratorLogger = createSpecificLogger(
	LOG_TCP_DIR,
	CONFIGURATOR_LABEL
);

/**
 • Логгер для сессий.
 • @constant
 */
const sessionLogger = createSpecificLogger(LOG_SESSION_DIR, SESSION_LABEL);

/**
 • Добавляет транспорты консоли для логгеров, если приложение запущено не в production.
 */
if (process.env.NODE_ENV !== 'production') {
	tcpDeviceLogger.add(new transports.Console());
	tcpConfiguratorLogger.add(new transports.Console());
	sessionLogger.add(new transports.Console());
}

export default {
	tcp: {
		devLogger: tcpDeviceLogger,
		confLogger: tcpConfiguratorLogger,
	},
	sessionLogger: sessionLogger,
	appLogger: appLogger,
	formatBufferMsg: formatBufferMsg,
	formatSocketEventMsg: formatSockEventMsg,
};
