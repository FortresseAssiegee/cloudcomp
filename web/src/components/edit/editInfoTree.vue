<template>
  <v-container fluid>
    <snackerMsg ref="msg"></snackerMsg>
    <v-row>
      <v-col lg="6">
        <h4 class="text-h4 text-typo font-weight-bold mb-3">修改以下信息</h4>
        <p class="text-body font-weight-light">
          We’re constantly trying to express ourselves and actualize our dreams.
          If you have the opportunity to play.
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
          >保存</v-btn
        ></v-col
      >
    </v-row>
    <v-row>
      <v-col cols="4">
        <v-card
          class="position-sticky top-1 card-shadow border-radius-xl px-4 py-4"
          elevation="5"
        >
          <v-card-title class="text-h5">
            目录
            <v-spacer></v-spacer>
            <v-btn icon outlined @click="showAddNode(0)">
              <v-icon dark> mdi-plus </v-icon></v-btn
            ></v-card-title
          >
          <v-divider></v-divider>
          <v-treeview
            :items="infoTree"
            :active.sync="active"
            activatable
            color="warning"
            open-on-click
            transition
            item-text="title"
            class="text-left"
            :v-if="infoTree.length > 0"
          >
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
        <v-row>
          <v-col cols="12">
            <v-card
              class="p card-shadow border-radius-xl px-10 py-4"
              elevation="5"
            >
              <v-text-field v-model="title" counter="25"></v-text-field>
            </v-card>
          </v-col>
          <v-col cols="12">
            <v-card
              class="p card-shadow border-radius-xl px-8 py-4"
              elevation="5"
            >
              <html-editor v-model="value"></html-editor>
            </v-card>
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
import { mapState, mapGetters } from "vuex";
import HtmlEditor from "../commont/HtmlEditor.vue";
import snackerMsg from "../commont/snackerMsg.vue";

export default {
  components: {
    HtmlEditor,
    snackerMsg,
  },
  data() {
    return {
      active: [],

      infoTree: [
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
      value: "", //tips
      urlObj: {},

      // 传输文件的相关配置

      // ppt
      pType: "",
      pTip: "",
      pPath: "",
      // video
      vType: "",
      vTip: "",
      vPath: "",
      // other
      oType: "",
      oTip: "",
      oPath: "",

      pUrl: "",
      vUrl: "",
      oUrl: "",
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
            userId: that.userInfo.userId,
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
      // console.log("showDelNode:", id);
    },
    showAddNode(id) {
      let that = this;
      that.editId = id;
      that.addDialog = true;
      // console.log("showDelNode:", id);
    },
    addNode(id) {
      let that = this;

      // console.log("addOne:", {
      //   actvId: that.actvId,
      //   userId: that.userInfo.userId,
      //   title: that.addNodeTitle,
      //   parentId: id,
      // });
      that.axios
        .post(
          "/infoTree/addOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            title: that.addNodeTitle,
            parentId: id,
          },
          that.config
        )
        .then((res) => {
          // console.log("addOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getInfoTree();
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

      // console.log("delOne:", {
      //   actvId: that.actvId,
      //   userId: that.userInfo.userId,
      //   id: id,
      // });
      that.axios
        .post(
          "/infoTree/delOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            id: id,
          },
          that.config
        )
        .then((res) => {
          // console.log("delOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getInfoTree();
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
      //   id: that.active[0],
      // });
      that.axios
        .post(
          "/infoTree/getOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            id: that.active[0],
          },
          that.config
        )
        .then((res) => {
          // console.log("getOne获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.value = res.data.nodeObj.content;
            that.title = res.data.nodeObj.title;
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
        that.$refs.msg.setMsg("请选择修改的节点", "error");
      }
      // console.log("getOne:", {
      //   actvId: that.actvId,
      //   userId: that.userInfo.userId,
      //   id: that.active[0],
      //   title: that.title,
      //   content: that.value,
      // });
      that.axios
        .post(
          "/infoTree/editOne",
          {
            actvId: that.actvId,
            userId: that.userInfo.userId,
            id: that.active[0],
            title: that.title,
            content: that.value,
          },
          that.config
        )
        .then((res) => {
          // console.log("getOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getInfoTree();
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
    that.getInfoTree();
  },
  watch: {
    active() {
      let that = this;
      // console.log("active改变:", that.active);
      // 写一个getone的接口
      that.getOne();
    },
  },
};
</script>
<style lang="less"></style>
