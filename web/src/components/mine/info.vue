<template>
  <v-container fluid class="py-6">
    <snackerMsg ref="msg"></snackerMsg>
    <v-row align="center" justify="center">
     
      <v-col lg="5" md="6" cols="12" class="position-relative">
        <v-card class="h-100" elevation="0">
          <div class="px-4 pt-4">
            <h6 class="mb-0 text-h6 text-typo">
              详细资料
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn text @click="editInfoDialog = true">
                    <v-icon color="info" dark v-bind="attrs" v-on="on">
                      mdi-pencil
                    </v-icon></v-btn
                  >
                </template>
                <span>修改信息</span> </v-tooltip
              >|
              <v-tooltip bottom>
                <template v-slot:activator="{ on, attrs }">
                  <v-btn text color="error" @click="editPwdDialog = true"
                    ><v-icon color="info" dark v-bind="attrs" v-on="on">
                      mdi-eye
                    </v-icon></v-btn
                  >
                </template>
                <span>修改密码</span>
              </v-tooltip>
            </h6>
          </div>
          <div class="px-4 py-4">
            <p class="text-sm font-weight-light text-body">
              {{ setIntro(userInfo.intro) }}
            </p>
            <hr class="horizontal dark mt-6 mb-3" />
            <v-list>
              <v-list-item-group class="border-radius-sm">
                <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">用户名:</strong>
                      &nbsp; {{ userInfo.userName }}
                    </div>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">电话号码:</strong>
                      &nbsp; {{ userInfo.tel }}
                    </div>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">Email:</strong>
                      &nbsp; {{ userInfo.email }}
                    </div>
                  </v-list-item-content>
                </v-list-item>
                <!-- <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">是否认证:</strong>
                      &nbsp; {{ userInfo.grade }}
                      <v-icon large color="green darken-2" v-if="userInfo.grade<1"> mdi-domain </v-icon>
                    </div>
                  </v-list-item-content>
                </v-list-item> -->
                <!-- <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">Social:</strong>
                      &nbsp;
                      <v-icon color="#344e86" class="mt-n1 ps-1 pe-2"
                        >fab fa-facebook fa-lg</v-icon
                      >
                      <v-icon color="#3ea1ec" class="mt-n1 ps-1 pe-2"
                        >fab fa-twitter fa-lg</v-icon
                      >
                      <v-icon color="#0e456d" class="mt-n1 ps-1 pe-2"
                        >fab fa-instagram fa-lg</v-icon
                      >
                    </div>
                  </v-list-item-content>
                </v-list-item> -->
              </v-list-item-group>
            </v-list>
          </div>
        </v-card>
        <hr class="vertical dark" />
      </v-col>

    </v-row>


    <v-dialog v-model="editInfoDialog" width="500">
      <v-card>
        <v-card-title class="text-h5 grey lighten-2"> 更改信息 </v-card-title>

        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <uploadImg
                  :uploadtype="uploadtype"
                  :dirPath="dirPath"
                  v-bind:geturl="changeAvt"
                  v-on:puturl="changeAvt = $event"
                ></uploadImg>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="用户名"
                  hint="用户名不超过16个字符"
                  persistent-hint
                  required
                  v-model="userInfo.userName"
                ></v-text-field>
              </v-col>

              <v-col cols="12">
                <v-text-field
                  label="邮箱"
                  hint="用户名不超过16个字符"
                  persistent-hint
                  required
                  v-model="userInfo.email"
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-textarea
                  outlined
                  counter
                  label="个人介绍"
                  :rules="rules"
                  v-model="userInfo.intro"
                ></v-textarea>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="editInfoDialog = false"
            >关闭
          </v-btn>
          <v-btn color="primary" text @click="saveInfo">保存 </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog v-model="editPwdDialog" width="500">
      <v-card>
        <v-card-title class="text-h5 grey lighten-2"> 更改密码 </v-card-title>

        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="初始密码"
                  type="password"
                  required
                  v-model="changePwd.oldPwd"
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="新密码"
                  id="pwd"
                  dense
                  append-icon="mdi-close"
                  prepend-icon="mdi-lock"
                  @click:append="pwdClearR"
                  v-model="changePwd.newPwd"
                  type="password"
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                  label="再次输入密码"
                  type="password"
                  dense
                  append-icon="mdi-close"
                  prepend-icon="mdi-lock"
                  @click:append="rePwdClearR"
                  v-model="changePwd.reNewPwd"
                  @blur="checkRePwd"
                  :error-messages="errorMessages"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="editPwdDialog = false"
            >关闭
          </v-btn>
          <v-btn color="primary" text @click="savePwd">保存 </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
import { mapState, mapGetters, mapMutations } from "vuex";
import uploadImg from "../../components/commont/uploadImg.vue";
import snackerMsg from "../commont/snackerMsg.vue";
import LineChart from "../charts/LineChart";
export default {
  name: "Profile-Overview",
  components: {
    uploadImg,
    snackerMsg,
    LineChart,
  },
  data: function () {
    return {
      rules: [(v) => v.length <= 200 || "请勿超过200字"],
      tagColor: [
        "deep-purple lighten-1",
        "pink darken-1",
        "blue lighten-1",
        "green lighten-1",
        "orange lighten-1",
      ],
      accountSettings: [
        {
          switchState: true,
          text: "Email me when someone follows me",
        },
        {
          switchState: false,
          text: "Email me when someone answers on...",
        },
        {
          switchState: true,
          text: "Email me when someone mentions me...",
        },
      ],
      applicationSettings: [
        {
          switchState: false,
          text: "New launches and projects",
        },
        {
          switchState: true,
          text: "Monthly product updates",
        },
        {
          switchState: false,
          text: "Subscribe to newsletter",
        },
      ],
      conversations: [
        {
          avatar: require("../../assets/哈尔.png"),
          user: "linux操作系统",
          message: ["考试模式", "计算机", "无奖品", "一个月内"],
        },
        {
          avatar: require("../../assets/哈尔.png"),
          user: "Python程序设计",
          message: ["考试模式", "计算机", "无奖品", "一个月内"],
        },
        {
          avatar: require("../../assets/哈尔.png"),
          user: "庆祝党史100周年",
          message: ["限时答题", "历史", "无奖品", "一年内"],
        },
        // {
        //   avatar: require("../../assets/哈尔.png"),
        //   user: "Peterson",
        //   message: "Have a great afternoon..",
        // },
        // {
        //   avatar: require("../../assets/哈尔.png"),
        //   user: "Nick Daniel",
        //   message: "Hi! I need more information..",
        // },
      ],
      cards: [
        {
          avatar: require("../../assets/jidou.jpg"),
          title: "Project #2",
          style: "Modern",
          description:
            "As Uber works through a huge amount of internal management turmoil.",
          avatars: [
            {
              image: require("../../assets/哈尔.png"),
              name: "Elena Morison",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Peterson",
            },
          ],
        },
        {
          image: require("../../assets/哈尔.png"),
          title: "Project #1",
          style: "Scandinavian",
          description:
            "Music is something that every person has his or her own specific opinion about.",
          avatars: [
            {
              image: require("../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Peterson",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Elena Morison",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
          ],
        },
        {
          image: require("../../assets/哈尔.png"),
          title: "Project #3",
          style: "Minimalist",
          description:
            "Different people have different taste, and various types of music.",
          avatars: [
            {
              image: require("../../assets/哈尔.png"),
              name: "Peterson",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Elena Morison",
            },
          ],
        },
        {
          image: require("../../assets/哈尔.png"),
          title: "Project #4",
          style: "Gothic",
          description:
            "Why would anyone pick blue over pink? Pink is obviously a better color.",
          avatars: [
            {
              image: require("../../assets/哈尔.png"),
              name: "Peterson",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Elena Morison",
            },
          ],
        },
      ],

      compeInfo: {},
      // 上传头像所需参数
      uploadtype: "img",
      dirPath: "/avt",
      changeAvt: "", //接收

      // 更改信息
      editInfoDialog: false,
      // 更改密码
      editPwdDialog: false,
      // 注册表单的错误提示
      errorMessages: "",

      changePwd: {
        oldPwd: "",
        newPwd: "",
        reNewPwd: "",
      },
    };
  },
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    ...mapMutations(["setUserInfo"]),
    check() {
      let that = this;
      // console.log("changeAvt:", that.changeAvt);
      that.$refs.msg.setMsg("xxx", "error");
    },
    // 修改信息
    saveInfo() {
      let that = this;
      let obj = {
        userId: sessionStorage.getItem("userId") * 1,
        userName: that.userInfo.userName,
        avt: !that.changeAvt ? that.userInfo.avt : that.changeAvt,
        tel: "",
        intro: that.userInfo.intro,
        email: that.userInfo.email,
      };
      // console.log("修改的信息:", obj);
      that.axios
        .post("/userInfo/editInfo", obj, that.config)
        .then((res) => {
          // console.log("mineCenter获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.editInfoDialog = false;
            that.getUserInfo();
            // that.userInfo = res.data.userInfo;
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
      if (that.changePwd.rePwd != that.changePwd.pwd) {
        that.errorMessages = "密码错误";
        that.disRegis = true;
      } else {
        that.errorMessages = "";
        that.disRegis = false;
      }
    },
    pwdClearR() {
      let that = this;
      that.changePwd.newPwd = "";
    },
    rePwdClearR() {
      let that = this;
      that.changePwd.reNewPwd = "";
    },
    // 修改密码
    savePwd() {
      let that = this;
      // console.log("savePwd", that.changePwd);
      that.axios
        .post(
          "/userInfo/editPwd",
          {
            userId: sessionStorage.getItem("userId") * 1,
            oldPwd: that.changePwd.oldPwd,
            newPwd: that.changePwd.reNewPwd,
          },
          that.config
        )
        .then((res) => {
          // console.log("mineCenter获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            // console.log("newPwd:", that.userInfo.pwd);
            that.editPwdDialog = false;
            // that.userInfo = res.data.userInfo;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    setIntro(intro) {
      if (intro != "") {
        return intro;
      } else {
        return "这个人很懒什么都没有留下......";
      }
    },

    getUserInfo() {
      let that = this;

      that.client = new OSS({
        region: "oss-cn-beijing", //
        accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
        accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
        bucket: "cloud-compe", //OSS 存储区域名
        timeout: "1800000", //30min
      });
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
  },
  mounted() {},
  watch: {},
};
</script>
