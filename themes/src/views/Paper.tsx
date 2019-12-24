import React, { useState } from 'react';
import marked from 'marked';
import hljs from 'highlight.js'
import Header from '../components/Header'
import api from "../api"

// marked use highlight Plugin
marked.setOptions({ highlight: (code, lang) => hljs.highlight(lang, code).value });

function createMarkup(data: string) {
    return {__html: marked(data)};
  }

function Main() {
    const [data, setData] = useState<string>("Load...");
    if(data === "Load...") api.getDetail(window.location.pathname).then(res => { setData(res.data)});
    return (
        <article style={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
            <section style={{  maxWidth:"1200px", width:"95%", padding: "20px 10px" }}>
                <div dangerouslySetInnerHTML={createMarkup(data)} ></div>
            </section>
        </article>
    )
}

export default class Paper extends React.Component {
    render() {
        return (
            <div id="paper"><Header /><Main/> </div>
        )
    }
}