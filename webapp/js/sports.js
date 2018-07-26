class Sports{
    constructor(){

    }

    // Do some initialize when page load.
    onLoad(options) {
        
        //this.mask = mui.createMask(function(){});//callback为用户点击蒙版时自动执行的回调；
        //this.mask.show();//显示遮罩

        var that = this
        //that.getGameList()
        
        mui.init({
            
            pullRefresh: {
                container: '#game-items-scroll',
                down: {
                    style:'circle',
                    callback: function(){
                        that.getGameList()
                    }
                },
                
                up: {
                    auto:true,
                    contentrefresh: '正在加载...',
                    callback: function(){
                        that.getGameList()
                    }
                }
            }
            
        });
        
        that.getGameList()

        this.onEvent()
       
    };

    onEvent(){
        $("#game-items").on("click", ".game-item", function () {
            var game_id = $(this).attr("game_id");
            console.log(game_id)
    
            $(location).attr('href', 'game.html?game_id=' + game_id);
        })
    }

    onReady() {
    // Do something when page ready.
    };

    // Do something when page show.
    onShow() {
        //this.mask.close();

    };

    // Do something when page hide.
    onHide() {

    };


    getGameList(){
        mui.showLoading("正在加载..","div"); 
        var that = this
        var url = "http://35.180.103.230:8080/v1/game/list";
        var data = {};
        $.ajax({
            type: "POST",
            url: url,
            data: data,
            dataType: "json",
            success: function (res) {
                mui.hideLoading(function(){});
                console.log(res)
                //mui('#game-items-main').pullRefresh().endPulldownToRefresh(true);
                that.showGameHtml(res.data.List)
                
                //that.onShow();
            }
        });
    };

    showGameHtml(datas){
        //$("#game-items").html("");
        $("#game-item-tpl").tmpl(datas).appendTo("#game-items");

    }
}

/*
$(function () {
   
})
*/
