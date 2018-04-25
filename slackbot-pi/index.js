var env = require('./.env.json');
var winston = require('winston');
var slack = require('./slack');
var webserver = require('./webserver');
var updates = require('./updates');
var mongoose = require('mongoose');

mongoose.connection.openUri('mongodb://raspibot:raspibot1@ds257589.mlab.com:57589/raspibot')

winston.add(winston.transports.File, {filename: 'log/chaos.log'});

var currentTime = new Date().toISOString();
winston.info('Started L40 Chaos Bot - ' + currentTime);

// Main loop
if(env.tasks[0].enabled){
  setInterval(updates.checkUpdates, 300000); // Retry every 5 min
}
