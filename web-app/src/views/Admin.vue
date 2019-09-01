<template>
  <div class="table">
    <Header />
    <article>
      <Table :data="dir" :deleteAt="deleteAt" />
    </article>
    <Footer />
  </div>
</template>



<script lang="ts">
import Vue from "vue";
import IResp, { IArchive } from "../types/resp";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import Table from "../components/Table.vue";
import api from "../api/index";

interface Data {
  dir: IArchive[];
}

export default Vue.extend({
  data(): Data {
    return {
      dir: []
    };
  },
  created() {
    this.fethData();
  },
  methods: {
    async fethData() {
      let res: IResp = await api.getArchiveAll();
      this.dir = res.data.dir.reverse();
    },
    async deleteAt(id: number) {
      let res: IResp = await api.delArchive(id);
      window.location.reload();
    }
  },
  components: {
    Header,
    Footer,
    Table
  }
});
</script>

<style lang="scss" scoped>
article {
  min-height: 780px;
  padding: 20px;
  background: #fafafa;
}
</style>
