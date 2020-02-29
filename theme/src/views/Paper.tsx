import React from 'react';
import marked from 'marked';
import hljs from 'highlight.js'
import Header from '../components/Header'
import api from "../api"

// marked use highlight Plugin
marked.setOptions({ highlight: (code, lang) => hljs.highlight(lang, code).value });

function createMarkup(data: string) {
    return { __html: marked(data) };
}

class PaperDetail extends React.Component<{},{content:string}> {
    constructor(props:{}){
        super(props)
        this.state = {
            content: ""
        }
    }
    componentDidMount(){
        const key = window.location.pathname.substring(1);
        api.fetchPaper(key).then(res => { 
            this.setState({ content: res.data.content}) 
        });
    }
    render() {
        return (
            <article className="d-flex flex-column flex-items-center">
                <section className="py-2 px-4" style={{ maxWidth: "980px", width: "95%"}}>
                    <div className="markdown-body" dangerouslySetInnerHTML={createMarkup(this.state.content)} ></div>
                </section>
            </article>
        )
    }
}

export default class Paper extends React.Component {
    render() {
        return (
            <div id="paper"><Header /><PaperDetail /> </div>
        )
    }
}