const env = require('./.env.json');
const express = require('express');
const winston = require('winston');
const app = express();
const readLastLines = require('read-last-lines');

app.get('/', function (req, res) {
  res.send('Howdy, world.');
});

app.get('/logs', function (req, res) {
  readLastLines.read('log/chaos.log', 100).then((lines) => {
    res.send(lines);
  });
});

app.get('/errors', function(req, res) {
  readLastLines.read('/home/ec2-user/.pm2/logs/index-error-0.log', 100).then((lines) => {
    res.send(lines);
  });
});

app.get('/pm2', function(req, res) {
  readLastLines.read('/home/ec2-user/.pm2/logs/index-out-0.log', 100).then((lines) => {
    res.send(lines);
  });
});

app.listen(env.webserver.port, function () {
  winston.info(`Example app listening on port ${env.webserver.port}!`);
});

module.exports = app;
