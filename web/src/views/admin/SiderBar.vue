<template>
  <v-app id="inspire2">
    <snackerMsg ref="msg"></snackerMsg>

    <!-- 消息框 -->
    <!-- <v-app-bar color="deep-purple" dark>
      <v-toolbar-title>线上知识竞赛管理后台</v-toolbar-title>
      <template v-slot:extension>
        <v-tabs align-with-title>
          <v-tab to="/admin/alys">数据分析</v-tab>
          <v-tab to="/admin/editHome">首页信息修改</v-tab>
          <v-tab to="/admin/check">认证审核</v-tab>
        </v-tabs>
      </template>
    </v-app-bar> -->

    <v-navigation-drawer
      app
      class="mx-auto"
      data-color="blue-grey darken-4"
      data-theme="dark"
      dark
    >
      <v-list class="">
        <v-list-item>
          <v-list-item-avatar>
            <v-img :src="userInfo.avt"></v-img>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title class="text-h6">
              {{ userInfo.userName }}
            </v-list-item-title>
            <v-list-item-subtitle>{{ userInfo.tel }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>

        <v-divider dark></v-divider>

        <div v-for="(item, i) in siderList" :key="i">
          <v-list-item v-if="item.children == null" link :to="item.path">
            <v-list-item-icon>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-icon>

            <v-list-item-title>{{ item.label }}</v-list-item-title>
          </v-list-item>

          <v-list-group link :prepend-icon="item.icon" v-else>
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title v-text="item.label"></v-list-item-title>
              </v-list-item-content>
            </template>

            <v-list-item
              v-for="child in item.children"
              :key="child.title"
              :to="child.path"
              link
            >
              <v-list-item-content>
                <v-list-item-title v-text="child.label"></v-list-item-title>
              </v-list-item-content>
            </v-list-item>
          </v-list-group>
        </div>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar app  dark>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>后台管理</v-toolbar-title>
    </v-app-bar>

    <v-main class="" app>
      <router-view />
    </v-main>
  </v-app>
</template>
<script>
import { mapState } from "vuex";
import snackerMsg from "../../components/commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  computed: {
    // 用户信息
    ...mapState(["userInfo", "config"]),
  },
  data() {
    return {
      drawer: false,
      siderList: [
        // {
        //   label: "数据分析",
        //   path: "/admin/alys",
        //   icon: "mdi-email",
        //   children: null,
        // },
        {
          label: "信息修改",
          path: "",
          icon: "mdi-email",
          children: [
            {
              label: "轮播图",
              path: "/admin/editImgs",
              icon: "mdi-email",
              children: null,
            },
            {
              label: "标签管理",
              path: "/admin/editTag",
              icon: "mdi-email",
              children: null,
            },
          ],
        },
        {
          label: "认证审核",
          path: "/admin/check",
          icon: "mdi-email",
          children: null,
        },
      ],
    };
  },
  methods: {
    getAdmin() {
      let that = this;
      // console.log("admin config", that.config);
      that.axios
        .post(
          "/userInfo/getInfo",
          {
            userId: sessionStorage.getItem("userId") * 1,
          },
          that.config
        )
        .then((res) => {
          // console.log("获取到的信息:", res.data);
          if (res.data.state == "win") {
            // console.log("userInfo", res.data.userInfo);
            this.$store.state.userInfo.avt = res.data.userInfo.avt;
            this.$store.state.userInfo.tel = res.data.userInfo.tel;
            this.$store.state.userInfo.userName = res.data.userInfo.userName;
            this.$store.state.userInfo.createTime =
              res.data.userInfo.createTime;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    let that = this;
    that.getAdmin();
  },
};
</script>
<style lang="less">
@media screen and (min-width: 1264px) {
  .v-application--is-ltr .v-main:not(.auth-pages) {
    // padding-left: 274px !important;
  }
}
</style>
