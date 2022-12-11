<template>
  <v-navigation-drawer
    v-model="getdrawer"
    absolute
    temporary
    dark
    data-color="pink lighten-2"
    data-theme="dark"
  >
    <v-list-item class="px-2">
      <v-list-item-avatar>
        <v-img :src="userInfo.avt"></v-img>
      </v-list-item-avatar>
      <v-list-item-title>{{ userInfo.userName }}</v-list-item-title>

      <v-btn icon @click="toMineCenter">
        <v-icon>mdi-home</v-icon>
      </v-btn>
    </v-list-item>
    <!-- <v-list-item class="px-2">
          <v-list-item-title>2892296496@qq.com</v-list-item-title>
        </v-list-item> -->

    <v-divider></v-divider>

    <v-list>
      <v-subheader>列表</v-subheader>
      <v-list-item-group active-class="pink--text">
        <v-list-item v-for="(item, i) in sideBar" :key="i" :to="item.path">
          <v-list-item-icon>
            <v-icon v-text="item.icon"></v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ item.name }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </v-navigation-drawer>
</template>
<script>
import { mapState, mapGetters } from "vuex";
export default {
  name: "drawer",
  props: {
    getdrawer: {
      type: Boolean,
      default: null,
    },
  },
  data() {
    return {
      sideBar: [
        {
          path: "/home",
          name: "首页",
          icon: "mdi-home",
        },
        {
          path: "/user/join/TempAll",
          name: "加入的竞赛",
          icon: "mdi-clock",
        },
        {
          path: "/user/create/TempAll",
          name: "创建的竞赛",
          icon: "mdi-account",
        },
      ],
    };
  },
  computed: {
    // 用户信息
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    toMineCenter() {
      let that = this;
      that.$router.push("/user/mineCenter");
    },
  },

  mounted() {
    // console.log("userInfo", this.userInfo);
  },
  watch: {
    getdrawer(val) {
      let that = this;
      // console.log("getdrawer", that.getdrawer);
      that.$emit("putdrawer", val);
    },
  },
};
</script>
<style lang=""></style>
