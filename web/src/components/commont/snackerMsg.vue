<template>
  <v-snackbar
    v-model="snackbarObj.show"
    :timeout="snackbarObj.time"
    top
    shaped
    :color="snackbarObj.type"
  >
    <v-icon>{{ snackbarObj.icon }}</v-icon>
    {{ snackbarObj.text }}
    <template v-slot:action="{ attrs }">
      <v-btn icon v-bind="attrs" @click="snackbarObj.show = false">
        <v-icon>mdi-close</v-icon>
      </v-btn>
    </template>
  </v-snackbar>
</template>
<script>
export default {
  data() {
    return {
      // 消息提醒
      snackbarObj: {
        time: 2000,
        show: false,
        text: "null",
        icon: "", //mdi-alert-outline,mdi-check-circle-outline,mdi-information-outline,mdi-alert-circle-outline
        type: "", //warning,success,info,error
      },
    };
  },
  methods: {
    // 消息提醒框
    setMsg(text, type) {
      let that = this;
      that.snackbarObj.show = true;
      that.snackbarObj.text = text;
      that.snackbarObj.type = type;
      if (type == "warning") {
        that.snackbarObj.icon = "mdi-alert-outline";
      } else if (type == "success") {
        that.snackbarObj.icon = "mdi-check-circle-outline";
      } else if (type == "info") {
        that.snackbarObj.icon = "mdi-information-outline";
      } else if (type == "error") {
        that.snackbarObj.icon = "mdi-alert-circle-outline";
      }
    },
  },
};
</script>
<style lang=""></style>
