<template>
  <v-app id="inspire">
    <!-- 消息框 -->
    <snackerMsg ref="msg"></snackerMsg>

    <v-app-bar
      dark
      scroll-target="#scrolling-techniques-3"
      src="https://picsum.photos/1920/1080?random"
      fade-img-on-scroll
      absolute
      color="#6A76AB"
      shrink-on-scroll
      prominent
    >
      <template v-slot:img="{ props }">
        <v-img
          v-bind="props"
          gradient="to top right, rgba(100,115,201,.7), rgba(25,32,72,.7)"
        ></v-img>
      </template>
      <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>

      <v-toolbar-title>线上知识竞赛平台</v-toolbar-title>
      <v-spacer></v-spacer>

      <!-- 退出 -->

      <!-- <v-menu transition="slide-y-transition" offset-y offset-x min-width="150">
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon :ripple="false" v-bind="attrs" v-on="on">
            <v-icon class="material-icons-round text-white"
              >mdi-dots-vertical</v-icon
            >
          </v-btn>
        </template> -->

      <v-tooltip bottom>
        <template v-slot:activator="{ on, attrs }">
          <v-btn icon :ripple="false" v-bind="attrs" v-on="on" @click="toLogin">
            <v-icon class="material-icons-round text-white" color="red">
              mdi-exit-to-app</v-icon
            >
          </v-btn>
        </template>
        <span> 退出登录</span>
      </v-tooltip>
      <!-- </v-menu> -->

      <template v-slot:extension>
        <v-tabs align-with-title>
          <v-tab to="/Home">发现</v-tab>
          <v-tab to="/AllShow">题库总览</v-tab>
          <!-- <v-tab to="/RankShow">排行榜单</v-tab> -->
          <v-tab to="/SchoolCompe">合作企校</v-tab>
        </v-tabs>
      </template>
    </v-app-bar>

    <v-sheet
      id="scrolling-techniques-3"
      class="overflow-y-auto"
      max-height="100vh"
    >
      <router-view />
    </v-sheet>
    <drawer v-bind:getdrawer="drawer" v-on:putdrawer="drawer = $event"></drawer>
    <!-- <v-container > -->
    <!-- </v-container> -->
  </v-app>
</template>

<script>
import { mapMutations, mapState } from "vuex";
import snackerMsg from "../../components/commont/snackerMsg.vue";
import drawer from "../../components/commont/drawer.vue";
export default {
  components: {
    snackerMsg,
    drawer,
  },
  data: () => ({
    drawer: false,
    createDoingList: [{}], //创建成功
    createDoneList: [{}], //创建完成
    items: [
      {
        text: "Dashboard",
        disabled: false,
        href: "breadcrumbs_dashboard",
      },
      {
        text: "Link 1",
        disabled: false,
        href: "breadcrumbs_link_1",
      },
      {
        text: "Link 2",
        disabled: true,
        href: "breadcrumbs_link_2",
      },
    ],
    playShow: true,
  }),
  computed: {
    // 用户信息
    ...mapState(["userInfo", "config"]),
  },
  methods: {
    ...mapMutations(["setUserInfo"]),

    toLogin() {
      setTimeout(() => {
        this.$router.push("/login");
        window.sessionStorage.clear(); // ​清空sessionStorage中所有信息
      }, 1000);
    },
  },
  mounted() {
    let that = this;
    // 获取用户信息
    // console.log("home config", that.config);
    that.axios
      .post(
        "/userInfo/getInfo",
        {
          userId: sessionStorage.getItem("userId") * 1,
        },
        that.config
      )
      .then((res) => {
        // console.log(" /userInfo/getInfo获取到的信息:", res.data);
        if (res.data.state == "win") {
          console.log("userInfo", res.data.userInfo);
          that.setUserInfo(res.data.userInfo);
        } else {
          that.$refs.msg.setMsg(res.data.commont, "error");
        }
      })
      .catch((error) => {
        console.log(error);
      });

  },
};
</script>
<style lang="less" scoped>
#inspire {
  background: #f0f2f5;
  .v-application .v-application--is-ltr .theme--light.v-tabs-items {
    background-color: transparent !important;
  }
}
nav {
  padding: 0;
}
.container {
  padding: 0px;
}

.listItem {
  .v-list-item {
    padding: 0px;
  }
}
.drawer {
  position: relative;
  top: 0rem;
  left: 0rem;
  // background: rgba(172, 167, 175, 0.5);
}
</style>
