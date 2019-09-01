<template>
  <div>
    <Header />
    <article v-loading="loading">
      <!-- <section></section> -->
      <section class="cent">
        <Contents :contents="archive.contents" class="markdown-body"></Contents>
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
  computed:{
    paramsID(){
      return parseInt(this.$route.params.id)
    },
    loading(){
      if(this.archive.ID < 0 ) return true
      return false
    }
  },
  methods: {
    async getData() {
      const res: IResp = await api.getArchive(this.paramsID);
      this.archive = res.data.archive;
    },
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
.cent {
  padding: 10px;
}
.markdown-body {
  box-sizing: border-box;
  min-width: 200px;
  max-width: 980px;
  margin: 0 auto;
  padding: 45px;
}

@media (max-width: 767px) {
  .markdown-body {
    padding: 15px;
  }
}
</style>
