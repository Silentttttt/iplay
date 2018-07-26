var weconstants = require('./weconstants');
var qcloud = require('../vendor/wafer2-client-sdk/index');
var config = require('../config');


var _request = function (url, needLogin, data, cb) {
  var method = url.indexOf("list")>0?'GET':'POST'
  method = 'POST'
  qcloud.request({
    url: url,
    login: needLogin,
    method: method,
    data: data,
    header: { 'content-type': 'application/json' },
    success: (response) => {
      var res = response.data;
      console.log(res);
      if ((typeof res.code) != 'undefined') {
        typeof cb == "function" && cb(res.code, res.msg, res.data);
      } else {
        typeof cb == "function" && cb(-1, weconstants.WE_DAKA_ERROR_SERVICE, null);
      }
    },
    fail: (error) => {
      typeof cb == "function" && cb(-2, weconstants.WE_DAKA_ERROR_NETWORK, error);
    }
  });
}


// 获取比赛列表 
var getGameList = function (params, cb) {
  _request(config.service.gameUrl+'list', false, params, cb)
}

// 获取比赛列表 
var getGameDetail = function (params, cb) {
  _request(config.service.gameUrl +'quizzes', false, params, cb)
}

module.exports = {
  getGameList: getGameList,
  getGameDetail: getGameDetail,
};