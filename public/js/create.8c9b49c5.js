(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["create"],{"242f":function(t,e,r){"use strict";var a=r("4ff9"),n=r.n(a);n.a},"4ff9":function(t,e,r){},"5f53":function(t,e,r){},b2ae:function(t,e,r){"use strict";var a=r("5f53"),n=r.n(a);n.a},ceb0:function(t,e,r){"use strict";var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{attrs:{id:"editor"}},[r("div",{staticClass:"article-info"},[r("div",[r("label",[t._v("Author")]),r("input",{directives:[{name:"model",rawName:"v-model",value:t.author,expression:"author"}],attrs:{type:"text",placeholder:"作者"},domProps:{value:t.author},on:{input:function(e){e.target.composing||(t.author=e.target.value)}}})]),r("div",[r("label",[t._v("Title")]),r("input",{directives:[{name:"model",rawName:"v-model",value:t.title,expression:"title"}],attrs:{type:"text",placeholder:"标题"},domProps:{value:t.title},on:{input:function(e){e.target.composing||(t.title=e.target.value)}}})]),r("div",[r("label",[t._v("target")]),r("input",{directives:[{name:"model",rawName:"v-model",value:t.target,expression:"target"}],attrs:{type:"text",placeholder:"标签"},domProps:{value:t.target},on:{input:function(e){e.target.composing||(t.target=e.target.value)}}})]),r("div",[r("label",[t._v("cover")]),r("input",{directives:[{name:"model",rawName:"v-model",value:t.cover,expression:"cover"}],attrs:{type:"text",placeholder:"封面"},domProps:{value:t.cover},on:{input:function(e){e.target.composing||(t.cover=e.target.value)}}})]),r("div",[r("label",[t._v("Summer")]),r("input",{directives:[{name:"model",rawName:"v-model",value:t.summary,expression:"summary"}],attrs:{type:"text",placeholder:"摘要"},domProps:{value:t.summary},on:{input:function(e){e.target.composing||(t.summary=e.target.value)}}})]),r("button",{staticClass:"btn",attrs:{type:"submit"},on:{click:t.submit}},[r("span",[t._v("提交")])])]),r("mavon-editor",{attrs:{id:"markdown-editor"},model:{value:t.contents,callback:function(e){t.contents=e},expression:"contents"}})],1)},n=[],i=(r("96cf"),r("3b8d")),s=r("2b0e"),o=r("79f6"),c=s["default"].extend({props:["handle","article"],data:function(){return{ID:0,author:"",title:"",target:"",summary:"",contents:"",cover:""}},created:function(){},watch:{article:function(){this.author=this.article.author,this.title=this.article.title,this.target=this.article.target,this.summary=this.article.summary,this.contents=this.article.contents,this.cover=this.article.cover}},computed:{},methods:{submit:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(){return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:"create"==this.handle?this.apiCreate():"update"==this.handle&&this.apiUpdate();case 1:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}(),apiCreate:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(){var e,r;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,o["a"].githubRenderAPI({text:this.contents});case 2:return e=t.sent,t.t0=o["a"],t.next=6,this.createArticle(e);case 6:return t.t1=t.sent,t.next=9,t.t0.createArticle.call(t.t0,t.t1);case 9:r=t.sent,"success"==r.code&&alert("成功");case 11:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}(),apiUpdate:function(){var t=Object(i["a"])(regeneratorRuntime.mark(function t(){var e,r;return regeneratorRuntime.wrap(function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,o["a"].githubRenderAPI({text:this.contents});case 2:return e=t.sent,t.next=5,o["a"].updateArticle(this.article.ID,this.createArticle(e));case 5:r=t.sent,"success"==r.code&&alert("成功");case 7:case"end":return t.stop()}},t,this)}));function e(){return t.apply(this,arguments)}return e}(),createArticle:function(t){return{title:this.title,author:this.author,target:this.target,cover:this.cover,summary:this.summary,contents:t}}}}),u=c,l=(r("242f"),r("2877")),d=Object(l["a"])(u,a,n,!1,null,"914817a2",null);e["a"]=d.exports},d879:function(t,e,r){"use strict";r.r(e);var a=function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",[r("Header"),r("article",[r("section",[r("Editor",{attrs:{handle:"create"}})],1)]),r("Footer")],1)},n=[],i=r("2b0e"),s=r("0418"),o=r("ceb0"),c=r("fd2d"),u=i["default"].extend({data:function(){return{}},methods:{},components:{Header:s["a"],Footer:c["a"],Editor:o["a"]}}),l=u,d=(r("b2ae"),r("2877")),p=Object(d["a"])(l,a,n,!1,null,"7bcdf6b3",null);e["default"]=p.exports}}]);
//# sourceMappingURL=create.8c9b49c5.js.map