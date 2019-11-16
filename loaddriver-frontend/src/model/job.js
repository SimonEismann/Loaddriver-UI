export default class Job {
  constructor(id, slaves, scriptName, intensityFile,
    warmupDuration = 30, warmupPause = 5,
    warmupRate = 0, threads = 128, timeout = 0, randomizeUsers = false) {
    this.id = id;
    this.slaves = slaves;
    this.scriptName = scriptName;
    this.intensityFile = intensityFile;
    this.warmupDuration = warmupDuration;
    this.warmupPause = warmupPause;
    this.warmupRate = warmupRate;
    this.threads = threads;
    this.timeout = timeout;
    this.randomizeUsers = randomizeUsers;
  }
}