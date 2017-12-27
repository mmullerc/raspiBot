var RtmClient = require('@slack/client').RtmClient;
var WebClient = require('@slack/client').WebClient;
var CLIENT_EVENTS = require('@slack/client').CLIENT_EVENTS;
var RTM_EVENTS = require('@slack/client').RTM_EVENTS;
var winston = require('winston');
var env = require('./.env.json');
var updates = require('./updates');
var pi = require('./raspberry/raspberry.js');

var botToken = env.SLACK_BOT_TOKEN || '';
var web = new WebClient(botToken);
var channel;
var userid;
var connected;
var rtm;
var users;

var tryToConnect = function() {
  if (botToken) {
    rtm = new RtmClient(botToken);

    rtm.on(CLIENT_EVENTS.RTM.AUTHENTICATED, (rtmStartData) => {
      users = rtmStartData.users;
      for (var c of rtmStartData.groups) {
        if (c.name ==='lot40') {
          channel = c.id;
        }
      }

      userid = rtmStartData.self.id;
      channel = 'D5W4WB8KU';
      winston.info('I am user # ' + userid);
    });

    // you need to wait for the client to fully connect before you can send messages
    rtm.on(CLIENT_EVENTS.RTM.RTM_CONNECTION_OPENED, function () {
      connected = true;
      winston.log('Connected to slack.');
      rtm.sendMessage("RaspberryPi car initialized", channel);
    });

    rtm.on(RTM_EVENTS.MESSAGE, function handleRtmMessage(message) {
      winston.log(message);

      if (message.text) {

        //Send commands to an pi board
        if (message.text.toLowerCase().startsWith('pi ')) {
          pi.raspberry(message, web);
        }
      }
    });

    rtm.start();
  }
};
tryToConnect();

exports.sendMessage = function (message) {
  if (connected) {
    rtm.sendMessage(message, channel);
  } else {
    winston.info('Unable to send message, not connected, trying to connect now.');
    tryToConnect();
  }
};
