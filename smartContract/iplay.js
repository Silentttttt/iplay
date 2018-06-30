'use strict';
//NRC20
'use strict';

var Allowed = function (obj) {
    this.allowed = {};
    this.parse(obj);
}

Allowed.prototype = {
    toString: function () {
        return JSON.stringify(this.allowed);
    },

    parse: function (obj) {
        if (typeof obj != "undefined") {
            var data = JSON.parse(obj);
            for (var key in data) {
                this.allowed[key] = new BigNumber(data[key]);
            }
        }
    },

    get: function (key) {
        return this.allowed[key];
    },

    set: function (key, value) {
        this.allowed[key] = new BigNumber(value);
    }
}

var StandardToken = function () {
    LocalContractStorage.defineProperties(this, {
        _name: null,
        _symbol: null,
        _decimals: null,
        _totalSupply: {
            parse: function (value) {
                return new BigNumber(value);
            },
            stringify: function (o) {
                return o.toString(10);
            }
        }
    });

    LocalContractStorage.defineMapProperties(this, {
        "balances": {
            parse: function (value) {
                return new BigNumber(value);
            },
            stringify: function (o) {
                return o.toString(10);
            }
        },
        "allowed": {
            parse: function (value) {
                return new Allowed(value);
            },
            stringify: function (o) {
                return o.toString();
            }
        }
    });
};

StandardToken.prototype = {
    init: function (name, symbol, decimals, totalSupply) {
        this._name = name;
        this._symbol = symbol;
        this._decimals = decimals | 0;
        this._totalSupply = new BigNumber(totalSupply).mul(new BigNumber(10).pow(decimals));

        var from = Blockchain.transaction.from;
        this.balances.set(from, this._totalSupply);
        this._transferEvent(true, from, from, this._totalSupply);
    },

    // Returns the name of the token
    name: function () {
        return this._name;
    },

    // Returns the symbol of the token
    symbol: function () {
        return this._symbol;
    },

    // Returns the number of decimals the token uses
    decimals: function () {
        return this._decimals;
    },

    totalSupply: function () {
        return this._totalSupply.toString(10);
    },

    balanceOf: function (owner) {
        var balance = this.balances.get(owner);

        if (balance instanceof BigNumber) {
            return balance.toString(10);
        } else {
            return "0";
        }
    },

    transfer: function (to, value) {
        value = new BigNumber(value);
        if (value.lt(0)) {
            throw new Error("invalid value.");
        }

        var from = Blockchain.transaction.from;
        var balance = this.balances.get(from) || new BigNumber(0);

        if (balance.lt(value)) {
            throw new Error("transfer failed.");
        }

        this.balances.set(from, balance.sub(value));
        var toBalance = this.balances.get(to) || new BigNumber(0);
        this.balances.set(to, toBalance.add(value));

        this._transferEvent(true, from, to, value);
    },

    transferFrom: function (from, to, value) {
        var spender = Blockchain.transaction.from;
        var balance = this.balances.get(from) || new BigNumber(0);

        var allowed = this.allowed.get(from) || new Allowed();
        var allowedValue = allowed.get(spender) || new BigNumber(0);
        value = new BigNumber(value);

        if (value.gte(0) && balance.gte(value) && allowedValue.gte(value)) {

            this.balances.set(from, balance.sub(value));

            // update allowed value
            allowed.set(spender, allowedValue.sub(value));
            this.allowed.set(from, allowed);

            var toBalance = this.balances.get(to) || new BigNumber(0);
            this.balances.set(to, toBalance.add(value));

            this._transferEvent(true, from, to, value);
        } else {
            throw new Error("transfer failed.");
        }
    },

    _transferEvent: function (status, from, to, value) {
        Event.Trigger(this.name(), {
            Status: status,
            Transfer: {
                from: from,
                to: to,
                value: value
            }
        });
    },

    approve: function (spender, currentValue, value) {
        var from = Blockchain.transaction.from;

        var oldValue = this.allowance(from, spender);
        if (oldValue != currentValue.toString()) {
            throw new Error("current approve value mistake.");
        }

        var balance = new BigNumber(this.balanceOf(from));
        var value = new BigNumber(value);

        if (value.lt(0) || balance.lt(value)) {
            throw new Error("invalid value.");
        }

        var owned = this.allowed.get(from) || new Allowed();
        owned.set(spender, value);

        this.allowed.set(from, owned);

        this._approveEvent(true, from, spender, value);
    },

    _approveEvent: function (status, from, spender, value) {
        Event.Trigger(this.name(), {
            Status: status,
            Approve: {
                owner: from,
                spender: spender,
                value: value
            }
        });
    },

    allowance: function (owner, spender) {
        var owned = this.allowed.get(owner);

        if (owned instanceof Allowed) {
            var spender = owned.get(spender);
            if (typeof spender != "undefined") {
                return spender.toString(10);
            }
        }
        return "0";
    }
};
//---------------------------------------END NRC20-----------------------------------------
/*
=======================================Game & Market=======================================
*/
var assert = function(expression, info) {
    if (!expression) {
        throw info;
    }
};

var assertNumber = function(n) {
    assert(typeof(n) == "number", "param is not a number: " + n);
};

var assertInteger = function(n) {
    assertNumber(n);
    assert(parseInt(n) == n, "param is not a integer: " + n);
};

var assertPosInteger = function(n) {
    assertInteger(n);
    assert(n > 0, "param is not a positive integer: " + n); 
};

var Ticket = function(ticketId, amount, gameId, optionNo, odd) { // 票据
    this.ticketId = ticketId; 
    this.amount = amount;   //ATP
    this.gameId = gameId;   
    this.optionNo = optionNo;  
    this.odd = odd;
    this.status = 0; // 1 -> 买入完成 2 -> 完成兑奖
};

var Option = function(description, odd) { // 每个game有多个option（投注项）
    this.description = description; //string
    this.odd = odd  //number = 真实赔率＊100
    this.bets = 0; //压这个选项的赌注 
    this.expectReward = 0; //如果这个选项正确，需要返回的奖金
};

var Game = function(id, owner, deadLine, type, theme, payType) {
    assertPosInteger(id);
    assertPosInteger(deadLine);
    //0->1->2->3->4
    this.id = id;
    this.deadLine = deadLine;
    this.owner = owner; 
    this.type = type;
    this.theme = theme;
    this.payType = payType;

    this.status = 0;
    this.nextOptionCount = 1; //number
    this.optionVersion = 1; //number
    this.result = -1;

    this.ownerDeposit = 0; // 庄家的押金
    this.bets = 0; // 总下注资金
    this.rewardBets = 0;
    this.deposit = 0;// 总共的押金 ＝ 庄家的押金 ＋ 总下注资金 
    this.depositAtLast = 0; 

    LocalContractStorage.defineMapProperty(this, "Options");
};

Game.prototype = {
    assertOwner: function(address) {
        assert(address == this.owner, "not the onwer of this game");
    },

    _getOptionKey: function(index) {
        return this.id.toString(10) + "#" + index.toString(10);
    },
    
    setOption: function(index, option) {
        var optionKey = this._getOptionKey(index);
        this.Options.set(optionKey, option);
    },

    getOption: function(index) {
        var option = this.Options.get(this._getOptionKey(index));
        if (!option) {
            throw "invalid index, failed to get option";
        }
        return option;
    },

    addOption: function(option) {
        var index = this.nextOptionCount;
        this.nextOptionCount += 1;
        this.setOption(index, option);
    },
};

var Market = function() {
    LocalContractStorage.defineProperties(this, {
        nextGameCount: null,
        nextTicketCount: null,
        admin: null,
        isMarketOpen: null, 
    });

    //map GameId => Game
    LocalContractStorage.defineMapProperty(this, "Games", {
        parse: function(value) {
            var params = JSON.parse(value);
            var game =  new Game(params.id, params.owner, params.deadLine, params.type, params.theme, params.payType);
            game.status = params.status;
            game.nextOptionCount = params.nextOptionCount; 
            game.optionVersion = params.optionVersion;
            game.result = params.result;
        
            game.ownerDeposit = params.ownerDeposit; 
            game.bets = params.bets; 
            game.deposit = params.deposit;
            game.depositAtLast = params.depositAtLast;
            game.rewardBets = params.rewardBets;
            return game;
        },
        stringify: function(obj) {
            return JSON.stringify(obj);
        }
    });


    //only admins can open new games
    LocalContractStorage.defineMapProperty(this, "Admins");

    //map ticketId => ticket
    LocalContractStorage.defineMapProperty(this, "Tickets");

    //map ticketId => owner
    LocalContractStorage.defineMapProperty(this, "TicketIdToOwner");
};

Market.prototype = {
    init: function(name, symbol, decimals, totalSupply) {
        this.admin = Blockchain.transaction.from;
        this.Admins.set(this.admin, 1);
        this.isMarketOpen = 0;
        this.nextGameCount = 1;
        this.nextTicketCount = 1;
        this._getTokenMgr().init(name, symbol, decimals, totalSupply);
      	console.log(Blockchain.transaction.value);
    },

    openMarket: function() {
        if (!this.isMarketOpen) {
            this.isMarketOpen = 1;
        }
    },

    closeMarket: function() {
        if (this.isMarketOpen) {
            this.isMarketOpen = 0;
        }
    },

    //for game owmer
    createGame: function(payType, type, deadLine, theme, options, amount) {//args: {[odd1, describe1], [odd2, descrobe2]...}
        //先转账，确保amount没有溢出
        assert(payType === 1 || payType === 2, "invlaid pay type");
        assertPosInteger(deadLine);
        amount = this._transfer(Blockchain.transaction.from, Blockchain.transaction.to, amount, payType);

        //check args
        assert(typeof(theme) == "string", "theme should be string");
        var creator = Blockchain.transaction.from;
        if (!this.isMarketOpen) {
            assert(this._isMarketAdmin(creator), "only market admin can create game");
        }
        assert(type === 1 || type === 2, "invalid type");
        if (deadLine <= Date.now()) {
            throw "deadLineTime shoud after current time";
        }
        
        //create and init Game
        var gameId = this._getGameId();
        var newGame = new Game(gameId, creator, deadLine, type, theme, payType);
        
        var optionCount = options.length;
        for (var i = 0; i < optionCount; i++) {
            //"check args"
            var odd = options[i].odd;
            var description = options[i].description;
            
            assert(typeof(description) == "string", "description should be a string");
            assert(typeof(odd) == "number", "odd should be a number");
            assert(odd > 1, "odd should larger than 1");
            assertPosInteger(odd * 100);//只允许两位小数

            var option = new Option(description, odd);
            newGame.addOption(option);
        }
        newGame.ownerDeposit = amount;
        newGame.deposit = amount;
        this._setGame(gameId, newGame);
        this._createGameEvent(newGame);
        return {gameId: gameId};
    },

    delGameBfStart: function(gameId) {
        var game = this._getGame(gameId);
        game.assertOwner(Blockchain.transaction.from);
        assert(game.status === 0, "game has started");
        this.transfer(Blockchain.transaction.to, Blockchain.transaction.from, game.ownerDeposit);
        game.ownerDeposit = 0;
        game.deposit = 0;
        game.status = 4;
        game._setGame(gameId, game);
        this._delGameEvent(game);
    },

    sendDeposit: function(gameId, amount) {
        var game = this._getGame(gameId);
        game.assertOwner(Blockchain.transaction.from);
        amount = this._transfer(Blockchain.transaction.from, Blockchain.transaction.to, amount, game.payType);

        assert(game.status === 0 || game.status === 1, "game has ended");
        game.ownerDeposit += amount;
        game.deposit += amount;
        console.log(game); //TODO: debug
        this._setGame(gameId, game);
    },

    startGame: function(gameId) {
        assertPosInteger(gameId);
        var game = this._getGame(gameId);

        assert(game.owner === Blockchain.transaction.from, "only owner of the game can start the game");
        assert(game.status === 0, "game has started or ended");
        game.status = 1;
        this._setGame(gameId, game);
    },

    createAndStartGame: function(payType, type, deadLine, theme, options, amount){
        var gameId = this.createGame(payType, type, deadLine, theme, options, amount).gameId;
        this.startGame(gameId);
        return {gameId: gameId};
    },

    changeOdds: function(gameId, odds, optionVersion) {
        //should keep atomic
        var game = this._getGame(gameId);

        assert(game.owner == Blockchain.transaction.from, "only owner of the gane can change odd");
        assert(game.status == 1 || games.status == 0, "game has ended");
        assert(game.optionVersion == optionVersion, "optionVersion has changed");
        assert(typeof(odds) == "object", "Odds should be a array");

        
        for(var i = 0; i < odds.length; i++) {
            var odd = odds[i];
            assert(typeof(odd) == "number", "odd should be a number");
            assert(odd > 1, "odd should larger than 1");
            assertPosInteger(odd * 100);

            var option = game.getOption(i + 1);
            option.odd = odd;
            game.setOption(i + 1, option);
        }
           
        game.optionVersion += 1;
        console.log(game); //TODO: Debug;
        this._setGame(gameId, game);
    },

    previewResult: function(gameId, resultIndex) {
        var game = this._getGame(gameId);
        game.assertOwner(Blockchain.transaction.from);
        assert(game.status == 1 || game.status == 2, "can not preview result in this status");
        assertPosInteger(resultIndex);

        assert(Date.now() > game.deadLine, "too early to preview result");
        assert(resultIndex < game.nextOptionCount, "invalid result index");
    
        var option = game.getOption(resultIndex);

        game.result = resultIndex;
        game.status = 2;
        game.rewardBets = option.bets;
        game.depositAtLast = game.deposit;
        console.log(game); //TODO: debug;
        this._setGame(gameId, game);
    },

    confirmResult: function(gameId) {
        var game = this._getGame(gameId);
        assert(this._isMarketAdmin(Blockchain.transaction.from), "only admins can confirm result");
        assert(game.status == 2, "game is not in status of previewing result");

        game.status = 3;
        
        if (game.type === 1) {
            var option = game.getOption(game.result);
            //庄家领回总资金减去赔付额后的资金
            var balance = game.deposit - option.expectReward;
            game.deposit = option.expectReward;
            game.ownerDeposit -= balance;
            this._transfer(Blockchain.transaction.to, game.owner, balance, game.payType);
            
        }
        this._setGame(gameId, game);
    },

    //for users
    buyTicket: function(gameId, optionNo, optionVersion, amount) {
        var game = this._getGame(gameId);
        assert(game.status == 1, "No bet in this status of game");
        assert(game.optionVersion === optionVersion, "option version has been updated");        

        //扣费
        amount = this._transfer(Blockchain.transaction.from, Blockchain.transaction.to, amount, game.payType);

        assertInteger(optionNo);
        assertPosInteger(optionVersion);
        var date = Date.now();
        assert(date <= game.deadLine, "deadline is passed");

        var option = game.getOption(optionNo);
        var expectReward = amount * option.odd;
        //TODO:debug
        console.log(optionNo);
        console.log(JSON.stringify(option));
        console.log(JSON.stringify(game.getOption(1)));
        console.log(JSON.stringify(game.getOption(2)));
        console.log(JSON.stringify(game.getOption(3)));
        console.log(expectReward);
        console.log(option.expectReward);
        console.log(game.deposit);

        assert(expectReward + option.expectReward <= game.deposit + amount, "remaining deposit is not enough");
        
        //生成token
        var ticketId = this.nextTicketCount;
        this.nextTicketCount += 1;
        var ticket = new Ticket(ticketId, amount, gameId, optionNo, option.odd);
        ticket.status = 1;
        this._setTicket(ticketId, ticket);
        this._setTicketOwner(ticketId, Blockchain.transaction.from);

        //更新数据
        option.bets = option.bets + amount;
        option.expectReward = option.expectReward + expectReward;
        game.setOption(optionNo, option);
        game.bets = game.bets + amount;
        game.deposit = game.deposit + amount;
        this._setGame(gameId, game);
        
        //返回TicketId
        return {ticketId: ticketId};
    },

    getReward: function(ticketId) { //主动调用， server存储 gameid->ticketIds 的表
        var ticket = this._getTicket(ticketId);
        var owner = this._getTicketOwner(ticketId);
        var gameId = ticket.gameId;
        var game = this._getGame(gameId);
        assert(game.status === 3, "the game result has not been confirm");
        assert(ticket.status === 1, "the ticket has been awarded");
        assert(ticket.optionNo === game.result, "loose the game!");

        var expectReward;
        if (game.type == 1) {
            expectReward = ticket.amount * ticket.odd;
        } else if (game.type === 2) {
            var option = game.getOption(game.result);
            expectReward = ticket.amount / option.bets * game.depositAtLast;
            expectReward = parseInt(expectReward * 100) / 100;
        } else {
            throw "unexpected error, game type is undifined";
        }
        game.rewardBets = game.rewardBets - ticket.amount;
        if (game.rewardBets == 0) {
            game.status = 4; //all tickets have been rewarded
        }
        game.deposit = game.deposit - expectReward;
        ticket.status = 2;
        this._transfer(Blockchain.transaction.to, owner, expectReward, game.payType);
        this._setTicket(ticketId, ticket);
        this._setGame(gameId,game);
    },

    getRemainDeposit: function(gameId){
        var game = this._getGame(gameId);
        game.assertOwner(Blockchain.transaction.from);
        assert(game.status == 4, "game is not in status to get remaining deposit");
        assert(game.rewardBets == 0, "unexpected error, bets shoud be 0 when status is 4");

        this._transfer(Blockchain.transaction.to, Blockchain.transaction.from, game.deposit, game.payType);
        game.deposit = 0;
    },

    transferTicket: function(ticketId, to) {
        assertInteger(ticketId);
        assert(Blockchain.verifyAddress(to) != 0, "to address is invalid");

        var owner = this._getTicketOwner(ticketId);

        assert(Blockchain.transaction.from == owner, "only ticket owner can transfer ticket");
        this._setTicketOwner(ticketId, to);
        this._transferTicketEvent(Blockchain.transaction.from, to, ticketId);
    },

/*
================================查询接口===============================
*/
    getGame: function(id) {
        assertPosInteger(id);
        var game  = this._getGame(id);
        var options = new Array();
        for (var i = 1; i < game.nextOptionCount; i++) {
            var option = game.getOption(i);
            options.push(option);
        }
        var result = {
            game: game,
            options: options,
        };

        return result;
    },

    getTicket: function(ticketId) {
        assertPosInteger(ticketId);
        var ticket = this._getTicket(ticketId);
        var owner = this._getTicketOwner(ticketId);
        var result = {
            ticket: ticket,
            owner: owner,
        };

        return result;
    },

/*
================================私有方法===============================
*/
    _isMarketAdmin: function(address) {
        return this.Admins.get(address) === 1;
    },

    _getGameId: function() {
        var id = this.nextGameCount;
        this.nextGameCount++;
        return id;
    },

    _getGame: function(gameId) {
        assertPosInteger(gameId);
        var game = this.Games.get(gameId);
        if (game) {
            assert(gameId === game.id, "unexpected error, data not consistency");
            return game;
        } else {
            throw "invalid gameId";
        }
    },

    _setGame: function(gameId, game) {
        console.log(game);//TODO: debug
        this.Games.set(gameId, game);
    },

    _getTicket: function(ticketId) {
        console.log(ticketId);//TODO: debug
        assertPosInteger(ticketId);
        var ticket = this.Tickets.get(ticketId);
        if (!ticket) {
            throw "invalid ticketId, failed to get ticket";
        }
        return ticket;
    },

    _setTicket: function(ticketId, ticket) {
        this.Tickets.set(ticketId, ticket);
    },

    _getTicketOwner: function(ticketId) {
        var owner = this.TicketIdToOwner.get(ticketId);
        if (!owner) {
            throw "invalid ticketId, failed to get ticket";
        }
        return owner;
    },

    _setTicketOwner: function(ticketId, owner) {
        this.TicketIdToOwner.set(ticketId, owner);
    },

    _transfer: function(from, to, amount, payType) {// 兑换bet, 并交易
        if (payType === 1) {
            assert(!Blockchain.transaction.value.gt(0), "this game is not support for nas");
            assertInteger(amount);
            assert(amount >= 0, "amount shoud large or equal than 0");
            var tokenMgr = this._getTokenMgr();
            var value = new BigNumber(amount);
            value = value.mul(1000000000000000000);// 10^18
            tokenMgr.transferByContract(from, to, value);
            return amount;
        } else if (payType === 2) {
            if (to === Blockchain.transaction.to) {
                amount = Blockchain.transaction.value.div(100000000000000000);// 10^17
                assert(amount.gt(0), "invalid amount, amount should big than 0");
                assert(amount.isInt(), "invalid amount, amount should be a multiple of 0.1 nas");
                return parseInt(amount.toString());
            } else if (from === Blockchain.transaction.to) {
                var value = new BigNumber(100000000000000000);//10^17
                value = value.mul(amount);
                Blockchain.transfer(to, value);
                return amount;
            } else {
                throw "unexpected error, form and to are both not contract address";
            }
        } else {
            throw "invalid pay type";
        }
    },

    _getTokenMgr: function() {
        return new StandardToken();
    },

    _createGameEvent: function(game) {
        Event.Trigger(this.name(), {
            creator: game.owner,
            gameId: game.id,
        });
    },

    _delGameEvent: function(game) {
        Event.Trigger(this.name(), {
            deletor: game.owner,
            id: game.id, 
        });
    },

    _transferTicketEvent: function(from, to, ticketId) {
        Event.Trigger(this.name(), {
            from: from,
            to: to,
            ticketId: ticketId,
        });
    },

/*
=====================================interfaces for NRC20=====================================
*/
    // Returns the name of the ticket
    name: function () {
        return this._getTokenMgr().name();
    },

    // Returns the symbol of the ticket
    symbol: function () {
        return this._getTokenMgr().symbol();
    },

    // Returns the number of decimals the ticket uses
    decimals: function () {
        return this._getTokenMgr().decimals();
    },

    totalSupply: function () {
        return this._getTokenMgr().totalSupply();
    },

    balanceOf: function (owner) {
        return this._getTokenMgr().balanceOf(owner);
    },

    transfer: function (to, value) {
        return this._getTokenMgr().transfer(to, value);
    },

    transferFrom: function (from, to, value) {
        return this._getTokenMgr().transferFrom(from, to, value);
    },

    approve: function (spender, currentValue, value) {
        return this._getTokenMgr().approve(spender, currentValue, value);
    },

    allowance: function (owner, spender) {
        return this._getTokenMgr().approve(owner, spender);
    }
};

module.exports = Market;

