<template>
  <div>
   
   

    <template >
      <v-footer color="transparent" >
        <v-card class="flex" color="transparent">
          <v-card-text class="px-0">
            <v-row>
              <v-col cols="12" md="6">
                <div class="copyright text-body-2 ls-0 text-muted">
                  © {{ new Date().getFullYear() }}, made with
                  <i class="fa fa-heart" aria-hidden="true"></i> by
                  <a
                    href="https://www.creative-tim.com"
                    class="
                      text-decoration-none text-typo text-body-2
                      ls-0
                      font-weight-bold
                    "
                    target="_blank"
                    >Creative Tim</a
                  >
                  for a better web.
                </div>
              </v-col>

              <v-col cols="10" md="6" class="d-flex justify-end">
                <ul
                  v-for="item in footer"
                  :key="item.linkName"
                  class="d-flex list-style-none"
                >
                  <li>
                    <a
                      :href="item.link"
                      class="
                        text-decoration-none text-muted text-body-2
                        ls-0
                        btn-dark-hover
                        no-default-hover
                      "
                      target="_blank"
                      >{{ item.linkName }}</a
                    >
                  </li>
                </ul>
              </v-col>
            </v-row>
          </v-card-text>
        </v-card>
      </v-footer>
    </template>
  </div>
</template>
<script>
export default {
  name: "content-footer",
  props: {
    auth: Boolean,
  },
  data() {
    return {
      footer: [
        {
          linkName: "Creative Tim",
          link: "https://www.creative-tim.com/",
        },
        {
          linkName: "About Us",
          link: "https://www.creative-tim.com/presentation",
        },
        {
          linkName: "Blog",
          link: "http://blog.creative-tim.com/",
        },
        {
          linkName: "License",
          link: "https://www.creative-tim.com/license",
        },
      ],
      icons: [
        {
          icon: "fab fa-dribbble",
        },
        {
          icon: "fab fa-twitter",
        },
        {
          icon: "fab fa-instagram",
        },
        {
          icon: "fab fa-pinterest",
        },
        {
          icon: "fab fa-github",
        },
      ],
    };
  },
};
</script>
