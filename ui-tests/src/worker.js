import { parentPort, workerData } from 'worker_threads';
import fetch from 'node-fetch';

async function performTask(task) {
	return new Promise((resolve) => {
		setTimeout(() => {
			resolve(task);
		}, 0);
	});
}

async function main() {
	const { tasks } = workerData;

	let COMPLETED_TASK = 0;

	for (const task of tasks) {
		const response = await fetch('https://google.com/', {
			method: 'GET',
			headers: { 'Content-Type': 'application/json' },
		});

		if (!response.ok) {
			console.error(
				`Error on fetch ${response.status} ${response.statusText}`
			);
		} else {
			const data = await response.json();
		}

		await performTask(task);
		COMPLETED_TASK++;
	}

	parentPort.postMessage({ COMPLETED_TASK });
}

main();
