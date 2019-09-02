<template>
  <div>
    <Header />
    <article v-loading="loading">
      <!-- <section></section> -->
      <section class="pagehead"></section>
      <section class="repository-content">
        <Contents :contents="archive.contents"></Contents>
      </section>
    </article>
    <Footer />
  </div>
</template>


<script lang="ts">
import Vue from "vue";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import Contents from "../components/Contents.vue";
import api from "../api/index";
import IResp, { IArchive, Archive } from "../types/resp";

export default Vue.extend({
  data(): { archive: IArchive } {
    return {
      archive: new Archive()
    };
  },
  created() {
    this.getData();
  },
  computed: {
    paramsID() {
      return parseInt(this.$route.params.id);
    },
    loading() {
      if (this.archive.ID < 0) return true;
      return false;
    }
  },
  methods: {
    async getData() {
      const res: IResp = await api.getArchive(this.paramsID);
      this.archive = res.data.archive;
    }
  },
  components: {
    Header,
    Footer,
    Contents
  }
});
</script>

<style lang="scss" scoped>
article {
  min-height: 820px;
  background-color: #fff;
}
.repository-content{
  max-width: 1012px;
  widows: 90%;
  margin: auto;
  padding: 16px;
}

@media (min-width: 767px) {
  .pagehead {
    height: 108px;
    background-color: #fafbfc;
    border-bottom: 1px solid #e1e4e8;
  }
}
@media (max-width: 767px) {
  .pagehead {
    height: 58px;
    background-color: #fafbfc;
    border-bottom: 1px solid #e1e4e8;
  }
  .markdown-body {
    padding: 15px;
  }
}
</style>
