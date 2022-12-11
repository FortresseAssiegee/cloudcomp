<template>
  <div>
    <v-file-input
      v-model="files"
      color="deep-purple accent-4 "
      :label="tips"
      placeholder="请选择你的头像"
      prepend-icon="mdi-paperclip"
      outlined
      id="avtupload"
      show-size
      @change="uploadAvatars"
      :counter="count"
      :dirPath="dirPath"
      multiple
      :dark="dark"
      :hide-input="icon"
    >
      <!-- v-slot:selection="{ index, text }" -->
      <template v-slot:selection>
        <v-img :src="geturl" max-height="100" max-width="100"> </v-img>
      </template>
    </v-file-input>
  </div>
</template>
<script>
import OSS from "ali-oss";
import { mapState, mapGetters } from "vuex";
export default {
  name: "uploadImg",
  props: {
    // 上传的类型

    // 传上去的地址
    geturl: {
      type: String,
      require: true,
    },
    //上传到oss的地址
    dirPath: {
      type: String,
      default: "/avt",
      require: true,
    },
    // 上传的数量
    count: {
      type: Number,
      default: 1,
    },
    // 提示语言
    tips: {
      type: String,
      default: "上传头像",
    },
    dark: {
      type: Boolean,
      default: false,
    },
    icon: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  data() {
    return {
      // oss
      client: null,
      //   上传的文件
      files: [],

      // 返回的地址
      url: "",
    };
  },
  methods: {
    // 判断类型

    // 上传图片
    uploadAvatars() {
      let that = this;

      // console.log("上传的头像文件:", that.files[0], that.files[0].name);
      let file = that.files[0];
      // if(!(/^\S+\.mp4$/.test(file.name))){
      //   return this.$message.error('请上传视频文件')
      // }
      /**
       * 文件的类型，判断是否是视频
       */
      let param = new FormData();
      param.append("file", file, file.name);
      console.log("开始上传");
      that.put(file.name, file);
    },

    async put(name, file) {
      let that = this;
      // console.log("上传到oss的地址:", that.dirPath);
      try {
        var fileName = that.dirPath + "/" + new Date().getTime() + name;
        // console.log("fileName", fileName, that.dirPath);
        //object-name可以自定义为文件名（例如file.txt）或目录（例如abc/test/file.txt）的形式，实现将文件上传至当前Bucket或Bucket下的指定目录。
        let result = await that.client.put(fileName, file);
        that.geturl = result.url; //返回的上传视频地址
        // console.log("返回的结果：", result);
      } catch (e) {
        console.log(e);
      }
    },

    //
  },
  mounted() {
    let that = this;

    // console.log("子组件接收到的值:", that.geturl);
    that.axios
      .post("/avt/oss", {}, that.config)
      .then((res) => {
        if (res.data.state == "win") {
          // console.log("获取签名:", res);

          // 连接oss
          that.client = new OSS({
            region: res.data.region,
            accessKeyId: res.data.accessKeyId, //阿里云产品的通用id
            accessKeySecret: res.data.accessKeySecret, //密钥
            bucket: res.data.bucket, //OSS 存储区域名
          });
        }
      })
      .catch((error) => {
        console.log(error);
      });
  },
  watch: {
    geturl(val) {
      let that = this;
      // console.log("geturl:", val);
      that.$emit("puturl", val); // 通过 this.$emit() 向父组件传值
    },
  },
};
</script>
<style lang=""></style>
