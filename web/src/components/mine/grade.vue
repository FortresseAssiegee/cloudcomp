<template>
  <v-container fluid class="py-6 px-0">
    <snackerMsg ref="msg"></snackerMsg>

    <v-row>
      <v-col cols="8">
        <v-card class="pt-8 px-1 shadow-blur">
          <div class="mt-n4">
            <div
              class="bg-gradient-info shadow-info border-radius-lg px-3 py-2 mx-3"
            >
              <v-container>
                <v-row>
                  <v-col md="10">
                    <div class="d-flex align-items-center">
                      <!-- <v-avatar size="48">
                        <img src="../../assets/jidou.jpg" alt="Avatar" />
                      </v-avatar> -->
                      <div class="ms-3">
                        <h6 class="mb-0 text-h6 d-block text-white">
                          认证申请
                        </h6>
                        <span class="text-sm text-white opacity-8"
                          >认证只适用于企业或者学校</span
                        >
                      </div>
                    </div>
                  </v-col>
                </v-row>
              </v-container>
            </div>
          </div>
          <div class="max-vh-50 overflow-scroll pt-6 px-4">
            <div v-for="chat in chatsList" :key="chat.gradeId">
              <v-row
                class="justify-end text-end mb-3"
                v-show="chat.startTime != ''"
              >
                <v-col cols="auto">
                  <v-card class="py-2 px-3 shadow border-radius-xl">
                    <p
                      class="mb-1 text-body font-weight-light"
                      v-if="chat.applyIntro"
                    >
                      {{ chat.applyIntro }}
                    </p>

                    <div
                      class="d-flex align-center text-sm text-body opacity-6"
                    >
                      <i class="ni ni-check-bold text-sm me-1"></i>
                      <small>{{ chat.startTime }}</small>
                    </div>
                  </v-card>
                </v-col>
              </v-row>
              <v-row class="justify-start mb-3" v-if="chat.endTime != ''">
                <v-col cols="auto">
                  <v-card
                    class="py-2 px-3 shadow bg-gradient-info border-radius-xl"
                  >
                    <p
                      class="mb-1 font-weight-light text-white"
                      v-if="chat.replyIntro"
                    >
                      {{ chat.replyIntro }}
                    </p>
                    <div
                      class="d-flex align-center justify-end text-sm text-white opacity-6"
                      v-if="chat.endTime"
                    >
                      <i class="ni ni-check-bold text-sm me-1"></i>
                      <small>{{ chat.endTime }}</small>
                    </div>
                  </v-card>
                </v-col>
              </v-row>
              <v-row class="justify-center mb-3" v-else>
                <v-col cols="auto">
                  <v-card
                    class="py-2 px-3 shadow bg-gradient-info border-radius-xl"
                  >
                    <div
                      class="d-flex align-center justify-end text-sm text-white opacity-6"
                    >
                      <small>还在审核中,请耐心等待...</small>
                    </div>
                  </v-card>
                </v-col>
              </v-row>
            </div>
          </div>

          <div class="py-4 px-4">
            <v-form class="d-flex w-100" id="navbar-search-main">
              <v-text-field
                rounded-sm
                outlined
                dense
                hide-details
                background-color="transparent"
                color="#e91e63"
                label="请输入申请理由或者原因"
                class="font-size-input input-style"
                :disabled="!sendFlag"
                v-model="applyIntro"
              >
              </v-text-field>
              <v-btn
                elevation="0"
                class="bg-gradient-info ms-2"
                height="40"
                @click="applyGrade"
                :disabled="!sendFlag"
              >
                <v-icon class="material-icons-round text-white" v-if="sendFlag"
                  >mdi-send</v-icon
                >
                <v-icon class="material-icons-round text-white" v-else
                  >mdi-send-lock</v-icon
                >
              </v-btn>
            </v-form>
          </div>
        </v-card>
      </v-col>
      <v-col cols="4">
        <v-card class="pt-8 px-3 shadow-blur">
          <h6 class="text-h6 text-typo mb-2">认证资料</h6>
          <v-form
            @submit.prevent="submit"
            class="navbar-search navbar-search-light d-inline-block w-100 mb-4"
            id="navbar-search-main"
          >
            <v-divider class="my-3" v-if="!showPro"></v-divider>
            <v-progress-linear
              color="light-blue"
              :value="progress"
              v-else
              class="my-3"
            ></v-progress-linear>

            <v-row>
              <v-col cols="10">
                <v-file-input
                  accept="zip/*"
                  label="上传认证资料"
                  v-model="files"
                  outlined
                  dense
                  multiple
                  rounded-sm
                  hide-details
                  background-color="transparent"
                  color="info"
                  class="font-size-input input-style"
                  @change="upload('com')"
                  counter="1"
                >
                  <template v-slot:selection="{ text }">
                    <v-chip small label color="info">
                      {{ text }}
                    </v-chip>
                  </template>
                </v-file-input></v-col
              >
              <v-col cols="2">
                <v-icon @click="uploadCanncel" dark class="mt-3"
                  >mdi-cancel</v-icon
                >
              </v-col>
            </v-row>

            <v-list rounded>
              <v-subheader>只能上传zip或者rarg格式文件</v-subheader>
              <v-list-item-group color="info">
                <v-list-item v-show="getName != ''">
                  <v-list-item-icon>
                    <v-icon>mdi-zip-box-outline </v-icon>
                  </v-list-item-icon>
                  <v-list-item-content>
                    <v-list-item-title>{{ getName }}</v-list-item-title>
                  </v-list-item-content>

                  <v-list-item-action>
                    <a :href="signUrl">
                      <v-icon color="yellow darken-3">mdi-download</v-icon>
                    </a>
                  </v-list-item-action>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import OSS from "ali-oss";
import { mapState, mapGetters } from "vuex";
import snackerMsg from "../commont/snackerMsg.vue";
import { getNowDate } from "../diyTools/tools.js";
export default {
  name: "Teams",
  components: {
    snackerMsg,
  },
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  data: function () {
    return {
      chatsList: [],

      applyIntro: "",
      sendFlag: false,
      files: [],
      filesMine: [],
      showPro: false,
      progress: 0,
      getName: "", //文件名字
      getNameImg: "", //文件名字
      signUrl: "", //下载地址
      signUrlImg: "", //下载地址

      sure: false,
    };
  },
  methods: {
    getGradeList() {
      let that = this;
      // console.log(
      //   "getGradeList userId:",
      //   sessionStorage.getItem("userId") * 1,
      //   "\nconfig",
      //   that.config
      // );
      that.axios
        .post(
          "/grade/GetOneGrade",
          {
            userId: sessionStorage.getItem("userId") * 1,
          },
          that.config
        )
        .then((res) => {
          // console.log("getGradeList res:", res.data);
          if (res.data.state == "win") {
            if (res.data.gradeInfo.length == 0) {
              that.sendFlag = true;
            }

            that.chatsList = res.data.gradeInfo;
            that.chatsList = that.chatsList.sort(
              that.compare("startTime", true)
            );
            let num = that.chatsList.length - 1; //endTime

            if (that.chatsList[num].endTime) {
              // console.log(
              //   "that.chatsList[num].endTime",
              //   that.chatsList[num].endTime
              // );
              that.sendFlag = true;
            } else {
              that.sendFlag = false;
            }
            // console.log("after sort() chatsList:", that.chatsList);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    addZero(s) {
      return s < 10 ? "0" + s : s;
    },

    getNowTime() {
      let that = this;
      let d = new Date();
      let nowTime =
        d.getFullYear() +
        "-" +
        (d.getMonth() + 1) +
        "-" +
        d.getDate() +
        " " +
        that.addZero(d.getHours()) +
        ":" +
        that.addZero(d.getMinutes()) +
        ":" +
        that.addZero(d.getSeconds());
      return nowTime;
    },
    applyGrade() {
      let that = this;

      // console.log("applyGrade:", {
      //   userId: sessionStorage.getItem("userId") * 1,
      //   applyIntro: that.applyIntro,
      //   startTime: that.getNowTime(),
      // });
      that.axios
        .post(
          "/grade/applyGrade",
          {
            userId: sessionStorage.getItem("userId") * 1,
            applyIntro: that.applyIntro,
            startTime: that.getNowTime(),
          },
          that.config
        )
        .then((res) => {
          // console.log("applyGrade res:", res.data);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.applyIntro = "";
            that.getGradeList();
            that.addGrade();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    compare(property, bol) {
      //property是你需要排序传入的key,bol为true时是升序，false为降序
      return function (a, b) {
        var value1 = a[property];
        var value2 = b[property];
        if (bol) {
          // 升序
          return Date.parse(value1) - Date.parse(value2);
        } else {
          // 降序
          return Date.parse(value2) - Date.parse(value1);
        }
      };
    },

    // oss

    uploadCanncel() {
      let that = this;
      that.client.cancel();
      that.showPro = false;
      that.$refs.msg.setMsg("取消上传", "success");
    },
    // 上传图片
    upload(flag) {
      let that = this;
      //   判断类型进行分流
      // console.log("上传的ppt文件:", that.files[0], that.files[0].name);
      let file = that.files[0];
      if (flag == "self") {
        if (
          !/^\S+\.png$/.test(file.name) &&
          !/^\S+\.jpe{0,1}g$/.test(file.name)
        ) {
          return that.$refs.msg.setMsg("请上传png,jpeg,jpg格式", "error");
        }
      } else {
        if (!/^\S+\.zip$/.test(file.name) && !/^\S+\.rar$/.test(file.name)) {
          return that.$refs.msg.setMsg("请上传zip或rar格式", "error");
        }
      }

      that.showPro = true;
      that.$refs.msg.setMsg("开始上传", "success");

      let param = new FormData();
      param.append("file", file, file.name);
      that.put(file.name, file);
    },

    async put(name, file) {
      let that = this;
      let result = "";
      var dirName = "";
      // console.log("that.getName", that.getName);
      try {
        if (that.getName) {
          result = await that.client.delete(that.getName);
          // console.log("del getName", result);
        }
        dirName = "/grade/" + that.userInfo.userId + "/" + name;

        // 填写Object完整路径。Object完整路径中不能包含Bucket名称。
      } catch (error) {
        console.log("del getName", error);
      }

      try {
        if (!that.userInfo.userId) {
          that.$refs.msg.setMsg("上传前请登录", "error");
          that.showPro = false;

          return;
        }

        // console.log("dirName", dirName);
        //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。

        // 指定待删除的Object的名称。Object名称需填写不包含Bucket名称在内的Object的完整路径，例如exampledir/exampleobject.txt。
        await that.client
          .multipartUpload(dirName, file, {
            ...that.option,
          })
          .then((result) => {
            // console.log("multipartUpload res", result);

            that.progress = 0;
            that.files = null;
            that.showPro = false;
            that.getList();
          })
          .catch((err) => {
            console.log("multipartUpload catch:", err);
          });
      } catch (e) {
        console.log(e);
        // that.$refs.msg.setMsg("上传超时", "error");
      }
    },

    // oss
    async delOne(name) {
      let that = this;
      try {
        // 填写Object完整路径。Object完整路径中不能包含Bucket名称。
        let result = await that.client.delete(name);
        console.log("delOne getName", result);
      } catch (error) {
        console.log("delOne getName", error);
      }
    },
    async getList() {
      let that = this;
      let path = "";

      path = "grade/" + that.userInfo.userId;

      // console.log("getList path", path);
      // path 前面不要加"/"
      that.client
        .list({
          prefix: path,
          //  delimiter: '/'
        })
        .then((res) => {
          // console.log("getList", res.objects[0]);

          that.getName = res.objects[0].name;
          that.signUrl = that.client.signatureUrl(that.getName);
          // console.log("getName:", that.getName, "\nsignUrl:", that.signUrl);
        })
        .catch((err) => {
          console.log("list catch:", err);
        });
    },

    addGrade() {
      let that = this;
      let nowDate = getNowDate();
      that.axios
        .post("/dayWeb/addApplyGrade", {
          date: nowDate,
        })
        .then((res) => {
          console.log("down获取到的信息:", res);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    let that = this;
    that.getGradeList();
    // 连接oss
    that.client = new OSS({
      region: "oss-cn-beijing", //
      accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
      accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
      bucket: "cloud-compe", //OSS 存储区域名
      timeout: "1800000", //30min
    });

    that.option = {
      // 获取分片上传进度、断点和返回值。
      progress: (p, cpt, res) => {
        that.progress = p * 100;
        // console.log("option progress:", p, cpt, res);
      },
      // 设置并发上传的分片数量。
      parallel: 20,
      // 设置分片大小。默认值为1 MB，最小值为100 KB。
      partSize: 1024 * 1024,
      // headers,
      // 自定义元数据，通过HeadObject接口可以获取Object的元数据。
      meta: { year: 2020, people: "test" },
      mime: ["application/zip", "application/rar"],
    };

    that.getList();
  },
  watch: {
    progress(val) {
      let that = this;
      console.log("pro", val);
      if (val * 1 == 100) {
        that.showPro = false;
      }
    },
  },
};
</script>
<style lang="less" scoped>
a {
  text-decoration: none;
}
</style>
