<template>
  <v-container fluid class="py-6 po">
    <snackerMsg ref="msg"></snackerMsg>
    <v-row align="center" class="ma-3" justify="center">
      <v-col cols="12" align="center">
        <v-img src="../../assets/学习 (1).png" width="150"></v-img>
        <div class="textFont text-center indigo--text text--darken-2">
          云知识竞赛平台
        </div>
      </v-col>
      <v-col lg="8" cols="12">
        <v-toolbar flat color="transparent">
          <!-- <v-text-field
            append-icon="mdi-magnify"
            flat
            hide-details
            label="请输入活动的名称"
            solo-inverted
          ></v-text-field> -->

          <template v-slot:extension>
            <v-tabs v-model="tabs" centered>
              <v-tab v-for="item in tagList" :key="'tab' + item.tagId">
                {{ item.label }}
              </v-tab>
            </v-tabs>
          </template>
        </v-toolbar>
      </v-col>
      <v-col lg="8" cols="12">
        <v-tabs-items v-model="tabs" class="px-5 py-5 border-radius-lg">
          <v-tab-item
            v-for="item in tagList"
            :key="'tab' + item.tagId"
            color="transparent"
          >
            <div v-for="(item, i) in otherTagList" :key="'other' + i">
              <h2 class="text-h6 mb-2 text-left">{{ item.name }}</h2>
              <v-chip-group
                v-model="search[item.label]"
                column
                multiple
                @change="searchBtn"
                :active-class="item.color"
              >
                <v-chip
                  filter
                  outlined
                  v-for="tag in item.children"
                  :key="tag.tagId"
                  :value="tag.tagId"
                >
                  {{ tag.label }}
                </v-chip>
              </v-chip-group>
            </div>
            <div>
              <h2 class="text-h6 mb-2 text-left">类型</h2>
              <v-chip-group
                v-model="search.tagCode"
                column
                multiple
                @change="searchBtn"
                active-class="deep-purple--text text--accent-4"
              >
                <v-chip
                  filter
                  outlined
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
    <v-divider class="mb-15"></v-divider>
    <v-row align="center" class="list">
      <v-col lg="3" md="4" v-for="item in searchList" :key="item.actvId">
        <v-card
          class="card card-shadow border-radius-xl py-8 text-center"
          data-animation="true"
        >
          <div class="mt-n11 mx-4 card-header position-relative z-index-2">
            <div class="d-block blur-shadow-image">
              <img
                :src="item.pstr"
                class="img-fluid shadow border-radius-lg"
                :alt="item.pstr"
                style="height: 250px"
              />
            </div>
            <div
              class="colored-shadow"
              style="height: 250px"
              v-bind:style="{ backgroundImage: 'url(' + item.pstr + ')' }"
            ></div>
          </div>
          <div class="d-flex mx-auto mt-n8">
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  v-bind="attrs"
                  v-on="on"
                  class="material-icons-round text-primary ms-auto px-5"
                  size="18"
                  text
                  @click="toInfo(item.actvId)"
                >
                  了解详情
                </v-btn>
              </template>
              <span>了解详情</span>
            </v-tooltip>
            <v-tooltip bottom>
              <template v-slot:activator="{ on, attrs }">
                <v-btn
                  v-bind="attrs"
                  v-on="on"
                  class="material-icons-round text-info me-auto px-5"
                  size="18"
                  text
                  @click="toJoin(item.actvId, item.userId, item)"
                >
                  立即加入
                </v-btn>
              </template>
              <span> 了解详情</span>
            </v-tooltip>
          </div>
          <h5 class="font-weight-normal text-typo text-h5 mt-7 mb-2 px-4">
            <a href="javascript:;" class="text-decoration-none text-default">{{
              item.actvName
            }}</a>
          </h5>
          <p
            class="mb-0 text-body font-weight-light px-5 txtEl"
            v-html="item.intro"
          ></p>
          <p>... ...</p>
          <hr class="horizontal dark my-6" />
          <div class="d-flex text-body mx-2 px-4">
            <p class="mb-0 font-weight-normal text-body">
              <span class="text-success text-lg"> 报名时间：</span
              ><span class="text-sm"
                >{{ item.applyStartTime }}至{{ item.applyEndTime }}</span
              >
            </p>
          </div>
          <div class="d-flex text-body mx-2 px-4">
            <p class="mb-0 font-weight-normal text-body">
              <span class="text-success text-lg"> 主办方: </span
              ><span class="text-sm">{{ item.createUserName }}</span>
            </p>
          </div>
        </v-card>
      </v-col>
    </v-row>
    <div class="text-center">
      <v-container>
        <v-row justify="center">
          <v-col cols="8">
            <v-container class="max-width">
              <v-pagination
                v-model="search.nowPage"
                class="my-4"
                :length="search.lenPage"
                @next="searchBtn()"
                @previous="searchBtn()"
                @input="searchBtn()"
              ></v-pagination>
            </v-container>
          </v-col>
        </v-row>
      </v-container>
    </div>
  </v-container>
</template>
<script>
import { mapState, mapGetters } from "vuex";
// import { mapMutations } from "vuex";
import snackerMsg from "../../components/commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  data() {
    return {
      tabs: null,
      text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",

      searchBtnShow: "true",
      // 搜寻条件
      search: {
        periodCode: [41, 44, 42, 43, 45, 46],
        awardCode: [30, 31],
        tagCode: [],
        examCode: [50, 51, 52, 53],
        count: 8,
        nowPage: 1,
        lenPage: null,
      },
      // 搜索结果
      searchList: [],

      tagList: [],

      otherTagList: [
        {
          name: "比赛模式",
          children: [
            {
              label: "混合模式",
              tagId: 50,
            },
            {
              label: "考试模式",
              tagId: 51,
            },
            {
              label: "闯关模式",
              tagId: 52,
            },
            {
              label: "限时答题模式",
              tagId: 53,
            },
          ],
          label: "examCode",
          allId: [50, 51, 52, 53],
          color: "red--text text--darken-2",
        },
        {
          name: "奖励",
          children: [
            {
              label: "无奖品",
              tagId: 30,
            },
            {
              label: "有奖品",
              tagId: 31,
            },
          ],
          label: "awardCode",
          allId: [30, 31],
          color: "orange--text text--darken-1",
        },

        {
          name: "比赛周期",
          children: [
            {
              label: "半个月内",
              tagId: 41,
            },
            {
              label: "一个月以内",
              tagId: 42,
            },
            {
              label: "三个月以内",
              tagId: 43,
            },
            {
              label: "半年以内",
              tagId: 44,
            },
            {
              label: "一年以内",
              tagId: 45,
            },
            {
              label: "一年以上",
              tagId: 46,
            },
          ],
          label: "periodCode",
          allId: [41, 44, 42, 43, 45, 46],
          color: "green--text text--darken-2",
        },
      ],
      searchAll: {
        periodCode: [41, 44, 42, 43, 45, 46],
        awardCode: [30, 31],
        tagCode: [],
        examCode: [50, 51, 52, 53],
      },
    };
  },
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    searchBtn() {
      let that = this;

      if (!that.search.periodCode.length) {
        that.search.periodCode = that.searchAll.periodCode;
      }
      if (!that.search.awardCode.length) {
        that.search.awardCode = that.searchAll.awardCode;
      }
      if (!that.search.tagCode.length) {
        // console.log("that.search.tagCode.length", that.search.tagCode.length);
        that.search.tagCode = that.searchAll.tagCode;
      }
      if (!that.search.examCode.length) {
        that.search.examCode = that.searchAll.examCode;
      }

      // console.log("选中：", that.search);

      that.axios
        .post(
          "/baseInfo/getTagShowList",
          that.search,
          that.config
        )
        .then((res) => {
          // console.log("getTagShowList 获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.searchList = res.data.baseInfo;

            that.search.lenPage = res.data.pageObj.lenPage;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
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
          console.log("getTagList 获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.tagList = res.data.TagInfo;
            for (let i = 0; i < that.tagList.length; i++) {
              let list = [];
              for (let j = 0; j < that.tagList[i].children.length; j++) {
                list.push(that.tagList[i].children[j].tagId);
              }
              that.tagList[i].allId = list;

              // that.search.tagCode = list;
            }
            that.searchAll.tagCode = that.tagList[0].allId;
            that.searchBtn();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    toInfo(actvId) {
      let that = this;
      this.$store.state.actvId = actvId;
      that.$router.push("/info");
    },
    toJoin(actvId, userId, baseInfo) {
      let that = this;
      console.log("checkJoin", {
        actvId: actvId,
        userCreateId: userId,
        userJoinId: this.userInfo.userId,
      });

      if (that.userInfo.userId === baseInfo.userId) {
        that.$refs.msg.setMsg("主办方不能参与报名!", "error");
        return;
      }
      let nowTime = new Date();
      let endTime = Date.parse(baseInfo.applyEndTime);
      let startTime = Date.parse(baseInfo.applyStartTime);
      if (nowTime < startTime || nowTime > endTime) {
        that.$refs.msg.setMsg("不在报名时间内!", "error");
        return;
      }

      that.axios
        .post(
          "/joinInfo/join",
          {
            actvId: actvId,
            userCreateId: userId,
            userJoinId: this.userInfo.userId,
          },
          that.config
        )
        .then((res) => {
          console.log("join 获取到的信息:", res.data);
          if (res.data.state == "win") {
            this.$store.state.actvId = actvId;
            that.$router.push("/user/join/goingTemp");
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
    that.getTagList();
  },
  watch: {
    tabs(val) {
      let that = this;
      that.searchAll.tagCode = that.tagList[val].allId;
      that.search.tagCode = that.tagList[val].allId;
      that.searchBtn();
      // console.log("tabs", val, that.searchAll.tagCode);
    },
  },
};
</script>
<style lang="less">
.txtEl {
  height: 64px;
  font-size: 14px;
  line-height: 16px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 5;
  -webkit-box-orient: vertical;
}
.po {
  margin-top: 200px;
  overflow: hidden;
}
.list {
  padding-left: 100px;
  padding-right: 100px;
}
.textFont {
  font-family: "方正舒体", Georgia, Serif;
  font-size: 3rem;
}
</style>
