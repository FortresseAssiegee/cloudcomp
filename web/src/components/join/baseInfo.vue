<template>
  <v-row>
    <snackerMsg ref="msg"></snackerMsg>
    <v-col lg="8" cols="12">
      <v-card>
        <div class="border-bottom d-flex align-center px-4 py-4">
          <div class="d-flex align-center">
            <!-- <a href="javascript:;" class="text-decoration-none">
              <v-avatar size="48" class="rounded-circle">
                <v-img src="../../assets/哈尔.png" alt="profile-image"></v-img>
              </v-avatar>
            </a> -->
            <div class="mx-4">
              <h3
                href="javascript:;"
                class="text-dark font-weight-600 text-decoration-none"
              >
                {{ baseInfo.actvName }}
              </h3>
              <h5 class="d-block text-muted">{{ baseInfo.periodCode }}</h5>
            </div>
          </div>
          <div class="text-end ms-auto">
            <v-btn
              outlined
              color="#fff"
              class="font-weight-bolder bg-gradient-info py-4 px-7"
              small
              @click="outActv"
            >
              取消报名
            </v-btn>
          </div>
        </div>
        <div class="px-4 py-4">
          <p
            class="mb-6 text-body font-weight-light"
            v-html="baseInfo.intro"
          ></p>

          <!-- 参加人数？ joinTable -->
          <v-row class="align-center px-2 mt-6 mb-2">
            <v-col sm="6">
              <div class="d-flex">
                <div class="d-flex align-center">
                  <v-icon
                    size="14"
                    class="material-icons-round me-1 text-body cursor-pointer"
                    >thumb_up</v-icon
                  >
                  <span class="text-sm text-body me-4">150</span>
                </div>
                <div class="d-flex align-center">
                  <v-icon
                    size="14"
                    class="material-icons-round me-1 text-body cursor-pointer"
                    >mode_comment</v-icon
                  >
                  <span class="text-sm text-body me-4">36</span>
                </div>
                <div class="d-flex align-center">
                  <v-icon
                    size="14"
                    class="material-icons-round me-1 text-body cursor-pointer"
                    >forward</v-icon
                  >
                  <span class="text-sm text-body me-4">12</span>
                </div>
              </div>
            </v-col>
            <v-col sm="6">
              <div class="d-flex align-center">
                <div class="d-flex align-items-center ms-auto">
                  <v-avatar size="24" class="border border-white ms-n2">
                    <img src="../../assets/哈尔.png" alt="Avatar" />
                  </v-avatar>
                  <v-avatar size="24" class="border border-white ms-n2">
                    <img src="../../assets/哈尔.png" alt="Avatar" />
                  </v-avatar>
                  <v-avatar size="24" class="border border-white ms-n2">
                    <img src="../../assets/哈尔.png" alt="Avatar" />
                  </v-avatar>
                </div>
                <small class="ps-2 font-weight-bold text-body text-sm"
                  >and 30+ more</small
                >
              </div>
            </v-col>
          </v-row>
          <div class="mb-1">
            <hr class="horizontal dark mt-1 mb-5" />

            <v-card elevation="0">
              <v-card-title>
                <v-icon large left color="cyan" dark flat>
                  mdi-file-multiple
                </v-icon>
                <span class="text-h6 font-weight-light">详情查看</span>
              </v-card-title>
              <v-card-text class="text-left">
                <div id="treeHtml"></div>

                <v-treeview open-all :items="infoTree" class="text-left">
                  <template v-slot:prepend="{ item }">
                    <h6 class="mb-0 text-h6 text-typo font-weight-bold">
                      <a :id="'infoTree' + item.id">{{ item.title }}</a>
                    </h6>
                    <p v-html="item.content" class="mb-0 text-body"></p>
                  </template>
                </v-treeview>
              </v-card-text>
            </v-card>
          </div>
        </div>
      </v-card>
    </v-col>
    <v-col lg="4" cols="12">
      <v-card class="mb-6">
        <v-img :src="baseInfo.pstr" class="border-radius-lg shadow-lg"> </v-img>
      </v-card>
      <v-card class="mb-6">
        <div class="px-4 py-4">
          <v-row class="align-center">
            <v-col cols="9">
              <h5 class="text-h5 font-weight-bold mb-1 text-gradient text-info">
                <a href="javascript:;" class="text-decoration-none">目录</a>
              </h5>
            </v-col>
          </v-row>
          <v-treeview activatable :items="infoTree" class="text-left">
            <template v-slot:prepend="{ item }">
              <a :href="'#infoTree' + item.id">{{ item.title }}</a>
            </template>
          </v-treeview>
        </div>
      </v-card>
      <v-card class="mb-6">
        <div class="px-4 py-4">
          <v-row class="align-center">
            <v-col cols="9">
              <h5 class="text-h5 font-weight-bold mb-1 text-gradient text-info">
                <a href="javascript:;" class="text-decoration-none">时间线</a>
              </h5>
            </v-col>
          </v-row>
          <v-timeline dense align-top class="timeline-line-color pt-0">
            <v-timeline-item
              v-for="(item, i) in timeline"
              :key="item.title + i"
              small
              class="timeline"
            >
              <template v-slot:icon>
                <v-avatar size="30" color="#ffffff">
                  <v-icon
                    :class="`material-icons-round text-white bg-` + item.btn"
                    size="14"
                    >{{ item.icon }}</v-icon
                  >
                </v-avatar>
              </template>
              <v-card elevation="0">
                <v-card-title class="px-0 pt-0 pb-1 d-block">
                  <h6 class="text-dark text-sm font-weight-bold mb-0 d-block">
                    {{ item.title }}
                  </h6>
                  <p class="text-secondary font-weight-light text-xs mb-0">
                    {{ item.time }}
                  </p>
                </v-card-title>
              </v-card>
            </v-timeline-item>
          </v-timeline>
        </div>
      </v-card>
    </v-col>
  </v-row>
</template>
<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import snackerMsg from "../commont/snackerMsg.vue";
export default {
  name: "Teams",

  components: {
    snackerMsg,
  },
  data: function () {
    return {
      baseInfo: {},
      timeline: [],
      infoTree: [],
      joinId: "",
    };
  },
  computed: {
    ...mapState([
      "config",
      "code",
      "boolToString",
      "boolToCode",
      "typeCode",
      "actvId",
      "typeCodeSele",
      "awardCode",
    ]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    ...mapMutations(["setActvBaseInfo"]),
    getInfo() {
      let that = this;
      // console.log("actv,user", that.actvId, that.userInfo.userId);
      that.axios
        .post(
          "/baseInfo/getOne",
          {
            userId: that.userInfo.userId,
            actvId: that.actvId,
          },
          that.config
        )
        .then((res) => {
          // console.log("获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.baseInfo = res.data.baseInfo;
            that.setActvBaseInfo(that.baseInfo);
            that.baseInfo.typeCode = that.code.get(that.baseInfo.typeCode);
            that.baseInfo.awardCode = that.awardCode.get(
              that.baseInfo.awardCode
            );
            that.baseInfo.periodCode = that.code.get(that.baseInfo.periodCode);
            that.setTimeLine(
              { label: "开始", time: that.baseInfo.startTime },
              { label: "报名", time: that.baseInfo.applyStartTime },
              { label: "报名截止", time: that.baseInfo.applyEndTime },
              { label: "考试", time: that.baseInfo.examStartTime },
              { label: "考试截止", time: that.baseInfo.examStartTime },
              { label: "结束", time: that.baseInfo.endTime }
            );
            that.checkJoin();
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    setTimeLine(...timeList) {
      let that = this;
      let nowTime = new Date();
      let nowStr =
        nowTime.getFullYear() +
        "-" +
        (nowTime.getMonth() + 1) +
        "-" +
        nowTime.getDate();

      timeList.push({
        label: "现在",
        time: nowStr,
      });
      timeList = timeList.sort(that.dateData("time", true));
      // console.log("getTimeList", timeList, "now", nowStr);

      for (let i = 0; i < timeList.length; i++) {
        if (timeList[i].label == "现在") {
          let now = {
            title: "现在",
            color: "#b0eed3",
            iconColor: "#1aae6f",
            icon: "ni-bell-55",
            btn: "success",
            time: timeList[i].time,
          };
          that.timeline.push(now);
        } else {
          let b = {
            title: timeList[i].label,
            color: "#fdd1da",
            iconColor: "#f80031",
            icon: "mdi-timer",
            btn: "default",
            time: timeList[i].time,
          };

          that.timeline.push(b);
        }
      }
    },
    dateData(property, bol) {
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
    getInfoTree() {
      let that = this;
      // console.log(
      //   "getInfoTree",
      //   that.userInfo.userId,
      //   that.actvId,
      //   that.config
      // );
      that.axios
        .post(
          "/infoTree/get",
          {
            actvId: that.actvId,
        
          },
          that.config
        )
        .then((res) => {
          // console.log("infoTree获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            if (res.data.infoTree == null) {
              that.infoTree = [];
            } else {
              that.infoTree = res.data.infoTree;
              // console.log("getInfoTree", that.infoTree);
            }
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    outActv() {
      let that = this;
      let d = new Date();
      that.axios
        .post(
          "/joinInfo/out",
          {
            joinId: that.joinId,
          },
          that.config
        )
        .then((res) => {
          // console.log("获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            setTimeout(function () {
              that.$router.push("/user/create/TempAll");
            }, 1500);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    downActv() {
      let that = this;
      that.axios
        .post(
          "/baseInfo/down",
          {
            userId: that.userInfo.userId,
            actvId: that.actvId,
          },
          that.config
        )
        .then((res) => {
          // console.log("down获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            setTimeout(function () {
              that.$router.push("/user/create/TempAll");
            }, 1500);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    checkJoin() {
      let that = this;
      // console.log("checkJoin", {
      //   actvId: that.actvId,
      //   userJoinId: that.userInfo.userId,
      //   userCreateId: that.baseInfo.userId,
      // });
      that.axios
        .post(
          "/joinInfo/checkJoin",
          {
            actvId: that.actvId,
            userJoinId: that.userInfo.userId,
            userCreateId: that.baseInfo.userId,
          },
          that.config
        )
        .then((res) => {
          // console.log("infoTree获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.joinId = res.data.joinId;
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
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
    that.getInfoTree();
  },
};
</script>
<style lang="less" scoped>
a {
  text-decoration: none;
  color: rgba(0, 0, 0, 0.87);
}
</style>
