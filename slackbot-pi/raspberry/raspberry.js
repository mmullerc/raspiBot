const request = require('request');
const Components = require('../models/components');
const Users = require('../models/users');
const navigateUrl = 'http://10.28.6.22:8080/navigate';

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
  raspberry: async (message, web) => {
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

    if(params[0] == 'come' && params[1] == 'here'){
      //Car should go to user's location
      let user = await findUser({"botId": message.user});

      // Configure the request
      var options = {
        uri:     navigateUrl,
        method:  'POST',
        json: {'color': user.direction}
      }

      // Start the request
      request(options, function (error, response, body) {
        if (!error && response.statusCode == 200) {
          web.chat.postMessage(message.channel, 'Going to your desk', opts);
        }
      })
      
    }
     
    //Validations
    if(params.length < 2) {
      web.chat.postMessage(message.channel, 'Not enough parameters', opts);
      return;
    }

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