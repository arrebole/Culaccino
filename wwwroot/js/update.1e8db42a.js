(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["update"],{"51e5":function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("Header"),a("article",[a("section",[a("Editor",{attrs:{handle:"update",archive:e.archive}})],1)]),a("Footer")],1)},n=[],c=(a("96cf"),a("3b8d")),i=a("2b0e"),s=a("0418"),o=a("ceb0"),u=a("fd2d"),l=a("79f6"),d=i["default"].extend({name:"Update",data:function(){return{archive:null}},created:function(){this.getData()},methods:{getData:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,l["a"].getRepo({domain:this.$route.params.domain,repoName:this.$route.params.repo});case 2:t=e.sent,this.archive=t.data.archive;case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()},components:{Header:s["a"],Footer:u["a"],Editor:o["a"]}}),p=d,h=(a("9dd8"),a("2877")),m=Object(h["a"])(p,r,n,!1,null,"40fd83b0",null);t["default"]=m.exports},"9dd8":function(e,t,a){"use strict";var r=a("fdd7"),n=a.n(r);n.a},aae7:function(e,t,a){"use strict";var r=a("d9bc"),n=a.n(r);n.a},ceb0:function(e,t,a){"use strict";var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{attrs:{id:"editor"}},[a("div",{staticClass:"archive-editor-nav"},[a("div",[a("label",[e._v("Repositly")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.title,expression:"cache.title"}],attrs:{type:"text",placeholder:"仓库名"},domProps:{value:e.cache.title},on:{input:function(t){t.target.composing||e.$set(e.cache,"title",t.target.value)}}})]),a("div",[a("label",[e._v("Area")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.area,expression:"cache.area"}],attrs:{type:"text",placeholder:"分区"},domProps:{value:e.cache.area},on:{input:function(t){t.target.composing||e.$set(e.cache,"area",t.target.value)}}})]),a("div",[a("label",[e._v("Summer")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.summary,expression:"cache.summary"}],attrs:{type:"text",placeholder:"摘要"},domProps:{value:e.cache.summary},on:{input:function(t){t.target.composing||e.$set(e.cache,"summary",t.target.value)}}})]),a("el-upload",{staticClass:"image-uploader",attrs:{action:"/api/static/upload","http-request":e.uploadImg,"show-file-list":!1}},[e.imageBlobUrl?a("img",{staticClass:"cover",attrs:{src:e.imageBlobUrl}}):a("i",{staticClass:"el-icon-plus image-uploader-icon"})]),a("button",{staticClass:"btn",attrs:{type:"submit"},on:{click:e.submit}},[a("span",[e._v("提交")])])],1),a("div",{staticClass:"archive-contents"},[a("textarea",{directives:[{name:"model",rawName:"v-model",value:e.cache.contents,expression:"cache.contents"}],domProps:{value:e.cache.contents},on:{input:function(t){t.target.composing||e.$set(e.cache,"contents",t.target.value)}}})])])},n=[],c=(a("96cf"),a("3b8d")),i=a("2b0e"),s=a("79f6"),o=a("4dd8"),u=i["default"].extend({props:["handle","archive"],data:function(){return{imageBlobUrl:"",cache:new o["b"]}},created:function(){},watch:{archive:function(){Object.assign(this.cache,this.archive)}},methods:{submit:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:"create"==this.handle?this.apiCreate():"update"==this.handle&&this.apiUpdate();case 1:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),apiCreate:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,s["a"].newRepo(this.cache);case 2:t=e.sent,0==t.code&&alert("成功");case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),apiUpdate:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,s["a"].commitRepo({domain:this.$route.params.domain,repoName:this.cache.title},this.cache);case 2:t=e.sent,0==t.code&&alert("成功");case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),setURL:function(e,t){this.cache.cover=e,this.imageBlobUrl=URL.createObjectURL(t)},uploadImg:function(e){var t=this;s["a"].staticUpload(e.file).then(function(a){0==a.code&&t.setURL(a.data.file.url,e.file)})}}}),l=u,d=(a("aae7"),a("2877")),p=Object(d["a"])(l,r,n,!1,null,"3ba5991b",null);t["a"]=p.exports},d9bc:function(e,t,a){},fdd7:function(e,t,a){}}]);
//# sourceMappingURL=update.1e8db42a.js.map