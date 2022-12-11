<template>
  <v-container fluid class="py-6 px-8 po">
    <!-- align="center" justify="center" -->

    <v-row class="hotRank my-8" justify="center">
      <v-col
        lg="4"
        md="6"
        cols="12"
        class="position-relative"
        v-for="(item, num) in rankImg"
        :key="num"
      >
        <v-hover v-slot="{ hover }">
          <v-card
            :elevation="hover ? 12 : 2"
            :class="{ 'on-hover': hover }"
            id="card"
          >
            <v-img :src="item.img">
              <v-card-title class="text-h6 white--text">
                <v-row class="fill-height flex-column" justify="space-between">
                  <p class="mt-4 subheading text-left">
                    {{ item.title }}
                  </p>

                  <div>
                    <p
                      class="ma-0 text-body-1 font-weight-bold font-italic text-left"
                    >
                      {{ item.subText }}
                    </p>
                  </div>
                </v-row>
              </v-card-title>
              <v-list class="px-5">
                <v-list-item-group class="border-radius-sm">
                  <v-list-item
                    :ripple="false"
                    v-for="(conversation, index) in conversations"
                    :key="conversation.text"
                    class="px-0 border-radius-sm mb-2"
                    @click="toIntro()"
                  >
                    <v-list-item-avatar>
                      <v-icon class="red darken-4" dark v-if="index == 0">
                        {{ index + 1 }}
                      </v-icon>
                      <v-icon
                        class="orange darken-3"
                        dark
                        v-else-if="index == 1"
                      >
                        {{ index + 1 }}
                      </v-icon>
                      <v-icon
                        class="amber lighten-1"
                        dark
                        v-else-if="index == 2"
                      >
                        {{ index + 1 }}
                      </v-icon>
                      <v-icon class="grey lighten-1" dark v-else>
                        {{ index + 1 }}
                      </v-icon>
                    </v-list-item-avatar>

                    <v-avatar
                      width="48"
                      height="48"
                      class="shadow border-radius-lg me-4"
                    >
                      <v-img
                        :src="conversation.image"
                        alt="Avatar"
                        class="border-radius-lg"
                      />
                    </v-avatar>
                    <v-list-item-content>
                      <div class="d-flex align-center">
                        <div>
                          <h6
                            class="mb-0 text-sm text-sm text-typo font-weight-bold"
                          >
                            {{ conversation.user }}
                          </h6>
                          <p class="mb-0 text-xs text-body font-weight-light">
                            {{ conversation.message }}
                          </p>
                        </div>
                      </div>
                    </v-list-item-content>
                  </v-list-item>
                </v-list-item-group>
              </v-list>
            </v-img>
          </v-card>
        </v-hover>
      </v-col>
    </v-row>
  </v-container>
</template>
<script>
export default {
  data() {
    return {
      conversations: [
        {
          image: require("../../assets/product-11.jpg"),
          user: "Sophie B.",
          message: "Hi! I need more information..",
        },
        {
          image: require("../../assets/product-11.jpg"),
          user: "Anne Marie",
          message: "Awesome work, can you..",
        },
        {
          image: require("../../assets/product-11.jpg"),
          user: "Ivanna",
          message: "About files I can..",
        },
        {
          image: require("../../assets/product-11.jpg"),
          user: "Peterson",
          message: "Have a great afternoon..",
        },
        {
          image: require("../../assets/product-11.jpg"),
          user: "Nick Daniel",
          message: "Hi! I need more information..",
        },
      ],

      rankImg: [
        {
          title: "本周排行榜",
          subText: "站内根据参与度前5名进行排名，每周更新",
          img: "https://images.unsplash.com/photo-1429514513361-8fa32282fd5f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=3264&q=80",
        },
        {
          title: "排行榜",
          subText: "综合企业/学校的所发行的所有赛事的热度之和",
          img: "https://images.unsplash.com/photo-1498038432885-c6f3f1b912ee?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=2100&q=80",
        },
        {
          title: "每日推荐",
          subText: "我们将根据你所收藏的题库，进行推荐",
          img: "https://images.unsplash.com/photo-1542320868-f4d80389e1c4?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=3750&q=80",
        },
      ],
    };
  },
};
</script>
<style lang="less">
.po {
  margin-top: 250px;
}
.hotRank {
  padding-left: 200px;
  padding-right: 200px;
}

#card {
  transition: opacity 0.4s ease-in-out;
}

#card:not(.on-hover) {
  opacity: 0.6;
}

.show-btns {
  color: rgba(255, 255, 255, 1) !important;
}
</style>
