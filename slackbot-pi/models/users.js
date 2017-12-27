var mongoose = require('mongoose');

var usersSchema = mongoose.Schema({
  name: String,
  botId: String,
  direction: String,
});

var Users = mongoose.model('users', usersSchema);

module.exports = Users;
