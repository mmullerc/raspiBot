var mongoose = require('mongoose');
var moment = require('moment');

var banSchema = mongoose.Schema({
  userID: Number,
  until: {type: Date, default: moment.utc().add(1, 'hour').toDate()},
});

var Ban = mongoose.model('Ban', banSchema);

module.exports = Ban;
