<template>
  <v-app id="exam">
    <v-main class="auth-pages">
      <div
        class="header-auth position-relative ma-4 pb-16 pt-16 border-radius-xl min-height-200"
        :style="`background-image:  url('https://images.unsplash.com/photo-1545569341-9eb8b30979d9?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1950&q=80'); background-size: cover; background-position: center;`"
      >
        <span
          class="mask bg-gradient-default border-radius-xl opacity-4"
        ></span>
      </div>
      <fade-transition :duration="200" origin="center top" mode="out-in">
        <router-view ></router-view>
      </fade-transition>
    </v-main>
  </v-app>
</template>
<script>
import { FadeTransition } from "vue2-transitions";

export default {
  name: "page-layout",
  components: {
    FadeTransition,
  },
  data() {
    return {
      paragraphs: "",
      tab: null,
    };
  },
  methods: {},
};
</script>
<style lang="less" scoped>
#exam {
  background: #f0f2f5;
  .v-application .v-application--is-ltr .theme--light.v-tabs-items {
    background-color: transparent !important;
  }
}
</style>
