var env = require('./.env.json');
var winston = require('winston');
var slack = require('./slack');
var webserver = require('./webserver');
var updates = require('./updates');
var mongoose = require('mongoose');

mongoose.connect('mongodb://10.28.6.16/raspiBot');

winston.add(winston.transports.File, {filename: 'log/chaos.log'});

var currentTime = new Date().toISOString();
winston.info('Started L40 Chaos Bot - ' + currentTime);

// Main loop
if(env.tasks[0].enabled){
  setInterval(updates.checkUpdates, 300000); // Retry every 5 min
}
