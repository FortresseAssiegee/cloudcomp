<template>
  <v-app>
    <snackerMsg ref="msg"></snackerMsg>
    <!-- <v-btn @click="setHtmlTree()">test</v-btn>
    <div id="test">
      
    </div> -->

    <div class="cover d-flex justify-center align-center">
      <v-card shaped v-if="loginShow" class="log">
        <v-row>
          <v-col cols="6" class="d-flex justify-center align-center">
            <video
              v-if="crsbnListLog.length == 0"
              width="400px"
              autoplay="autoplay"
              loop
              muted="muted"
              src="../../assets/加速开发中.mp4"
            ></video>

            <v-carousel
              cycle
              hide-delimiter-background
              show-arrows-on-hover
              v-else
            >
              <v-carousel-item
                v-for="(slide, i) in crsbnListLog"
                :key="i"
                :src="slide.url"
                width="400px"
              >
              </v-carousel-item>
            </v-carousel>
          </v-col>

          <v-col cols="6">
            <v-card-title>Login · 登录</v-card-title>

            <v-form ref="loginFormRel" style="padding: 0 20px">
              <v-text-field
                label="用户名"
                id="id"
                append-icon="mdi-close"
                prepend-icon="mdi-phone"
                @click:append="phoneClear"
                v-model="loginForm.userName"
              ></v-text-field>
              <v-text-field
                label="密码"
                id="pwd"
                append-icon="mdi-close"
                prepend-icon="mdi-lock"
                @click:append="pwdClear"
                v-model="loginForm.pwd"
                type="password"
              ></v-text-field>

              <v-btn color="success" @click="submit" style="margin-right: 10px"
                >登录</v-btn
              >
              <v-btn color="warning" @click="reset">清空</v-btn>

              <v-card-text class="toRegister"
                >没有账号？<button @click="toRegister">
                  注册
                </button></v-card-text
              >
            </v-form>
          </v-col>
        </v-row>
      </v-card>

      <!-- 注册 -->
      <v-card shaped v-else class="log">
        <v-row>
          <v-col cols="6">
            <v-card-title>Register · 注册</v-card-title>
            <v-form ref="registerFormRel" style="padding: 0 20px">
              <v-text-field
                label="用户名"
                id="userName"
                append-icon="mdi-close"
                prepend-icon="mdi-account"
                @click:append="userNameClear"
                v-model="registerForm.userName"
                :rules="userNameRules"
                dense
              ></v-text-field>

              <v-text-field
                dense
                label="电话"
                id="tel"
                append-icon="mdi-close"
                prepend-icon="mdi-phone"
                @click:append="phoneClearR"
                v-model="registerForm.tel"
              ></v-text-field>
              <v-text-field
                label="密码"
                id="pwd"
                dense
                append-icon="mdi-close"
                prepend-icon="mdi-lock"
                @click:append="pwdClearR"
                v-model="registerForm.pwd"
                type="password"
              ></v-text-field>

              <v-text-field
                label="再次输入密码"
                type="password"
                dense
                append-icon="mdi-close"
                prepend-icon="mdi-lock"
                @click:append="rePwdClearR"
                v-model="registerForm.rePwd"
                @blur="checkRePwd"
                :error-messages="errorMessages"
              ></v-text-field>

              <v-btn
                color="success"
                @click="Register"
                style="margin-right: 10px"
                :disabled="disRegis"
                >注册</v-btn
              >
              <v-btn color="warning" @click="resetR">清空</v-btn>

              <v-card-text
                >已有账号？<button @click="toLogin">登录</button></v-card-text
              >
            </v-form>
          </v-col>
          <v-col cols="6" class="d-flex justify-center align-center">
            <video
              v-if="crsbnListLog.length == 0"
              width="400px"
              autoplay="autoplay"
              loop
              muted="muted"
              src="../../assets/加速开发中.mp4"
            ></video>
            <v-carousel
              cycle
              hide-delimiter-background
              show-arrows-on-hover
              v-else
            >
              <v-carousel-item
                v-for="(slide, i) in crsbnListLog"
                :key="i"
                :src="slide.url"
                width="400px"
              >
              </v-carousel-item>
            </v-carousel>
          </v-col>
        </v-row>
      </v-card>
    </div>
    <Vcode :show="isShow" @success="Login" @close="close" />
  </v-app>
</template>

<script>
import { mapMutations } from "vuex";
import Vcode from "vue-puzzle-vcode";
import snackerMsg from "../../components/commont/snackerMsg.vue";
import uploadImg from "../../components/commont/uploadImg.vue";
import OSS from "ali-oss";
import { getNowDate } from "../../components/diyTools/tools.js";
export default {
  name: "",
  components: {
    snackerMsg,
    uploadImg,
    Vcode,
  },
  data() {
    return {
      isShow: false, // 验证码模态框是否出现
      userNameRules: [
        (v) => !!v || "必填",
        (v) => (v && v.length <= 15) || "不能超过15位",
      ],
      telRules: [
        (v) => !!v || "必填",
        (v) => /.+@.+\..+/.test(v) || "手机号码格式不正确",
      ],
      pwdRules: [
        (v) => !!v || "必填",
        (v) => /.+@.+\..+/.test(v) || "不低于8位并且带字母和数字",
      ],

      // 登录页面显示
      loginShow: true,
      // 登录表单
      loginForm: {},
      // 注册表单
      registerForm: {},
      // 注册表单的错误提示
      errorMessages: "",
      // shop注册按钮
      disRegis: false,

      uploadtype: "img",
      dirPath: "/avt",
      changeAvt: "", //接收

      colors: [
        "indigo",
        "warning",
        "pink darken-2",
        "red lighten-1",
        "deep-purple accent-4",
      ],
      slides: ["海量竞赛，等你参与", "多种模式", "Third", "Fourth", "Fifth"],
      testData: [
        {
          title: "1ewrdwer",
          content: "saaSAS",
          id: "1",
          child: [
            {
              title: "11sadsad",
              content: "11sasadsadaSAS",
              id: "11",
              child: [
                {
                  title: "111ewrdwer",
                  content: "111saaSAS",
                  id: "111",
                  child: [
                    {
                      title: "1111sadsad",
                      content: "1111sasadsadaSAS",
                      id: "1111",
                      child: [],
                    },
                  ],
                },
              ],
            },
          ],
        },
        {
          title: "2ewrdwer",
          content: "saaSAS",
          id: "2",
          child: [
            {
              title: "22sadsad",
              content: "22asadsadaSAS",
              id: "22",
              child: [],
            },
          ],
        },
      ],
      crsbnListLog: [], //轮播图列表
    };
  },
  methods: {
    ...mapMutations(["setConfig", "setUserId"]),
    // --------------页面相关功能函数----------
    // 清空
    reset() {
      this.$refs.loginFormRel.reset();
    },
    resetR() {
      this.$refs.registerFormRel.resetR();
    },

    // close-icon功能
    phoneClear() {
      let that = this;
      that.loginForm.tel = "";
    },
    pwdClear() {
      let that = this;
      that.loginForm.pwd = "";
    },

    // RegisterForm
    userNameClear() {
      let that = this;
      that.registerForm.userName = "";
    },
    phoneClearR() {
      let that = this;
      that.registerForm.tel = "";
    },
    pwdClearR() {
      let that = this;
      that.registerForm.pwd = "";
    },
    rePwdClearR() {
      let that = this;
      that.registerForm.rePwd = "";
    },

    // 显示登录卡片
    toLogin() {
      let that = this;
      that.loginShow = true;
    },
    toRegister() {
      let that = this;
      that.loginShow = false;
    },

    // 登录
    Login(msg) {
      let that = this;
      // console.log("loginForm", that.loginForm);
      that.isShow = true;
      if (msg) {
        // console.log("puzzle msg", msg);
        that.axios
          .post("/login", {
            //"/offline/teacher/login"
            userName: that.loginForm.userName,
            pwd: that.loginForm.pwd + "",
          })
          .then((res) => {
            // console.log("登录:", res);
            if (res.data.state == "win") {
              sessionStorage.setItem("accessToken", res.data.accessToken);
              sessionStorage.setItem("userId", res.data.userId);
              that.setConfig(res.data.accessToken);
              that.setUserId(res.data.userId);
              that.$refs.msg.setMsg(res.data.commont, "success");

              if (res.data.grade < 0) {
                setTimeout(() => {
                  that.$router.push("/admin");
                }, 1000);
              } else {
                that.addLogin();
                setTimeout(() => {
                  that.$router.push("/Home");
                }, 1000);
              }
            } else {
              that.$refs.msg.setMsg(res.data.commont, "error");
            }
          })
          .catch((error) => {
            console.log(error);
          });
        that.msg = null;
      }
    },

    // 注册
    Register() {
      let that = this;
      // console.log("注册:", that.registerForm);

      that.axios
        .post("/sign", {
          userName: that.registerForm.userName,
          tel: that.registerForm.tel,
          pwd: that.registerForm.rePwd,
          avt: that.registerForm.url,
        })
        .then((res) => {
          // console.log("sign:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            sessionStorage.setItem("accessToken", res.data.accessToken);
            sessionStorage.setItem("userId", res.data.userId);
            that.setConfig(res.data.accessToken);
            that.setUserId(res.data.userId);

            that.$router.push("/Home");
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    checkRePwd() {
      let that = this;
      if (that.registerForm.rePwd != that.registerForm.pwd) {
        that.errorMessages = "密码错误";
        that.disRegis = true;
      } else {
        that.errorMessages = "";
        that.disRegis = false;
      }
    },

    submit() {
      this.isShow = true;
    },
    // 用户通过了验证
    success(msg) {
      let that = this;

      // console.log("puzzle msg", msg);
      this.isShow = false; // 通过验证后，需要手动隐藏模态框
    },
    // 用户点击遮罩层，应该关闭模态框
    close() {
      this.isShow = false;
    },
    // // test()
    // setHtmlTree() {
    //   // var el = document.getElementById("treeHtml");
    //   this.createTree(data, 0);
    // },
    // createTree(data, num) {
    //   for (let i = 0; i < data.length; i++) {
    //     console.log(num, data[i]);
    //     if (data[i].child.length != 0) {
    //       this.createTree(data[i].child, num + 1);
    //     }
    //   }
    // },

    async getListLog() {
      let that = this;
      // path 前面不要加"/"
      that.client
        .list({
          prefix: "LogCrsbn",
          // delimiter: "/",
        })
        .then((res) => {
          that.crsbnListLog = [];
          for (let i = 0; i < res.objects.length; i++) {
            const element = res.objects[i];
            element.url = that.client.signatureUrl(element.name);
            that.crsbnListLog.push(res.objects[i]);
          }
          // console.log("Log list res", res);
        })
        .catch((err) => {
          console.log("Log list catch:", err);
        });
    },
    lookup() {
      let that = this;
      let nowDate = getNowDate();
      that.axios
        .post(
          "/dayWeb/addLookup",
          {
            date: nowDate,
          }
        )
        .then((res) => {
          // console.log("down获取到的信息:", res);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    addLogin() {
      let that = this;

      let nowDate = getNowDate();
      that.axios
        .post(
          "/dayWeb/addLogin",
          {
            date: nowDate,
          }
        )
        .then((res) => {
          // console.log("down获取到的信息:", res);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    let that = this;
    that.client = new OSS({
      region: "oss-cn-beijing", //
      accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
      accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
      bucket: "cloud-compe", //OSS 存储区域名
      timeout: "1800000", //30min
    });
    that.getListLog();
    that.lookup();
  },
  watch: {
    changeAvt(val) {
      let that = this;
      that.registerForm.url = val;
      // console.log("获取到的路径:", val);
    },
  },
};
</script>

<style lang="less" scoped>
.cover {
  width: 100vw;
  height: 100vh;
  background-image: url("../../assets/login/bg.png");

  background-repeat: no-repeat;
  background-position: center;
  background-attachment: fixed;
  background-size: 200%;
  -webkit-background-size: 200%;
  -moz-background-size: 200%;
  -o-background-size: 200%;
}

.loginCard {
  width: 60vw;
  margin-left: 20vw;
  margin-top: 10vh;
  overflow: hidden;
}

.registerCard {
  width: 60vw;
  margin-left: 20vw;
  margin-top: 5vh;
  overflow: hidden;
}

.infoTitle {
  margin-top: 20px;
}
</style>
