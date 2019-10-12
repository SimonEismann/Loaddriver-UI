export default class Job {
  constructor(id, slaves, scriptName, intensityFile) {
    this.id = id;
    this.slaves = slaves;
    this.scriptName = scriptName;
    this.intensityFile = intensityFile;
  }
}