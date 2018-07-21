// pages/game/index.js
var wegame = require('../../models/wegame')

var sliderWidth = 96; // 需要设置slider的宽度，用于计算中间位置

Page({

  /**
   * 页面的初始数据
   */
  data: {
    tabs: ["竞猜", "我的", "好友"],
    activeIndex: 0,
    sliderOffset: 0,
    sliderLeft: 0,
    beting:'hidden',
    betdata:{}
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    var that = this;

    var game_id = parseInt(options.game_id)
    that.setData({
      game_id:game_id
    })

    wx.getSystemInfo({
      success: function (res) {
        console.log(res)
        that.setData({
          sliderLeft: (res.windowWidth / that.data.tabs.length - sliderWidth) / 2,
          sliderOffset: res.windowWidth / that.data.tabs.length * that.data.activeIndex
        });
      }
    });

    this.getGameDetail(game_id);
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {
  
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
  
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {
  
  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {
  
  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {
  
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {
  
  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {
  
  },

  tabClick: function (e) {
    this.setData({
      sliderOffset: e.currentTarget.offsetLeft,
      activeIndex: e.currentTarget.id
    });
  },

  onBetEventTap:function(e){
    console.log(e)

    var betdata = this.data.betdata
    betdata.choiceId = e.currentTarget.dataset.choiceId
    betdata.choiceName = e.currentTarget.dataset.choiceName
    betdata.choiceDesc = e.currentTarget.dataset.choiceDesc
    betdata.choiceOdds = e.currentTarget.dataset.choiceOdds
    betdata.choiceType = e.currentTarget.dataset.choiceType
    betdata.reward = "-"

    this.setData({
      beting:'show',
      betdata: betdata
    })
  },

  onHiddenBetEventTap:function(e){
    this.setData({
      beting: 'hidden'
    })
  },

  onChangeBetEventTap:function(e){

    var betdata = this.data.betdata

    betdata.value = parseInt(e.currentTarget.dataset.value)
    betdata.reward = parseInt(betdata.value * betdata.choiceOdds)+'QB'
    
    this.setData({
      beting: 'show',
      betdata: betdata
    })
  },

  getGameDetail(game_id) {

    var that = this
    var params = {
      game_id: game_id
    }
    wegame.getGameDetail(params, function (code, msg, datas) {
      if (code != 200) {

      }
      console.log(datas)
      that.setData({
        game: datas.game,
        quizzes: datas.quizzes
      })
    })
  }

})