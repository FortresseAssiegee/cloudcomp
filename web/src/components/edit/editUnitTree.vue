<template>
  <v-container fluid>
    <snackerMsg ref="msg"></snackerMsg>
    <v-row>
      <v-col lg="6">
        <h4 class="text-h4 text-typo font-weight-bold mb-3">修改以下信息</h4>
        <p class="text-body font-weight-light">
          单元的类型分为：
          <a class="text-decoration-none">学习模式</a>将可以上传学习资料，<a
            class="text-decoration-none"
            >考试模式</a
          >可以编辑题库。
        </p>
      </v-col>
      <v-col cols="6" class="text-end my-auto">
        <v-btn
          elevation="0"
          height="43"
          class="font-weight-bold text-uppercase btn-default py-2 px-6 me-2"
          color="#5e72e4"
          small
          @click="editOne"
          >保存
        </v-btn>
        <v-btn
          v-show="!(node.model == 71)"
          elevation="0"
          height="43"
          class="font-weight-bold py-2 px-6 me-2"
          small
          @click="toQuest"
        >
          <v-icon>mdi-import</v-icon>
          题库管理
        </v-btn>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="4">
        <v-card class="p card-shadow border-radius-xl px-10 py-4" elevation="5">
          <v-card-title class="text-h5">
            目录
            <v-spacer></v-spacer>
            <v-btn icon outlined @click="showAddNode(0)">
              <v-icon dark> mdi-plus </v-icon></v-btn
            ></v-card-title
          >
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
            <template v-slot:append="{ item, open }">
              <!-- <v-icon v-if="!item.file">
                  {{ open ? "mdi-folder-open" : "mdi-folder" }}
                </v-icon>
                <v-icon v-else>
                  {{ files[item.file] }}
                </v-icon> -->
              <v-btn icon @click="showDelNode(item.id)">
                <v-icon>mdi-delete</v-icon>
              </v-btn>
              <!-- v-if="item.child.length > 0" -->
              <v-btn icon @click="showAddNode(item.id)">
                <v-icon>mdi-plus</v-icon>
              </v-btn>
            </template>
          </v-treeview>
        </v-card>
      </v-col>
      <v-col cols="8">
        <v-row class="my-10" v-show="active[0] != ''">
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
                    单元名称
                  </p>
                  <h4 class="text-h4 text-typo font-weight-bolder mb-0">
                    {{ node.title }}
                  </h4>
                </v-col>
              </v-row>

              <hr class="dark horizontal mt-3 mb-4" />
              <v-row class="px-4">
                <v-col cols="12">
                  <v-text-field
                    v-model="node.title"
                    counter="25"
                  ></v-text-field>
                </v-col>
              </v-row>
              <v-row class="px-4">
                <v-col cols="6">
                  <div>
                    <v-select
                      :items="modelCode"
                      label="请选择单元类型"
                      color="#e91e63"
                      class="font-size-input input-style placeholder-light border-radius-md select-style mb-0"
                      multi-line
                      v-model="node.model"
                      item-text="label"
                      item-value="id"
                    >
                    </v-select>
                    <v-row v-show="!(node.model == 71)">
                      <v-col cols="11" sm="12">
                        <v-dialog
                          v-model="sTimeshow"
                          transition="dialog-bottom-transition"
                          max-width="600"
                        >
                          <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                              v-model="sTime.date"
                              label="考试日期"
                              prepend-icon="mdi-clock-time-four-outline"
                              readonly
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>

                            <v-text-field
                              v-model="sTime.time"
                              label="考试时间"
                              prepend-icon="mdi-clock-time-four-outline"
                              readonly
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>
                          </template>
                          <v-card>
                            <v-card-title>开始时间 </v-card-title>
                            <v-card-text>
                              <v-row>
                                <v-col ms="12" cols="6">
                                  <v-date-picker
                                    v-model="sTime.date"
                                    scrollable
                                    locale="zh-cn"
                                  >
                                  </v-date-picker>
                                </v-col>
                                <v-col ms="12" cols="6">
                                  <v-time-picker
                                    v-model="sTime.time"
                                    scrollable
                                    use-seconds
                                    format="24hr"
                                  >
                                  </v-time-picker>
                                </v-col>
                              </v-row>
                            </v-card-text>
                            <v-card-actions>
                              <v-spacer></v-spacer>
                              <v-btn
                                color="primary"
                                text
                                @click="sTimeshow = false"
                              >
                                确定
                              </v-btn>
                              <v-btn
                                text
                                @click="
                                  sTime.date = '';
                                  sTime.time = '';
                                "
                              >
                                清空
                              </v-btn>
                            </v-card-actions>
                          </v-card>
                        </v-dialog>
                      </v-col>
                      <v-spacer></v-spacer>
                      <v-col cols="11" sm="12">
                        <v-dialog
                          v-model="eTimeshow"
                          transition="dialog-bottom-transition"
                          max-width="600"
                        >
                          <template v-slot:activator="{ on, attrs }">
                            <v-text-field
                              v-model="eTime.date"
                              label="考试日期"
                              prepend-icon="mdi-clock-time-four-outline"
                              readonly
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>

                            <v-text-field
                              v-model="eTime.time"
                              label="考试时间"
                              prepend-icon="mdi-clock-time-four-outline"
                              readonly
                              v-bind="attrs"
                              v-on="on"
                            ></v-text-field>
                          </template>
                          <v-card>
                            <v-card-title>开始时间 </v-card-title>
                            <v-card-text>
                              <v-row>
                                <v-col ms="12" cols="6">
                                  <v-date-picker
                                    v-model="eTime.date"
                                    scrollable
                                    locale="zh-cn"
                                  >
                                  </v-date-picker>
                                </v-col>
                                <v-col ms="12" cols="6">
                                  <v-time-picker
                                    v-model="eTime.time"
                                    scrollable
                                    use-seconds
                                    format="24hr"
                                  >
                                  </v-time-picker>
                                </v-col>
                              </v-row>
                            </v-card-text>
                            <v-card-actions>
                              <v-spacer></v-spacer>
                              <v-btn
                                color="primary"
                                text
                                @click="eTimeshow = false"
                              >
                                确定
                              </v-btn>
                              <v-btn
                                text
                                @click="
                                  eTime.date = '';
                                  eTime.time = '';
                                "
                              >
                                清空
                              </v-btn>
                            </v-card-actions>
                          </v-card>
                        </v-dialog>
                      </v-col>
                    </v-row>
                  </div>
                  <div></div>
                </v-col>
                <v-col cols="6" v-show="!(node.model == 71)">
                  <div>
                    <v-select
                      :disabled="actvBaseInfo.examCode * 1 != '50'"
                      :items="unitTypeCode"
                      label="请选择比赛模式"
                      color="#e91e63"
                      class="font-size-input input-style placeholder-light border-radius-md select-style mb-0"
                      multi-line
                      v-model="node.examCode"
                      item-text="label"
                      item-value="id"
                    >
                    </v-select>
                  </div>
                  <div>
                    <v-switch
                      v-model="node.face"
                      :label="boolToStringFace.get(Boolean(node.face))"
                    ></v-switch>
                  </div>
                  <div v-if="node.examCode == '50' || node.examCode == '51'">
                    <v-slider
                      v-model="node.during"
                      label="考试时间（分钟）"
                      thumb-color="#5e72e4"
                      thumb-label="always"
                    ></v-slider>
                  </div>
                  <div>
                    <v-slider
                      v-model="node.count"
                      label="题目数量"
                      thumb-color="#5e72e4"
                      thumb-label="always"
                      :max="maxExam"
                    ></v-slider>
                  </div>
                  <div></div>
                </v-col>
              </v-row>
            </v-card>
          </v-col>

          <v-col cols="12" v-show="node.model == 72">
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
                      >mdi-chart-box</v-icon
                    >
                  </v-avatar>
                </v-col>
                <v-col sm="8" class="text-end">
                  <p
                    class="text-sm mb-0 text-capitalize text-body font-weight-light"
                  >
                    Rank
                  </p>
                  <h4 class="text-h4 text-typo font-weight-bolder mb-0">
                    排名
                  </h4>
                </v-col>
              </v-row>

              <hr class="dark horizontal mt-3 mb-4" />
              <v-row align="center" justify="center">
                <snackerMsg ref="msg"></snackerMsg>
                <v-col  cols="12">
                  <v-card elevation="0">
                    <v-card-text>
                      <v-data-table
                        :headers="headers"
                        :items="list"
                        hide-default-footer
                        class="elevation-0"
                      >
                        <template v-slot:item.avt="{ item }">
                          <v-avatar
                            size="45"
                            class="shadow border-radius-lg my-5"
                          >
                            <img
                              :src="item.avt"
                              alt="Avatar"
                              class="border-radius-lg"
                            />
                          </v-avatar>
                        </template>
                
                        <!-- <template v-slot:item.actions="{ item }">
              <v-btn small @click="delJoin(item)" color="info"> 删除 </v-btn>
            </template> -->
                      </v-data-table>
                    </v-card-text>
                  </v-card>
                </v-col>
                <v-col lg="8" cols="12">
                  <v-pagination
                    v-model="page.now"
                    class="my-4"
                    :length="page.len"
                    @next="getList()"
                    @previous="getList()"
                    @input="getList()"
                  ></v-pagination>
                </v-col>
              </v-row>
            </v-card>
          </v-col>
          <v-col cols="6" v-show="!(node.model == 72)">
            <uploadMultiPart
              :fileName="fileName"
              :type="umpZip"
            ></uploadMultiPart>
          </v-col>

          <!-- <v-col cols="12" v-show="!(node.model == 72)">
            <uploadMultiPart
              :type="umpVideo"
              :fileName="fileName"
            ></uploadMultiPart>
          </v-col> -->
          <v-col cols="12" v-show="!(node.model == 72)">
            <uploadMultiPart :fileName="fileName"></uploadMultiPart>
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
import rankUnit from "../lookup/rankUnit";
export default {
  components: {
    HtmlEditor,
    snackerMsg,
    rankUnit,
    uploadMultiPart,
  },
  data() {
    return {
      active: [],
      unitTree: [
        {
          id: 1,
          name: "xxx",
          children: [{ id: 11, name: "ddd" }],
        },
      ],

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

      umpVideo: "video", //uploadMultiPart
      umpZip: "other",

      fileName: "",
      maxExam: 10,
      startTime: null,

      sTime: {
        date: "",
        time: "",
      },
      sTimeshow: null,
      eTime: {
        date: "",
        time: "",
      },
      eTimeshow: null,
      //rank
      headers: [
        { text: "排名" ,value:"rank"},
        { text: "头像", value: "avt" },
        {
          text: "用户名",
          align: "start",
          sortable: false,
          value: "UserName",
        },
        { text: "分数", value: "score" },
        { text: "考试所用时间", value: "during" },
        { text: "考试状态描述", value: "intro" },
        { text: "电话号码", value: "tel" },
        // { text: "操作", value: "actions", sortable: false },
      ],

      page: {
        now: 1,
        count: 8,
        len: 1,
      },
      list: [],
    };

    //
  },
  computed: {
    ...mapState([
      "config",
      "actvId",

      "code",
      "boolToStringFace",
      "boolToCode",
      "typeCode",
      "typeCodeSele",
      "awardCode",
      // "compeTypeCode",
      "unitTypeCode",
      "modelCode",
    ]),
    ...mapGetters(["userInfo", "actvBaseInfo"]),
  },
  methods: {
    // setUnitInfo
    ...mapMutations(["setUnitInfo"]),

    toQuest() {
      let that = this;
      that.$router.push("/user/create/editExam");
    },

    getUnitTree() {
      let that = this;
      console.log(
        "getUnitTree",
        that.userInfo.userId,
        that.actvId,
        that.config
      );
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
          console.log("infoTree获取到的信息:", res);
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

    // 添加父节点
    showDelNode(id) {
      let that = this;
      that.editId = id;
      that.delDialog = true;
      console.log("showDelNode:", id);
    },
    showAddNode(id) {
      let that = this;
      that.editId = id;
      that.addDialog = true;
      console.log("showDelNode:", id);
    },
    addNode(id) {
      let that = this;

      console.log("addOne:", {
        actvId: that.actvId,
        userId: that.userInfo.userId,
        title: that.addNodeTitle,
        parentId: id,
      });
      that.axios
        .post(
          "/unitTree/addOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            title: that.addNodeTitle,
            parentId: id,
          },
          that.config
        )
        .then((res) => {
          console.log("addOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getUnitTree();
            that.addDialog = false;
            that.addNodeTitle = "";
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    delNode(id) {
      let that = this;

      console.log("delOne:", {
        actvId: that.actvId,
        userId: that.userInfo.userId,
        unitId: id,
      });
      that.axios
        .post(
          "/unitTree/delOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            UnitId: id,
          },
          that.config
        )
        .then((res) => {
          console.log("delOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getUnitTree();
            that.delDialog = false;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
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
            that.node = res.data.node;

            let s = that.node.startTime.split(" ");
  
            that.sTime.date = s[0];
            that.sTime.time = s[1];
            let e = that.node.endTime.split(" ");
            that.eTime.date = e[0];
            that.eTime.time = e[1];

            if (that.actvBaseInfo.examCode * 1 != 50) {
              that.node.examCode = that.actvBaseInfo.examCode;
              // console.log("node.examCode", that.node.examCode);
            }
            that.maxExam = res.data.maxExam;
            that.fileName = res.data.node.fileName;
            // 再调用查看排名
            that.getList();
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // editOne
    editOne() {
      let that = this;
      if (!that.active[0]) {
        that.$refs.msg.setMsg("请选择节点", "error");
      }
      let oneNode = {
        title: that.node.title,
        tips: that.node.tips,
        model: that.node.model,

        during: that.node.during * 1,
        face: that.node.face * 1,
        count: that.node.count * 1,
        examCode: that.node.examCode * 1,
        startTime: that.sTime.date + " " + that.sTime.time,
        endTime: that.eTime.date + " " + that.eTime.time,
      };
      // console.log("getOne:", {
      //   actvId: that.actvId,
      //   userId: that.userInfo.userId,
      //   UnitId: that.active[0],
      //   oneNode: oneNode,
      // });

      that.axios
        .post(
          "/unitTree/editOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            UnitId: that.active[0],
            oneNode: oneNode,
          },
          that.config
        )
        .then((res) => {
          // console.log("getOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getUnitTree();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    getList() {
      let that = this;
      // console.log("getRankList:", {
      //   unitId: that.active[0],
      //   userCreateId: that.userInfo.userId,
      //   count: that.page.count,
      //   nowPage: that.page.now,
      // });
      let id = that.active[0] * 1;
      that.axios
        .post(
          "/joinUnitInfo/getRankList",
          {
            unitId: id,
            userCreateId: that.userInfo.userId,
            count: that.page.count,
            nowPage: that.page.now,
          },
          that.config
        )
        .then((res) => {
          that.list = res.data.list;

          that.page.len = res.data.pageObj.lenPage;
          // console.log("getRankList获取到的信息:", res);
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    let that = this;
    that.getUnitTree();
    // console.log("info", that.actvBaseInfo, that.node);
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
    sTimeshow(val) {
      // console.log("sTime", this.sTime);
    },
  },
};
</script>
<style lang="less" scoped></style>
