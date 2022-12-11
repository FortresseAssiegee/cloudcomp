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
            <v-btn text small @click="downActv"> 下架活动 </v-btn>
            <v-btn
              outlined
              color="#fff"
              class="font-weight-bolder bg-gradient-info py-4 px-7"
              small
              @click="outActv"
            >
              结束活动
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

                <v-treeview open-all :items="infoTree">
                  <template v-slot:prepend="{ item }">
                    <h6
                      class="mb-0 text-h6 text-typo font-weight-bold"
                      v-show="true"
                    >
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
        <!-- <div class="px-4 py-4">
          <v-row class="align-center">
            <v-col cols="9">
              <h5
                class="text-h5 font-weight-bold mb-1 text-gradient text-primary"
              >
                <a href="javascript:;" class="text-decoration-none">海报</a>
              </h5>
            </v-col>
          </v-row>
        </div> -->
        <v-img :src="baseInfo.pstr" class="border-radius-lg shadow-lg"> </v-img>
      </v-card>

      <v-card class="mb-6">
        <div class="px-4 py-4">
          <v-row class="align-center">
            <v-col cols="9">
              <h5 class="text-h5 font-weight-bold mb-1 text-gradient text-info">
                <a href="javascript:;" class="text-decoration-none">基础信息</a>
              </h5>
            </v-col>
          </v-row>
          <v-list>
            <v-list-item-group class="border-radius-sm">
              <v-list-item :ripple="false" class="px-0 border-radius-sm">
                <v-list-item-content class="py-0">
                  <div class="text-body text-sm">
                    <strong class="text-dark">比赛类型:</strong>
                    &nbsp;
                    <v-chip-group active-class="primary--text">
                      <v-chip
                        class="d-block text-muted white--text"
                        color="deep-purple accent-4"
                      >
                        {{ baseInfo.periodCode }}
                      </v-chip>
                      <!-- <v-chip
                          class="d-block text-muted white--text"
                          color="deep-purple accent-4"
                        >
                          {{ baseInfo.tagCode }}
                        </v-chip> -->
                      <v-chip
                        class="d-block text-muted white--text"
                        color="orange darken-4"
                      >
                        {{ baseInfo.examCode }}
                      </v-chip>
                      <v-chip
                        class="d-block text-muted white--text"
                        color="green darken-3"
                      >
                        {{ baseInfo.awardCode }}
                      </v-chip>
                    </v-chip-group>
                  </div>
                </v-list-item-content>
              </v-list-item>
              <v-list-item :ripple="false" class="px-0 border-radius-sm">
                <v-list-item-content class="py-0">
                  <div class="text-body text-sm">
                    <strong class="text-dark">联系电话:</strong>
                    &nbsp; {{ baseInfo.createUserTel }}
                  </div>
                </v-list-item-content>
              </v-list-item>
              <v-list-item :ripple="false" class="px-0 border-radius-sm">
                <v-list-item-content class="py-0">
                  <div class="text-body text-sm">
                    <strong class="text-dark">联系邮箱:</strong>
                    &nbsp; {{ baseInfo.createUserEmail }}
                  </div>
                </v-list-item-content>
              </v-list-item>
              <v-list-item :ripple="false" class="px-0 border-radius-sm">
                <v-list-item-content class="py-0">
                  <div class="text-body text-sm">
                    <strong class="text-dark">比赛模式:</strong>
                    &nbsp; {{ baseInfo.examCode }}
                  </div>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </div>
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
      <!-- <v-card class="mb-6" v-for="card in cards" :key="card.title">
        <div class="px-4 py-4">
          <v-row class="align-center">
            <v-col cols="9">
              <h5
                class="text-h5 font-weight-bold mb-1 text-gradient text-primary"
              >
                <a href="javascript:;" class="text-decoration-none">{{
                  card.title
                }}</a>
              </h5>
            </v-col>
            <v-col cols="3" class="text-end">
              <v-menu
                transition="slide-y-transition"
                offset-y
                offset-x
                min-width="150"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                    icon
                    :ripple="false"
                    class="text-secondary"
                    v-bind="attrs"
                    v-on="on"
                    small
                  >
                    <v-icon size="16">fas fa-ellipsis-v</v-icon>
                  </v-btn>
                </template>

                <v-list class="pa-2">
                  <v-list-item class="list-item-hover-active border-radius-md">
                    <v-list-item-content class="pa-0">
                      <v-list-item-title
                        class="text-body-2 ls-0 text-body font-weight-600 py-2"
                      >
                        Edit Team
                      </v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item class="list-item-hover-active border-radius-md">
                    <v-list-item-content class="pa-0">
                      <v-list-item-title
                        class="text-body-2 ls-0 text-body font-weight-600 py-2"
                      >
                        Add Member
                      </v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                  <v-list-item class="list-item-hover-active border-radius-md">
                    <v-list-item-content class="pa-0">
                      <v-list-item-title
                        class="text-body-2 ls-0 text-body font-weight-600 py-2"
                      >
                        See Details
                      </v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                  <hr class="horizontal dark my-2" />
                  <v-list-item class="list-item-hover-active border-radius-md">
                    <v-list-item-content class="pa-0">
                      <v-list-item-title
                        class="text-danger ls-0 font-weight-600 mb-0"
                      >
                        Remove Team
                      </v-list-item-title>
                    </v-list-item-content>
                  </v-list-item>
                </v-list>
              </v-menu>
            </v-col>
          </v-row>
          <p class="text-body font-weight-light mt-3">
            {{ card.description }}
          </p>
          <div class="d-flex">
            <p class="mb-0 font-weight-light text-body">Industry:</p>
            <v-btn
              elevation="0"
              small
              :ripple="false"
              height="21"
              class="border-radius-md font-weight-bolder px-3 py-3 badge-font-size ms-auto text-body"
              color="#e4e8ed"
              >{{ card.industry }}</v-btn
            >
          </div>
          <hr class="horizontal dark my-4" />
          <div class="d-flex">
            <p class="mb-0 font-weight-light text-body">Rating:</p>
            <div class="rating ms-auto">
              <v-icon size="16">fas fa-star text-body</v-icon>
              <v-icon size="16">fas fa-star text-body ms-1</v-icon>
              <v-icon size="16">fas fa-star text-body ms-1</v-icon>
              <v-icon size="16">fas fa-star text-body ms-1</v-icon>
              <v-icon size="16" v-if="card.rating == 'partial'"
                >fas fa-star-half-alt text-body ms-1</v-icon
              >
              <v-icon size="16" v-if="card.rating == 'full'"
                >fas fa-star text-body ms-1</v-icon
              >
            </div>
          </div>
          <hr class="horizontal dark my-4" />
          <div class="d-flex">
            <p class="mb-0 font-weight-light text-body">Members:</p>
            <span class="avatar-group d-flex ms-auto">
              <v-tooltip
                bottom
                color="#212529"
                v-for="avatar in card.avatars"
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

      <v-card class="mb-6" v-for="(card, i) in brands" :key="i">
        <div class="px-4 py-4">
          <div class="d-flex">
            <v-avatar>
              <v-img :src="card.img"></v-img>
            </v-avatar>
            <div class="ms-2 my-auto">
              <h6 class="mb-0 text-h6 text-typo font-weight-bold">
                {{ card.title }}
              </h6>
              <p class="text-xs mb-0 text-body">{{ card.hour }}</p>
            </div>
          </div>
          <p class="text-body font-weight-light mt-3">
            {{ card.description }}
          </p>
          <div class="d-flex">
            <p class="mb-0 text-body">
              <b>Meeting ID</b>:
              <span class="font-weight-light">{{ card.id }}</span>
            </p>
          </div>
          <hr class="horizontal dark my-4" />
          <div class="d-flex">
            <v-btn
              :elevation="0"
              color="#cb0c9f"
              class="font-weight-bolder btn-success bg-gradient-default py-4 px-4"
              small
            >
              Join
            </v-btn>
            <span class="avatar-group d-flex ms-auto">
              <v-tooltip
                bottom
                color="#212529"
                v-for="avatar in card.avatars"
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
      </v-card> -->
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
      stories: [
        {
          image: require("../../assets/哈尔.png"),
          user: "Abbie W",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Boris U",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Kay R",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Tom M",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Nicole N",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Marie P",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Bruce M",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Sandra A",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Katty L",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Emma O",
        },
        {
          image: require("../../assets/哈尔.png"),
          user: "Tao G",
        },
      ],
      cards: [
        {
          title: "Digital Marketing",
          description:
            "A group of people who collectively are responsible for all of the work necessary to produce working, validated assets.",
          industry: "Martketing Team",
          rating: "partial",
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
          title: "Design",
          description:
            "Because it's about motivating the doers. Because I’m here to follow my dreams and inspire other people to follow their dreams, too.",
          industry: "Design Team",
          rating: "full",
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
              name: "Elena Morison",
            },
            {
              image: require("../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
          ],
        },
      ],
      brands: [
        {
          img: require("../../assets/哈尔.png"),
          title: "Slack Meet",
          hour: "11:00 AM",
          id: "902-128-281",
          description: "You have an upcoming meet for Marketing Planning",
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
          img: require("../../assets/哈尔.png"),
          title: "Invision",
          hour: "4:50 PM",
          id: "111-968-981",
          description:
            "You have an upcoming video call for Soft Design at 5:00 PM.",
          avatars: [
            {
              image: require("../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
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
              name: "Peterson",
            },
          ],
        },
      ],
      baseInfo: {},
      timeline: [],
      infoTree: [
        {
          id: 1,
          infoName: "竞赛详情",
          children: [
            {
              id: 2,
              infoName: "竞赛背景",
              content:
                "2022 年是党的二十大召开之年，也是中国共产主义青年团成立 100 周年，为持续深入推动党史学习，以赛促学，检验学习成果，激励广大师生“立大志、明大德、成大才、担大任”，为实现中华民族伟大复兴的中国梦凝聚起强大力量，根据四川省普通高等学校图书馆情报工作指导委员会（省图工委„2022‟4 号）通知要求，组织我校师生积极参加四川省高校“青春逢盛世〃奋斗正当时”党史知识竞赛活动。请校属各党委广泛宣传动员，积极组织本部门或单位师生踊跃参加。",
            },
            {
              id: 4,
              infoName: "竞赛目的",
              content:
                "2022 年是党的二十大召开之年，也是中国共产主义青年团成立 100 周年，为持续深入推动党史学习，以赛促学，检验学习成果，激励广大师生“立大志、明大德、成大才、担大任”，为实现中华民族伟大复兴的中国梦凝聚起强大力量，根据四川省普通高等学校图书馆情报工作指导委员会（省图工委„2022‟4 号）通知要求，组织我校师生积极参加四川省高校“青春逢盛世〃奋斗正当时”党史知识竞赛活动。请校属各党委广泛宣传动员，积极组织本部门或单位师生踊跃参加。",
            },
            {
              id: 3,
              infoName: "活动相关单位",
              content:
                "2022 年是党的二十大召开之年，也是中国共产主义青年团成立 100 周年，为持续深入推动党史学习，以赛促学，检验学习成果，激励广大师生“立大志、明大德、成大才、担大任”，为实现中华民族伟大复兴的中国梦凝聚起强大力量，根据四川省普通高等学校图书馆情报工作指导委员会（省图工委„2022‟4 号）通知要求，组织我校师生积极参加四川省高校“青春逢盛世〃奋斗正当时”党史知识竞赛活动。请校属各党委广泛宣传动员，积极组织本部门或单位师生踊跃参加。",
            },
          ],
        },
        {
          id: 5,
          infoName: "活动内容",
          children: [
            { id: 7, infoName: "竞赛内容" },
            { id: 8, infoName: "活动流程" },
            { id: 9, infoName: "培训流程" },
          ],
        },
        {
          id: 10,
          infoName: "竞赛评奖",
          children: [
            { id: 12, infoName: "竞赛模式" },
            { id: 13, infoName: "奖项设置" },
          ],
        },
        {
          id: 15,
          infoName: "联系方式",
        },
      ],
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
            that.baseInfo.awardCode = that.code.get(that.baseInfo.awardCode);
            that.baseInfo.periodCode = that.code.get(that.baseInfo.periodCode);
            that.baseInfo.examCode = that.code.get(that.baseInfo.examCode);

            that.setTimeLine(
              { label: "开始", time: that.baseInfo.startTime },
              { label: "报名", time: that.baseInfo.applyStartTime },
              { label: "报名截止", time: that.baseInfo.applyEndTime },
              { label: "考试", time: that.baseInfo.examStartTime },
              { label: "考试截止", time: that.baseInfo.examStartTime },
              { label: "结束", time: that.baseInfo.endTime }
            );
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
              // that.setHtmlTree(that.infoTree);
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
      let d=new Date()
      let nowTime =
        d.getFullYear() + "-" + (d.getMonth() + 1) + "-" + d.getDate();
      that.axios
        .post(
          "/baseInfo/out",
          {
            userId: that.userInfo.userId,
            actvId: that.actvId,
            endTime: nowTime,
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
