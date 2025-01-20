import { Worker, workerData } from 'worker_threads';
import os from 'os';

class WorkerJS {
	constructor(numThreads, numTasksPerThread) {
		this.numThreads = numThreads;
		this.numTasksPerThread = numTasksPerThread;

		this.activeThreads = 0;
		this.completedTasks = 0;

		this.startTime = Date.now();
	}

	createTask(i) {
		return new Promise((resolve) => {
			setTimeout(() => {
				resolve(`Task ${i} completed`);
			}, 0);
		});
	}

	runWorker(tasks) {
		return new Promise((resolve, reject) => {
			const worker = new Worker('./worker.js', {
				workerData: { tasks },
			});

			this.activeThreads++;

			worker.on('message', (message) => {
				this.completedTasks += message.completedTasks;
			});

			worker.on('error', (error) => {
				this.activeThreads--;
				reject(`Worker error: ${error}`);
			});

			worker.on('exit', (code) => {
				activeThreads--;

				if (code !== 0) {
					reject(`Worker stopped with exit code ${code}`);
				}

				resolve();
			});
		});
	}

	async main() {
		const cpuCount = os.cpus().length;

		console.log('num threads: ', this.numThreads);
		console.log('num tasks per thread: ', this.numTasksPerThread);
		console.log('number of cpu cores: ', cpuCount);

		const promises = [];
		for (let i = 0; i < this.numThreads; i++) {
			const tasks = [];
			for (let j = 0; j < this.numTasksPerThread; j++) {
				tasks.push({ id: j, task: j });
			}

			promises.push(this.runWorker(tasks));
		}

		await Promise.all(promises);

		const endTime = Date.now();
		const elapsedTime = (endTime - this.startTime) / 1000;

		console.log(`Total completed tasks: ${this.completedTasks}`);
		console.log(`Elapsed time: ${elapsedTime} seconds`);
	}
}

const worker = new WorkerJS(10, 1000);
worker.main();
