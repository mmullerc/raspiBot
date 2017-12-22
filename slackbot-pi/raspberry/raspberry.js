var request = require('request');
var Components = require('../models/components');

let opts = {
 'as_user': true,
};
let component = {
  name: "",
  state: "",
  direction: "",
  speed: "",
};

module.exports = {
   raspberry: (message, web) => {
     let command = message.text.replace('pi ', '');
     let params = command.split(' ');

    //Validations
    if(params.length < 2) {
      web.chat.postMessage(message.channel, 'Not enough parameters', opts);
      return;
    }

    if(['motor', 'lights'].indexOf(params[0]) < 0) {
      web.chat.postMessage(message.channel, 'Sorry, ' + params[0] + ' is not yet active.', opts);
      return;
    }

    if(['on', 'off'].indexOf(params[1]) < 0) {
      web.chat.postMessage(message.channel, 'Sorry, I dont know this action for ' + params[0], opts);
      return;
    }

    if(params[2] && ['forward', 'backward', 'left', 'right'].indexOf(params[2]) < 0) {
      web.chat.postMessage(message.channel, 'Sorry, I dont know this action for ' + params[0], opts);
      return;
    }

    if(params[3] && ['slow', 'fast'].indexOf(params[3]) < 0) {
      web.chat.postMessage(message.channel, 'Sorry, I cant get that speed.', opts);
      return;
    }

    component = {
      name: params[0],
      state: params[1],
      direction: params[2],
      speed: params[3],
    };

    updateComponent(component);

    let response;
    if (component.name === 'motor') {
      if(component.state === 'on') {
        request(`http://10.28.6.68:8080/setUpMotors`);
        response = 'Car is moving '+component.direction;
      }

      else if(component.state === 'off') {
        request(`http://10.28.6.68:8080/stopMotor`);
        response = 'Car has stopped';
      }
    }

    else if(component.name === 'lights') {
      if(component.state === 'blink') {
        request(`http://10.28.6.68:8080/startLed`);
        response = 'LED is blinking';
      }

      else if(component.state === 'off') {
        response = 'LED is turned off';
      }

      else if(component.state === 'on') {
        response = 'LED is turned on';
      }
    }
    web.chat.postMessage(message.channel, response, opts);
   },
 };

function updateComponent(component) {
  var query = {"name": component.name};
  Components.findOneAndUpdate(query, component, {upsert:true}, function(err, doc) {
    let result = doc.name + " succesfully saved";
    if(err) result="Error: " + err;
    console.log(result);
  });
}
