<template>
  <div id="editor">
    <div class="archive-editor-nav">
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
        <label>Summer</label>
        <input type="text" v-model="cache.summary" placeholder="摘要" />
      </div>

      <el-upload
        class="image-uploader"
        action="/api/static/upload"
        :http-request="uploadImg"
        :show-file-list="false"
      >
        <img v-if="imageBlobUrl" :src="imageBlobUrl" class="cover" />
        <i v-else class="el-icon-plus image-uploader-icon"></i>
      </el-upload>

      <button type="submit" @click="submit" class="btn">
        <span>提交</span>
      </button>
    </div>

    <div class="archive-contents">
      <textarea v-model="cache.contents"></textarea>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import api from "../api/index";
import IResp, {
  IArchiveBase,
  IArchive,
  Archive,
  ArchiveBase
} from "../types/resp";

export default Vue.extend({
  props: ["handle", "archive"],
  data(): { cache: IArchiveBase; imageBlobUrl: string } {
    return {
      imageBlobUrl: "",
      cache: new ArchiveBase()
    };
  },
  created() {},
  watch: {
    archive() {
      Object.assign(this.cache, this.archive);
    }
  },
  methods: {
    async submit() {
      if (this.handle == "create") this.apiCreate();
      else if (this.handle == "update") this.apiUpdate();
    },
    async apiCreate() {
      let res: IResp = await api.createArchive(this.cache);
      if (res.code == 0) alert("成功");
    },
    async apiUpdate() {
      let res: IResp = await api.updateArchive(
        this.$route.params.id,
        this.cache
      );
      if (res.code == 0) alert("成功");
    },
    setURL(url:string,f:File){
      this.cache.cover = url;
      this.imageBlobUrl= URL.createObjectURL(f)
    },
    uploadImg(item: { file: File }) {
      api.staticUpload(item.file).then(res =>{
        if (res.code == 0) this.setURL(res.data.file.url,item.file)
      });
      
    }
  }
});
</script>


<style lang="scss" scoped>
.image-uploader {
  border: 1px dashed #4b4949;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.image-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 60px;
  height: 60px;
  line-height: 60px;
  text-align: center;
}

.cover {
  border-radius: 6px;
  width: 80px;
  overflow: hidden;
}
.archive-editor-nav {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-items: center;
  height: 90px;
  width: 80%;
  margin: auto;
}
.archive-contents {
  padding: 30px;
}
.archive-contents > textarea {
  width: 100%;
  min-height: 600px;
}
</style>
