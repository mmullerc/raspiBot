var mongoose = require('mongoose');

var componentsSchema = mongoose.Schema({
  name: String,
  state: String,
  direction: String,
  speed: String,
});

var Components = mongoose.model('components', componentsSchema);

module.exports = Components;
