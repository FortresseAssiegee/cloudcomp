<template>
  <div>
    <Vcode :show="isShow" @success="success" @close="close" />
    <el-button @click="submit">登录</el-button>
  </div>
</template>

<script>
import Vcode from "vue-puzzle-vcode";
export default {
  data() {
    return {
      isShow: true, // 验证码模态框是否出现
    };
  },
  components: {
    Vcode,
  },
  methods: {
    submit() {
      this.isShow = true;
    },
    // 用户通过了验证
    success(msg) {
      this.isShow = false; // 通过验证后，需要手动隐藏模态框
    },
    // 用户点击遮罩层，应该关闭模态框
    close() {
      this.isShow = false;
    },
  },
};
</script>
