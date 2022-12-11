<template>
  <v-row align="center" justify="center">
    <snackerMsg ref="msg"></snackerMsg>
    <v-col lg="8" cols="12">
      <v-card>
        <v-card-text>
          <v-data-table
            :headers="headers"
            :items="list"
            hide-default-footer
            class="elevation-0"
          >
            <template v-slot:item.avt="{ item }">
              <v-avatar size="45" class="shadow border-radius-lg my-5">
                <img :src="item.avt" alt="Avatar" class="border-radius-lg" />
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

    <v-dialog
      v-model="delDialogShow"
      max-width="600"
      transition="dialog-bottom-transition"
    >
      <v-card>
        <v-toolbar color="primary" dark>删除？</v-toolbar>
        <v-card-text>
          <div class="text-h5 pa-12">
            确定要取消<span class="text-h4 text-success">{{
              delItem.userName
            }}</span
            >的收藏资格？
          </div>
        </v-card-text>
        <v-divider></v-divider>

        <v-card-actions>
          <!-- <v-spacer></v-spacer> -->
          <v-btn color="green darken-1" text @click="deleteItem"> 确定 </v-btn>
          <v-btn color="" text @click="delDialogShow = false"> 取消 </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-row>
</template>
<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import snackerMsg from "../commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  data() {
    return {
      headers: [
        { text: "头像", value: "avt" },
        {
          text: "用户名",
          align: "start",
          sortable: false,
          value: "userName",
        },
        { text: "分数", value: "score" },
        { text: "考试所用时间", value: "during" },
        { text: "考试状态描述", value: "Intro" },
        { text: "电话号码", value: "tel" },
        // { text: "操作", value: "actions", sortable: false },
      ],
      desserts: [
        {
          name: "Donut",
          calories: 452,
          fat: 26.0,
        },
        {
          name: "KitKat",
          calories: 518,
          fat: 26.0,
        },
      ],

      page: {
        now: 1,
        count: 8,
        len: 1,
      },
      list: [],
      delItem: {},
      delDialogShow: false,
    };
  },
  computed: {
    ...mapState(["config", "code", "actvId", "typeCodeSele", "awardCode"]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
    getList() {
      let that = this;
      // console.log("getAllListJoin:", {
      //   actvId: that.userInfo.unitId,
      //   count: that.page.count,
      //   nowPage: that.page.now,
      // });
      that.axios
        .post(
          "/joinUnitInfo/getRankList",
          {
            unitId: that.userInfo.unitId,
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
    deleteItem() {
      let that = this;
      // console.log("delItem:", that.delItem);
      that.axios
        .post(
          "/joinInfo/out",
          {
            joinId: that.delItem.joinId,
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
    delJoin(item) {
      let that = this;

      that.delDialogShow = true;
      that.delItem = item;
    },
  },
  mounted() {
    let that = this;

    that.getList();
  },
};
</script>
<style lang=""></style>
