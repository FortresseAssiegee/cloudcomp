<template>
  <v-container>
    <snackerMsg ref="msg"></snackerMsg>
    <v-data-table
      :headers="headers"
      :items="crsbnListLog"
      hide-default-footer
      class="elevation-1 my-10"
    >
      <template v-slot:top>
        <v-toolbar color="success" dark>
          <div class="card-header-padding">
            <h5 class="font-weight-bold text-h5 text-typo mb-0 white--text">
              登录/注册轮播图
            </h5>
            <p class="text-sm text-body font-weight-light mb-0 white--text">
              上传海报需要是为4/3尺寸
            </p>
          </div>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>

          <div>
            <v-file-input
              v-model="filesLog"
              prepend-icon="mdi-paperclip"
              @change="uploadLog"
              counter="1"
              multiple
              :hide-input="true"
              v-if="!showProLog"
            >
            </v-file-input>
            <v-icon @click="uploadCanncel" dark v-else>mdi-cancel </v-icon>
          </div>
        </v-toolbar>
        <v-progress-linear
          color="success"
          :value="progressLog"
          v-show="progressLog"
        ></v-progress-linear>
      </template>
      <template v-slot:item.url="{ item }">
        <v-img width="150px" class="my-2" :src="item.url"></v-img>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-icon small @click="delOne(item.name)"> mdi-delete </v-icon>
      </template>
    </v-data-table>
    <!--  -->
    <v-data-table
      :headers="headers"
      :items="crsbnList"
      hide-default-footer
      class="elevation-1 my-10"
    >
      <template v-slot:top>
        <v-toolbar color="warning" dark>
          <div class="card-header-padding">
            <h5 class="font-weight-bold text-h5 text-typo mb-0 white--text">
              轮播图管理
            </h5>
            <p class="text-sm text-body font-weight-light mb-0 white--text">
              上传海报需要是为16/9尺寸,上传海报前先输入海波相绑定的
              <b>活动ID </b>
            </p>
          </div>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-text-field
            label="活动ID"
            hint="请输入活动的Id号"
            hide-details="auto"
            v-model="crsbnId"
            color="white"
          ></v-text-field>
          <div>
            <v-file-input
              v-model="files"
              prepend-icon="mdi-paperclip"
              @change="upload"
              counter="1"
              multiple
              :hide-input="true"
              v-if="!showPro"
            >
            </v-file-input>
            <v-icon @click="uploadCanncel" dark v-else>mdi-cancel </v-icon>
          </div>
        </v-toolbar>
        <v-progress-linear
          color="warning"
          :value="progress"
          v-show="progress"
        ></v-progress-linear>
      </template>
      <template v-slot:item.url="{ item }">
        <v-img width="150px" class="my-2" :src="item.url"></v-img>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-icon small @click="delOne(item.name)"> mdi-delete </v-icon>
      </template>
    </v-data-table>
  </v-container>
</template>
<script>
import OSS from "ali-oss";
import snackerMsg from "../../../components/commont/snackerMsg.vue";
export default {
  components: {
    snackerMsg,
  },
  data() {
    return {
      // 登录管理
      progressLog: 0,
      showProLog: false,

      filesLog: [], //上传文件
      crsbnListLog: [], //轮播图列表

      // 首页管理
      headers: [
        {
          text: "文件名称",
          value: "name",
        },
        { text: "路径", value: "url" },
        // { text: "活动结束时间", value: "endTime" },
        { text: "时间", value: "lastModified" }, //20进行中 21结束 22未发布

        { text: "操作", value: "actions", sortable: false },
      ],
      progress: 0,
      showPro: false,

      files: [], //上传文件
      crsbnList: [], //轮播图列表

      crsbnId: "", //id活动号

      client: null,
    };
  },
  methods: {
    // 上传图片
    upload() {
      let that = this;
      if (that.crsbnId == "" || that.crsbnList == null) {
        return that.$refs.msg.setMsg("请输入海报要绑定的Id", "error");
      }
      //   判断类型进行分流
      // console.log("上传的文件:", that.files[0], that.files[0].name);
      let file = that.files[0];

      if (
        !/^\S+\.png$/.test(file.name) ||
        !!/^\S+\.jpe{0,1}g$/.test(file.name)
      ) {
        return that.$refs.msg.setMsg("请上传png或jpeg或jpg格式", "error");
      }

      that.showPro = true;
      that.$refs.msg.setMsg("开始上传", "success");

      let param = new FormData();
      param.append("file", file, file.name);
      that.put(file.name, file);
    },

    uploadLog() {
      let that = this;

      //   判断类型进行分流
      // console.log("上传的文件:", that.filesLog[0], that.filesLog[0].name);
      let file = that.filesLog[0];

      if (
        !/^\S+\.png$/.test(file.name) ||
        !!/^\S+\.jpe{0,1}g$/.test(file.name)
      ) {
        return that.$refs.msg.setMsg("请上传png或jpeg或jpg格式", "error");
      }

      that.showProLog = true;
      that.$refs.msg.setMsg("开始上传", "success");

      let param = new FormData();
      param.append("file", file, file.name);
      that.putLog(file.name, file);
    },

    async put(name, file) {
      let that = this;

      try {
        var dirName = "/crsbn/" + that.crsbnId + "/" + name;

        // console.log("dirName", dirName);
        //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。

        // 指定待删除的Object的名称。Object名称需填写不包含Bucket名称在内的Object的完整路径，例如exampledir/exampleobject.txt。
        await that.client
          .multipartUpload(dirName, file, {
            ...that.option,
          })
          .then((result) => {
            // console.log("multipartUpload res", result);

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
      }
    },

    async putLog(name, file) {
      let that = this;

      try {
        var dirName = "/LogCrsbn/" + name;

        // console.log("dirName", dirName);
        //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。

        // 指定待删除的Object的名称。Object名称需填写不包含Bucket名称在内的Object的完整路径，例如exampledir/exampleobject.txt。
        await that.client
          .multipartUpload(dirName, file, {
            ...that.option,
          })
          .then((result) => {
            // console.log("multipartUpload res", result);

            that.progressLog = 0;

            that.filesLog = null;
            that.showProLog = false;
            that.getListLog();
          })
          .catch((err) => {
            console.log("multipartUpload catch:", err);
          });
      } catch (e) {
        console.log(e);
      }
    },
    async getList() {
      let that = this;
      // path 前面不要加"/"
      that.client
        .list({
          prefix: "crsbn",
          //    delimiter: '/'
        })
        .then((res) => {
          that.crsbnList = [];
          for (let i = 1; i < res.objects.length; i++) {
            const element = res.objects[i];
            element.url = that.client.signatureUrl(element.name);
            that.crsbnList.push(res.objects[i]);
          }
          // console.log("list res", res);
        })
        .catch((err) => {
          console.log("list catch:", err);
        });
    },

    async getListLog() {
      let that = this;
      // path 前面不要加"/"
      that.client
        .list({
          prefix: "LogCrsbn",
          // delimiter: "/",
        })
        .then((res) => {
          that.crsbnListLog = [];
          for (let i = 0; i < res.objects.length; i++) {
            const element = res.objects[i];
            element.url = that.client.signatureUrl(element.name);
            that.crsbnListLog.push(res.objects[i]);
          }
          // console.log("Log list res", res);
        })
        .catch((err) => {
          console.log("Log list catch:", err);
        });
    },

    async delOne(name) {
      let that = this;
      try {
        // 填写Object完整路径。Object完整路径中不能包含Bucket名称。
        let result = await that.client.delete(name);
        // console.log("delOne getName", result);
        that.getList();
      } catch (error) {
        console.log("delOne getName", error);
      }
    },
    uploadCanncel() {
      let that = this;
      that.client.cancel();
      that.showPro = false;
      that.$refs.msg.setMsg("取消上传", "success");
    },
  },
  mounted() {
    let that = this;
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
    };
    that.getList();
    that.getListLog();
  },
  watch: {
    progress(val) {
      let that = this;
      // console.log("pro", val);
      if (val * 1 == 100) {
        that.showPro = false;
      }
    },
    progressLog(val) {
      let that = this;
      // console.log("pro", val);
      if (val * 1 == 100) {
        that.showProLog = false;
      }
    },
  },
};
</script>
<style lang="less"></style>
