<template>
  <!-- 引入的 -->
  <v-container fluid>
    <snackerMsg ref="msg"></snackerMsg>

    <v-row>
      <v-col lg="6">
        <h4 class="text-h4 text-typo font-weight-bold mb-3">修改以下信息</h4>
        <p class="text-body font-weight-light">
          该信息为活动的基础信息，均为必填项
        </p>
      </v-col>
      <v-col lg="6" class="text-end my-auto">
        <v-btn
          elevation="0"
          height="43"
          class="font-weight-bold text-uppercase btn-default py-2 px-6 me-2"
          color="#5e72e4"
          small
          @click="editKeep"
          >保存</v-btn
        >
        <!-- <v-btn
          elevation="0"
          height="43"
          class="font-weight-bold py-2 px-6 me-2"
          small
          @click="getInfo"
          >取消</v-btn
        > -->
        <v-btn
          elevation="0"
          height="43"
          class="font-weight-bold py-2 px-6 me-2"
          small
          @click="pubulish"
          >发布</v-btn
        >
      </v-col>
    </v-row>
    <v-row>
      <v-col md="4" cols="12">
        <v-card
          class="card card-shadow border-radius-xl py-5 text-center mt-6 pa-5"
          data-animation="true"
          elevation="5"
        >
          <div class="mt-n15 mx-4 card-header position-relative z-index-2">
            <div class="d-block blur-shadow-image">
              <v-img
                :src="pstr"
                class="img-fluid shadow border-radius-lg"
                alt=""
              />
            </div>
          </div>
          <div class="mt-n12 mx-5">
            <uploadImg
              v-bind:geturl="pstr"
              v-on:puturl="pstr = $event"
              :tips="tip"
              :dirPath="dirPath"
            ></uploadImg>
            <p class="text-body font-weight-light">
              请上传活动的宣传海报，推荐比例是1/1，上传后记得保存
            </p>
          </div>
          <!-- <div class="mt-6 mx-5">
              <h5 class="text-h5 text-typo font-weight-normal mt-4 mb-2">
                Product Image
              </h5>
              <p class="mb-0 text-body font-weight-light">
                The place is close to Barceloneta Beach and bus stop just 2 min
                by walk and near to "Naviglio" where you can enjoy the main
                night life in Barcelona.
              </p>
            </div> -->
        </v-card>
      </v-col>
      <v-col md="8" cols="12">
        <v-card class="card-shadow border-radius-xl pa-5" elevation="5">
          <div class="card-padding">
            <h5 class="text-h5 text-typo font-weight-bolder">活动信息</h5>
            <v-row class="mt-0">
              <v-col sm="12" cols="12">
                <v-text-field
                  label="活动名称"
                  placeholder="Minimal Bar Stool"
                  color="#e91e63"
                  required
                  v-model="baseInfo.actvName"
                  class="font-size-input input-style"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="mt-0">
              <v-col sm="6">
                <label class="text-sm text-body">比赛模式</label>
                <v-select
                  :items="compeTypeCode"
                  value="请选择比赛模式"
                  color="#e91e63"
                  class="font-size-input input-style placeholder-light border-radius-md select-style mb-0"
                  multi-line
                  height="36"
                  v-model="baseInfo.examCode"
                  item-text="label"
                  item-value="id"
                >
                </v-select>
                <v-switch
                  v-model="baseInfo.awardCode"
                  :label="baseInfo.awardCode ? '有奖品' : '无奖品'"
                ></v-switch>
              </v-col>
              <v-col cols="6">
                <v-tabs
                  v-model="tabs"
                  background-color="deep-purple accent-4"
                  center-active
                  dark
                  next-icon="mdi-arrow-right-bold-box-outline"
                  prev-icon="mdi-arrow-left-bold-box-outline"
                  show-arrows
                >
                  <v-tab v-for="item in tagList" :key="item.tagId">
                    {{ item.label }}
                  </v-tab>
                </v-tabs>
                <v-tabs-items v-model="tabs">
                  <v-tab-item v-for="item in tagList" :key="item.tagId">
                    <div class="pa-4">
                      <v-chip-group
                        active-class="primary--text"
                        column
                        v-model="baseInfo.tagCode"
                      >
                        <v-chip
                          v-for="tag in item.children"
                          :key="tag.tagId"
                          :value="tag.tagId"
                        >
                          {{ tag.label }}
                        </v-chip>
                      </v-chip-group>
                    </div>
                  </v-tab-item>
                </v-tabs-items>
              </v-col>
            </v-row>
            <v-row>
              <v-col sm="12">
                <label class="text-body text-sm font-weight-normal ms-1"
                  >活动简介</label
                >
                <!-- <p class="form-text text-muted text-xs ms-1 d-inline">
                    (optional)
                  </p> -->
                <html-editor v-model="value"></html-editor>
              </v-col>
            </v-row>
          </div>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col lg="4" cols="12">
        <v-card class="card-shadow border-radius-xl h-100 pa-5" elevation="5">
          <div class="card-padding">
            <h5 class="text-h5 text-typo font-weight-bolder">联系方式</h5>
            <v-row>
              <v-col cols="12">
                <v-text-field
                  label="主办方"
                  placeholder="@soft"
                  color="#e91e63"
                  required
                  class="font-size-input input-style"
                  disabled
                  v-model="userInfo.userName"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="mt-0">
              <v-col cols="12" class="py-0">
                <v-text-field
                  label="联系电话"
                  placeholder="请输入电话号码"
                  color="#e91e63"
                  required
                  class="font-size-input input-style"
                  v-model="baseInfo.createUserTel"
                ></v-text-field>
              </v-col>
            </v-row>
            <v-row class="mt-0">
              <v-col cols="12">
                <v-text-field
                  label="联系邮箱"
                  placeholder="请输入邮箱"
                  color="#e91e63"
                  required
                  class="font-size-input input-style"
                  v-model="baseInfo.createUserEmail"
                ></v-text-field>
              </v-col>
            </v-row>
          </div>
        </v-card>
      </v-col>
      <v-col lg="8">
        <v-card class="card-shadow border-radius-xl pa-5" elevation="5">
          <div class="card-padding">
            <h5 class="text-h5 text-typo font-weight-bolder">时间</h5>
            <v-row class="mt-2">
              <v-col cols="6">
                <v-menu
                  v-model="startMenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="baseInfo.startTime"
                      label="发布时间"
                      :hint="timeHint.startTime"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="baseInfo.startTime"
                    locale="zh-cn"
                    :allowed-dates="allowedDates"
                    :events="functionEvents"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="6">
                <v-menu
                  v-model="endMenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="baseInfo.endTime"
                      label="结束时间"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :hint="timeHint.endTime"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="baseInfo.endTime"
                    locale="zh-cn"
                    :allowed-dates="allowedDates"
                    :events="functionEvents"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="6">
                <v-menu
                  v-model="applyStartMenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="baseInfo.applyStartTime"
                      label="报名时间"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :hint="timeHint.applyStartTime"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="baseInfo.applyStartTime"
                    locale="zh-cn"
                    :allowed-dates="applyAllowedDates"
                    :events="functionEvents"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="6">
                <v-menu
                  v-model="applyEndMenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="baseInfo.applyEndTime"
                      label="报名截止"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :hint="timeHint.applyEndTime"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="baseInfo.applyEndTime"
                    @input="applyEndMenu = false"
                    locale="zh-cn"
                    :allowed-dates="applyAllowedDates"
                    :events="functionEvents"
                  ></v-date-picker> </v-menu
              ></v-col>
              <v-col cols="6">
                <v-menu
                  v-model="examStartMenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="baseInfo.examStartTime"
                      label="比赛开始"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :hint="timeHint.examStartTime"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="baseInfo.examStartTime"
                    locale="zh-cn"
                    :allowed-dates="examAllowedDates"
                    :events="functionEvents"
                  ></v-date-picker>
                </v-menu>
              </v-col>
              <v-col cols="6">
                <v-menu
                  v-model="examEndMenu"
                  :close-on-content-click="false"
                  :nudge-right="40"
                  transition="scale-transition"
                  offset-y
                  min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                      v-model="baseInfo.examEndTime"
                      label="比赛截止"
                      prepend-icon="mdi-calendar"
                      readonly
                      v-bind="attrs"
                      v-on="on"
                      :hint="timeHint.examEndTime"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                    v-model="baseInfo.examEndTime"
                    locale="zh-cn"
                    :allowed-dates="examEndAllowedDates"
                    :events="functionEvents"
                  ></v-date-picker>
                </v-menu>
              </v-col>
            </v-row>
            <!-- <v-row class="mt-2">
                <v-col cols="12">
                  <label class="text-sm text-body ms-1">Tags</label>
                  <v-select
                    :items="tags"
                    class="font-size-input input-style mt-0"
                    chips
                    multiple
                    height="46"
                  >
                    <template v-slot:selection="{ item }">
                      <v-chip label color="bg-default" class="py-1 px-2 my-0">
                        <span class="text-white text-caption ls-0">{{
                          item
                        }}</span>
                      </v-chip>
                    </template>
                  </v-select>
                </v-col>
              </v-row> -->
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import snackerMsg from "../commont/snackerMsg.vue";
import uploadImg from "../commont/uploadImg.vue";
import HtmlEditor from "../commont/HtmlEditor.vue";
export default {
  components: {
    snackerMsg,
    uploadImg,
    HtmlEditor,
  },
  data() {
    return {
      e1: 1, //所在那一步
      applyStartMenu: false, //
      applyEndMenu: false,
      examStartMenu: false, //
      examEndMenu: false,
      startMenu: false,
      endMenu: false,
      timeHint: {
        startTime: "发布之后基本信息不可修改",
        endTime: "比赛结束后将不会再被推荐",
        applyStartTime: "请选择发布时间和结束时间",
        applyEndTime: "",
        examStartTime: "",
        examEndTime: "",
      },

      rules: [(v) => v.length <= 25 || "Max 25 characters"],

      pstr: "", //海报

      dirPath: "/pics", //上传的地址
      tip: "上传海报",

      date: new Date(Date.now() - new Date().getTimezoneOffset() * 60000)
        .toISOString()
        .substr(0, 10), //格式

      baseInfo: {
        awardCode: false,
      },

      value: "请输入内容",
      tagList: [],
      tabs: 0,
    };
  },
  computed: {
    ...mapState([
      "config",
      "code",
      "boolToString",
      "boolToCode",
      "actvId",
      "typeCodeSele",
      "awardCode",
      "compeTypeCode",
    ]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    ...mapMutations(["setActvBaseInfo"]),
    pubulish() {
      let that = this;
      var d = new Date();
      let nowTime =
        d.getFullYear() + "-" + (d.getMonth() + 1) + "-" + d.getDate();
      // console.log("发布:", {
      //   userId: that.userInfo.userId,
      //   actvId: that.actvId,
      //   startTime: nowTime,
      // });

      that.axios
        .post(
          "/baseInfo/pubulish",
          {
            userId: that.userInfo.userId,
            actvId: that.actvId,
            startTime: nowTime,
          },
          that.config
        )
        .then((res) => {
          // console.log("pubulish 获取到的信息:", res.data);
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
    // 时间选择器
    allowedDates: (val) => {
      var time = new Date(val);
      return time > Date.now() - 8.64e7;
    },
    applyAllowedDates(val) {
      let that = this;
      if (that.baseInfo.startTime && that.baseInfo.endTime) {
        var time = new Date(val);
        var startTime = new Date(that.baseInfo.startTime);
        var endTime = new Date(that.baseInfo.endTime);

        return time - 8.64e7 < endTime && time > startTime - 8.64e7;
      }
    },
    examAllowedDates(val) {
      let that = this;
      if (that.baseInfo.endTime && that.baseInfo.applyStartTime) {
        var time = new Date(val);
        var startTime = new Date(that.baseInfo.applyStartTime);
        var endTime = new Date(that.baseInfo.endTime);

        return time - 8.64e7 < endTime && time > startTime - 8.64e7;
      }
    },
    examEndAllowedDates(val) {
      let that = this;
      if (that.baseInfo.endTime && that.baseInfo.applyEndTime) {
        var time = new Date(val);
        var startTime = new Date(that.baseInfo.applyEndTime);
        var endTime = new Date(that.baseInfo.endTime);

        return time - 8.64e7 < endTime && time > startTime - 8.64e7;
      }
    },
    // 时间节点的颜色
    functionEvents(date) {
      let that = this;

      if (date == that.baseInfo.startTime || date == that.baseInfo.endTime) {
        return "red";
      } else if (
        date == that.baseInfo.applyStartTime ||
        date == that.baseInfo.applyEndTime
      ) {
        return "yellow";
      } else if (
        date == that.baseInfo.examStartTime ||
        date == that.baseInfo.examEndTime
      ) {
        return "blue";
      }
    },

    // 获取时间码
    getPeriodCode(start, end) {
      var startTime = new Date(start);
      var endTime = new Date(end);
      let days = (endTime - startTime) / 8.64e7;
      // console.log("getPeriodCode", days);

      if (days <= 15) {
        return 41;
      } else if (days > 15 && days <= 31) {
        return 42;
      } else if (days > 31 && days <= 92) {
        return 43;
      } else if (days > 92 && days <= 183) {
        return 44;
      } else if (days > 183 && days <= 366) {
        return 45;
      } else {
        return 46;
      }
    },

    // 接口
    editKeep() {
      let that = this;
      // 修改
  
      that.baseInfo.userId = sessionStorage.getItem("userId") * 1;
      that.baseInfo.pstr = that.pstr;
      that.baseInfo.awardCode = that.baseInfo.awardCode ? 31 : 30;
      that.baseInfo.stateCode = 10;
      that.baseInfo.periodCode = that.getPeriodCode(
        that.baseInfo.startTime,
        that.baseInfo.endTime
      );
      that.baseInfo.intro = that.value;
      // console.log("baseInfo", that.baseInfo);

      that.axios
        .post(
          "/baseInfo/edit",
          {
            userId: sessionStorage.getItem("userId") * 1,
            baseInfo: that.baseInfo,
          },
          that.config
        )
        .then((res) => {
          // console.log("添加新的:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
          that.getInfo();
        })
        .catch((error) => {
          console.log(error);
          that.getInfo();
        });
    },
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
            that.value = that.baseInfo.intro;
            that.setActvBaseInfo(that.baseInfo);
            that.pstr = res.data.baseInfo.pstr;
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getTagList() {
      let that = this;
      that.axios
        .post("/tags/getTagList", {}, that.config)
        .then((res) => {
          // console.log("getTagList 获取到的信息:", res.data);
          if (res.data.state == "win") {
            // console.log("tagList", res.data.TagInfo);
            that.tagList = res.data.TagInfo;
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
    that.getInfo();
    that.getTagList();
  },
  watch: {
    pstr(val) {
      // console.log("pstr:", val);
    },
  },
};
</script>

<style lang="less" scoped>
a {
  text-decoration: none;
  color: rgba(0, 0, 0, 0.87);
}
.theme--light.v-tabs-items {
  background-color: transparent;
}
</style>
