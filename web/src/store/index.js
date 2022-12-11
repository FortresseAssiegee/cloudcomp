import Vue from 'vue'
import Vuex from 'vuex'
import {
  Base64
} from 'js-base64';

Vue.use(Vuex)

export default new Vuex.Store({
  // 静态
  state: {
    userId: sessionStorage.getItem(`userId`) || ``,
    // auth配置
    config: {
      headers: {
        // "Content-Type": "application/json",
        Authorization: sessionStorage.getItem(`accessToken`) || ``,
      },
    },
    // 所在的活动id
    actvId: sessionStorage.getItem(`actvId`) || ``,

    // 用户信息
    userInfo: sessionStorage.getItem(`userInfo`) || ``,
    actvBaseInfo: sessionStorage.getItem(`actvBaseInfo`) || ``,
    unitInfo: sessionStorage.getItem(`unitInfo`) || ``,
    examInfo: sessionStorage.getItem(`examInfo`) || ``,

    // 码
    code: new Map([
      [0, "不存在"],
      // stateCode
      [10, "未发布"],
      [11, "进行中"],
      [12, "已结束"],
      [13, "未报名"],
      [14, "审核"],
      [15, "未通过"],

      // awardCode
      [30, "无奖品"],
      [31, "有奖品"],
      // periodCode
      [40, " 不确定"],
      [41, "半个月内"],
      [42, "一个月内"],
      [43, "三个月内"],
      [44, "半年内"],
      [45, "一年内"],
      [46, "一年以上"],

      [50, '混合模式'],
      [51, '考试模式'],
      [52, '闯关模式'],
      [53, '限时答题模式'],

      [70, "学习&考试"],
      [71, "学习"],
      [72, "考试"]

    ]),

    boolToStringFace: new Map([
      [true, "开启人脸识别"],
      [false, "不开人脸识别"],
    ]),
    boolToString: new Map([
      [true, "有奖品"],
      [false, "无奖品"],
    ]),
    boolToCode: new Map([
      [true, 31],
      [false, 30],
    ]),
    awardCode: new Map([
      [31, true],
      [30, false]
    ]),

    compeTypeCode: [{
      label: "混合模式",
      id: 50
    }, {
      label: '考试模式',
      id: 51
    }, {
      label: "闯关模式",
      id: 52
    }, {
      label: "限时答题模式",
      id: 53,
    }],

    unitTypeCode: [{
      label: '考试模式',
      id: 51
    }, {
      label: "闯关模式",
      id: 52
    }, {
      label: "限时答题模式",
      id: 53,
    }],
    modelCode: [{
      label: "学习&考试",
      id: 70
    }, {
      label: "学习",
      id: 71
    }, {
      label: "考试",
      id: 72
    }, ]

  },

  // 不会修改数据，只是进行包装，类似于计算数学
  getters: {
    // 用户信息
    userInfo(state) {
      let obj = Base64.decode(sessionStorage.getItem(`userInfo`) || ``)
      // console.log("解码", obj);
      if (state.userInfo) {
        // console.log("parse", JSON.parse(obj));
        return JSON.parse(obj)

      } else {
        return obj
      }
    },
    actvBaseInfo(state) {
      let obj = Base64.decode(sessionStorage.getItem(`actvBaseInfo`) || ``)

      if (state.actvBaseInfo) {
        return JSON.parse(obj)

      } else {
        return obj
      }
    },
    unitInfo(state) {
      let obj = Base64.decode(sessionStorage.getItem(`unitInfo`) || ``)
      if (state.unitInfo) {
        return JSON.parse(obj)
      } else {
        return obj
      }
    },
    examInfo(state) {
      let obj = Base64.decode(sessionStorage.getItem(`examInfo`) || ``)
      if (state.examInfo) {
        return JSON.parse(obj)
      } else {
        return obj
      }
    },

  },


  // 用做修改state
  mutations: {
    setConfig: (state, token) => {
      sessionStorage.setItem(`accessToken`, token)
      state.config.headers.Authorization = token
      // console.log("mutations setConfig:", state.config.headers.Authorization);
    },
    setUserId: (state, id) => {
      sessionStorage.setItem(`userId`, id)
      state.userId = id
      // console.log("mutations setUserId:", state.userId);
    },

    // 加密
    setUserInfo: (state, info) => {
      let objStr = JSON.stringify(info)
      sessionStorage.setItem(`userInfo`, Base64.encode(objStr))
      state.userInfo = info
      // console.log("mutations userInfo:", state.userInfo);
    },
    setActvBaseInfo: (state, baseInfo) => {
      let objStr = JSON.stringify(baseInfo)
      sessionStorage.setItem(`actvBaseInfo`, Base64.encode(objStr))
      state.actvBaseInfo = baseInfo
      // console.log("mutations setActvBaseInfo:", state.actvBaseInfo);
    },
    setUnitInfo: (state, unitInfo) => {
      let objStr = JSON.stringify(unitInfo)
      sessionStorage.setItem(`unitInfo`, Base64.encode(objStr))
      state.unitInfo = unitInfo
      // console.log("mutations setUnitInfo:", state.unitInfo);
    },
    setExamInfo: (state, examInfo) => {
      let objStr = JSON.stringify(examInfo)
      sessionStorage.setItem(`examInfo`, Base64.encode(objStr))
      state.examInfo = examInfo
      // console.log("mutations setExamInfo:", state.examInfo);
    }

  },
  // 执行异步操作
  actions: {

  },
  modules: {}
})