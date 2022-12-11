<template>
  <v-app id="inspire">
    <!-- 消息框 -->
 

    <v-container style="height: 1000px">
      <router-view />
    </v-container>

    <v-navigation-drawer v-model="drawer" absolute temporary>
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
        <v-list-item-group v-model="selectedItem" color="primary">
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
  </v-app>
</template>

<script>

import snackerMsg from "../../components/commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  data: () => ({
    sideBar: [
      {
        path: "/user/join/TempAll",
        name: "参与的竞赛总览",
        icon: "mdi-clock",
      },
      {
        path: "/user/create/TempAll",
        name: "创建的竞赛总览",
        icon: "mdi-account",
      },
    ],
 
  }),
  computed: {

  },
  methods: {

  },
  mounted() {
   
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
</style>
