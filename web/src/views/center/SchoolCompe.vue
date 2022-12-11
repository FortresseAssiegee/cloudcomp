<template>
  <v-container fluid class="py-6 px-8 po">
    <!-- align="center" justify="center" -->

    <v-row class="my-8" justify="center">
      <v-col lg="4" md="6" cols="12" v-for="(item, index) in list" :key="index">
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 12 : 2"
            :class="{ 'on-hover': hover }"
            id="card"
            dark
          >
            <v-img
              height="100%"
              src="https://cdn.vuetifyjs.com/images/cards/server-room.jpg"
            >
              <v-row>
                <v-col justify="center" align="center" class="pa-0" cols="12">
                  <v-avatar class="profile" size="150">
                    <v-img :src="item.avt"></v-img>
                  </v-avatar>
        
                </v-col>
                <v-col class="py-0" >
                  <v-list-item color="rgba(0, 0, 0, .4)" dark>
                    <v-list-item-content>
                      <v-list-item-title class="text-h6">
                        {{ item.userName }}
                      </v-list-item-title>
                      <v-list-item-subtitle>{{
                        item.intro
                      }}</v-list-item-subtitle>
                    </v-list-item-content>
                  </v-list-item>
                </v-col>
              </v-row>
            </v-img>
          </v-card>
        </v-hover>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
import { mapState, mapGetters } from "vuex";
export default {
  data() {
    return {
      list: [],
    };
  },
  methods: {},
  computed: {
    ...mapState(["config"]),
    ...mapGetters(["userInfo"]),
  },
  mounted() {
    let that = this;
    that.axios
      .post(
        "/grade/GetAllGrade",
        {
          num: 20,
        },
        that.config
      )
      .then((res) => {
        console.log("grade/GetAllGrade:", res);
        if (res.data.state == "win") {
          that.list = res.data.userList;
        }
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>
<style lang="less">
.po {
  margin-top: 250px;
}
</style>
