<template>
  <v-container>
    <snackerMsg ref="msg"></snackerMsg>
    <v-row class="align-center d-flex h-100 mt-md-n16 pb-0">
      <v-col lg="4" md="8" class="mx-auto">
        <v-card class="card-shadow border-radius-xl py-1">
          <div
            class="card-padding text-center mt-n7 bg-gradient-info shadow-info border-radius-lg mx-4"
          >
            <h4 class="text-h4 font-weight-bolder text-white mb-2">准考证</h4>
            <p class="mb-1 text-white font-weight-light text-sm">
              进入考试之后：不可切换屏幕，鼠标不可离开考卷，不可返回上一页面，否则将会强制交卷。若需调用摄像头，请先检查设备。
            </p>
          </div>

          <div class="px-4 pt-6 pb-1">
            <v-list-item
              class="px-0 py-3 bg-gray-100 border-radius-lg p-6 mb-6"
            >
              <v-list-item-content class="px-6">
                <div class="d-flex flex-column">
                  <h6 class="mb-3 text-sm text-typo font-weight-bold">
                    考生信息：
                  </h6>
                  <span class="mb-2 text-xs text-body"
                    >姓名:
                    <span class="text-dark font-weight-bold ms-sm-2">{{
                      userInfo.userName
                    }}</span></span
                  >
                  <span class="mb-2 text-xs text-body"
                    >联系方式:
                    <span class="text-dark ms-sm-2 font-weight-bold">{{
                      userInfo.tel
                    }}</span></span
                  >
                </div>
              </v-list-item-content>
            </v-list-item>
            <v-list-item
              class="px-0 py-3 bg-gray-100 border-radius-lg p-6 mb-6"
            >
              <v-list-item-content class="px-6">
                <div class="d-flex flex-column">
                  <h6 class="mb-3 text-sm text-typo font-weight-bold">
                    考试信息
                  </h6>
                  <span class="mb-2 text-xs text-body"
                    >考试内容:
                    <span class="text-dark font-weight-bold ms-sm-2"
                      >{{ unitInfo.actvName }}-{{ unitInfo.unitName }}</span
                    ></span
                  >
                  <span class="mb-2 text-xs text-body"
                    >比赛时间:
                    <span class="text-dark font-weight-bold ms-sm-2"
                      >{{ unitInfo.startTime }}到{{ unitInfo.endTime }}</span
                    ></span
                  >
                  <span class="mb-2 text-xs text-body"
                    >考试模式：
                    <span class="text-dark ms-sm-2 font-weight-bold">{{
                      code.get(unitInfo.examCode)
                    }}</span></span
                  >
                  <span class="mb-2 text-xs text-body"
                    >考试时间:
                    <span class="text-dark ms-sm-2 font-weight-bold"
                      >{{ unitInfo.during }}分钟</span
                    ></span
                  >
                  <span class="mb-2 text-xs text-body"
                    >是否开启人脸识别:
                    <span class="text-dark ms-sm-2 font-weight-bold">{{
                      unitInfo.face ? "是" : "否"
                    }}</span></span
                  >
                </div>
              </v-list-item-content>
            </v-list-item>
          </div>
          <div class="text-center">
            <v-btn
              elevation="0"
              :ripple="false"
              height="43"
              class="font-weight-bold text-uppercase btn-default bg-gradient-info py-2 px-6 me-2 mb-2 w-90"
              color="#5e72e4"
              small
              dark
              @click="toRouter"
              :disabled="sure"
              >开始考试</v-btn
            >
            <v-btn
              elevation="0"
              :ripple="false"
              height="43"
              class="font-weight-bold text-uppercase py-2 px-6 me-2 mb-2 w-90"
              small
              @click="backRouter"
              >返回</v-btn
            >
          </div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import OSS from "ali-oss";
import { mapState, mapGetters } from "vuex";
import snackerMsg from "../../components/commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  data() {
    return { checkbox: false, client: null, sure: false };
  },
  computed: {
    ...mapState(["config", "code", "actvId"]),
    ...mapGetters(["userInfo", "actvBaseInfo", "unitInfo", "examInfo"]),
  },
  methods: {
    toRouter() {
      let that = this;
      // console.log("node", that.unitInfo.examCode);

      switch (that.unitInfo.examCode) {
        case 51:
          that.$router.push("/exam/test");
          break;
        case 52:
          that.$router.push("/exam/breakThrough");
          break;
        case 53:
          that.$router.push("/exam/timeLimit");
          break;
        default:
          break;
      }
    },
    backRouter() {
      this.$router.back();
    },
    getGradeSelf() {
      let that = this;
      let path = "gradeSelf/" + that.unitInfo.userJoinId;
      that.client
        .list({
          prefix: path,
          //  delimiter: '/'
        })
        .then((res) => {
          // console.log("getList", res.objects.length);
          if (res.objects.length == 0) {
            that.sure = true;
            that.$refs.msg.setMsg("未上传个人证件照，无法进入考试！", "warning");
          } else {
            that.sure = false;
          }
        })
        .catch((err) => {
          console.log("list catch:", err);
        });
    },
  },
  mounted() {
    let that = this;
    that.client = new OSS({
      region: "oss-cn-beijing", //
      accessKeyId: "LTAI5tGBLNYrTJgL7BjEbzJG", //阿里云产品的通用id
      accessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b", //密钥
      bucket: "cloud-compe", //OSS 存储区域名
      timeout: "1800000", //30min
    });
    that.getGradeSelf();
  },
};
</script>
<style lang=""></style>
