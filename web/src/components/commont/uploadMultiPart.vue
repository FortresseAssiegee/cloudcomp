<template>
  <v-card class="mb-6 card-shadow border-radius-xl py-4">
    <snackerMsg ref="msg"></snackerMsg>
    <v-row no-gutters class="px-4">
      <v-col sm="2">
        <v-avatar
          :color="diyType.get(this.type).color"
          :class="diyType.get(this.type).shadow"
          height="64"
          width="64"
        >
          <v-file-input
            v-model="files"
            class="pl-4"
            prepend-icon="mdi-paperclip"
            @change="upload"
            counter="1"
            multiple
            :dark="true"
            :hide-input="true"
            v-if="!showPro"
            :disabled="miss"
          >
          </v-file-input>
          <v-icon @click="uploadCanncel" dark v-else class="mt-3"
            >mdi-cancel</v-icon
          >
        </v-avatar>
      </v-col>
      <v-col sm="8" class="text-end">
        <p class="text-sm mb-0 text-capitalize text-body font-weight-light">
          {{ diyType.get(this.type).subTitle }}
        </p>
        <h4 class="text-h4 text-typo font-weight-bolder mb-0">
          {{ diyType.get(this.type).title }}
        </h4>
      </v-col>
    </v-row>
    <hr class="dark horizontal" />
    <v-progress-linear
      :color="diyType.get(this.type).lineColor"
      :value="progress"
      v-show="showPro"
    ></v-progress-linear>
    <v-row class="px-4">
      <v-col cols="12">
        <div v-if="type == 'ppt'">
          <iframe
            frameborder="0"
            height="500"
            width="100%"
            :src="signUrl"
            v-if="signUrl"
          ></iframe>
          <video
            width="100%"
            height="200"
            autoplay="autoplay"
            loop
            muted="muted"
            src="../../assets/加速开发中.mp4"
            v-else
          ></video>
        </div>
        <div v-else-if="type == 'video'">
          <div v-show="signUrl">
            <!-- <video
              id="my-video"
              class="video-js"
              controls
              preload="auto"
              width="640"
              height="264"
              data-setup="{}"
            >
              <source :src="signUrl" type="video/mp4" />
              <source :src="signUrl" type="video/webm" />
              <p class="vjs-no-js">
                To view this video please enable JavaScript, and consider
                upgrading to a web browser that
                <a
                  href="https://videojs.com/html5-video-support/"
                  target="_blank"
                  >supports HTML5 video</a
                >
              </p>
            </video> -->

            <video width="100%" controls :src="signUrl">
              抱歉，你的浏览器不支持video标签。。
            </video>
          </div>
          <div v-show="!signUrl">
            <video
              width="100%"
              height="200"
              autoplay="autoplay"
              loop
              muted="muted"
            >
              <source src="../../assets/加速开发中.mp4" type="video/mp4" />
            </video>
          </div>
        </div>
        <div v-else-if="type == 'other'" align="center" justify="center">
          <!-- :poster="posterUrl" -->
          <v-simple-table fixed-header>
            <template v-slot:default>
              <thead>
                <tr>
                  <th class="text-center">名字</th>
                  <th class="text-center">下载</th>
                </tr>
              </thead>
              <tbody class="text-center">
                <tr v-for="item in otherList" :key="item.name">
                  <td>{{ setName(item.name) }}</td>
                  <td>
                    <a :href="item.url" style="text-decoration: none">
                      <v-icon color="grey darken-4">mdi-download</v-icon></a
                    >
                    <v-btn @click="delOne(item.name)" icon :disabled="miss">
                      <v-icon>mdi-trash-can </v-icon></v-btn
                    >
                  </td>
                </tr>
              </tbody>
            </template>
          </v-simple-table>
          <!-- <span v-else class="text-body font-weight-light">请上传资料... ...</span> -->
        </div>
      </v-col>
    </v-row>
  </v-card>
</template>
<script>
import OSS from "ali-oss";
import snackerMsg from "./snackerMsg.vue";
import videojs from "video.js";
import "videojs-contrib-hls";


export default {
  name: "uploadMultiPart",
  components: {
    snackerMsg,
  },
  props: {
    // 上传的类型
    // uploadtype: {
    //   type: String,
    //   default: null,
    //   require: true,
    // },

    // 上传类型
    type: {
      type: String,
      default: "ppt", //ppt video other
    },
    fileName: {
      type: String,
      default: "not",
      require: true,
    },
    miss: {
      type: Boolean,
      default: false,
      require: false,
    },
  },

  data() {
    return {
      // oss
      client: null,
      option: {}, //分片上传

      //   上传的文件
      files: [],

      //   进度条
      progress: 0,
      showPro: false,
      signUrl: "",
      videoUrl: "",

      getName: "", //ppt,video name

      diyType: new Map([
        [
          "ppt",
          {
            color: "bg-gradient-primary border-radius-xl pb-3 mt-n8 ",
            shadow: "shadow-primary",
            lineColor: "pink darken-1",
            title: "课件",
            subTitle: "Unit's PPT",
            mime: "text/plain",
            process: "imm/previewdoc,copy_1",
            dirPath: "ppt/",
          },
        ],
        [
          "video",
          {
            color: "bg-gradient-success border-radius-xl pb-3 mt-n8 ",
            shadow: "shadow-success",
            lineColor: "success",
            title: "视频",
            subTitle: "Unit's Video",
            mime: "video/mp4",
            process: "",
            dirPath: "video/",
          },
        ],
        [
          "other",
          {
            color: "bg-gradient-default border-radius-xl pb-3 mt-n8 ",
            shadow: "shadow-default",
            lineColor: "grey darken-4",
            title: "附件",
            subTitle: "Unit's Other",
            mime: ["application/zip", "application/rar"],
            process: "",
            dirPath: "other/",
          },
        ],
      ]),

      otherList: [],
    };
  },
  methods: {
    setName(name) {
      let that = this;
      let arr = name.split("/");
      let newName = arr.pop();
      // that.getList();
      return newName;
    },
    uploadCanncel() {
      let that = this;
      that.client.cancel();
      that.showPro = false;
      that.$refs.msg.setMsg("取消上传", "success");
    },

    async delOne(name) {
      let that = this;
      try {
        // 填写Object完整路径。Object完整路径中不能包含Bucket名称。
        let result = await that.client.delete(name);
        // console.log("delOne getName", result);
      } catch (error) {
        console.log("delOne getName", error);
      }
    },

    // 上传图片
    upload() {
      let that = this;
      //   判断类型进行分流
      // console.log("上传的ppt文件:", that.files[0], that.files[0].name);
      let file = that.files[0];
      if (that.type == "ppt") {
        // console.log("file", /^\.pptx{0,1}$/.test("第1章Linux操作系统概述.ppt"));
        if (!/^\S+\.pptx{0,1}$/.test(file.name)) {
          that.files = [];
          return that.$refs.msg.setMsg(
            "请上传ppt或pptx格式,且名称不能有空格",
            "error"
          );
        }
      } else if (that.type == "video") {
        if (!/^\S+\.mp4$/.test(file.name) && !/^\S+\.webm$/.test(file.name)) {
          that.files = [];
          return that.$refs.msg.setMsg(
            "请上传mp4或webm格式,且名称不能有空格",
            "error"
          );
        }
      } else if (that.type == "other") {
        if (!/^\S+\.zip$/.test(file.name) && !/^\S+\.rar$/.test(file.name)) {
          that.files = [];
          return that.$refs.msg.setMsg(
            "请上传zip或rar格式,且名称不能有空格",
            "error"
          );
        }
      }

      that.showPro = true;
      that.$refs.msg.setMsg("开始上传", "success");

      let param = new FormData();
      param.append("file", file, file.name);
      that.put(file.name, file);
    },

    async put(name, file) {
      let that = this;

      if (that.type != "other") {
        try {
          // console.log("getName", that.getName);
          // 填写Object完整路径。Object完整路径中不能包含Bucket名称。
          let result = await that.client.delete(that.getName);
          console.log("del getName", result);
        } catch (error) {
          console.log("del getName", error);
        }
      }

      try {
        var dirName =
          that.diyType.get(that.type).dirPath +
          that.fileName +
          "/" +
          //   new Date().getTime() +
          //   "/" +
          name;

        // console.log("dirName", dirName);
        //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。

        // 指定待删除的Object的名称。Object名称需填写不包含Bucket名称在内的Object的完整路径，例如exampledir/exampleobject.txt。
        await that.client
          .multipartUpload(dirName, file, {
            ...that.option,
          })
          .then((result) => {
            // console.log("multipartUpload res", result);
            that.geturl = result.name;
            // that.signUrl = that.client.signatureUrl(result.name, {
            //   process: that.diyType.get(that.type).process,
            // });
            // console.log("返回的结果：", result.name, that.signUrl);
            that.progress = 0;
            that.files = null;
            that.showPro = false;
            that.getList();
          })
          .catch((err) => {
            console.log("multipartUpload catch:", err);
          });
      } catch (e) {
        console.log(e);
        // that.$refs.msg.setMsg("上传超时", "error");
      }
    },

    getList() {
      let that = this;
      let path = that.diyType.get(that.type).dirPath + that.fileName + "/";
      // console.log("getList path", path);
      // path 前面不要加"/"
      that.client
        .list({
          prefix: path,
          //  delimiter: '/'
        })
        .then((res) => {
          // console.log(that.type + " list res", res);
          if (res.objects.length == 0) {
            that.signUrl = "";
            that.getName = "";
          } else {
            if (that.type == "other") {
              that.otherList = res.objects;
            } else {
              that.getName = res.objects[0].name;
              // console.log("getName", that.getName);
              that.signUrl = that.client.signatureUrl(that.getName, {
                process: that.diyType.get(that.type).process,
              });
            }
          }
        })
        .catch((err) => {
          console.log("list catch:", err);
        });
    },
    // setVideo() {
    //   this.player = videojs(
    //     "cameraMonitoringVideo",
    //     {
    //       bigPlayButton: false,
    //       textTrackDisplay: false,
    //       posterImage: true,
    //       errorDisplay: false,
    //       controlBar: true,
    //       muted: true, //静音模式 、、 解决首次页面加载播放。
    //     },
    //     function () {
    //       this.play(); // 自动播放
    //     }
    //   );
    // },
  },
  mounted() {
    let that = this;
    // console.log("获取信息:", that.userInfo.userId, that.actvId);
    // console.log("moun that.geturl", that.geturl, "getName", that.getName);
    // 连接oss
    that.client = new OSS({
      region: "oss-cn-beijing", //
      accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
      accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
      bucket: "cloud-compe", //OSS 存储区域名
      timeout: "1800000", //30min
    });

    that.option = {
      // 获取分片上传进度、断点和返回值。
      progress: (p, cpt, res) => {
        that.progress = p * 100;
        // console.log("option progress:", p, cpt, res);
      },
      // 设置并发上传的分片数量。
      parallel: 20,
      // 设置分片大小。默认值为1 MB，最小值为100 KB。
      partSize: 1024 * 1024,
      // headers,
      // 自定义元数据，通过HeadObject接口可以获取Object的元数据。
      meta: { year: 2020, people: "test" },
      mime: that.diyType.get(that.type).mime,
    };

    that.getName = "";
    that.getList();
    // that.setVideo();
  },
  watch: {
    fileName(val) {
      let that = this;
      // console.log("watch fileName:", that.fileName);
      that.otherList = "";
      that.signUrl = "";
      that.getList();
    },
    signUrl() {
      // let that = this;
      // console.log("watch signUrl:", that.signUrl);
    },
    progress(val) {
      let that = this;
      // console.log("pro", val);
      if (val * 1 == 100) {
        that.showPro = false;
      }
    },
  },
  destroyed() {
    this.signUrl = "";
    this.getName = "";
    // window.removeEventListener("fileName", this.fileName);
  },
};
</script>
<style lang=""></style>
