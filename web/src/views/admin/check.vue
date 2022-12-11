<template>
  <v-container>
    <snackerMsg ref="msg"></snackerMsg>
    <v-card class="card-shadow border-radius-xl">
      <div class="card-header-padding">
        <div class="d-flex align-center">
          <div>
            <h5 class="font-weight-bold text-h5 text-typo mb-0">认证申请</h5>
            <p class="text-sm font-weight-light text-body mb-0">
              认证等级越高,主页被分享的机会就会更多
            </p>
          </div>
        </div>
      </div>
      <v-tabs v-model="tab" align-with-title>
        <v-tabs-slider color="yellow"></v-tabs-slider>

        <v-tab v-for="item in items" :key="item.id">
          {{ item.label }}
        </v-tab>
      </v-tabs>
      <v-tabs-items v-model="tab">
        <v-tab-item>
          <v-card-text class="px-0 py-0">
            <v-data-table
              :headers="headersNew"
              :items="newGradeInfo"
              class="table"
              :page.sync="page"
              hide-default-footer
              @page-count="pageCount = $event"
              :items-per-page="itemsPerPage"
              mobile-breakpoint="0"
            >
              <template v-slot:item.startTime="{ item }">
                <div class="d-flex align-center ms-2">
                  <span class="text-sm font-weight-normal text-body">
                    {{ item.startTime }}
                  </span>
                </div>
              </template>

              <template v-slot:item.userName="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.userName }}
                </span>
              </template>

              <template v-slot:item.userTel="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.userTel }}
                </span>
              </template>

              <template v-slot:item.applyIntro="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.applyIntro }}
                </span>
              </template>

              <template v-slot:item.actions="{ item }">
                <v-btn
                  @click="
                    replyShow = true;
                    replyItem = item;
                  "
                  icon
                  elevation="0"
                  :ripple="false"
                  height="28"
                  min-width="36"
                  width="36"
                  class="btn-ls me-2 my-2 rounded-sm"
                  color="#67748e"
                >
                  <v-icon size="14" class="material-icons-round"
                    >mdi-reply</v-icon
                  >
                </v-btn>

                <v-btn
                  icon
                  @click="getDownLoad(item)"
                  elevation="0"
                  :ripple="false"
                  height="28"
                  min-width="36"
                  width="36"
                  class="btn-ls me-2 my-2 rounded-sm"
                  color="#67748e"
                >
                  <v-icon size="14" class="material-icons-round"
                    >mdi-download</v-icon
                  >
                </v-btn>
              </template>
            </v-data-table>
          </v-card-text>
          <v-card-actions class="card-padding">
            <v-row>
              <v-col cols="6" class="ml-auto d-flex justify-end">
                <v-pagination
                  class="pagination"
                  color="#D81B60"
                  v-model="page"
                  :length="pageCount"
                  circle
                ></v-pagination>
              </v-col>
            </v-row>
          </v-card-actions>
        </v-tab-item>
        <v-tab-item>
          <v-card-text class="px-0 py-0">
            <v-data-table
              :headers="headersPass"
              :items="passGradeInfo"
              class="table"
              :page.sync="page"
              hide-default-footer
              @page-count="pageCount = $event"
              :items-per-page="itemsPerPage"
              mobile-breakpoint="0"
            >
              <template v-slot:item.startTime="{ item }">
                <div class="d-flex align-center ms-2">
                  <span class="text-sm font-weight-normal text-body">
                    {{ item.startTime }}
                  </span>
                </div>
              </template>

              <template v-slot:item.endTime="{ item }">
                <div class="d-flex align-center ms-2">
                  <span class="text-sm font-weight-normal text-body">
                    {{ item.endTime }}
                  </span>
                </div>
              </template>
              <template v-slot:item.gradeCode="{ item }">
                <div class="d-flex align-center ms-2">
                  <v-chip
                    :color="item.gradeCode == 61 ? 'success' : 'red'"
                    dark
                  >
                    {{ item.gradeCode == 61 ? "认证通过" : "认证未通过" }}
                  </v-chip>
                </div>
              </template>

              <template v-slot:item.userName="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.userName }}
                </span>
              </template>

              <template v-slot:item.userTel="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.userTel }}
                </span>
              </template>

              <template v-slot:item.applyIntro="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.applyIntro }}
                </span>
              </template>

              <template v-slot:item.replyIntro="{ item }">
                <span class="text-sm font-weight-normal text-body">
                  {{ item.replyIntro }}
                </span>
              </template>
              <template v-slot:item.actions="{ item }">
                <v-btn
                  @click="getDownLoad(item)"
                  icon
                  elevation="0"
                  :ripple="false"
                  height="28"
                  min-width="36"
                  width="36"
                  class="btn-ls me-2 my-2 rounded-sm"
                  color="#67748e"
                >
                  <v-icon size="14" class="material-icons-round"
                    >mdi-download</v-icon
                  >
                </v-btn>
              </template>
            </v-data-table>
          </v-card-text>
          <v-card-actions class="card-padding">
            <v-row>
              <v-col cols="6" class="ml-auto d-flex justify-end">
                <v-pagination
                  class="pagination"
                  color="#D81B60"
                  v-model="page"
                  :length="pageCount"
                  circle
                ></v-pagination>
              </v-col>
            </v-row>
          </v-card-actions>
        </v-tab-item>
      </v-tabs-items>
    </v-card>

    <v-dialog v-model="replyShow" persistent max-width="600px">
      <v-card>
        <v-card-title>
          <span class="text-h5">认证回复</span>
        </v-card-title>
        <v-card-text>
          <div class="px-4 py-4">
            <v-list>
              <v-list-item-group class="border-radius-sm">
                <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">用户名:</strong>
                      &nbsp; {{ replyItem.userName }}
                    </div>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">电话号码:</strong>
                      &nbsp; {{ replyItem.userTel }}
                    </div>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item :ripple="false" class="px-0 border-radius-sm">
                  <v-list-item-content class="py-0">
                    <div class="text-body text-sm">
                      <strong class="text-dark">申请理由:</strong>
                      &nbsp; {{ replyItem.applyIntro }}
                    </div>
                  </v-list-item-content>
                </v-list-item>
              </v-list-item-group>
            </v-list>
          </div>
          <v-divider></v-divider>
          <v-container>
            <v-row>
              <v-col cols="12" sm="6">
                <v-subheader class="pl-0"> 认证PASS?</v-subheader>
                <v-switch
                  class="mt-0"
                  v-model="replyInfo.gradeCode"
                  color="indigo darken-3"
                  value="indigo darken-3"
                  hide-details
                  :label="replyInfo.gradeCode ? '通过' : '不通过'"
                ></v-switch>
                <v-subheader class="pl-0"> 认证等级</v-subheader>
                <v-slider
                  v-model="replyInfo.grade"
                  :thumb-size="24"
                  thumb-label="always"
                  step="1"
                  max="6"
                  min="0"
                  :tick-labels="satisfactionEmojis"
                  :disabled="!replyInfo.gradeCode"
                >
                  <template v-slot:thumb-label="{ value }">
                    {{ value }}
                  </template>
                </v-slider>
              </v-col>
              <v-col cols="12" sm="6">
                <v-subheader class="pl-0"> 认证回复</v-subheader>
                <v-textarea
                  outlined
                  name="input-7-4"
                  label="请输入拒绝原因或者感谢加入......"
                  v-model="replyInfo.replyIntro"
                ></v-textarea
              ></v-col>
            </v-row>
          </v-container>
          <small>*默认认证等级是1</small>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" text @click="replyShow = false">
            关闭
          </v-btn>
          <v-btn color="blue darken-1" text @click="reply()"> 确定 </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
<script>
import { mapState } from "vuex";
import snackerMsg from "../../components/commont/snackerMsg.vue";
export default {
  name: "paginated-tables",
  components: {
    snackerMsg,
  },
  data() {
    return {
      page: 1,
      pageCount: 0,
      itemsPerPage: 10,
      dialog: false,
      dialogDelete: false,
      search: "",
      editedIndex: -1,
      editedItem: {
        name: "",
        email: "",
        age: "",
        salary: "",
      },
      defaultItem: {
        name: "",
        email: "",
        age: "",
        salary: "",
      },
      headersNew: [
        {
          text: "申请时间",
          align: "start",
          cellClass: "border-bottom",
          sortable: true,
          value: "startTime",
          class:
            "text-secondary font-weight-bolder opacity-7 border-bottom ps-6",
        },
        {
          text: "用户名",
          value: "userName",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "电话号码",
          value: "userTel",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "申请理由",
          value: "applyIntro",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "操作",
          value: "actions",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
      ],
      headersPass: [
        {
          text: "申请时间",
          align: "start",
          cellClass: "border-bottom",
          sortable: true,
          value: "startTime",
          class:
            "text-secondary font-weight-bolder opacity-7 border-bottom ps-6",
        },
        {
          text: "处理时间",
          value: "endTime",
          sortable: true,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "状态",
          value: "gradeCode",
          sortable: true,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "用户名",
          value: "userName",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "电话号码",
          value: "userTel",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "申请理由",
          value: "applyIntro",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "回复信息",
          value: "replyIntro",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
        {
          text: "操作",
          value: "actions",
          sortable: false,
          class: "text-secondary font-weight-bolder opacity-7",
        },
      ],

      newGradeInfo: [],
      passGradeInfo: [],
      tab: null,
      items: [
        {
          label: "未审核",
          id: 1,
        },
        {
          id: 2,
          label: "审核记录",
        },
      ],
      replyShow: false,
      replyItem: {}, //回复的对象
      replyInfo: { gradeCode: false, garde: 1 }, //回复的具体信息
      satisfactionEmojis: ["😭", "😐", "🙂", "😊", "😁", "😄", "😍"],
    };
  },
  computed: {
    ...mapState(["userInfo", "config"]),
    pages() {
      return this.pagination.rowsPerPage
        ? Math.ceil(this.items.length / this.pagination.rowsPerPage)
        : 0;
    },
  },
  methods: {
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

    // /grade/getNewGradeInfo
    getNewGradeInfo() {
      let that = this;
      // console.log("获取config", this.config, this.userInfo);
      that.axios
        .post("/grade/getNewGradeInfo", {}, that.config)
        .then((res) => {
          // console.log("getNewGradeInfo获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.newGradeInfo = res.data.gradeInfo;
            // console.log("newGradeInfo", that.newGradeInfo);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getPassGradeInfo() {
      let that = this;
      that.axios
        .post("/grade/getPassGradeInfo", {}, that.config)
        .then((res) => {
          // console.log("getPassGradeInfo获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.passGradeInfo = res.data.gradeInfo;
            // console.log("passGradeInfo", that.passGradeInfo);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    reply() {
      let that = this;

      // console.log("申请:", {
      //   replyIntro: that.replyInfo.replyIntro,
      //   endTime: that.getNowTime(),
      //   gradeId: that.replyItem.gradeId,
      //   gradeCode: that.replyInfo.gradeCode ? 61 : 62,
      // });
      that.axios
        .post(
          "/grade/replyGrade",
          {
            replyIntro: that.replyInfo.replyIntro,
            endTime: that.getNowTime(),
            gradeId: that.replyItem.gradeId,
            gradeCode: that.replyInfo.gradeCode ? 61 : 62,
            garde: that.replyInfo.garde,
          },
          that.config
        )
        .then((res) => {
          // console.log("getPassGradeInfo获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.getNewGradeInfo();
            that.replyShow = false;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getDownLoad(item) {},

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
  },
  watch: {
    dialog(val) {
      val || this.close();
    },
    dialogDelete(val) {
      val || this.closeDelete();
    },
  },
  mounted() {
    let that = this;
    that.getNewGradeInfo();
    that.getPassGradeInfo();
  },
};
</script>
