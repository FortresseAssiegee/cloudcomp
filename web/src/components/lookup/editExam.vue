<template>
  <v-row class="text-left mt-5">
    <snackerMsg ref="msg"></snackerMsg>

    <v-col md="2">
      <v-card class="shadow-success">
        <v-toolbar color="success" dark>
          <!-- <v-app-bar-nav-icon></v-app-bar-nav-icon> -->
          <v-toolbar-title>题目总览</v-toolbar-title>

          <v-spacer></v-spacer>

          <!-- <v-btn icon>
              <v-icon @click="addQuest">mdi-plus</v-icon>
            </v-btn> -->
        </v-toolbar>
        <v-divider></v-divider>
        <v-list shaped height="70vh" style="overflow: auto">
          <v-list-item-group>
            <template>
              <!-- <v-divider  :key="`divider-${i}`"></v-divider> -->
              <v-list-item :value="addItem" @click="getOne(null)">
                <template v-slot:default="{ active }">
                  <v-list-item-content class="text-left light-green--text">
                    <v-list-item-title>新增</v-list-item-title>
                  </v-list-item-content>
                </template>
              </v-list-item>
              <v-divider class="my-3"></v-divider>
              <v-list-item
                v-for="item in examList"
                :key="item.examId"
                :value="item"
                active-class="success--text text--accent-4"
                @click="getOne(item.examId)"
              >
                <template v-slot:default="{ active }">
                  <v-list-item-content class="text-left">
                    <v-list-item-title v-text="item.title"></v-list-item-title>
                  </v-list-item-content>

                  <v-list-item-action>
                    <!-- <v-checkbox
                        :input-value="active"
                        color="deep-purple accent-4"
                      ></v-checkbox> -->

                    <v-icon @click="delOne(item.examId)">mdi-delete</v-icon>
                  </v-list-item-action>
                </template>
              </v-list-item>
            </template>
          </v-list-item-group>
        </v-list>
      </v-card>
    </v-col>
    <v-col md="10" class="side">
      <v-card class="pt-8 px-1 shadow-blur px-2">
        <div class="mt-n15">
          <div
            class="bg-gradient-success shadow-success border-radius-lg px-3 py-2 mx-3"
          >
            <v-container>
              <v-row>
                <v-col md="10">
                  <div class="d-flex align-items-center">
                    <v-textarea
                      rows="1"
                      auto-grow
                      outlined
                      v-model="model.title"
                      dark
                    ></v-textarea>
                    <!-- <span class="text-sm text-white opacity-8"
                          >last seen today at 1:53am</span
                        > -->
                  </div>
                </v-col>
                <v-spacer></v-spacer>
                <v-col cols="1" class="my-auto">
                  <v-tooltip top v-if="model.examId != null">
                    <template v-slot:activator="{ on, attrs }">
                      <v-icon
                        v-bind="attrs"
                        v-on="on"
                        size="36"
                        class="material-icons-round text-white"
                        @click="editOne()"
                      >
                        mdi-upload
                      </v-icon>
                    </template>
                    <span>保存</span>
                  </v-tooltip>
                  <v-tooltip top v-else>
                    <template v-slot:activator="{ on, attrs }">
                      <v-icon
                        v-bind="attrs"
                        v-on="on"
                        size="36"
                        class="material-icons-round text-white"
                        @click="addObj()"
                      >
                        mdi-plus
                      </v-icon>
                    </template>
                    <span>添加</span>
                  </v-tooltip>

                  <!-- <v-tooltip top v-if="model.examId != null">
                      <template v-slot:activator="{ on, attrs }">
                        <v-icon
                          v-bind="attrs"
                          v-on="on"
                          size="36"
                          class="material-icons-round text-white"
                        >
                          mdi-delete
                        </v-icon>
                      </template>
                      <span>删除</span>
                    </v-tooltip> -->
                </v-col>
                <!-- <v-col cols="1" class="my-auto">
                   
                  </v-col> -->
              </v-row>
            </v-container>
          </div>
        </div>
        <!-- <div class="max-vh-50 overflow-scroll pt-6 px-4">
         
          </div> -->
        <v-row class="my-3 px-5">
          <v-col cols="6" class="px-10">
            <div :class="titleClass + ''">选项</div>
            <div :class="subTitleClass + ''">添加选项</div>
            <v-row class="my-3">
              <v-col cols="11">
                <v-textarea
                  outlined
                  label="选项"
                  dense
                  rows="1"
                  auto-grow
                  color="success"
                  v-model="addQuestText"
                ></v-textarea>
              </v-col>
              <v-col cols="1">
                <v-icon
                  color="grey lighten-1"
                  @click="addQuestItem"
                  class="mt-2"
                >
                  mdi-plus
                </v-icon>
              </v-col>
            </v-row>
            <div :class="subTitleClass + ''">选项内容</div>
            <v-list class="pa-2" dense>
              <v-list-item
                class="list-item-hover-active border-radius-md text--grey"
                v-for="(item, index) in model.quest"
                :key="'quest' + index"
              >
                <v-list-item-content class="pa-0">
                  <!-- <v-list-item-title
                      class="text-body ls-0 text-typo font-weight-600 mb-0"
                      >{{ item }}
                    </v-list-item-title> -->
                  <v-text-field
                    :value="item"
                    @input="chgQuest($event, index)"
                    dense
                  >
                  </v-text-field>
                </v-list-item-content>

                <v-list-item-action>
                  <v-icon
                    color="success lighten-1"
                    @click="delQuestItem(index)"
                  >
                    mdi-delete
                  </v-icon>

                  <v-icon
                    color="success lighten-1"
                    @click="intoAnswerItem(index)"
                  >
                    mdi-arrow-right-bottom
                  </v-icon>
                </v-list-item-action>
              </v-list-item>
            </v-list>
          </v-col>
          <v-divider vertical></v-divider>
          <v-col cols="6" class="px-10">
            <div :class="titleClass + ''">答案</div>
            <div :class="subTitleClass + ''">添加答案</div>
            <v-row class="my-3">
              <v-col cols="11">
                <v-textarea
                  outlined
                  label="答案"
                  dense
                  rows="1"
                  color="success"
                  auto-grow
                  v-model="addAnswerText"
                ></v-textarea>
              </v-col>
              <v-col cols="1">
                <v-icon
                  color="grey lighten-1"
                  @click="addAnswerItem"
                  class="mt-2"
                >
                  mdi-plus
                </v-icon>
              </v-col>
            </v-row>
            <div :class="subTitleClass + ''">答案内容</div>
            <v-list class="pa-2" dense>
              <v-list-item
                class="list-item-hover-active border-radius-md text--grey"
                v-for="(item, index) in model.answer"
                :key="'answer' + index"
              >
                <v-list-item-content class="pa-0">
                  <v-text-field
                    :value="item"
                    @input="chgQuest($event, index)"
                    dense
                  >
                  </v-text-field>
                </v-list-item-content>
                <v-list-item-action>
                  <v-icon color="success lighten-1" @click="delAnswerItem(i)">
                    mdi-delete
                  </v-icon>
                </v-list-item-action>
              </v-list-item>
            </v-list>
          </v-col>

          <v-col cols="12" class="mt-10 px-10" v-if="model.examId != null">
            <div :class="titleClass + ''">上传图片</div>
            <div :class="subTitleClass + ''">添加图片</div>

            <v-file-input
              class="my-5 mx-15 px-15"
              v-model="files"
              counter="20"
              multiple
              @change="uploadImgs"
              outlined
              prepend-icon="mdi-paperclip"
            ></v-file-input>

            <v-slide-group
              v-model="showUrl"
              class="pa-4"
              active-class="success"
              show-arrows
            >
              <v-slide-item v-for="item in imgUrl" :key="item">
                <v-card class="ma-4" width="200px">
                  <v-img :src="item.url" height="200px"> </v-img>
                  <v-card-actions>
                    <v-btn
                      text
                      color="deep-purple accent-4"
                      @click="delImg(item.name)"
                    >
                      删除
                    </v-btn>
                    <v-btn
                      text
                      color="deep-purple accent-4"
                      @click="preview(item.url)"
                    >
                      预览
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-slide-item>
            </v-slide-group>
          </v-col>
        </v-row>
      </v-card>
    </v-col>

    <v-dialog
      v-model="isShowImg"
      transition="dialog-bottom-transition"
      width="80vh"
    >
      <v-img :src="showImg"></v-img>
    </v-dialog>
  </v-row>
</template>
<script>
import uploadMultiImg from "../commont/uploadMultiImg.vue";
import { mapState, mapGetters } from "vuex";
import snackerMsg from "../commont/snackerMsg.vue";
import OSS from "ali-oss";
export default {
  name: "Teams",
  components: {
    uploadMultiImg,
    snackerMsg,
  },
  data: function () {
    // getAll getOne editOne delOne
    return {
      // class
      titleClass: "text-h6 mb-2 font-weight-black",
      subTitleClass: "subtitle-2 text--secondar text--secondary",


      model: [],
      model: {},
      examList: [
        {
          title: "",
          img: "",
          quest: [],
          answer: [],
          method: 50,
          examId: null,
        },
      ],
      addItem: {
        title: "",
        img: "",
        quest: [],
        answer: [],
        method: 50,
        examId: null,
      },
      addQuestText: "",
      addAnswerText: "",

      showUrl: [],

      client: null,
      files: [],
      imgUrl: [],

      showImg: "", //要查看的
      isShowImg: false,
    };
  },
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  methods: {
  
    preview(url) {
      let that = this;
      that.isShowImg = true;
      that.showImg = url;
    },

    chgQuest(e, index) {
      let that = this;
      that.model.quest[index] = e;
      // console.log("chgText", e, index, that.model.quest[index]);
    },
    chganswer(e, index) {
      let that = this;
      that.model.answer[index] = e;
      // console.log("chganswer", e, index, that.model.answer[index]);
    },

    addAnswerItem() {
      let that = this;
      // console.log("addAnswerItem", that.addAnswerText);
      that.model.answer.push(that.addAnswerText);
      that.addAnswerText = "";
    },
    addQuestItem() {
      let that = this;
      that.model.quest.push(that.addQuestText);
      that.addQuestText = "";
      // if (that.addQuestText=="") {
      //   this
      // }
      
    },
    delAnswerItem(i) {
      let that = this;
      that.model.answer.splice(i, 1);
    },
    delQuestItem(i) {
      let that = this;
      that.model.quest.splice(i, 1);
    },
    intoAnswerItem(i) {
      let that = this;
      that.model.answer.push(that.model.quest[i]);
    },

    addObj() {
      let that = this;
      // console.log("model", that.model);
      let addObj = {
        title: that.model.title,
        img: that.model.img,
        quest: that.model.quest.join("@.@"),
        answer: that.model.answer.join("@.@"),
        score: that.model.score,
        method: 50,
      };

      // console.log("addObj", {
      //   actvId: that.userInfo.actvId,
      //   userId: that.userInfo.userId,
      //   unitId: that.userInfo.unitId,
      //   questOneList: addObj,
      // });

      that.axios
        .post(
          "/examInfo/addOne",
          {
            actvId: that.userInfo.actvId,
            userId: that.userInfo.userId,
            unitId: that.userInfo.unitId,
            questOneList: addObj,
          },
          that.config
        )
        .then((res) => {
          // console.log("addOne获取到的信息:", res);
          that.getAll();
          that.addItem = {
            title: "",
            img: "",
            quest: [],
            answer: [],
            method: 50,
            examId: null,
          };
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getAll() {
      let that = this;
      that.axios
        .post(
          "/examInfo/get",
          {
            actvId: that.userInfo.actvId,
            userId: that.userInfo.userId,
            unitId: that.userInfo.unitId,
          },
          that.config
        )
        .then((res) => {
          // console.log("get获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");
            that.examList = res.data.questOne;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getOne(id) {
      let that = this;
      if (id == null) {
        that.model = that.addItem;
        return;
      }
      that.axios
        .post(
          "/examInfo/getOne",
          {
            actvId: that.userInfo.actvId,
            userId: that.userInfo.userId,
            unitId: that.userInfo.unitId,
            examId: id,
          },
          that.config
        )
        .then((res) => {
          // console.log("getOne获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            let obj = {
              title: res.data.questOneList.title,
              method: res.data.questOneList.method,
              quest: res.data.questOneList.quest.split("@.@"),
              img: res.data.questOneList.img,
              answer: res.data.questOneList.answer.split("@.@"),
              examId: res.data.questOneList.examId,
            };

            that.model = obj;
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    editOne() {
      let that = this;
      let editObj = {
        title: that.model.title,
        img: that.model.img,
        quest: that.model.quest.join("@.@"),
        answer: that.model.answer.join("@.@"),
        method: 50,
      };
      // console.log("editOne", editObj, that.model.examId);
      that.axios
        .post(
          "/examInfo/editOne",
          {
            actvId: that.userInfo.actvId,
            userId: that.userInfo.userId,
            unitId: that.userInfo.unitId,
            examId: that.model.examId,
            questOneList: editObj,
          },
          that.config
        )
        .then((res) => {
          // console.log("editOne获取到的信息:", res);
          if (res.data.state == "win") {
            // that.$refs.msg.setMsg(res.data.commont, "success");
            that.getAll();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    delOne(id) {
      let that = this;
      that.axios
        .post(
          "/examInfo/delOne",
          {
            actvId: that.userInfo.actvId,
            userId: that.userInfo.userId,
            unitId: that.userInfo.unitId,
            examId: id,
          },
          that.config
        )
        .then((res) => {
          // console.log("delOne获取到的信息:", res);
          if (res.data.state == "win") {
            that.getAll();
          } else {
            that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    getOss() {
      let that = this;
      that.client = new OSS({
        region: "oss-cn-beijing", //
        accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
        accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
        bucket: "cloud-compe", //OSS 存储区域名
      });
      // that.option = {
      //   // 获取分片上传进度、断点和返回值。
      //   progress: (p, cpt, res) => {
      //     that.progress = p * 100;
      //     // console.log("option progress:", p, cpt, res);
      //   },
      //   // 设置并发上传的分片数量。
      //   parallel: 20,
      //   // 设置分片大小。默认值为1 MB，最小值为100 KB。
      //   partSize: 1024 * 1024,
      //   // headers,
      //   // 自定义元数据，通过HeadObject接口可以获取Object的元数据。
      //   meta: { year: 2020, people: "test" },
      //   // mime: that.diyType.get(that.type).mime,
      // };
    },

    // 上传图片
    async uploadImgs() {
      let that = this;

      // console.log("上传的头像文件:", that.files[0], that.files[0].name);
      for (let i = 0; i < that.files.length; i++) {
        let file = that.files[i];
        let param = new FormData();
        param.append("file", file, file.name);
        console.log("开始上传");
        that.put(file.name, file);
      }
    },

    async put(name, file) {
      let that = this;

      try {
        var fileName = "/exam/" + that.model.img + "/" + name;
        // console.log("fileName", fileName, that.model.img);
        //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。
        let result = await that.client.put(fileName, file);
        console.log("返回的结果：", result);
        that.getImg();
      } catch (e) {
        console.log(e);
      }
    },

    //
    getImg() {
      let that = this;
      // console.log("getImg", that.model.img);
      if (that.model.img == "") {
        return;
      }
      that.client
        .list({
          //   prefix: "my-",
          prefix: "exam/" + that.model.img,
          //    delimiter: '/'
        })
        .then((res) => {
          // console.log("list res", res);
          that.imgUrl = res.objects;
        })
        .catch((err) => {
          console.log("list catch:", err);
        });
    },
    delImg(name) {
      let that = this;
      // 指定待删除的Object的名称。Object名称需填写不包含Bucket名称在内的Object的完整路径，例如exampledir/exampleobject.txt。
      that.client
        .delete(name)
        .then((res) => {
          // console.log("delImg", res);
          that.getImg();
        })
        .catch((e) => {
          console.log(e);
        });
    },
  },
  mounted() {
    let that = this;
    that.getOss();
    that.getAll();
    that.model = that.examList[0];
    // console.log("userInfo", that.userInfo);
  },
  watch: {
    model() {
      let that = this;
      // console.log("model", that.model);
      if (that.model.examId != null) {
        //   that.dirPath =
        //     that.userInfo.userId.toString() +
        //     that.userInfo.actvId.toString() +
        //     that.userInfo.unitId.toString() +
        //     that.model.examId.toString();
        //   console.log("dirPath:", that.dirPath);

        that.getImg();
      }
    },
  },
};
</script>

<style lang="less" scoped>
#bg {
  background: rgb(rgba(51, 51, 51, 0));
}
</style>
