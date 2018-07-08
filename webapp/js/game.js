var getUrlParameter = function getUrlParameter(sParam) {
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
};

jQuery(document).ready(function ($) {

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



    var game_id = getUrlParameter("game_id");
    game_id = parseInt(game_id)

    $(".game-card-tab-item").click(function () { 
        $(".game-card-tab-item").removeClass("mui-active")
        $(this).addClass("mui-active");
       
        $(".game-card-content-item").addClass("hide")
        var tab_item = $(this).attr("tab-item");
        $("#"+tab_item).removeClass("hide");
        console.log("#"+tab_item)
    })


    var choice_opt = {
        choice_id: 0,
        quizzes_id: 0,
        odds: 1,
        name: "",
        bet: 0,
    }

    var odds = 1;
    var name = "";
    $("#guess-items").on("click", ".quizzes-choice-item", function () {
        choice_opt.name = $(this).attr("quizzes-choice-item-name");
        choice_opt.odds = $(this).attr("quizzes-choice-item-odds");
        choice_opt.choice_id = $(this).attr("quizzes-choice-item-id");

        choice_opt.quizzes_id = $(this).parents(".quizzes-choice-items").attr("quizzes-id");

        console.log($(this).parents(".quizzes-choice-items"))
        choice_opt.bet = 0;
        console.log(name, odds, choice_opt)

        $('#game-bet-name').html("投注: " + choice_opt.name);
        $('#game-bet-odds').html(choice_opt.odds);
        $('#game-bet-reward').html(choice_opt.odds * choice_opt.bet + 'QB');

        $('#bet_bgc').fadeIn(100);
        $('#bet_banner').slideDown(100);
    })

    
    $(".game-bet-change-money").click(function () {

        choice_opt.bet = $(this).attr("bet-money");

        $(this).addClass("mui-active");

        console.log(this)

        $('#game-bet-reward').html(parseInt(choice_opt.odds * choice_opt.bet) + 'QB');
    })
    


    $("#game-quizz-bet").click(function () {

        var user_id = parseInt($.cookie('user_id'));
        if (user_id == 0) {
            alert("请登录");
            window.location.href = "login.html";
            return;
        }
        if (choice_opt.bet<=0){
            alert("请选择下注额度");
            return;
        }
        var params = {
            game_id: game_id ? parseInt(game_id) : 0,
            "auth_token": $.cookie('auth_token'),
            "bet_amount": parseInt(choice_opt.bet),
            "choice_opt_id": parseInt(choice_opt.choice_id),
            "quizzes_id": parseInt(choice_opt.quizzes_id),
            "user_id": user_id
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
    })

    function load() {
        var params = {
            game_id: game_id ? parseInt(game_id) : 1
        };
        var url = "http://35.180.103.230:8080/v1/game/quizzes";

        $.ajax({
            type: "POST",
            url: url,
            data: JSON.stringify(params),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (res) {
                console.log(res)
                showGame(res.data.game)

                showQuizzes(res.data.quizzes)
            }
        });
    }


    function showMyQuizzesChoice(game_id) {
        var user_id = parseInt($.cookie('user_id'));

        var params = {
            "auth_token": $.cookie('auth_token'),
            "user_id": user_id,
            "game_id": game_id,
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

    function showGame(game) {

        document.title = game.home_team.name + 'VS' + game.visit_team.name;

        $("#game-item-tpl").tmpl(game).appendTo(".game-item");

        $(".game-card-tab").show();
    }

    function showQuizzes(quizzes) {
        $("#game-quizz-tpl").tmpl(quizzes).appendTo("#guess-items");
    }

    $('.theme-popover-mask').click(function () {
        $('.theme-popover-mask').fadeOut(100);
        $('.bet-theme-popover').slideUp(100);
    })

    load()
    showMyQuizzesChoice(game_id)
})