/**
 * 小程序配置文件
 */

// 域名
var host = 'http://35.180.103.230:8080/v1/';

var config = {

    // 下面的地址配合云端 Demo 工作
    service: {
        host,

        gameUrl: `${host}game/`,

        // 登录地址，用于建立会话
        loginUrl: `${host}/weapp/login`,

        // 测试的请求地址，用于测试会话
        requestUrl: `${host}/weapp/user`,
    }
};

module.exports = config;
