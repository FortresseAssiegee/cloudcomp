<template>
  <v-app id="inspire">
    <!-- 消息框 -->

    <snackerMsg ref="msg"></snackerMsg>

    <drawer
      v-bind:getdrawer="drawer"
      v-on:putdrawer="drawer = $event"
      app
    ></drawer>
    <!-- <v-app-bar color="transparent" class="elevation-0 " app>
    </v-app-bar> -->

    <v-container fluid class="px-15">
      <v-toolbar color="transparent" class="elevation-0">
        <v-app-bar-nav-icon @click="drawer = true"></v-app-bar-nav-icon>
        <v-toolbar-title>{{ title }}</v-toolbar-title>
      </v-toolbar>
      <router-view />
    </v-container>
  </v-app>
</template>

<script>
import { mapMutations } from "vuex";
import { mapState } from "vuex";
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
    title: window.sessionStorage.getItem("title"),
  }),
  computed: {
    // 用户信息
    ...mapState(["userInfo", "config"]),
  },
  methods: {
    ...mapMutations(["setConfig", "setUserInfo"]),
    // 获取用户信息
    getInfo() {
      let that = this;
      // console.log("user config", that.config);
      that.axios
        .post(
          "/userInfo/getInfo",
          {
            userId: sessionStorage.getItem("userId") * 1,
          },
          that.config
        )
        .then((res) => {
          // console.log("mine getInfo获取到的信息:", res.data);
          if (res.data.state == "win") {
            // console.log("userInfo", res.data.userInfo);
            that.setUserInfo(res.data.userInfo);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    // 获取创建的列表
    getBaseInfo() {
      let that = this;
      that.axios
        .post(
          "/baseInfo/getList",
          {
            userId: sessionStorage.getItem("userId") * 1,
          },
          that.config
        )
        .then((res) => {
          // console.log("createDoingList获取到的信息:", res);
          if (res.data.state == "win") {
            that.createDoingList = res.data.baseInfo.filter((item) => {
              return item.stateCode == 10;
            });
            that.createDoneList = res.data.baseInfo.filter((item) => {
              return item.stateCode == 11;
            });
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    let that = this;
    that.getInfo();
  },
  watch: {

  },
};
</script>
<style lang="less" scoped>
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
}
#inspire {
  background: #f0f2f5;
  .v-application .v-application--is-ltr .theme--light.v-tabs-items {
    background-color: transparent !important;
  }
}
</style>
