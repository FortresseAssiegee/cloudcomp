<template>
  <v-container fluid>
    <snackerMsg ref="msg"></snackerMsg>
    <v-row>
      <v-col cols="4">
        <v-card
          class="card-shadow border-radius-xl px-10 py-4 position-sticky top-1"
          elevation="5"
        >
          <v-card-title class="text-h5">
            目录
            <v-spacer></v-spacer>
          </v-card-title>
          <v-divider></v-divider>
          <v-treeview
            :items="unitTree"
            :active.sync="active"
            activatable
            color="warning"
            open-on-click
            transition
            item-children="child"
            item-text="title"
            class="text-left"
            :v-if="unitTree.length > 0"
          >
            <template v-slot:prepend="{ item, open }">
              <v-icon v-if="item.model == 72"> mdi-script-text </v-icon>
              <v-icon v-else-if="item.model == 71"> mdi-school </v-icon>
            </template>
          </v-treeview>
        </v-card>
      </v-col>
      <v-col cols="8">
        <v-row class="my-10" v-show="active != ''">
          <!-- xxx -->
          <v-col cols="12">
            <v-card class="mb-6 card-shadow border-radius-xl py-4">
              <v-row no-gutters class="px-4">
                <v-col sm="2">
                  <v-avatar
                    color="bg-gradient-info shadow border-radius-xl mt-n8"
                    class="shadow-dark"
                    height="64"
                    width="64"
                  >
                    <v-icon class="material-icons-round text-white" size="24"
                      >mdi-home</v-icon
                    >
                  </v-avatar>
                </v-col>
                <v-col sm="8" class="text-end">
                  <p
                    class="text-sm mb-0 text-capitalize text-body font-weight-light"
                  >
                    Today's Users
                  </p>
                  <h4 class="text-h4 text-typo font-weight-bolder mb-0">
                    {{ node.title }}
                  </h4>
                </v-col>
              </v-row>
              <hr class="dark horizontal mt-3 mb-4" />
              <v-row class="px-4" v-if="!(node.model == 71)">
                <v-col cols="6">
                  <v-list-item
                    class="px-0 py-3 bg-gray-100 border-radius-lg p-6 mb-6"
                  >
                    <v-list-item-content class="px-6">
                      <div class="d-flex flex-column">
                        <h6 class="mb-3 text-sm text-typo font-weight-bold">
                          比赛信息
                        </h6>

                        <span class="mb-2 text-xs text-body"
                          >比赛时间:
                          <span class="text-dark font-weight-bold ms-sm-2"
                            >{{ node.startTime }}到{{ node.endTime }}</span
                          ></span
                        >
                        <span class="mb-2 text-xs text-body"
                          >考试模式：
                          <span class="text-dark ms-sm-2 font-weight-bold">{{
                            code.get(node.examCode)
                          }}</span></span
                        >
                        <span class="mb-2 text-xs text-body"
                          >比赛时长:
                          <span class="text-dark ms-sm-2 font-weight-bold"
                            >{{ node.during }}分钟</span
                          ></span
                        >
                        <span class="mb-2 text-xs text-body"
                          >是否开启人脸识别:
                          <span class="text-dark ms-sm-2 font-weight-bold">{{
                            node.face ? "是" : "否"
                          }}</span></span
                        >
                      </div>
                    </v-list-item-content>
                  </v-list-item>
                </v-col>
                <v-col cols="6" v-if="!etf">
                  <v-list-item
                    class="px-0 py-3 bg-gray-100 border-radius-lg p-6 mb-6"
                  >
                    <v-list-item-content class="px-6">
                      <div class="d-flex flex-column">
                        <h6 class="mb-3 text-sm text-typo font-weight-bold">
                          比赛成绩
                        </h6>

                        <span class="mb-2 text-xs text-body"
                          >比赛时间:
                          <span class="text-dark font-weight-bold ms-sm-2"
                            >{{ unitJoinInfo.startTime }}到{{
                              unitJoinInfo.endTime
                            }}</span
                          ></span
                        >
                        <span class="mb-2 text-xs text-body"
                          >分数：
                          <span class="text-dark ms-sm-2 font-weight-bold">{{
                            unitJoinInfo.score
                          }}</span></span
                        >
                        <span class="mb-2 text-xs text-body"
                          >排名:
                          <span class="text-dark ms-sm-2 font-weight-bold">{{
                            unitJoinInfo.rank
                          }}</span></span
                        >
                      </div>
                    </v-list-item-content>
                  </v-list-item>
                </v-col>
                <v-col cols="12"
                  ><v-btn
                    :disabled="!etf"
                    @click="toExam"
                    color="deep-purple lighten-2"
                    text
                    justify="end"
                    >进入考场</v-btn
                  ></v-col
                >
              </v-row>
            </v-card>
          </v-col>

          <v-col cols="6" v-if="!(node.model == 72)">
            <uploadMultiPart
              :fileName="fileName"
              :type="umpZip"
              :miss="true"
            ></uploadMultiPart>
          </v-col>

          <v-col cols="12" v-if="!(node.model == 72)">
            <uploadMultiPart
              :type="umpVideo"
              :fileName="fileName"
              :miss="true"
            ></uploadMultiPart>
          </v-col>
          <v-col cols="12" v-if="!(node.model == 72)">
            <uploadMultiPart
              :fileName="fileName"
              :miss="true"
            ></uploadMultiPart>
          </v-col>
        </v-row>
      </v-col>
    </v-row>

    <!-- 添加父节点 -->
    <v-dialog
      transition="dialog-top-transition"
      max-width="500"
      v-model="addDialog"
    >
      <v-card class="px-4 py-4">
        <v-card-title>添加 </v-card-title>
        <v-card-text>
          <v-form>
            <v-text-field
              v-model="addNodeTitle"
              counter="25"
              hint="请输入添加的标题"
              label="标题"
            ></v-text-field>
          </v-form>
        </v-card-text>
        <v-card-actions class="justify-end">
          <v-btn text @click="addNode(editId)">确定</v-btn>
          <v-btn text @click="addDialog = false">取消</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- 添加子节点 -->
    <v-dialog
      transition="dialog-top-transition"
      max-width="500"
      v-model="delDialog"
    >
      <v-card>
        <v-card-title>删除</v-card-title>

        <v-card-text> 确定删除吗？</v-card-text>

        <v-card-actions class="justify-end">
          <v-btn text @click="delNode(editId)">确定</v-btn>
          <v-btn text @click="delDialog = false">取消</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import HtmlEditor from "../commont/HtmlEditor.vue";
import snackerMsg from "../commont/snackerMsg.vue";
import uploadMultiPart from "../commont/uploadMultiPart";

export default {
  components: {
    HtmlEditor,
    snackerMsg,
    uploadMultiPart,
  },
  data() {
    return {
      active: [],
      unitTree: [],

      addNodeTitle: "",
      delNodeTitle: "",

      addDialog: false, //父节点
      delDialog: false,

      parentId: "",
      editId: null,

      // 查看
      title: "",
      value: "",
      node: {},

      pptUrl: "",
      videoUrl: "",
      otherUrl: "",

      umpVideo: "video", //uploadMultiPart
      umpZip: "other",

      fileName: "",
      maxExam: 10,
      unitJoinInfo: {},
      etf: true,
    };
  },
  computed: {
    ...mapState([
      "config",
      "code",

      "boolToStringFace",
      "boolToCode",
      "typeCode",
      "actvId",
      "typeCodeSele",
      "awardCode",
      // "compeTypeCode",
      "unitTypeCode",
    ]),
    ...mapGetters(["userInfo", "actvBaseInfo"]),
  },
  methods: {
    // setUnitInfo
    ...mapMutations(["setUnitInfo", "setExamInfo"]),

    getUnitTree() {
      let that = this;
      // console.log(
      //   "getUnitTree",
      //   that.userInfo.userId,
      //   that.actvId,
      //   that.config
      // );
      that.axios
        .post(
          "/unitTree/get",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
          },
          that.config
        )
        .then((res) => {
          // console.log("infoTree获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            if (res.data.unitTree == null) {
              that.unitTree = [];
            } else {
              that.unitTree = res.data.unitTree;
            }
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    // getOne
    getOne() {
      let that = this;
      // console.log("getOne:", {
      //   actvId: that.actvId,
      //   userId: that.userInfo.userId,
      //   unitId: that.active[0],
      // });
      that.axios
        .post(
          "/unitTree/getOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            unitId: that.active[0],
          },
          that.config
        )
        .then((res) => {
          // console.log("getOne获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            // that.value = res.data.node.tips;
            // that.title = res.data.node.title;
            that.node = res.data.node;
            if (that.actvBaseInfo.examCode * 1 != 50) {
              that.node.examCode = that.actvBaseInfo.examCode;
              // console.log("node.examCode", that.node.examCode);
            }
            that.maxExam = res.data.maxExam;
            that.fileName = res.data.node.fileName;
            that.checkExam();
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    toExam() {
      let that = this;
      let examInfo = {
        // 活动
        actvId: that.actvBaseInfo.actvId,
        actvName: that.actvBaseInfo.actvName,
        userCreateId: that.actvBaseInfo.userId,
        userCreateName: that.actvBaseInfo.userCreateName,
        // 单元信息
        unitId: that.active[0],
        unitName: that.node.title,
        examCode: that.node.examCode,
        face: that.node.face,
        during: that.node.during,
        unitCount: that.count,
        startTime: that.node.startTime,
        endTime: that.node.endTime,

        // 考试信息
        userJoinId: that.userInfo.userId,
      };
      // console.log("toExam:", examInfo);
      that.setUnitInfo(examInfo);
      that.$router.push("/exam");
    },
    checkExam() {
      let that = this;
      that.axios
        .post(
          "/joinUnitInfo/checkExam",
          {
            actvId: that.actvBaseInfo.actvId,
            unitId: that.active[0],
            userCreateId: that.actvBaseInfo.userId,
            userJoinId: that.userInfo.userId,
          },
          that.config
        )
        .then((res) => {
          // console.log("checkExam:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.unitJoinInfo = res.data.joinUnitInfo;
            if (that.unitJoinInfo.startTime == "") {
              that.etf = true;
              sessionStorage.setItem(`etf`, true); //没有考试
            } else {
              that.etf = false;
              sessionStorage.setItem(`etf`, false);
            }
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
    that.getUnitTree();
    // console.log("actvBaseInfo", that.actvBaseInfo);
  },
  watch: {
    active() {
      let that = this;
      // console.log("active改变:", that.active);
      that.userInfo.unitId = that.active[0];
      that.userInfo.actvId = that.actvId;
      // 写一个getone的接口
      that.getOne();
      // that.fileName =
      //   that.userInfo.userId.toString() +
      //   that.actvId.toString() +
      //   that.active.toString();
      // console.log("fileName:", that.fileName);

      that.getOne(that.node);
    },
  },
};
</script>
<style lang="less" scoped></style>
