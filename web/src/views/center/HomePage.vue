<template>
  <v-container fluid class="py-6 px-10 po">
    <v-carousel
      cycle
      show-arrows-on-hover
      hide-delimiter-background
      delimiter-icon="mdi-minus"
      align="center"
      justify="center"
    >
      <v-carousel-item
        v-for="(n, num) in crsbnList"
        :key="`card-${num}`"
        width="75vw"
        :src="n.url"
        @click="getActvId(n.url)"
      >
      </v-carousel-item>
    </v-carousel>

    <!-- 企业学校合作 -->
    <v-row class="align-center px-2 mt-6 mb-2">
      <v-col sm="6" cols="12">
        <div class="d-flex">
          <div class="d-flex align-center">
            <v-icon
              size="28"
              class="material-icons-round me-1 text-body cursor-pointer"
              @click="$router.push('/SchoolCompe')"
              router
              >mdi-arrow-right</v-icon
            >
            <h6 class="text-h6 text-typo">合作企校</h6>
          </div>
          <!-- <div class="d-flex align-center">
            <v-icon
              size="14"
              class="material-icons-round me-1 text-body cursor-pointer"
              >企业</v-icon
            >
            <span class="text-sm text-body me-4">150</span>
          </div>
          <div class="d-flex align-center">
            <v-icon
              size="14"
              class="material-icons-round me-1 text-body cursor-pointer"
              >学校</v-icon
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
          </div> -->
        </div>
      </v-col>
      <v-col sm="6" cols="12">
        <div class="d-flex align-center">
          <div class="d-flex align-items-center ms-auto">
            <span class="avatar-group d-flex ms-auto">
              <v-tooltip
                bottom
                color="#212529"
                v-for="(n, index) in scList"
                :key="index"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-avatar
                    v-bind="attrs"
                    v-on="on"
                    size="35"
                    class="border border-white ms-n3"
                  >
                    <img :src="n.avt" alt="Avatar" />
                  </v-avatar>
                </template>
                <span>{{ n.userName }}</span>
              </v-tooltip>
            </span>
          </div>
          <small class="ps-2 font-weight-bold text-body text-sm">100+</small>
        </div>
      </v-col>
    </v-row>

    <!-- 热门推荐 -->
    <v-row class="d-flex align-center px-2 mt-6 mb-2 justify-space-between">
      <v-col sm="6" cols="12">
        <div class="pt-4">
          <div class="d-flex align-center">
            <v-icon
              size="28"
              class="material-icons-round me-1 text-body cursor-pointer"
              @click="$router.push('/AllShow')"
              router
              >mdi-arrow-right</v-icon
            >
            <h6 class="text-h6 text-typo">热门活动</h6>
          </div>
        </div></v-col
      >

      <v-col sm="6" cols="12" class="d-flex align-end flex-column">
        <v-sheet
          elevation="0"
          class="py-4 px-1"
          color="transparent"
          max-width="500px"
          justify="end"
        >
          <v-chip-group
            active-class="deep-purple--text text--accent-4"
            color="primary"
          >
            <v-chip
              v-for="item in hotTagList"
              :key="item.tagId"
              color="info"
              @click="hotTags = item"
              outlined
            >
              {{ item.label }}
            </v-chip>
          </v-chip-group>
        </v-sheet>
      </v-col>
      <v-col cols="12">
        <v-slide-group show-arrows>
          <v-slide-item
            v-for="(dev, n) in hotTags.children"
            :key="'slides' + n"
            class="my-5"
          >
            <v-hover v-slot="{ hover }">
              <v-card
                :elevation="hover ? 10 : 2"
                class="mx-5 shadow border-radius-xl border border-info px-1 py-1"
                @click="toInfo(dev.actvId)"
              >
                <v-list three-line>
                  <template>
                    <v-list-item>
                      <v-avatar
                        width="80"
                        height="80"
                        class="shadow border-radius-lg me-4"
                      >
                        <v-img
                          :src="dev.pstr"
                          alt="Avatar"
                          class="border-radius-lg"
                        />
                      </v-avatar>
                      <v-list-item-content>
                        <v-list-item-title
                          v-html="dev.actvName"
                        ></v-list-item-title>
                        <v-list-item-subtitle
                          ><span>报名时间：</span> {{ dev.applyStartTime }} |{{
                            dev.applyEndTime
                          }}
                          <br />
                          <span>主办方：</span>
                          {{ dev.createUserName }}</v-list-item-subtitle
                        >
                      </v-list-item-content>
                    </v-list-item>
                  </template>
                </v-list>
              </v-card></v-hover
            >
          </v-slide-item>
        </v-slide-group>
      </v-col>
    </v-row>
    <v-row class="d-flex align-center px-2 mt-6 mb-2 justify-space-between">
      <v-col sm="6" cols="12">
        <div class="pt-4">
          <div class="d-flex align-center">
            <v-icon
              size="28"
              class="material-icons-round me-1 text-body cursor-pointer"
              router
              >mdi-arrow-right</v-icon
            >
            <h6 class="text-h6 text-typo">排行榜单</h6>
          </div>
          <!-- @click="$router.push('/RankShow')" -->
        </div></v-col
      >
    </v-row>

    <v-row class="hotRank my-8">
      <v-col
        lg="4"
        md="6"
        cols="12"
        class="position-relative"
        v-for="(item, num) in rankImg"
        :key="num"
      >
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 12 : 2"
            :class="{ 'on-hover': hover }"
            id="card"
          >
            <v-img :src="item.img">
              <v-card-title class="text-h6 white--text">
                <v-row class="fill-height flex-column" justify="space-between">
                  <p class="mt-4 subheading text-left">
                    {{ item.title }}
                  </p>

                  <div>
                    <p
                      class="ma-0 text-body-1 font-weight-bold font-italic text-left"
                    >
                      {{ item.subText }}
                    </p>
                  </div>
                </v-row>
              </v-card-title>

              <v-list class="px-5" color="transparent">
                <v-list-item-group class="border-radius-sm">
                  <v-list-item
                    :ripple="false"
                    v-for="(item, index) in item.list"
                    :key="item.actvId"
                    class="px-0 border-radius-sm mb-2"
                    @click="toIntro()"
                  >
                    <v-list-item-avatar>
                      <v-icon class="red darken-4" dark v-if="index == 0">
                        {{ index + 1 }}
                      </v-icon>
                      <v-icon
                        class="orange darken-3"
                        dark
                        v-else-if="index == 1"
                      >
                        {{ index + 1 }}
                      </v-icon>
                      <v-icon
                        class="amber lighten-1"
                        dark
                        v-else-if="index == 2"
                      >
                        {{ index + 1 }}
                      </v-icon>
                      <v-icon class="grey lighten-1" dark v-else>
                        {{ index + 1 }}
                      </v-icon>
                    </v-list-item-avatar>

                    <v-avatar
                      width="48"
                      height="48"
                      class="shadow border-radius-lg me-4"
                    >
                      <v-img
                        :src="item.pstr || item.avt"
                        alt="Avatar"
                        class="border-radius-lg"
                      />
                    </v-avatar>
                    <v-list-item-content>
                      <div class="d-flex align-center">
                        <div>
                          <h6
                            class="mb-0 text-sm text-sm text-typo font-weight-bold white--text"
                          >
                            {{ item.actvName || item.UserName }}
                          </h6>
                          <p
                            class="mb-0 text-xs text-body font-weight-light white--text"
                          >
                            {{ item.createUserName }}
                          </p>
                        </div>
                      </div>
                    </v-list-item-content>
                  </v-list-item>
                </v-list-item-group>
              </v-list>
            </v-img>
          </v-card>
        </v-hover>
      </v-col>
    </v-row>
    <v-row class="d-flex align-center px-2 mt-6 mb-2 justify-space-between">
      <v-col sm="6" cols="12">
        <div class="pt-4">
          <div class="d-flex align-center">
            <v-icon
              size="28"
              class="material-icons-round me-1 text-body cursor-pointer"
              >mdi-heart</v-icon
            >
            <h6 class="text-h6 text-typo">猜你喜欢</h6>
          </div>
        </div></v-col
      >
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-card elevation="0" color="transparent">
          <!-- <div class="px-4 pt-4">
            <h6 class="mb-1 text-typo text-h6 font-weight-bold">每日推荐</h6>
            <p class="text-sm text-body">每天分享十个</p>
          </div> -->
          <div class="py-4">
            <v-slide-group show-arrows>
              <v-slide-item
                v-for="(dev, n) in shareRankList"
                :key="'slide' + n"
                v-slot="{ active, toggle }"
              >
                <v-card elevation="5" width="300" class="mx-5 pa-5 my-5">
                  <div class="position-relative">
                    <a class="d-block shadow-xl border-radius-xl">
                      <v-img
                        :src="dev.pstr"
                        class="shadow border-radius-xl"
                        height="300px"
                      >
                      </v-img>
                    </a>
                  </div>
                  <div class="px-1 pt-6">
                    <p class="text-body font-weight-light mb-0 text-sm">
                      {{ dev.applyStartTime }} -- {{ dev.applyEndTime }}
                    </p>
                    <a href="javascript:;" class="text-decoration-none">
                      <h5 class="text-h5 font-weight-bold text-typo mb-2">
                        {{ dev.actvName }}
                      </h5>
                    </a>
                    <p
                      class="mb-6 text-sm text-body font-weight-light txtEl"
                      v-html="dev.intro"
                    ></p>
                    <p>... ...</p>
                    <div class="d-flex align-center justify-space-between">
                      <v-btn
                        outlined
                        color="#e91e63"
                        class="font-weight-bolder py-4 px-5"
                        small
                        @click="toInfo(dev.actvId)"
                      >
                        了解详情
                      </v-btn>
                      <span class="avatar-group d-flex">
                        <v-tooltip
                          bottom
                          color="#212529"
                          v-for="avatar in dev.avatars"
                          :key="avatar.name"
                        >
                          <template v-slot:activator="{ on, attrs }">
                            <v-avatar
                              v-bind="attrs"
                              v-on="on"
                              size="24"
                              class="border border-white ms-n3"
                            >
                              <img :src="avatar.image" alt="Avatar" />
                            </v-avatar>
                          </template>
                          <span>{{ avatar.name }}</span>
                        </v-tooltip>
                      </span>
                    </div>
                  </div>
                </v-card>
              </v-slide-item>
            </v-slide-group>
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState, mapGetters } from "vuex";
import OSS from "ali-oss";
export default {
  data() {
    return {
      model: "slid0",

      // 相关数据

      rankImg: [
        {
          title: "本周排行榜",
          list: [],
          subText: "站内根据参与度前5名进行排名，每周更新",
          img: "https://images.unsplash.com/photo-1429514513361-8fa32282fd5f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=3264&q=80",
        },
        {
          title: "企校排行榜",
          list: [],
          subText: "综合企业/学校的所发行的所有赛事的热度之和",
          img: "https://images.unsplash.com/photo-1498038432885-c6f3f1b912ee?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2100&q=80",
        },
        {
          title: "热门排行榜",
          list: [],
          subText: "我们将根据一周内题库的热度，进行推荐",
          img: "https://images.unsplash.com/photo-1542320868-f4d80389e1c4?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=3750&q=80",
        },
      ],
      crsbnList: [],
      hotTagList: [],
      hotTags: [],
      shareRankList: [],
      scList: [],
    };
  },
  computed: {
    ...mapState(["config", "code"]),
  },
  methods: {
 

    async getList() {
      let that = this;
      // path 前面不要加"/"
      that.client
        .list({
          prefix: "crsbn",
          // delimiter: "/",
        })
        .then((res) => {
          that.crsbnListLog = [];
          for (let i = 1; i < res.objects.length; i++) {
            const element = res.objects[i];
            element.url = that.client.signatureUrl(element.name);
            that.crsbnList.push(res.objects[i]);
          }
          // console.log("list res", res);
        })
        .catch((err) => {
          console.log("list catch:", err);
        });
    },
    getHotTagList() {
      let that = this;
      that.axios
        .post(
          "/baseInfo/getHotTagShowList",
          {},
          that.config
        )
        .then((res) => {
          // console.log("热门", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.hotTagList = res.data.hotTagList;
            that.hotTags = that.hotTagList[1];
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    getRankList() {
      let that = this;
      that.axios
        .post("/baseInfo/getRankAll", {}, that.config)
        .then((res) => {
          // console.log("getRankList", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");

            that.rankImg[0].list = res.data.weekLookUpRank;
            that.rankImg[1].list = res.data.schoolCompeRank;
            that.rankImg[2].list = res.data.weekHotRank;
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getShareRankList() {
      let that = this;
      that.axios
        .post("/baseInfo/getAllNoUser", {}, that.config)
        .then((res) => {
          // console.log("getShareRankList", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.shareRankList = res.data.baseInfo;
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getSchoolCompe() {
      let that = this;
      that.axios
        .post(
          "/grade/GetAllGrade",
          {
            num: 8,
          },
          that.config
        )
        .then((res) => {
          // console.log("grade/GetAllGrade:", res);
          if (res.data.state == "win") {
            that.scList = res.data.userList;
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    toInfo(actvId) {
      let that = this;
      // console.log("toInfo", actvId);
      this.$store.state.actvId = actvId;
      that.$router.push("/info");
    },
    getActvId(str) {
      let that = this;
      let arr = str.split("/");

      that.toInfo(arr[4] * 1);
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
    that.getSchoolCompe();
    that.getList();
    that.getHotTagList();
    that.getRankList();
    that.getShareRankList();
  },
};
</script>

<style lang="less" scoped>
.txtEl {
  height: 96px;
  font-size: 14px;
  line-height: 16px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 6;
  -webkit-box-orient: vertical;
}
.po {
  margin-top: 250px;
}
.hotRank {
  padding-left: 200px;
  padding-right: 200px;
}
.v-parallax {
  overflow: hidden;
}

#card {
  transition: opacity 0.4s ease-in-out;
}

#card:not(.on-hover) {
  opacity: 0.6;
}

.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}
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
</style>
