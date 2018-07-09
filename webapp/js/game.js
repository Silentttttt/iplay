
class Game{

    constructor(userObj){
        this.userObj = userObj;

        var game_id = this.getUrlParameter("game_id");
        this.game_id = parseInt(game_id)

        this.user_id = this.userObj.getUserId();

        this.choice_opt = {
            choice_id: 0,
            quizzes_id: 0,
            odds: 1,
            name: "",
            bet: 0,
        }
    }

    onLoad(){

        this.getGameInfo();
        this.showMyQuizzesChoice();
        this.onEvent()
    }

    getUrlParameter(sParam) {
        var sPageURL = decodeURIComponent(window.location.search.substring(1)),
            sURLVariables = sPageURL.split('&'),
            sParameterName,
            i;
    
        for (i = 0; i < sURLVariables.length; i++) {
            sParameterName = sURLVariables[i].split('=');
    
            if (sParameterName[0] === sParam) {
                return sParameterName[1] === undefined ? true : sParameterName[1];
            }
        }
    }


    onEvent(){
         // 打开  好友pk
        $("#friend_play").click(function () {
            $("#haoyoupk_bgc").fadeIn(100);
            $('#haoyoupk_banner').slideDown(200);
        })

        // 关闭 好友pk
        $('#friend_play .close').click(function () {
            $('#haoyoupk_bgc').fadeOut(100);
            $('#haoyoupk_banner').slideUp(200);
        })

        // 关闭下注
        $('#bet_banner .close').click(function () {
            $('#bet_bgc').fadeOut(100);
            $('#bet_banner').slideUp(200);
        })

        $(".game-card-tab-item").click(function () { 
            $(".game-card-tab-item").removeClass("mui-active")
            $(this).addClass("mui-active");
           
            $(".game-card-content-item").addClass("hide")
            var tab_item = $(this).attr("tab-item");
            $("#"+tab_item).removeClass("hide");
            console.log("#"+tab_item)
        })


        $('.theme-popover-mask').click(function () {
            $('.theme-popover-mask').fadeOut(100);
            $('.bet-theme-popover').slideUp(100);
        })

        this.onEventBet()
    }


    onEventBet(){
        var that = this
        var odds = 1;
        var name = "";
        $("#guess-items").on("click", ".quizzes-choice-item", function () {
            that.choice_opt.name = $(this).attr("quizzes-choice-item-name");
            that.choice_opt.odds = $(this).attr("quizzes-choice-item-odds");
            that.choice_opt.choice_id = $(this).attr("quizzes-choice-item-id");
    
            that.choice_opt.quizzes_id = $(this).parents(".quizzes-choice-items").attr("quizzes-id");
    
            console.log($(this).parents(".quizzes-choice-items"))
            that.choice_opt.bet = 0;
            console.log(name, odds, that.choice_opt)
    
            $('#game-bet-name').html("投注: " + that.choice_opt.name);
            $('#game-bet-odds').html(that.choice_opt.odds);
            $('#game-bet-reward').html(that.choice_opt.odds * that.choice_opt.bet + 'QB');
    
            $('#bet_bgc').fadeIn(100);
            $('#bet_banner').slideDown(100);
        })
    
        
        $(".game-bet-change-money").click(function () {
    
            that.choice_opt.bet = $(this).attr("bet-money");
    
            $(this).addClass("mui-active");
    
            console.log(this)
    
            $('#game-bet-reward').html(parseInt(that.choice_opt.odds * that.choice_opt.bet) + 'QB');
        })

        $("#game-quizz-bet").click(function () {
            that.doBet()
        })
    };


    doBet() {

        if(!this.userObj.isLogin()){
            var btnArray = ['否', '是'];
            mui.confirm('还未登录，请登录', '请登录', btnArray, function(e) {
                if (e.index == 1) {
                    window.location.href = "login.html";
                } else {
                    
                }
            })
            return 
        }
        
        if (this.choice_opt.bet<=0){
            mui.toast('请选择下注额度',{ duration:'long', type:'div' });
            return;
        }
        var params = {
            game_id: this.game_id,
            "auth_token": this.userObj.getAuthToken(),
            "bet_amount": parseInt(this.choice_opt.bet),
            "choice_opt_id": parseInt(this.choice_opt.choice_id),
            "quizzes_id": parseInt(this.choice_opt.quizzes_id),
            "user_id": this.user_id
        };
        var url = "http://35.180.103.230:8080/v1/user/do_quizzes";

        console.log(params)
        $.ajax({
            type: "POST",
            url: url,
            data: JSON.stringify(params),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (res) {
                console.log(res)

                if (res.code == 200) {
                    alert('下注成功');
                    
                    $(".game-card-tab-item").removeClass("mui-active");
                    $(".game-card-tab-item-myself").addClass("mui-active");

                    $(".game-card-content-item").addClass("hide")
                    $("#myself-items").removeClass("hide")

                    //showMyQuizzesChoice(game_id)
                    $('.theme-popover-mask').fadeOut(100);
                    $('.bet-theme-popover').slideUp(200);
                    
                } else if (res.code == 501) {
                    alert("请先登录");
                    window.location.href = "login.html"
                } else {
                    alert(res.msg);
                }
            }
        });
    }


    getGameInfo() {
        var params = {
            game_id: this.game_id
        };
        var url = "http://35.180.103.230:8080/v1/game/quizzes";

        var that = this
        $.ajax({
            type: "POST",
            url: url,
            data: JSON.stringify(params),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (res) {
                console.log(res)
                that.showGame(res.data.game)

                that.showQuizzes(res.data.quizzes)
            }
        });
    }

    showMyQuizzesChoice() {

        var params = {
            "auth_token": this.userObj.getAuthToken(),
            "user_id": this.user_id,
            "game_id": this.game_id,
        };
        var url = "http://35.180.103.230:8080/v1/user/quizzes_list";

        $.ajax({
            type: "POST",
            url: url,
            data: JSON.stringify(params),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (res) {
                console.log(res)

                if (res.code == 200 && res.data.length > 0) {
                    $("#myself-items").html("");
                    $("#game-quizz-my-tpl").tmpl(res.data[0].quizzes).appendTo("#myself-items");
                }

            }
        });
    }

    showGame(game) {

        document.title = game.home_team.name + 'VS' + game.visit_team.name;

        $("#game-item-tpl").tmpl(game).appendTo(".game-item");

        $(".game-card-tab").show();
    }

    showQuizzes(quizzes) {
        $("#game-quizz-tpl").tmpl(quizzes).appendTo("#guess-items");
    }

    
}