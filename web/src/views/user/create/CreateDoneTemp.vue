<template>
  <v-container fluid>
    <div
      class="min-height-300 border-radius-xl overflow-hidden position-relative mx-3"
      :style="`background-image: url('https://images.unsplash.com/photo-1531512073830-ba890ca4eba2?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1920&q=80'); background-position: 50%; background-size: cover;`"
    >
      <span class="mask bg-gradient-info opacity-6"></span>
    </div>
    <v-card
      class="shadow px-4 py-4 mx-9 overflow-hidden border-radius-xl mt-n16"
    >
      <v-row>
        <v-col cols="auto">
          <v-avatar width="74" height="74" class="shadow border-radius-lg">
            <img
              :src="userInfo.avt"
              alt="Avatar"
              class="border-radius-lg"
            />
          </v-avatar>
        </v-col>
        <v-col cols="auto" class="my-auto">
          <div class="h-100">
            <h5 class="mb-1 text-h5 text-typo font-weight-bold">
              {{ userInfo.userName }}
            </h5>
            <p class="mb-0 font-weight-light text-body text-sm">
              加入:{{ userInfo.joinNum }} / 创建:{{ userInfo.joinNum }}
            </p>
          </div>
        </v-col>

        <v-col lg="4" md="6" class="my-sm-auto ms-sm-auto me-sm-0 mx-auto mt-3">
          <v-tabs
            v-model="tab"
            class="mb-5"
            right
            background-color="transparent"
          >
            <v-tabs-slider color="yellow"></v-tabs-slider>
            <v-tab v-for="item in items" :key="item">
              {{ item }}
            </v-tab>
          </v-tabs>
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <v-tabs-items v-model="tab" class="rounded-lg">
            <v-tab-item> <edit-base-info></edit-base-info></v-tab-item>
            <v-tab-item> <edit-unit-tree></edit-unit-tree></v-tab-item>
            <v-tab-item> <lookupJoinTab></lookupJoinTab></v-tab-item>
          </v-tabs-items>
        </v-col>
      </v-row>
    </v-card>
  </v-container>
</template>

<script>
import editBaseInfo from "../../../components/lookup/lookupBaseInfo.vue";
import editUnitTree from "../../../components/edit/editUnitTree.vue";
import lookupJoinTab from "../../../components/lookup/lookupJoinTab.vue";
import { mapState,mapGetters } from "vuex";

export default {
  components: {
    editBaseInfo,
    editUnitTree,
    lookupJoinTab,
  },
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  data() {
    return {
      tab: null,
      items: ["基本信息", "教学流程", "报名情况"],
      stories: [
        {
          image: require("../../../assets/哈尔.png"),
          user: "Abbie W",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Boris U",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Kay R",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Tom M",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Nicole N",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Marie P",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Bruce M",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Sandra A",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Katty L",
        },
        {
          image: require("../../../assets/哈尔.png"),
          user: "Emma O",
        },
        {
          image: require("../../../assets/哈尔.png"),
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
              image: require("../../../assets/哈尔.png"),
              name: "Elena Morison",
            },
            {
              image: require("../../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
            {
              image: require("../../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
            {
              image: require("../../../assets/哈尔.png"),
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
              image: require("../../../assets/哈尔.png"),
              name: "Peterson",
            },
            {
              image: require("../../../assets/哈尔.png"),
              name: "Nick Daniel",
            },
            {
              image: require("../../../assets/哈尔.png"),
              name: "Elena Morison",
            },
            {
              image: require("../../../assets/哈尔.png"),
              name: "Ryan Milly",
            },
          ],
        },
      ],
      baseInfo: {},
    };
  },
  mounted() {
    // let that = this;
    // console.log("userInfo", that.userInfo);
  },
  methods: {},
};
</script>
<style lang="less" scoped>
.theme--light.v-tabs-items {
  background-color: transparent !important;
}
</style>
