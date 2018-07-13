class User{
    constructor(){
        var user_id =  $.cookie('user_id')

        if (user_id == null || user_id == 'null'){
            user_id = 0;
        }else{
            user_id = parseInt(user_id);
            user_id = isNaN(user_id)?0:user_id;
        }

        this._setUser(user_id, $.cookie('username'), $.cookie('auth_token'), $.cookie('balance'), $.cookie('hash_address'))
    }

    getUserId(){
        return this.user_id;
    }

    getAuthToken(){
        return this.auth_token;
    }

    getBalance(){
        return this.balance;
    }

    isLogin() {
        return this.user_id > 0 ? true : false;
    }

    Login(params, callback){
        params = JSON.stringify(params);

        $.ajax({
            type: "POST",
            url: 'http://35.180.103.230:8080/v1/user/login',
            data: loginData,
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function(res){
                if(res.code != 200 ){
                    callback(-1, res.msg)
                    return
                }
                var data = res.data;
                console.log(data)

                $.cookie('auth_token', data.auth_token);
                $.cookie('username', data.user.username);
                $.cookie('balance', data.user.balance);
                $.cookie('hash_address', data.user.hash_address);
                $.cookie('user_id', data.user.Id)

                this._setUser(data.user.Id, data.user.username, data.auth_token, data.user.balance, data.user.hash_address)

                callback(0, '', res.data)

            },
            error: function (err) {
                callback(-1, err)
            }
        })
    }


    Logout(){
        
        $.cookie('username', null);
        $.cookie('user_id', null);
        $.cookie('auth_token', null);
        $.cookie('balance', null);
        $.cookie('hash_address', null);

        this._setUser(0, null, null, null, null)
    }


    _setUser(user_id, user_name, auth_token, balance, hash_address){
        this.user_id = user_id;
        this.user_name = user_name;
        this.balance = balance;
        this.hash_address = hash_address;
        this.auth_token = auth_token;
    }
}