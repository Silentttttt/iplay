"use strict";

var Wallet = require("nebulas");
var HttpRequest = require("./node-request.js");
var TestNetConfig = require("./test_config.js");
var Neb = Wallet.Neb;
var Transaction = Wallet.Transaction;
var FS = require("fs");
var expect = require('chai').expect;
var Unit = Wallet.Unit;

// mocha cases/contract/xxx testneb2 -t 2000000
var args = process.argv.splice(2);
var env = args[1];
if (env == null){
    env = "local";
}
var testNetConfig = new TestNetConfig(env);

var neb = new Neb();
var ChainID = testNetConfig.ChainId;
var sourceAccount = testNetConfig.sourceAccount;
var coinbase = testNetConfig.coinbase;
var apiEndPoint = testNetConfig.apiEndPoint;
neb.setRequest(new HttpRequest(apiEndPoint));

var contractAddress;
var toAddress = Wallet.Account.NewAccount();
var nonce;
var contractNonce = 0;

/*
 * set this value according to the status of your testnet.
 * the smaller the value, the faster the test, with the risk of causing error
 */

var maxCheckTime = 30;
var checkTimes = 0;
var beginCheckTime;

console.log("env:", env);

function checkTransaction(hash, callback) {
    if (checkTimes === 0) {
        beginCheckTime = new Date().getTime();
    }
    checkTimes += 1;
    if (checkTimes > maxCheckTime) {
        console.log("check tx receipt timeout:" + hash);
        checkTimes = 0;
        callback();
        return;
    }

    neb.api.getTransactionReceipt(hash).then(function (resp) {

        console.log("tx receipt status:" + resp.status);
        if (resp.status === 2) {
            setTimeout(function () {
                checkTransaction(hash, callback);
            }, 2000);
        } else {
            checkTimes = 0;
            var endCheckTime = new Date().getTime();
            console.log("check tx time: : " + (endCheckTime - beginCheckTime) / 1000);
            callback(resp);
        }
    }).catch(function (err) {
        console.log("fail to get tx receipt hash: " + hash);
        console.log("it may becuase the tx is being packing, we are going on to check it!");
       // console.log(err);
        setTimeout(function () {
            checkTransaction(hash, callback);
        }, 2000);
    });
}

function checkNonce(accout, nonce, callback) {
    var interval = setInterval(function () {
        neb.api.getAccountState(accout.getAddressString()).then(function (resp) {
            console.log("check nonce")
            console.log(JSON.stringify(resp));
            if (resp.nonce >= nonce) {
                clearInterval(interval);
                callback();
            }
        });

    }, 2000);
}

var accounts = new Array();

describe('claim token', function () {

    before("0. claim Token", function (done) {
        neb.api.getAccountState(sourceAccount.getAddressString()).then(function(resp){
            console.log(JSON.stringify(resp));
            var nonce = parseInt(resp.nonce);

            for(var i = 0; i < 10; i++) {
                var account = Wallet.Account.NewAccount();
                accounts.push(account);
                var tx = new Transaction(ChainID, sourceAccount, account, Unit.nasToBasic(100), nonce + 1 +i, 1000000, 20000000);
                tx.signTransaction();
                neb.api.sendRawTransaction(tx.toProtoString());
            }
            checkNonce(sourceAccount, nonce + 10, function(){
                done();
            });
        });   
    });

    before('1. deploy contract', function (done) {
        try {
            var banker = accounts[0];
            neb.api.getAccountState(banker.getAddressString()).then(function(resp) {
                console.log("----step0. get source account state: " + JSON.stringify(resp));
                var contractSource = FS.readFileSync("iplay.js", "UTF-8");
                var contract = {
                    'source': contractSource,
                    "sourceType": "js",
                    "args": "[\"StandardToken\", \"ATP\", 18, \"1000000000\"]"
                };
                nonce = parseInt(resp.nonce);
                nonce = nonce + 1;
                var tx = new Transaction(ChainID, banker, banker, 0, nonce, 1000000, 20000000, contract);
                tx.signTransaction();
                return neb.api.sendRawTransaction(tx.toProtoString());
            }).then(function(resp) {
                console.log("----step1. deploy contract: " + JSON.stringify(resp));
                contractAddress = resp.contract_address;
                console.log(contractAddress);
                checkTransaction(resp.txhash, function(resp) {
                    try {
                        expect(resp).to.not.be.a('undefined');
                        expect(resp.status).equal(1);
                        console.log("----step2. have been on chain");
                        done()
                     } catch (err) {
                        console.log(err);
                        done(err);
                    }
                });
            }).catch(function(err) {
                console.log("unexpected err: " + err);
                done(err);
            });
        } catch (err) {
            console.log("unexpected err: " + err);
            done(err);
        }
    });


    var createGame = function(banker, payType, type, afterTime, theme, options, amount, nasValue, callback) {
        var args = new Array();
        var now = Date.now();
        var deadLine = now + afterTime;
        console.log("create game, deadLine: " + deadLine);
        var args = new Array();
        args.push(payType);
        args.push(type);
        args.push(deadLine);
        args.push(theme);
        args.push(options);
        args.push(amount);
        neb.api.getAccountState(banker.getAddressString()).then(function(resp) {
            console.log("---step2. get banker account state: " + JSON.stringify(resp));
            nonce = parseInt(resp.nonce);

            var contract = {
                'function': "createGame",
                "args": JSON.stringify(args),
            }
            console.log("===========", contractAddress);
            var tx = new Transaction(ChainID, banker, contractAddress, Wallet.Unit.nasToBasic(nasValue), ++nonce, 1000000, 20000000, contract)
            tx.signTransaction();
            return neb.api.sendRawTransaction(tx.toProtoString());
        }).then(callback);
    };

    it('1# create game', function (done) {
        try {
            var banker = accounts[0];
            var options = new Array();
            options.push({description:"胜", odd: 1.5});
            options.push({description:"平", odd: 2});
            options.push({description:"负", odd: 2.5});


            createGame(banker, 1, 1, 30 * 1000, "测试局 德国vs巴西", options, 1000, 10, function(resp) {
                checkTransaction(resp.txhash, function(resp) {
                    try{
                        console.log(JSON.stringify(resp));
                        expect(resp).to.not.be.a('undefined');
                        expect(resp.status).equal(1);
                        var gameId = resp.result.gameId;
                        
                        var contract = {
                            "function": "getGame",
                            "args": "[" + gameId + "]",
                        }
                        var tx = new Transaction(ChainID, banker, contractAddress, 0, ++nonce, 1000000, 20000000, contract);
                        tx.signTransaction();
                        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp) {
                            console.log(JSON.stringify(resp));
                            done();
                        });
                    } catch(err) {
                        console.log(err);
                        done(err);
                    }
                });
            });
        } catch (err) {
            console.log("unexpected err: " + err);
            done(err);
        }
    });
});