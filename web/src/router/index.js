import Vue from 'vue'
import VueRouter from 'vue-router'


import login from "../views/login/Login.vue"


// import siderBar from '../views/admin/SiderBar.vue'

Vue.use(VueRouter)

const routes = [{
    path: "/",
    redirect: "/login",
  },
  // 登录
  {
    path: "/login",
    name: "login",
    component: login
  },
  // ----------------------主页----------------
  {
    path: "/center",
    name: "center",

    component: () => import( /* webpackChunkName: "about" */ '../views/center/SiderBar.vue'),
    meta: {
      requireAuth: true,
    },
    children: [{
        path: "",
        redirect: "HomePage",
      },
      // 首页
      {
        path: "/Home",
        name: "HomePage",
        component: () => import("../views/center/HomePage.vue"),
        meta: {
          requireAuth: true,
        },
      },
      {
        path: '/allShow',
        name: "AllShow",
        component: () => import("../views/center/AllShow.vue"),
        meta: {
          requireAuth: true,
        },
      },
      {
        path: "/rankShow",
        name: "RankShow",
        component: () => import("../views/center/RankShow.vue"),
        meta: {
          requireAuth: true,
        },
      },
      {
        path: "/SchoolCompe",
        name: "SchoolCompe",
        component: () => import("../views/center/SchoolCompe.vue"),
        meta: {
          requireAuth: true,
        },
      },
      // 信息页
      {
        path: "/info",
        name: "info",
        component: () => import("../views/center/lookupInfo.vue"),
        meta: {
          requireAuth: true,
        },

      },
    ]
  },
  //用户
  {
    path: "/user",
    name: "user",
    component: () => import('../views/user/SiderBar.vue'),
    meta: {
      requireAuth: true,
    },
    children: [{
        path: "",
        redirect: "mineCenter",
      },
      {
        path: "/user/mineCenter",
        name: "mineCenter",
        component: () => import("../views/user/mineCenter.vue"),
        meta: {
          title: "个人中心",
          requireAuth: true,
        },
      },
      // ——————————————————————————————创建的竞赛————————————————————
      // 已完成的比赛
      {
        path: "/user/create/doneTemp",
        name: "doneTemp",
        component: () => import("../views/user/create/CreateDoneTemp.vue"),
        meta: {
          title: "创建的题库",
          requireAuth: true,
        },
      },
      // 正在进行的比赛
      {
        path: "/user/create/goingTemp",
        name: "goingTemp",
        component: () => import("../views/user/create/CreateGoingTemp.vue"),
        meta: {
          title: "创建的题库",
          requireAuth: true,
        },
      },
      // 进行中总览
      {
        path: "/user/create/TempAll",
        name: "goingTempAll",
        component: () => import("../views/user/create/CreateTempAll.vue"),
        meta: {
          title: "创建的题库",
          requireAuth: true,
        }
      },
      {
        path: '/user/create/editExam',
        name: "newTemp",
        component: () => import("../components/edit/editExam.vue"),
        meta: {
          title: "题库编辑",
          requireAuth: true,
        },
      },
      // ——————————————————————————————加入的竞赛————————————————————
      {
        path: "/user/join/doneTemp",
        name: "doneTemp",
        component: () => import("../views/user/join/JoinDoneTemp.vue"),
        meta: {
          title: "收藏的题库",
          requireAuth: true,
        },
      },
      // 正在进行的比赛
      {
        path: "/user/join/goingTemp",
        name: "goingTemp",
        component: () => import("../views/user/join/JoinGoingTemp.vue"),
        meta: {
          title: "收藏的题库",
          requireAuth: true,
        },
      },
      // 已完成总览
      {
        path: "/user/join/TempAll",
        name: "doneTempAll",
        component: () => import("../views/user/join/JoinTempAll.vue"),
        meta: {
          title: "收藏的题库",
          requireAuth: true,
        },
      },







    ]
  },




  // 管理员的
  {
    path: "/admin",
    name: "admin",
    component: () => import("../views/admin/SiderBar.vue"),
    meta: {
      requireAuth: true,
    },
    children: [{
        path: "/admin",
        redirect: "alys",
      }, {
        path: "/admin/alys",
        name: "alys",
        component: () => import("../views/admin/alys.vue"),
        meta: {
          title: "数据分析",
          requireAuth: true,
        },
      },

      {
        // 认证审核
        path: "/admin/check",
        name: "alys",
        component: () => import("../views/admin/check.vue"),
        meta: {
          title: "认证审核",
          requireAuth: true,
        },
      },
      // editHome
      {
        // 标签修改
        path: "/admin/editTag",
        name: "editTag",
        component: () => import("../views/admin/editHome/editTag.vue"),
        meta: {
          title: "标签管理",
          requireAuth: true,
        },
      },
      {
        // 标签修改
        path: "/admin/editImgs",
        name: "editImgs",
        component: () => import("../views/admin/editHome/editImgs.vue"),
        meta: {
          title: "轮播图管理",
          requireAuth: true,
        },
      },

    ]
  },


  // 考试的
  {
    path: "/exam",
    name: "examSider",
    component: () => import("../views/examPage/SiderBar.vue"),
    meta: {
      requireAuth: true,
    },
    children: [{
        path: "",
        redirect: "/exam/classRoom",
      },
      {

        path: "/exam/classRoom",
        name: "classRoom",
        component: () => import("../views/examPage/classRoom.vue"),
        meta: {
          title: "考场",
          requireAuth: true,
        },
      },
      // 考试模式
      {
        path: "/exam/test",
        name: "test",
        component: () => import("../views/examPage/exam51.vue"),
        meta: {
          title: "考场",
          requireAuth: true,
        },
      },
      // 闯关模式
      {
        path: "/exam/breakThrough",
        name: "breakThrough",
        component: () => import("../views/examPage/exam52.vue"),
        meta: {
          title: "考场",
          requireAuth: true,
        },
      },
      // 限时答题
      {
        path: "/exam/timeLimit",
        name: "timeLimit",
        component: () => import("../views/examPage/exam53.vue"),
        meta: {
          title: "考场",
          requireAuth: true,
        },
      }
    ]
  },




]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})



router.beforeEach((to, from, next) => {
  // console.log("to.meta.requireAuth", to.meta.requireAuth);
  // console.log("to", to);

  // console.log("next", next);
  if (to.meta.requireAuth) {
    if (window.sessionStorage.getItem("accessToken")) {
      next();
    } else {
      next({
        path: "/login"
      })
    }
  } else {
    next()
  }

  if (to.path == "/exam/breakThrough" || to.path == "/exam/test" || to.path == "/exam/timeLimit") {
    // etf 考试没有
    let flag = window.sessionStorage.getItem("etf")
    // console.log("etf flag", flag);
    if (flag) {
      // console.log("etf flag next");
      next();
    } else {
      next({
        path: "/user/join/goingTemp"
      })
    }
  }
})

export default router