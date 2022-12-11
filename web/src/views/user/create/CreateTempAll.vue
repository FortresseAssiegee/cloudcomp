<template>
  <v-container fluid>
    <snackerMsg ref="msg"></snackerMsg>

    <v-row class="mt-8">
      <v-col cols="12">
        <v-data-table
          :headers="headers"
          :items="createList"
          :loading="loading"
          hide-default-footer
          class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar flat dark>
              <v-toolbar-title>创建的竞赛表</v-toolbar-title>
              <v-divider class="mx-4" inset vertical></v-divider>
              <v-spacer></v-spacer>
              <!--  @click="toAddActv" -->
              <v-btn color="grey darken-2" dark class="mb-2" @click="addShow = true" >
                新增题库
              </v-btn>
            </v-toolbar>
          </template>
          <template v-slot:item.stateCode="{ item }">
            <v-chip dark color="red">
              {{ code.get(item.stateCode) }}
            </v-chip>
          </template>
          <!-- <template v-slot:item.awardCode="{ item }">
            <v-chip dark color="cyan">
              {{ code.get(item.awardCode) }}
            </v-chip>
          </template> -->
          <template v-slot:item.infoDegree="{ item }">
            <v-chip dark color="primary"> {{ item.infoDegree }}% </v-chip>
          </template>

          <template v-slot:item.actions="{ item }">
            <v-icon
              small
              class="mr-2"
              @click="editItem(item)"
              v-if="item.stateCode != 11"
            >
              mdi-pencil
            </v-icon>
            <v-icon small class="mr-2" @click="checkItem(item)" v-else>
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

    <v-dialog v-model="addShow" width="500">
      <v-card>
        <v-card-title class="text-h5 grey lighten-2"> 新增活动 </v-card-title>

        <v-card-text>
          <v-text-field
            v-model="actvName"
            :counter="30"
            label="活动名称"
            required
          ></v-text-field>
        </v-card-text>
        <v-divider></v-divider>

        <v-card-actions>
          <!-- <v-spacer></v-spacer> -->
          <v-btn color="primary" text @click="toAddActv"> 确定 </v-btn>
          <v-btn color="" text @click="addShow = false"> 取消 </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
<script>
import { mapState, mapGetters } from "vuex";
import snackerMsg from "../../../components/commont/snackerMsg.vue";
export default {
  components: { snackerMsg },
  data() {
    return {
      // 结构
      loading: "true",

      headers: [
        {
          text: "活动名称",
          value: "actvName",
        },
        { text: "活动开始时间", value: "startTime" },
        { text: "活动结束时间", value: "endTime" },
        // { text: "活动类型", value: "typeCode" }, //20进行中 21结束 22未发布
        { text: "信息完善度", value: "infoDegree" },
        { text: "状态", value: "stateCode" },
        { text: "操作", value: "actions", sortable: false },
      ],

      activeList: [],
      createList: [],

      cards: [
        {
          icon: "mdi-home",
          text: "活动参与人数最多",
          value: "30",
          percent: "Linux操作系统",
          color: "bg-gradient-info shadow border-radius-xl mt-n8",
        },
        {
          icon: "mdi-home",
          text: "结束活动",
          value: "3",
          percent: "有54人参与活动",
          color: "bg-gradient-success shadow border-radius-xl mt-n8",
        },
        {
          icon: "mdi-home",
          text: "进行活动",
          value: "8",
          percent: "有30人参与活动",
          color: "bg-gradient-warning shadow border-radius-xl mt-n8",
        },
        {
          icon: "mdi-home",
          text: "最佳分数",
          value: "90",
          percent: "testUser1",
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
      page: {
        now: 1,
        count: 8,
        len: 1,
      },
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
      // console.log("deleteItem:", item);
      that.axios
        .post(
          "/baseInfo/del",
          {
            actvId: item.actvId,
            userId: sessionStorage.getItem("userId") * 1,
          },
          that.config
        )
        .then((res) => {
          // console.log("删除", res);
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
    editItem(item) {
      let that = this;
      // console.log("editItem", item);
      // if (item.stateCode == 10) {
      this.$store.state.actvId = item.actvId;
      that.$router.push("/user/create/goingTemp");
      // }
    },
    checkItem(item) {
      let that = this;
      // console.log("editItem", item);
      this.$store.state.actvId = item.actvId;
      that.$router.push("/user/create/doneTemp");
    },

    // 接口
    toAddActv() {
      let that = this;
      let addInfo = {
        userId: sessionStorage.getItem("userId") * 1,
        actvName: that.actvName,
      };

      // console.log("addInfo", addInfo);
      that.axios
        .post("/baseInfo/add", addInfo, that.config)
        .then((res) => {
          // console.log("添加新的:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.addShow = false;
            that.getList();
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

      that.axios
        .post(
          "/baseInfo/getAll",
          {
            userId: sessionStorage.getItem("userId") * 1,
            count: that.page.count,
            nowPage: that.page.now,
          },
          that.config
        )
        .then((res) => {
          // console.log("baseInfo/getAll:", res);
          that.createList = res.data.baseInfo;
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
