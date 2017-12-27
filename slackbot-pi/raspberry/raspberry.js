var request = require('request');
var Components = require('../models/components');
var Users = require('../models/users');

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

     console.log(params)

     

      if (params[0] == "delete") {
        if (params.length > 1) {

          let userId = params[1].replace(/[^a-zA-Z0-9 ]/g, "")

          console.log(userId);

          web.chat.postMessage(message.channel, 'hello', opts);
          deleteUser({
                      botId : userId,
                  });

          return
        } else {
          web.chat.postMessage(message.channel, 'Not enough parameters', opts);
          return
        }
      }

      if (params[0] == 'add') {

        if (params.length > 2) {

          let userId = params[1].replace(/[^a-zA-Z0-9 ]/g, "");
          let direction = params[2];

          console.log(userId);

          web.chat.postMessage(message.channel, 'hello', opts);
          addUser({
                      botId : userId,
                      direction : direction,
                  });
          return
        } else {
          web.chat.postMessage(message.channel, 'Not enough parameters', opts);
          return
        }
      }
     

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

async function addUser(user) {
  const query = {"botId": user.botId};
  let userExist = await findUser(query);

  if (!userExist) {
    console.log("new user");
    const newUser = new Users(user);
    newUser.save()
  } else {
    console.log("user exist");
  }
  
}

async function findUser(query) {
  var x;
  await Users.findOne(query,function(err, doc) {
    x = doc;
  });
  return  x;
}

async function deleteUser(user) {
  const query = {"botId": user.botId};
  let userExist = await findUser(query);
    if(userExist) {
      userExist.remove()
      console.log("removed");
    } else {
      console.log("user dosn't exist");
    }
}