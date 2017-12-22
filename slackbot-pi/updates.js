var env = require('./.env.json');
var shell = require('shelljs');
var winston = require('winston');
var slack = require('./slack');
var RESTClient = require('node-rest-client').Client;

var optionsAuth = {user: env.username, password: env.password};
var client = new RESTClient(optionsAuth);
var moment = require('moment');
var Ban = require('./models/ban');

var mergeIt = (mergeReq) => {
  winston.info('Merging!');
  winston.info(mergeReq);

  var args = {
    headers: {'Content-Type': 'application/json'},
  };
  var prUrl = env.bitbucketAPI + 'projects/' + env.project + '/repos/' + env.repo + '/pull-requests';
  var buildCheck = 'https://bitbucket.kgportal.com/rest/build-status/latest/commits/stats/';

  let shouldMerge = Math.random() > 0.1;

  if (!shouldMerge) {
    client.post(prUrl + '/' + mergeReq.id + '/decline', args, function (data, response) {
      slack.sendMessage('Fuck you ' + mergeReq.author.user.displayName);
    });

    new Ban({userID: mergeReq.author.user.id}).save();

    return;
  }


  Ban.find({until: {$gt: moment.utc().toDate()}}).exec((err, users) => {
    var ban = users.find(x => x.userID === mergeReq.author.user.id);

    if (ban) {
      client.post(prUrl + '/' + mergeReq.id + '/decline', args, function (data, response) {
        slack.sendMessage('Fuck you ' + mergeReq.author.user.displayName + '. You will be unbanned ' + moment(ban.until).fromNow());
      });

      return;
    }

    client.post(prUrl + '/' + mergeReq.id + '/merge?version=' + mergeReq.version, args, function (data, response) {

      winston.info(data);
      if(data.errors){

        slack.sendMessage('Unable to merge: ' + mergeReq.title + ' Error: ' + data.errors[0].message);
        client.post(prUrl + '/' + mergeReq.id + '/decline', args, function (data, response) {
          slack.sendMessage('Fuck you ' + mergeReq.author.user.displayName);
        });

      }else{

        slack.sendMessage('Merged: ' + mergeReq.title + ' from ' + mergeReq.author.user.displayName);

        shell.exec('git pull');
        shell.exec('yarn install --no-lockfile');
        process.exit(0); // pm2 will restart the bot

      }
    });
  });

};

var checkUpdates = () => {
  winston.info('Checking for updates');

  // https://developer.atlassian.com/bitbucket/server/docs/latest/how-tos/command-line-rest.html
  // https://developer.atlassian.com/static/rest/bitbucket-server/latest/bitbucket-rest.html
  var args = {
    headers: {'Content-Type': 'application/json'},
  };
  var prUrl = env.bitbucketAPI + 'projects/' + env.project + '/repos/' + env.repo + '/pull-requests';
  var buildCheck = 'https://bitbucket.kgportal.com/rest/build-status/latest/commits/stats/';
  client.get(prUrl, args, function (data, response) {
    data.values.forEach((a) => {
      if(a.reviewers.find(r => r.user.name === env.username)){
        client.get(buildCheck + a.fromRef.latestCommit + '?includeUnique=true', args, function (data, response) {
          if(data.successful > 0){
            mergeIt(a);
          }else{
            slack.sendMessage('Cant merge: ' + a.title + ' from ' + a.author.user.displayName + ' it does not pass the bamboo test');
          }
        });
      }
    });
  });
};

exports.checkUpdates = checkUpdates;
