(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["update"],{"233c":function(e,t,a){"use strict";var r=a("42ba"),n=a.n(r);n.a},"2e8c":function(e,t,a){"use strict";a.r(t);var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("Header"),a("article",[a("section",[a("Editor",{attrs:{handle:"update",archive:e.archive}})],1)]),a("Footer")],1)},n=[],c=(a("96cf"),a("3b8d")),i=a("2b0e"),s=a("0418"),o=a("ceb0"),u=a("fd2d"),l=a("79f6"),h=i["default"].extend({name:"Update",data:function(){return{archive:null}},created:function(){this.getData()},methods:{getData:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){var t,a;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return t=parseInt(this.$route.params.articleID),e.next=3,l["a"].getArchive(t);case 3:a=e.sent,this.archive=a.archive;case 5:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()},components:{Header:s["a"],Footer:u["a"],Editor:o["a"]}}),d=h,p=(a("e699"),a("2877")),v=Object(p["a"])(d,r,n,!1,null,"0969d894",null);t["default"]=v.exports},"42ba":function(e,t,a){},"7c6f":function(e,t,a){},ceb0:function(e,t,a){"use strict";var r=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{attrs:{id:"editor"}},[a("div",{staticClass:"archive-info"},[a("div",[a("label",[e._v("Author")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.author,expression:"cache.author"}],attrs:{type:"text",placeholder:"作者"},domProps:{value:e.cache.author},on:{input:function(t){t.target.composing||e.$set(e.cache,"author",t.target.value)}}})]),a("div",[a("label",[e._v("Title")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.title,expression:"cache.title"}],attrs:{type:"text",placeholder:"标题"},domProps:{value:e.cache.title},on:{input:function(t){t.target.composing||e.$set(e.cache,"title",t.target.value)}}})]),a("div",[a("label",[e._v("target")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.target,expression:"cache.target"}],attrs:{type:"text",placeholder:"标签"},domProps:{value:e.cache.target},on:{input:function(t){t.target.composing||e.$set(e.cache,"target",t.target.value)}}})]),a("div",[a("label",[e._v("cover")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.cover,expression:"cache.cover"}],attrs:{type:"text",placeholder:"封面"},domProps:{value:e.cache.cover},on:{input:function(t){t.target.composing||e.$set(e.cache,"cover",t.target.value)}}})]),a("div",[a("label",[e._v("Summer")]),a("input",{directives:[{name:"model",rawName:"v-model",value:e.cache.summary,expression:"cache.summary"}],attrs:{type:"text",placeholder:"摘要"},domProps:{value:e.cache.summary},on:{input:function(t){t.target.composing||e.$set(e.cache,"summary",t.target.value)}}})]),a("button",{staticClass:"btn",attrs:{type:"submit"},on:{click:e.submit}},[a("span",[e._v("提交")])])]),a("mavon-editor",{attrs:{id:"markdown-editor"},model:{value:e.cache.contents,callback:function(t){e.$set(e.cache,"contents",t)},expression:"cache.contents"}})],1)},n=[],c=(a("96cf"),a("3b8d")),i=a("2b0e"),s=a("79f6"),o=a("4dd8"),u=i["default"].extend({props:["handle","archive"],data:function(){return{cache:new o["b"]}},created:function(){},watch:{archive:function(){Object.assign(this.cache,this.archive)}},methods:{submit:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:"create"==this.handle?this.apiCreate():"update"==this.handle&&this.apiUpdate();case 1:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),apiCreate:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,s["a"].createArchive(this.archive);case 2:t=e.sent,"success"==t.code&&alert("成功");case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}(),apiUpdate:function(){var e=Object(c["a"])(regeneratorRuntime.mark(function e(){var t;return regeneratorRuntime.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,s["a"].updateArchive(this.$route.params.id,this.cache);case 2:t=e.sent,"success"==t.code&&alert("成功");case 4:case"end":return e.stop()}},e,this)}));function t(){return e.apply(this,arguments)}return t}()}}),l=u,h=(a("233c"),a("2877")),d=Object(h["a"])(l,r,n,!1,null,"7441b877",null);t["a"]=d.exports},e699:function(e,t,a){"use strict";var r=a("7c6f"),n=a.n(r);n.a}}]);
//# sourceMappingURL=update.bc0a40d3.js.map