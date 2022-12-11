<template>
  <v-container fluid>
    <snackerMsg ref="msg"></snackerMsg>

    <v-row class="mt-8">
      <v-col cols="12">
        <v-data-table
          :headers="headers"
          :items="createList"
          hide-default-footer
          class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar flat dark>
              <v-toolbar-title>加入的竞赛表</v-toolbar-title>
              <v-divider class="mx-4" inset vertical></v-divider>
              <v-spacer></v-spacer>
            </v-toolbar>
          </template>
          <template v-slot:item.stateCode="{ item }">
            <v-chip dark color="red">
              {{ code.get(item.stateCode) }}
            </v-chip>
          </template>

          <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2" @click="checkItem(item)">
              mdi-eye
            </v-icon>
            <v-icon small @click="deleteItem(item)"> mdi-delete </v-icon>
          </template>
        </v-data-table>
        <v-container>
          <v-pagination
            v-model="page.now"
            class="my-4"
            :length="page.len"
            @next="getList()"
            @previous="getList()"
            @input="getList()"
          ></v-pagination>
        </v-container>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import { mapState, mapGetters } from "vuex";
import snackerMsg from "../../../components/commont/snackerMsg.vue";
import { getNowDate } from "../../../components/diyTools/tools";
export default {
  components: { snackerMsg },
  data() {
    return {
      // 结构
      loading: "true",

      page: {
        now: 1,
        count: 8,
        len: 1,
      },

      headers: [
        {
          text: "活动名称",
          value: "actvName",
        },
        { text: "主办方", value: "userCreateName" },
        { text: "当前排名", value: "rank" },
        { text: "比赛开始时间", value: "examStartTime" },
        { text: "比赛结束时间", value: "examEndTime" },
        // { text: "活动类型", value: "typeCode" }, //20进行中 21结束 22未发布
        { text: "状态", value: "stateCode" },
        { text: "操作", value: "actions", sortable: false },
      ],

      activeList: [],
      createList: [],

      cards: [
        {
          icon: "mdi-home",
          text: "历史最佳排名",
          value: "1",
          percent: "Linux操作系统",
          color: "bg-gradient-info shadow border-radius-xl mt-n8",
        },
        {
          icon: "mdi-home",
          text: "进行中",
          value: "3",
          percent: "牛刀小试",
          color: "bg-gradient-success shadow border-radius-xl mt-n8",
        },
        {
          icon: "mdi-home",
          text: "结束",
          value: "8",
          percent: "有30人参与活动",
          color: "bg-gradient-warning shadow border-radius-xl mt-n8",
        },
        {
          icon: "mdi-home",
          text: "历史最佳分数",
          value: "90",
          percent: "python程序设计",
          color: "bg-gradient-primary shadow border-radius-xl mt-n8",
        },
      ],

      items: [
        {
          text: "Dashboard",
          disabled: false,
          href: "breadcrumbs_dashboard",
        },
        {
          text: "Link 1",
          disabled: false,
          href: "breadcrumbs_link_1",
        },
        {
          text: "Link 2",
          disabled: true,
          href: "breadcrumbs_link_2",
        },
      ],

      // 新增活动
      addShow: false,
      actvName: "",
    };
  },
  computed: {
    // 用户信息
    ...mapState(["config", "code"]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    deleteItem(item) {
      let that = this;
      // console.log("outItem:", item);
      that.axios
        .post(
          "/joinInfo/out",
          {
            joinId: item.joinId,
          },
          that.config
        )
        .then((res) => {
          // console.log("退出活动", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getList();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    checkItem(item) {
      let that = this;
      // console.log("editItem", item);
      let nowDate = getNowDate();
      this.$store.state.actvId = item.actvId;

      that.axios
        .post(
          "/dayActv/addHot",
          {
            date: nowDate,
            userId: item.userCreateId,
            actvId: item.actvId,
          },
          that.config
        )
        .then((res) => {
          // console.log("/joinInfo/getAllListJoin:", res);
    
        })
        .catch((error) => {
          console.log(error);
        });

      if (item.stateCode == 11) {
        that.$router.push("/user/join/goingTemp");
      } else {
        that.$router.push("/user/join/doneTemp");
      }
    },

    getList() {
      let that = this;
      that.axios
        .post(
          "/joinInfo/getAllListJoin",
          {
            userId: sessionStorage.getItem("userId") * 1,
            count: that.page.count,
            nowPage: that.page.now,
          },
          that.config
        )
        .then((res) => {
          // console.log("/joinInfo/getAllListJoin:", res);
          that.createList = res.data.joinList;
          that.page.len = res.data.pageObj.lenPage;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    let that = this;
    that.getList();
  },
};
</script>
<style lang="less" scoped></style>
