<template>
  <v-container>
    <snackerMsg ref="msg"></snackerMsg>

    <v-card class="my-8" elevation="1">
      <v-card-title >
        <div class="">
          <h5 class="font-weight-bold text-h5 text-typo mb-0">标签管理</h5>
          <p class="text-sm text-body font-weight-light mb-0">
            增删改标签,使得分类更加精确,被选中的标签将会成为<b
              class="text-primary"
              ><v-icon>mdi-fire</v-icon>热门标签</b
            >
          </p>
        </div>
        <v-divider class="mx-4" inset vertical></v-divider>
        <v-spacer></v-spacer>
        <div>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on">
                <v-icon @click="tagAddShow = true">mdi-plus</v-icon>
              </v-btn>
            </template>
            <span>添加新类别</span>
          </v-tooltip>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn icon v-bind="attrs" v-on="on">
                <v-icon @click="keepHotTagList">mdi-fire</v-icon>
              </v-btn>
            </template>
            <span>保存热门标签</span>
          </v-tooltip>
        </div>
      </v-card-title>

      <v-divider class="mx-4"></v-divider>

      <v-card-text v-for="(item, i) in tagList" :key="item.id">
        <span class="subheading">{{ item.label }}</span>
        <v-chip-group
          v-model="hotTags"
          active-class="deep-purple--text text--accent-4"

          multiple
        >
          <v-chip
            filter
            outlined
            close
            @click:close="delTag(dev.tagId)"
            v-for="dev in item.children"
            :key="dev.id"
            :value="dev.tagId"
          >
            {{ dev.label }}
          </v-chip>
          <v-spacer></v-spacer>
          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <!-- addTag(item.tagId) -->
              <v-btn
                @click="
                  addChipsShow = true;
                  addChipsPid = item.tagId;
                "
                icon
                v-bind="attrs"
                v-on="on"
              >
                <v-icon>mdi-plus </v-icon>
              </v-btn>
            </template>
            <span
              >向<b>{{ item.label }}</b
              >类添加新标签</span
            >
          </v-tooltip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <v-btn @click="delTag(item.tagId)" icon v-bind="attrs" v-on="on">
                <v-icon>mdi-trash-can </v-icon>
              </v-btn>
            </template>
            <span
              >删除<b>{{ item.label }}</b
              >类</span
            >
          </v-tooltip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on, attrs }">
              <!-- addTag(item.tagId) -->
              <v-btn
                @click="
                  tagEditShow = true;
                  editTagId = item.tagId;
                  editTagLabel = item.label;
                "
                icon
                v-bind="attrs"
                v-on="on"
              >
                <v-icon>mdi-pencil </v-icon>
              </v-btn>
            </template>
            <span
              >修改<b>{{ item.label }}</b
              >类的名称</span
            >
          </v-tooltip>
        </v-chip-group>
      </v-card-text>
    </v-card>
    <!-- 添加新类 -->
    <v-dialog v-model="tagAddShow" width="500">
      <v-card>
        <v-card-title class="text-h5 grey lighten-2"> 标签类别 </v-card-title>

        <v-card-text>
          <v-text-field
            v-model="addTagLabel"
            :counter="8"
            required
          ></v-text-field>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="tagAdd(0)"> 添加</v-btn>
          <v-btn color="" text @click="tagAddShow = false"> 取消</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- 添加新标签 -->
    <v-dialog v-model="addChipsShow" width="500">
      <v-card>
        <v-card-title class="text-h5 grey lighten-2"> 添加新标签 </v-card-title>

        <v-card-text>
          <v-text-field
            v-model="addTagLabel"
            :counter="8"
            required
          ></v-text-field>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="tagAdd(addChipsPid)"> 添加</v-btn>
          <v-btn color="" text @click="addChipsShow = false"> 取消</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>

    <!-- 修改类名 -->
    <v-dialog v-model="tagEditShow" width="500">
      <v-card>
        <v-card-title class="text-h5 grey lighten-2">
          修改类别名称
        </v-card-title>

        <v-card-text>
          <v-text-field
            v-model="editTagLabel"
            :counter="8"
            required
          ></v-text-field>
        </v-card-text>

        <v-divider></v-divider>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="primary" text @click="editTag(editTagId)"> 修改</v-btn>
          <v-btn color="" text @click="tagEditShow = false"> 取消</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>
<script>
import { mapState } from "vuex";
import snackerMsg from "../../../components/commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  computed: {
    ...mapState(["userInfo", "config"]),
  },
  data() {
    return {
      crsbnList: [], //轮播图列表

      //   模式
      hotTags: [],
      tagList: [],
      tagAddShow: false,
      addTagLabel: "",

      tagEditShow: false,
      editTagLabel: "",
      editTagId: "",

      tagDelShow: false,

      //   chips
      addChipsShow: false,
      //   addChipsLabel: "",
      addChipsPid: "",

      
    };
  },
  methods: {

    // tag
    tagAdd(pid) {
      let that = this;
      // console.log("pid:", pid, "\nlabel:", that.addTagLabel);

      that.axios
        .post(
          "/tags/addOneTag",
          {
            parentId: pid,
            label: that.addTagLabel,
          },
          that.config
        )
        .then((res) => {
          // console.log("tagAdd 获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.addTagLabel = "";
            that.tagAddShow = false;
            that.addChipsShow = false;
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getTagList();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    delTag(tagId) {
      let that = this;
      // console.log("tagid:", tagId, "\nlabel:", that.delTagLabel);

      that.axios
        .post(
          "/tags/delOneTag",
          {
            tagId: tagId,
          },
          that.config
        )
        .then((res) => {
          // console.log("tagAdd 获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.delTagLabel = "";
            that.tagDelShow = false;
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getTagList();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    editTag(tagId) {
      let that = this;
      // console.log("tagid:", tagId, "\nlabel:", that.editTagLabel);

      that.axios
        .post(
          "/tags/editOneTag",
          {
            tagId: tagId,
            label: that.editTagLabel,
          },
          that.config
        )
        .then((res) => {
          // console.log("tagAdd 获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.editTagLabel = "";
            that.tagEditShow = false;
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.getTagList();
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
          // console.log("getTagList 获取到的信息:", res.data);
          if (res.data.state == "win") {
            // console.log("tagList", res.data.TagInfo);
            that.tagList = res.data.TagInfo;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },

    getHotTagList() {
      let that = this;
      that.axios
        .post("/tags/getHotTagList", {}, that.config)
        .then((res) => {
          // console.log("getHotTagList 获取到的信息:", res.data);
          if (res.data.state == "win") {
            that.hotTags = res.data.tagList;
            // console.log("hotTags", res.data.tagList);
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    keepHotTagList() {
      let that = this;

      that.axios
        .post(
          "/tags/keepHotTagList",
          {
            tagList: that.hotTags,
          },
          that.config
        )
        .then((res) => {
          // console.log("keepHotTagList 获取到的信息:", res.data);
          if (res.data.state == "win") {
            // console.log("hotTags", res.data.tagList);
            that.$refs.msg.setMsg(res.data.commont, "success");
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
    that.getHotTagList();
  },
  watch: {},
};
</script>
<style lang="less"></style>
