<template>
  <v-container>
    <snackerMsg ref="msg"></snackerMsg>
    <v-row class="d-flex h-100 mt-md-n16 pb-0">
      <v-col cols="3">
        <v-card class="card-shadow border-radius-xl position-sticky top-1">
          <v-app-bar dark color="deep-purple darken-2">
            <v-app-bar-nav-icon></v-app-bar-nav-icon>
            <h5 class="mb-1 text-h5 font-weight-bold">题目列表</h5>
          </v-app-bar>

          <v-card-text>
            <v-row justify="center">
              <v-col cols="4" v-for="(item, i) in examList" :key="item.examId"
                ><v-btn
                  :color="
                    item.answer ? 'purple darken-2' : 'deep-purple lighten-3'
                  "
                  fab
                  small
                  dark
                  :href="'#exam' + item.examId"
                >
                  {{ i + 1 }}
                </v-btn>
              </v-col>
            </v-row>
            <v-divider class="my-5"></v-divider>
            <v-row>
              <v-col cols="12">
                距离考试结束还有：
                <br />
                <v-chip
                  x-large
                  color="purple"
                  dark
                  class="text-center"
                  outlined
                >
                  <template v-slot:default>
                    <h1>{{ minNum }}</h1>
                    <br />分</template
                  >
                </v-chip>
                ：
                <v-chip
                  x-large
                  color="deep-purple lighten-1"
                  dark
                  class="text-center"
                  outlined
                >
                  <template v-slot:default>
                    <h3>{{ senNum }}</h3>
                    <br />秒
                  </template>
                </v-chip>
              </v-col>
            </v-row>
            <v-divider class="my-5"></v-divider>
            <div class="see">
              <video
                id="testVi"
                width="200px"
                height="200px"
                poster="../../assets/jidou.jpg"
                muted
                loop
                playsinline
                @loadedmetadata="fnRun"
              ></video>
              <canvas id="testCav" />

              <!-- 截图 -->
              <!-- <div class="target">
                <img height="200px" id="testImg2" :src="src" />
                <canvas id="testCav2" />
              </div> -->
            </div>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="9">
        <v-card class="card-shadow px-4 py-4 overflow-hidden border-radius-xl">
          <v-row>
            <v-col cols="auto">
              <v-avatar width="74" height="74" class="shadow rounded-circle">
                <img
                  alt="Avatar"
                  class="rounded-circle"
                  src="../../assets/jidou.jpg"
                />
              </v-avatar>
            </v-col>
            <v-col cols="auto" class="my-auto">
              <div class="h-100">
                <h5 class="mb-1 text-h5 text-typo font-weight-bold">
                  {{ unitInfo.actvName }}-{{ unitInfo.unitName }}
                </h5>
                <p class="mb-0 font-weight-light text-body text-sm">
                  考试时间：{{ sExam }}到{{ mabeExam }}
                </p>
              </div>
            </v-col>
            <v-col
              lg="4"
              md="6"
              class="my-sm-auto ms-sm-auto me-sm-0 mx-auto mt-3"
            >
              <div class="d-flex align-center justify-center">
                考试环境违规：{{ record.envirWarning }}次<br />
                考生非本人：{{ record.simVal }}次<br />
                <!-- <v-btn text>退出考试</v-btn> -->
                <v-divider vertical class="px-5"></v-divider>
                <v-btn color="deep-purple lighten-2" text @click="endExam()"
                  >交卷</v-btn
                >
              </div>
            </v-col>
          </v-row>
        </v-card>
        <v-card
          class="card-shadow my-5 px-4 py-4 overflow-hidden border-radius-xl"
          v-for="(item, i) in examList"
          :key="item.examId"
          :id="'exam' + item.examId"
        >
          <v-card-title>{{ i + 1 }}.{{ item.title }}</v-card-title>
          <v-card-text>
            <v-radio-group v-model="item.answer" column>
              <v-radio
                :label="n"
                :value="n"
                v-for="(n, num) in item.quest"
                :key="num"
              ></v-radio>
            </v-radio-group>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import snackerMsg from "../../components/commont/snackerMsg.vue";
import { getNowTime, GetDuring } from "../../components/diyTools/tools.js";
import * as faceapi from "face-api.js";
import OSS from "ali-oss";
export default {
  name: "exam51",
  components: { snackerMsg },
  data() {
    return {
      sExam: "",
      mabeExam: "",
      id: "",

      examList: [],
      time: null,
      number: null,
      minNum: 0,
      senNum: 0,
      record: {
        envirWarning: 0,
        simVal: 0,
      },
      name: "",

      // face 参数
      vi: null, //视频
      cav: null, //画图
      img: null,
      img2: null,

      options: null, //模型参数
      optionsImg: null,

      // 视频媒体参数配置
      constraints: {
        audio: false,
        video: {
          // ideal（应用最理想的）
          width: 200,
          height: 200,
          // frameRate受限带宽传输时，低帧率可能更适宜
          frameRate: {
            min: 15,
            ideal: 30,
            max: 60,
          },
          // 显示模式前置后置
          facingMode: "user",
        },
      },
      timeout: 0,
      timeoutImg: 0,

      src: "",

      clientOss: null,
      clientCore: null,

      // 监听
      blurNum: 1,
    };
  },
  computed: {
    ...mapState(["config", "code", "actvId"]),
    ...mapGetters(["userInfo", "actvBaseInfo", "unitInfo"]),
  },
  methods: {
    ...mapMutations(["setExamInfo"]),

    endExam() {
      let that = this;
      that.fnClose();
      // console.log("sExam", that.unitInfo);

      let aObj = [];

      for (let i = 0; i < that.examList.length; i++) {
        aObj[i] = {
          examId: that.examList[i].examId,
          answer: that.examList[i].answer || "",
        };
      }

      let endObj = {
        examId: that.id,
        actvId: that.unitInfo.actvId,
        unitId: that.unitInfo.unitId,
        userCreateId: that.unitInfo.userCreateId,
        userJoinId: that.unitInfo.userJoinId,

        startTime: that.sExam,
        endTime: getNowTime(),

        answerList: aObj,
        recordInfo: {
          intro: "",
          illegal: "",
          illegalIntro:
            "考试环境违规：" +
            that.record.envirWarning +
            "次；考生与本人不匹配：" +
            that.record.simVal +
            "次；",
        },
      };
      // console.log("提交：", endObj);
      that.axios
        .post("/joinUnitInfo/endExam", endObj, that.config)
        .then((res) => {
          // console.log("endExam获取到的信息:", res);
          if (res.data.state == "win") {
            that.$refs.msg.setMsg(res.data.commont, "success");

            setTimeout(() => {
              that.$router.push("/user/join/goingTemp");
            }, 1000);
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    startExam() {
      let that = this;
      // console.log("sExam", that.unitInfo);
      that.axios
        .post(
          "/joinUnitInfo/startExam",
          {
            actvId: that.unitInfo.actvId,
            unitId: that.unitInfo.unitId,
            userCreateId: that.unitInfo.userCreateId,
            userJoinId: that.unitInfo.userJoinId,
          },
          that.config
        )
        .then((res) => {
          // console.log("startExam获取到的信息:", res);
          if (res.data.state == "win") {
            that.examList = res.data.questList;
            that.id = res.data.examId;

            let tObj = GetDuring(that.unitInfo.during);
            that.sExam = tObj.now;
            that.mabeExam = tObj.end;

            that.setInfo();

            that.$refs.msg.setMsg(res.data.commont, "success");
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
    // 设置倒计时
    setInfo() {
      let that = this;
      //到计时 during min *60
      var maxtime = that.unitInfo.during * 60;
      that.time = setInterval(function () {
        if (maxtime >= 0) {
          let minutes = Math.floor(maxtime / 60);
          let seconds = Math.floor(maxtime % 60);
          that.minNum = minutes;
          that.senNum = seconds;
          // let msg = minutes + "分" + seconds + "秒";
          // that.number = msg;
          if (maxtime == 15 * 60) {
            that.$refs.msg.setMsg("注意考试时间，还剩15分钟", "warning");
          }

          --maxtime;
        } else {
          clearInterval(that.time);
          that.$refs.msg.setMsg("考试结束！", "info");
          that.endExam();
        }
      }, 1000);
    },

    // -----------------------face-------------------

    // 启动摄像头
    fnOpen() {
      console.log("执行fnOpen");
      if (typeof window.stream === "object") return;

      clearTimeout(this.timeout);
      this.timeout = setTimeout(() => {
        clearTimeout(this.timeout);

        navigator.mediaDevices
          .getUserMedia(this.constraints)
          .then((result) => {
            this.fuSuccess(result);
          })
          .catch((err) => {
            this.fnError(err);
          });
      }, 3000);
    },
    fnOpenImg() {
      clearTimeout(this.timeoutImg);
      this.timeoutImg = setTimeout(() => {
        clearTimeout(this.timeoutImg);
        this.fnImgRun();
      }, 30000);
    },
    // 调用成功后
    fuSuccess(stream) {
      console.log("执行fuSuccess");
      window.stream = stream;
      this.vi.srcObject = stream;
      this.vi.play();
    },
    // 调用失败后
    fnError(error) {
      console.log(error);
      alert("视频媒体流获取错误:" + error);
    },

    // 关闭摄像头
    fnClose() {
      console.log("执行fnClose");
      this.cav
        .getContext("2d")
        .clearRect(0, 0, this.cav.width, this.cav.height);
      //   this.cav2
      //     .getContext("2d")
      //     .clearRect(0, 0, this.cav2.width, this.cav2.height);
      this.vi.pause(); //暂停
      clearTimeout(this.timeout);
      clearTimeout(this.timeoutImg);
      if (typeof window.stream === "object") {
        window.stream.getTracks().forEach((track) => track.stop());
        window.stream = "";
        this.vi.srcObject = null;
      }
    },

    // 识别
    async fnInit() {
      let that = this;
      console.log("执行fnInit");

      await faceapi.nets["ssdMobilenetv1"].loadFromUri("/models");

      this.options = new faceapi.SsdMobilenetv1Options({
        minConfidence: 0.5, // 0.1 ~ 0.9
      });

      // 连接oss
      that.clientOss = new OSS({
        region: "oss-cn-beijing", //
        accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
        accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
        bucket: "cloud-compe", //OSS 存储区域名
      });

      // 节点属性化  视频人脸追踪
      this.vi = document.getElementById("testVi");
      this.cav = document.getElementById("testCav");
    },

    async fnRun() {
      //   console.log("Run");
      // 识别绘制人脸信息 detectAllFaces检测出现的所有人脸
      let that = this;
      const result = await faceapi["detectAllFaces"](this.vi, this.options);
      //   console.log("res", result);

      if (result && !this.vi.paused) {
        if (result.length > 1) {
          console.log("注意周围环境,有一个人以上！");
          that.record.envirWarning++;
        }
        if (result.length < 1) {
          console.log("摄像头没有检测到考生！");
          that.record.envirWarning++;
        }

        const dims = faceapi.matchDimensions(this.cav, this.vi, true);
        const resizeResults = faceapi.resizeResults(result, dims);

        faceapi.draw.drawDetections(this.cav, resizeResults);
      } else {
        // console.log("进入这个");
        this.cav
          .getContext("2d")
          .clearRect(0, 0, this.cav.width, this.cav.height);
      }
      // 1秒检查一次环境
      this.timeout = setTimeout(() => this.fnRun(), 7000);
    },

    fnImgRun() {
      let that = this;
      console.log("fnImgRun");

      var canvas = document.createElement("canvas");
      canvas.width = that.vi.width;
      canvas.height = that.vi.height;

      canvas
        .getContext("2d")
        .drawImage(that.vi, 0, 0, canvas.width, canvas.height);
      that.src = canvas.toDataURL("image/png");

      let data = that.base64ToFile(that.src, new Date().getTime());
      that.uploadImg(data);
  


      that.axios
        .post(
          "/joinUnitInfo/simOss",
          {
            actvId: that.unitInfo.actvId,
            unitId: that.unitInfo.unitId,
            userCreateId: that.unitInfo.userCreateId,
            userJoinId: that.unitInfo.userJoinId,
            imageUriA:
              "oss://cloud-compe/gradeSelf/" + that.unitInfo.userJoinId,
            imageUriB: "oss://cloud-compe/" + that.name,
          },
          that.config
        )
        .then((res) => {
          console.log("simOss获取到的信息:", res);
          if (res.data.state == "win") {
            if (res.data.simVal) {
              if (res.data.simVal < 60) {
                that.record.simVal++;
              }
            }
            // that.$refs.msg.setMsg(res.data.commont, "success");
          } else {
            // that.$refs.msg.setMsg(res.data.commont, "error");
          }
        })
        .catch((error) => {
          console.log(error);
        });

      // 5s 检查是否本人
      that.timeoutImg = setTimeout(() => that.fnImgRun(), 30000);
    },

    uploadImg(file) {
      let that = this;

      var fileName =
        "examIng/" + that.id.toString() + "/" + new Date().getTime();
      console.log("fileName", fileName, file);
      //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。
      that.clientOss
        .put(fileName, file)
        .then((res) => {
          console.log("clientOss获取到的信息:", res);
          that.name = res.name;
          console.log("name", that.name);
        })
        .catch((error) => {
          console.log(error);
        });
    },

    base64ToFile(dataUrl, name) {
      var arr = dataUrl.split(","),
        // mime = arr[0].match(/:(.*?);/)[1],
        bstr = atob(arr[1]),
        n = bstr.length,
        u8arr = new Uint8Array(n);
      while (n--) {
        u8arr[n] = bstr.charCodeAt(n);
      }
      return new File([u8arr], name + ".png", { type: "image/png" });
    },

    // 且ping
    clickEven(choose) {
      //屏蔽右击
      document.body.oncontextmenu =
        document.body.ondragstart =
        document.body.onselectstart =
        document.body.onbeforecopy =
          function () {
            return choose;
          };
      //屏蔽复制粘贴
      document.body.oncopy = document.body.oncut = function () {
        return choose;
      };
    },
    keyDown(choose) {
      let that = this;
      document.onkeydown = (e) => {
        console.log("e", e);
        let e1 =
          e || event || window.event || arguments.callee.caller.arguments[0];
        //禁用Ctrl+C
        if (e1.ctrlKey && e1.keyCode == 67) {
          return choose;
        }

        //禁用Ctrl+V
        if (e1.ctrlKey && e1.keyCode == 86) {
          return choose;
        }

        //禁用F5
        if (e1.keyCode == 116) {
          return choose;
        }
        //禁用F12
        if (e1.keyCode == 123) {
          return choose;
        }

        //当焦点位置为非文本框时，禁用BackSpace
        //IE下不兼容window.event.target，使用srcElement代替
        if (e1.keyCode == 8) {
          console.log("keyCode", e1.keyCode);
          var ev = that.window.event;
          var obj = ev.srcElement;
          if (
            obj.getAttribute("type") != "text" &&
            obj.getAttribute("type") != "textarea"
          ) {
            return choose;
          }
        }
      };
    },
    pageHidden(e = null, isReduce = 0, router = false) {
      let that = this;
      if (document.visibilityState === "hidden" || router) {
        that.endExam();
      }
    },
    handleScroll(e) {
      this.endExam();
    },
  },
  mounted() {
    let that = this;

    that.clickEven(false);
    that.keyDown(false);
    window.addEventListener("visibilitychange", this.pageHidden);
    window.document.addEventListener("mouseleave", this.handleScroll);

    that.$nextTick(() => {
      that.fnInit();
      that.fnOpen();
      that.fnOpenImg();
    });

    // 开始考试
    that.startExam();
  },
  beforeDestroy() {
    window.removeEventListener("visibilitychange", this.pageHidden);
    window.document.removeEventListener("mouseleave", this.handleScroll);
    this.keyDown(true);
    this.clickEven(true);
    this.fnClose();
    this.endExam();
  },
  watch: {
    blurNum() {
      console.log("blurNum", that.blurNum);
    },
  },
};
</script>
<style lang="less" scoped>
.see {
  position: relative;
  canvas {
    position: absolute;
    top: 0;
    left: 0;
  }
}

.target {
  position: relative;
  img {
    max-width: 600px;
    max-height: 400px;
  }
  canvas {
    position: absolute;
    top: 0;
    left: 0;
  }
}
</style>
