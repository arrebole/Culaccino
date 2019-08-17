<template>
  <div id="editor">
    <div class="archive-info">
      <div>
        <label>Author</label>
        <input type="text" v-model="cache.author" placeholder="作者" />
      </div>
      <div>
        <label>Title</label>
        <input type="text" v-model="cache.title" placeholder="标题" />
      </div>
      <div>
        <label>target</label>
        <input type="text" v-model="cache.target" placeholder="标签" />
      </div>
      <div>
        <label>cover</label>
        <input type="text" v-model="cache.cover" placeholder="封面" />
      </div>
      <div>
        <label>Summer</label>
        <input type="text" v-model="cache.summary" placeholder="摘要" />
      </div>
      <button type="submit" @click="submit" class="btn">
        <span>提交</span>
      </button>
    </div>
    <mavon-editor id="markdown-editor" v-model="cache.contents" />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import api from "../api/index";
import IResp, { IArchiveBase, IArchive, Archive,ArchiveBase } from "../types/resp";

export default Vue.extend({
  props: ["handle", "archive"],
  data(): { cache: IArchiveBase } {
    return {
      cache: new ArchiveBase()
    };
  },
  created() {},
  watch: {
    archive() {
      Object.assign(this.cache,this.archive)
    }
  },
  methods: {
    async submit() {
      if (this.handle == "create") this.apiCreate();
      else if (this.handle == "update") this.apiUpdate();
    },
    async apiCreate() {
      let res: IResp = await api.createArchive(this.cache);
      if (res.code == "success") alert("成功");
    },
    async apiUpdate() {
      let res: IResp = await api.updateArchive(this.$route.params.id, this.cache);
      if (res.code == "success") alert("成功");
    }
  }
});
</script>


<style lang="scss" scoped>
#editor {
  display: flex;
  flex-direction: column;
  height: 100%;
}
#markdown-editor {
  flex: 1;
}
.archive-info {
  padding: 10px;
  margin-bottom: 20px;
  div {
    display: inline;
    padding: 0px 10px;
  }
}
</style>
