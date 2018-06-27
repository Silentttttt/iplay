// "use strict";

var Wallet = require("nebulas");
var HttpRequest = require("./node-request.js");
var TestNetConfig = require("./test_config.js");
var Neb = Wallet.Neb;
var Transaction = Wallet.Transaction;
var FS = require("fs");
var expect = require('chai').expect;
var Unit = Wallet.Unit;
var BigNumber = require('bignumber.js');

// mocha cases/contract/xxx testneb2 -t 2000000
var args = process.argv.splice(2);
var env = args[1];
if (env == null) {
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
var nonces = new Array();

describe('pay type = 1 and game type = 1 ', function () {

    before("0. claim Token", function (done) {
        neb.api.getAccountState(sourceAccount.getAddressString()).then(function(resp){
            console.log(JSON.stringify(resp));
            var nonce = parseInt(resp.nonce);

            for(var i = 0; i < 10; i++) {
                var account = Wallet.Account.NewAccount();
                accounts.push(account);
                nonces.push(0);
                var tx = new Transaction(ChainID, sourceAccount, account, Unit.nasToBasic(100), nonce + 1 +i, 1000000, 20000000);
                tx.signTransaction();
                neb.api.sendRawTransaction(tx.toProtoString());
            }
            checkNonce(sourceAccount, nonce + 10, function() {
                done();
            });
        });   
    });

    before('1. deploy contract', function (done) {
        try {
            var banker = accounts[0];
                var contractSource = FS.readFileSync("iplay.js", "UTF-8");
                var contract = {
                    'source': contractSource,
                    "sourceType": "js",
                    "args": "[\"StandardToken\", \"ATP\", 18, \"1000000000\"]"
                };
                
                var tx = new Transaction(ChainID, banker, banker, 0, ++nonces[0], 1000000, 20000000, contract);
                tx.signTransaction();
                neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp) {
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
            });
        } catch (err) {
            console.log("unexpected err: " + err);
            done(err);
        }
    });

    it('2. claim nrc20 token', function (done) {
        try {
            for(var i = 1; i < 10 ; i++) {
                var args = new Array();
                args.push(accounts[i].getAddressString());
                args.push("1000000000000000000000");
                var contract = {
                    "function": "transfer",
                    "args": JSON.stringify(args),
                };
            
                var tx = new Transaction(ChainID, accounts[0], contractAddress, 
                    Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
                tx.signTransaction();
                neb.api.sendRawTransaction(tx.toProtoString());
            }
            checkNonce(accounts[0], nonces[0], function() {
                try {
                    var args = new Array();
                    args.push(accounts[1].getAddressString());
                    var contract = {
                    "function": "balanceOf",
                    "args": JSON.stringify(args),  
                    };

                    neb.api.call(accounts[1].getAddressString(),contractAddress, 
                        Unit.nasToBasic(0), 0, 1000000, 20000000, contract).then(function(resp){
                            console.log("========", resp);
                            expect(JSON.parse(resp.result)).equal("1000000000000000000000");
                            done();
                        }).catch(function(err) {
                            console.log(err);
                            done(err);
                        });
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            });  
        } catch (err) {
            console.log("unexpected err: " + err);
            done(err);
        }
    });

    var gameId;
    var deadLine;
    var createGame = function(banker, payType, type, afterTime, theme, options, amount, nasValue, callback) {
        var args = new Array();
        var now = Date.now();
        deadLine = now + afterTime;
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
            var tx = new Transaction(ChainID, banker, contractAddress, Wallet.Unit.nasToBasic(nasValue), ++nonces[0], 1000000, 20000000, contract)
            tx.signTransaction();
            return neb.api.sendRawTransaction(tx.toProtoString());
        }).then(callback);
    };

    it('create game', function (done) {
        try {
            var banker = accounts[0];
            var options = new Array();
            options.push({description:"胜", odd: 2});
            options.push({description:"平", odd: 3});
            options.push({description:"负", odd: 4});


            createGame(banker, 1, 1, 200 * 1000, "测试局 德国vs巴西", options, 100, 0, function(resp) {
                checkTransaction(resp.txhash, function(resp) {
                    try{
                        console.log(JSON.stringify(resp));
                        expect(resp).to.not.be.a('undefined');
                        expect(resp.status).equal(1);
                        gameId = JSON.parse(resp.execute_result).gameId;
                        console.log(typeof(gameId));
                        var contract = {
                            "function": "getGame",
                            "args": "[" + gameId.toString() + "]",
                        }
                        
                        neb.api.call(banker.getAddressString(), contractAddress, Wallet.Unit.nasToBasic(0),
                             nonce, 100000, 2000000, contract).then(function(resp) {
                            console.log("============", JSON.stringify(resp));
                            done();
                        }).catch(function(err){
                            console.log(err);
                            done(err);
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

    it("2# not owner startGame", function(done) {
        var contract = {
            "function": "startGame",
            "args": "[" + gameId.toString() + "]",
        }
        var tx = new Transaction(ChainID, accounts[1], contractAddress, 0, ++nonces[1], 1000000, 20000000, contract)
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try {
                    console.log(resp);
                    expect(resp.status).equal(0);
                    expect(resp.execute_result).equal("only owner of the game can start the game");
                    done();
                } catch (err) {
                    console.log(err);
                    done(err);
                }
            });
        })
    });

    it("buy token before game start[buyToken]", function(done){
        var args = new Array();
        args.push(gameId);
        args.push(1);
        args.push(1);
        args.push(10);
        var contract = {
            "function": "buyTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[0], contractAddress, 
                Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(0);
                    expect(resp.execute_result).equals("No bet in this status of game");
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("owner start game[startGame]", function(done) {
        var contract = {
            "function": "startGame",
            "args": "[" + gameId.toString() + "]",
        }
        var tx = new Transaction(ChainID, accounts[0], contractAddress, 0, ++nonces[0], 1000000, 20000000, contract)
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try {
                    console.log(resp);
                    expect(resp.status).equal(1);
                    done();
                } catch (err) {
                    console.log(err);
                    done(err);
                }
            });
        })
    });

    it("account 1 buy op[1] for 10 [buyToken]", function(done){
        var args = new Array();
        args.push(gameId);
        args.push(1);
        args.push(1);
        args.push(10);
        var contract = {
            "function": "buyTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[1], contractAddress, 
                Unit.nasToBasic(0), ++nonces[1], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    expect(JSON.parse(resp.execute_result).ticketId).equal(1);
                    var contract = {
                        "function": "getTicket",
                        "args": "[1]",
                    }
                    neb.api.call(accounts[1].getAddressString(), contractAddress, Wallet.Unit.nasToBasic(0),
                         nonce, 100000, 2000000, contract).then(function(resp) {
                        console.log(resp);
                        done();
                    });

                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("account 2 buy op[2] for 60[buyToken]", function(done){
        var args = new Array();
        args.push(gameId);
        args.push(2);
        args.push(1);
        args.push(60);
        var contract = {
            "function": "buyTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[2], contractAddress, 
                Unit.nasToBasic(0), ++nonces[2], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(0);
                    expect(resp.execute_result).equal("remaining deposit is not enough");
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("sendDeposit", function(done) {
        var args = new Array();
        args.push(gameId);
        args.push(20);
        var contract = {
            "function": "sendDeposit",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[0], contractAddress, 
            Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("account 2 buy op[2] for 60[buyToken]", function(done){
        var args = new Array();
        args.push(gameId);
        args.push(2);
        args.push(1);
        args.push(60);
        var contract = {
            "function": "buyTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[2], contractAddress, 
                Unit.nasToBasic(0), ++nonces[2], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    expect(JSON.parse(resp.execute_result).ticketId).equal(2);
                    var contract = {
                        "function": "getTicket",
                        "args": "[2]",
                    }
                    neb.api.call(accounts[2].getAddressString(), contractAddress, Wallet.Unit.nasToBasic(0),
                         nonce, 100000, 2000000, contract).then(function(resp) {
                        console.log(resp);
                        done();
                    });

                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("change Odds", function(done) {
        var odds = new Array();
        odds.push(3);
        odds.push(2);
        odds.push(4);
        var args = new Array();
        args.push(gameId);
        args.push(odds);
        args.push(1);
        var contract = {
            "function": "changeOdds",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[0], contractAddress, 
            Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("account 3 buy op[2] for 20[buyTicket]", function(done){
        var args = new Array();
        args.push(gameId);
        args.push(2);
        args.push(2);
        args.push(20);
        var contract = {
            "function": "buyTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[3], contractAddress, 
                Unit.nasToBasic(0), ++nonces[3], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(0);
                    expect(resp.execute_result).equal("remaining deposit is not enough");
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("account 3 buy op[2] for 10[buyToken]", function(done){
        var args = new Array();
        args.push(gameId);
        args.push(2);
        args.push(2);
        args.push(10);
        var contract = {
            "function": "buyTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[3], contractAddress, 
                Unit.nasToBasic(0), ++nonces[3], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    expect(JSON.parse(resp.execute_result).ticketId).equal(3);
                    var contract = {
                        "function": "getTicket",
                        "args": "[3]",
                    }
                    neb.api.call(accounts[1].getAddressString(), contractAddress, Wallet.Unit.nasToBasic(0),
                         nonce, 100000, 2000000, contract).then(function(resp) {
                        console.log(resp);
                        done();
                    });

                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("preview result[previewResult]", function(done) {
        var args = new Array();
        args.push(gameId);
        args.push(2);
        var contract = {
            "function": "previewResult",
            "args": JSON.stringify(args),
        };
        console.log(deadLine - Date.now());
        // while(Date.now() < deadLine) {
        //     sleep(1000);
        // }
        var tx = new Transaction(ChainID, accounts[0], contractAddress, 
            Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("get reward[getReward]", function(done) {
        var args = new Array();
        args.push(1);
        var contract = {
            "function": "getReward",
            "args": JSON.stringify(args),
        };
     
        var tx = new Transaction(ChainID, accounts[1], contractAddress, 
            Unit.nasToBasic(0), ++nonces[1], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(0);
                    done();
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("confirm result[confirmResult]", function(done) {
        var args = new Array();
        args.push(gameId);
        var contract = {
            "function": "confirmResult",
            "args": JSON.stringify(args),
        };
        console.log(accounts[0].getAddressString());
        console.log(contractAddress);

        var tx = new Transaction(ChainID, accounts[0], contractAddress, 
            Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            var txhash = resp.txhash;
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    neb.api.getEventsByHash(txhash).then(function(resp) {
                        console.log(JSON.stringify(resp));
                        done();
                    })
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("get reward[getReward]", function(done) {
        for(var i = 1; i<=3; i++) {
            var args = new Array();
            args.push(i);
            var contract = {
                "function": "getReward",
                "args": JSON.stringify(args),
            };
        
            var tx = new Transaction(ChainID, accounts[0], contractAddress, 
                Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
            tx.signTransaction();
            neb.api.sendRawTransaction(tx.toProtoString());
        }
        checkNonce(accounts[0], nonces[0], function(){
            done();
        });
    }); 

    it("get remain deposit result[getRemainDeposit]", function(done) {
        var args = new Array();
        args.push(gameId);
        var contract = {
            "function": "getRemainDeposit",
            "args": JSON.stringify(args),
        };

        var tx = new Transaction(ChainID, accounts[0], contractAddress, 
            Unit.nasToBasic(0), ++nonces[0], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            var txhash = resp.txhash;
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    neb.api.getEventsByHash(txhash).then(function(resp) {
                        console.log(JSON.stringify(resp));
                        done();
                    })
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            })
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("[transferTicket]", function(done) {
        var args = new Array();
        args.push(1);
        args.push(accounts[2].getAddressString());
        var contract = {
            "function": "transferTicket",
            "args": JSON.stringify(args),
        };
        var tx = new Transaction(ChainID, accounts[1], contractAddress, 
            Unit.nasToBasic(0), ++nonces[1], 1000000, 20000000, contract);
        tx.signTransaction();
        neb.api.sendRawTransaction(tx.toProtoString()).then(function(resp){
            var txhash = resp.txhash;
            checkTransaction(resp.txhash, function(resp) {
                try{
                    console.log(resp);
                    expect(resp.status).equal(1);
                    var contract = {
                        "function": "getTicket",
                        "args": "[1]",
                    }
                    
                    neb.api.call(accounts[0].getAddressString(), contractAddress, Wallet.Unit.nasToBasic(0),
                         nonce, 100000, 2000000, contract).then(function(resp) {
                        console.log("============", JSON.stringify(resp));
                        done();
                    }).catch(function(err){
                        console.log(err);
                        done(err);
                    });
                } catch(err) {
                    console.log(err);
                    done(err);
                }
            });
        }).catch(function(err) {
            console.log(err);
            done(err);
        });
    });

    it("check balance account 0", function(done) {
        var args = new Array();
        args.push(accounts[0].getAddressString());
        var contract = {
            "function": "balanceOf",
            "args": JSON.stringify(args),  
        };

        neb.api.call(accounts[0].getAddressString(),contractAddress, 
            Unit.nasToBasic(0), 0, 1000000, 20000000, contract).then(function(resp){
                console.log("========", resp);
                expect(JSON.parse(resp.result)).equal("999990880000000000000000000");
                done();
            }).catch(function(err) {
                console.log(err);
                done(err);
            });
    });

    it("check balance account 1", function(done) {
        var args = new Array();
        args.push(accounts[1].getAddressString());
        var contract = {
            "function": "balanceOf",
            "args": JSON.stringify(args),  
        };

        neb.api.call(accounts[0].getAddressString(),contractAddress, 
            Unit.nasToBasic(0), 0, 1000000, 20000000, contract).then(function(resp){
                console.log("========", resp);
                expect(JSON.parse(resp.result)).equal("990000000000000000000");
                done();
            }).catch(function(err) {
                console.log(err);
                done(err);
            });
    });

    it("check balance account 2", function(done) {
        var args = new Array();
        args.push(accounts[2].getAddressString());
        var contract = {
            "function": "balanceOf",
            "args": JSON.stringify(args),  
        };

        neb.api.call(accounts[0].getAddressString(),contractAddress, 
            Unit.nasToBasic(0), 0, 1000000, 20000000, contract).then(function(resp){
                console.log("========", resp);
                expect(JSON.parse(resp.result)).equal("1120000000000000000000");
                done();
            }).catch(function(err) {
                console.log(err);
                done(err);
            });
    });

    it("check balance account 3", function(done) {
        var args = new Array();
        args.push(accounts[3].getAddressString());
        var contract = {
            "function": "balanceOf",
            "args": JSON.stringify(args),  
        };

        neb.api.call(accounts[0].getAddressString(),contractAddress, 
            Unit.nasToBasic(0), 0, 1000000, 20000000, contract).then(function(resp){
                console.log("========", resp);
                expect(JSON.parse(resp.result)).equal("1010000000000000000000");
                done();
            }).catch(function(err) {
                console.log(err);
                done(err);
            });
    });

});