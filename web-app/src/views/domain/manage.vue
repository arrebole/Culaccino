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
import IResp, { IArchive } from "../../types/resp";
import Header from "../../components/Header.vue";
import Footer from "../../components/Footer.vue";
import Table from "../../components/Table.vue";
import api from "../../api/index";

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
      let res: IResp = await api.getReposOwner(this.$route.params.domain);
      if (res.data.dir){
        this.dir = res.data.dir.reverse();
      }
    },
    async deleteAt(item:{domain:string,repoName:string}) {
      let res: IResp = await api.delRepo(item) ;
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
