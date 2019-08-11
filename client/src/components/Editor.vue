<template>
  <div id="editor">
    <div class="article-info">
      <div>
        <label>Author</label>
        <input type="text" v-model="author" placeholder="作者" />
      </div>
      <div>
        <label>Title</label>
        <input type="text" v-model="title" placeholder="标题" />
      </div>
      <div>
        <label>target</label>
        <input type="text" v-model="target" placeholder="标签" />
      </div>
      <div>
        <label>cover</label>
        <input type="text" v-model="cover" placeholder="封面" />
      </div>
      <div>
        <label>Summer</label>
        <input type="text" v-model="summary" placeholder="摘要" />
      </div>
      <button type="submit" @click="submit" class="btn">
        <span>提交</span>
      </button>
    </div>
    <mavon-editor
      id="markdown-editor"
      v-model="contents"
    />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import api from "../api/index";
import IResp, { IArticleBase } from "../types/resp";

export default Vue.extend({
  props: ["handle", "article"],
  data() {
    return {
      ID: 0,
      author: "",
      title: "",
      target: "",
      summary: "",
      contents: "",
      cover:"",
    };
  },
  created() {},
  watch: {
    article() {
      this.author = this.article.author;
      this.title = this.article.title;
      this.target = this.article.target;
      this.summary = this.article.summary;
      this.contents = this.article.contents;
      this.cover = this.article.cover;
    }
  },
  computed: {},
  methods: {
    async submit() {
      if (this.handle == "create") this.apiCreate();
      else if (this.handle == "update") this.apiUpdate();
    },
    async apiCreate() {
      let md = await api.githubRenderAPI({ text: this.contents });
      let res: IResp = await api.createArticle(await this.createArticle(md));
      if (res.code == "success") alert("成功");
    },
    async apiUpdate() {
      let md = await api.githubRenderAPI({ text: this.contents });
      let res: IResp = await api.updateArticle(
        this.article.ID,
        this.createArticle(md)
      );
      if (res.code == "success") alert("成功");
    },
    createArticle(contents: string): IArticleBase {
      return {
        title: this.title,
        author: this.author,
        target: this.target,
        cover: this.cover,
        summary: this.summary,
        contents
      };
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
.article-info {
  padding: 10px;
  margin-bottom: 20px;
  div {
    display: inline;
    padding: 0px 10px;
  }
}
</style>
