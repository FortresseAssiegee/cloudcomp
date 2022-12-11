<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12">
        <v-tabs v-model="tab" class="mb-5" right background-color="transparent">
          <v-tabs-slider color="yellow"></v-tabs-slider>
          <v-tab v-for="item in items" :key="item">
            {{ item }}
          </v-tab>
        </v-tabs>
      </v-col>

      <v-col cols="12">
        <v-tabs-items v-model="tab" class="rounded-lg">
          <v-tab-item> <edit-base-info></edit-base-info></v-tab-item>
          <v-tab-item> <editInfoTree></editInfoTree></v-tab-item>
          <v-tab-item> <edit-unit-tree></edit-unit-tree></v-tab-item>
        </v-tabs-items>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import editBaseInfo from "../../../components/edit/editBaseInfo.vue";
import editInfoTree from "../../../components/edit/editInfoTree.vue";
import editUnitTree from "../../../components/edit/editUnitTree.vue";
export default {
  components: {
    editBaseInfo,
    editInfoTree,
    editUnitTree,
  },
  data() {
    return {
      tab: null,
      items: ["基本信息", "详细信息", "教学流程"],
    };
  },
};
</script>
<style lang="less" scoped>
.theme--light.v-tabs-items {
  background-color: transparent !important;
}
</style>
