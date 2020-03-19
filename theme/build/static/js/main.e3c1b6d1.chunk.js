(this.webpackJsonpthemes=this.webpackJsonpthemes||[]).push([[0],{235:function(e,t,a){},236:function(e,t,a){"use strict";a.r(t);var n=a(0),r=a.n(n),l=a(30),c=a.n(l),i=a(12),o=a(10),s=a(2),m=a(3),u=a(7),d=a(6),h=a(8),p=function(){return r.a.createElement("header",{className:"box-shadow-super bg-blue p-3"},r.a.createElement("nav",null,r.a.createElement("h1",{className:"m-0 f3"}," ",r.a.createElement("a",{className:"text-white text-decoration-none",href:"/"},"Culaccino"))))},f=function(e){return r.a.createElement("div",{className:"d-flex m-0 p-4 box-shadow-large rounded-2 bg-white flex-column"},r.a.createElement("div",{className:"text-align-center"},r.a.createElement("div",{className:"my-1"},r.a.createElement(i.b,{className:"h4 text-decoration-none text-blue",to:e.title},e.title)),r.a.createElement("div",{className:"f6 d-flex flex-justify-center text-gray-light"},r.a.createElement("div",{className:"px-1"},"Create on: ")," ",r.a.createElement("time",null,e.create_at.split(" ")[0]),r.a.createElement("div",{className:"px-1"},"|"),r.a.createElement("div",{className:"px-1"},"Update on: ")," ",r.a.createElement("time",null,e.update_at.split(" ")[0]))),r.a.createElement("div",{className:"p-2"},r.a.createElement("img",{style:{width:"100%"},src:e.cover,alt:"img"})),r.a.createElement("p",{className:"f4 text-gray text-align-center"},e.summary))},b=new(function(){function e(){Object(s.a)(this,e)}return Object(m.a)(e,[{key:"fetchPapers",value:function(){arguments.length>0&&void 0!==arguments[0]&&arguments[0];return fetch("/api/papers").then((function(e){return e.json()}))}},{key:"fetchPaper",value:function(e){return fetch("/api/paper?key=".concat(e)).then((function(e){return e.json()}))}},{key:"newPaper",value:function(e){return fetch("/api/new",{method:"POST",body:JSON.stringify(e)}).then((function(e){return e.json()}))}},{key:"updatePaper",value:function(e){return fetch("/api/update",{method:"POST",body:JSON.stringify(e)}).then((function(e){return e.json()}))}}]),e}()),v=function(e){function t(e){var a;return Object(s.a)(this,t),(a=Object(u.a)(this,Object(d.a)(t).call(this,e))).state={papers:[]},a}return Object(h.a)(t,e),Object(m.a)(t,[{key:"componentDidMount",value:function(){var e=this;b.fetchPapers().then((function(t){e.setState({papers:t.data})}))}},{key:"render",value:function(){return this.state.papers.map((function(e){return r.a.createElement("article",{className:"m-3",style:{width:"95%",maxWidth:"900px"},key:e.title},r.a.createElement(f,e))}))}}]),t}(r.a.Component),E=function(e){function t(){return Object(s.a)(this,t),Object(u.a)(this,Object(d.a)(t).apply(this,arguments))}return Object(h.a)(t,e),Object(m.a)(t,[{key:"render",value:function(){return r.a.createElement("div",{id:"home",className:"bg-gray-light"},r.a.createElement(p,null),r.a.createElement("main",{className:"d-flex flex-column flex-items-center"},r.a.createElement(v,null)))}}]),t}(r.a.Component),g=a(23),y=a.n(g),x=a(33),N=a.n(x);y.a.setOptions({highlight:function(e,t){return N.a.highlight(t,e).value}});var j=function(e){function t(e){var a;return Object(s.a)(this,t),(a=Object(u.a)(this,Object(d.a)(t).call(this,e))).state={paper:{title:"Load...",type:"",cover:"",create_at:"Load...",update_at:"Load...",summary:"",content:""}},a}return Object(h.a)(t,e),Object(m.a)(t,[{key:"componentDidMount",value:function(){var e=this,t=window.location.pathname.substring(1);b.fetchPaper(t).then((function(t){t.code<0||e.setState({paper:t.data})}))}},{key:"render",value:function(){return r.a.createElement("div",{id:"paper"},r.a.createElement(p,null),r.a.createElement("div",{className:"box-shadow-large m-4 py-3 border rounded-2 m-conter",style:{maxWidth:"980px",height:"60px"}},r.a.createElement("div",{className:"text-align-center h5"}," ",this.state.paper.title," "),r.a.createElement("div",{className:"f6 d-flex flex-justify-center text-gray-light m-1"},r.a.createElement("div",{className:"px-1"},"Create on: ")," ",r.a.createElement("time",null,this.state.paper.create_at.split(" ")[0]),r.a.createElement("div",{className:"px-1"},"|"),r.a.createElement("div",{className:"px-1"},"Update on: ")," ",r.a.createElement("time",null,this.state.paper.update_at.split(" ")[0]))),r.a.createElement("article",{className:"box-shadow-large m-4 border rounded-2 m-conter",style:{maxWidth:"980px"}},r.a.createElement("section",{className:"py-2 px-4",style:{minHeight:"500px"}},r.a.createElement("div",{className:"markdown-body",dangerouslySetInnerHTML:(e=this.state.paper.content,{__html:y()(e)})}))));var e}}]),t}(r.a.Component),w=a(11),O=function(e){function t(e){var a;return Object(s.a)(this,t),(a=Object(u.a)(this,Object(d.a)(t).call(this,e))).elem=null,a._onChange=a._onChange.bind(Object(w.a)(a)),a}return Object(h.a)(t,e),Object(m.a)(t,[{key:"_onChange",value:function(e){var t=null!=this.elem?this.elem.innerText:"";this.props.handleChange(t,"content")}},{key:"render",value:function(){var e=this;return r.a.createElement("div",{style:{minHeight:"500px",width:"100%",maxWidth:"980px",boxSizing:"border-box",whiteSpace:"pre-line"},className:"bg-white border p-2",ref:function(t){return e.elem=t},dangerouslySetInnerHTML:{__html:this.props.content},onBlur:this._onChange,contentEditable:!0})}}]),t}(r.a.Component);function C(e){var t=new RegExp("(^|&)"+e+"=([^&]*)(&|$)","i"),a=window.location.search.substr(1).match(t);return null!=a?unescape(a[2]):""}function k(e){return{title:e.title,cover:e.cover,summary:e.summary,content:e.content,type:"\u672a\u5206\u7c7b",create_at:"",update_at:""}}var _=function(e){function t(e){var a;return Object(s.a)(this,t),(a=Object(u.a)(this,Object(d.a)(t).call(this,e))).state={title:"",type:"",cover:"",summary:"",content:""},a.handleChange=a.handleChange.bind(Object(w.a)(a)),a.handleCommit=a.handleCommit.bind(Object(w.a)(a)),a}return Object(h.a)(t,e),Object(m.a)(t,[{key:"handleChange",value:function(e,t){switch(t){case"content":this.setState({content:e});break;case"summary":this.setState({summary:e.target.value});break;case"cover":this.setState({cover:e.target.value});break;case"title":this.setState({title:e.target.value})}}},{key:"handeCancel",value:function(){window.location.reload()}},{key:"handleCommit",value:function(){switch(C("tab")){case"":case"new":b.newPaper(k(this.state)).then((function(e){return alert(e.message)}));break;case"update":b.updatePaper(k(this.state)).then((function(e){return alert(e.message)}))}}},{key:"componentDidMount",value:function(){var e=this;""!==C("title")&&b.fetchPaper(C("title")).then((function(t){t.code<0||e.setState({title:t.data.title,type:t.data.type,cover:t.data.cover,summary:t.data.summary,content:t.data.content})}))}},{key:"render",value:function(){var e=this;return r.a.createElement("div",{id:"editor",className:"bg-gray-light"},r.a.createElement(p,null),r.a.createElement("main",{className:"d-flex flex-column flex-items-center"},r.a.createElement("div",{className:"py-2",style:{width:"100%",maxWidth:"980px"}},r.a.createElement("input",{onChange:function(t){return e.handleChange(t,"title")},value:this.state.title,className:"border p-2",placeholder:"Title",type:"text",style:{height:"16px"}})),r.a.createElement(O,{content:this.state.content,handleChange:this.handleChange}),r.a.createElement("div",{className:"bg-white border m-2 d-flex flex-column p-3",style:{width:"100%",maxWidth:"980px",boxSizing:"border-box"}},r.a.createElement("h3",{className:"f3 m-2"},"Commit changes"),r.a.createElement("input",{value:this.state.cover,onChange:function(t){return e.handleChange(t,"cover")},className:"m-2 border p-2 bg-gray-light box-shadow rounded-2 f5",type:"text",style:{outline:"none"},placeholder:"Cover URL"}),r.a.createElement("textarea",{value:this.state.summary,onChange:function(t){return e.handleChange(t,"summary")},className:"m-2 border p-2 bg-gray-light box-shadow rounded-2 f5",style:{minHeight:"100px",outline:"none"},placeholder:"Add an optional extended description\u2026"})),r.a.createElement("div",null,r.a.createElement("button",{className:"btn-primary f5 p-2 m-1 rounded-2 border",onClick:this.handleCommit},"Commit changes"),r.a.createElement("button",{className:"btn-danger f5 p-2 m-1 rounded-2 border",onClick:this.handeCancel},"Cancel"))),r.a.createElement("footer",{className:"p-3"}))}}]),t}(r.a.Component),S=function(e){var t=e.data.map((function(e){return r.a.createElement("li",{key:e.title,className:"d-grid grid-col-10 f5 p-2"},r.a.createElement("div",null,"Title: ",e.title),r.a.createElement("div",null,"Type: ",e.type),r.a.createElement("div",null,"Create: ",e.create_at),r.a.createElement("div",null,"Update: ",e.create_at),r.a.createElement("div",null,r.a.createElement(i.b,{to:"/editor?title=".concat(e.title,"&tab=update")},"\u4fee\u6539")))}));return r.a.createElement("ul",null,t)},P=function(e){function t(e){var a;return Object(s.a)(this,t),(a=Object(u.a)(this,Object(d.a)(t).call(this,e))).state={papers:[]},a}return Object(h.a)(t,e),Object(m.a)(t,[{key:"componentDidMount",value:function(){var e=this;b.fetchPapers().then((function(t){e.setState({papers:t.data})}))}},{key:"render",value:function(){return r.a.createElement("div",{id:"admin",className:"bg-gray-light"},r.a.createElement(p,null),r.a.createElement("main",null,r.a.createElement("div",{className:"d-flex flex-recolumn p-2 m-2"},r.a.createElement(i.b,{to:"/editor?tab=new"},r.a.createElement("button",{className:"btn-primary f5 p-2 m-1 rounded-2 border"},"New Paper"))),r.a.createElement(S,{data:this.state.papers})))}}]),t}(r.a.Component),T=function(){return r.a.createElement(i.a,null,r.a.createElement(o.d,null,r.a.createElement(o.b,{exact:!0,path:"/"}," ",r.a.createElement(E,null),"    "),r.a.createElement(o.b,{path:"/admin"},"  ",r.a.createElement(P,null),"   "),r.a.createElement(o.b,{path:"/editor"}," ",r.a.createElement(_,null),"  "),r.a.createElement(o.b,{path:"/:name"},"  ",r.a.createElement(j,null),"   "),r.a.createElement(o.b,{path:"*"},r.a.createElement(o.a,{to:{pathname:"/"}}))))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));a(232),a(233),a(234),a(235);c.a.render(r.a.createElement(T,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()}))},34:function(e,t,a){e.exports=a(236)}},[[34,1,2]]]);
//# sourceMappingURL=main.e3c1b6d1.chunk.js.map