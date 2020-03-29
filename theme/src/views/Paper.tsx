import React from 'react';
import marked from 'marked';
import hljs from 'highlight.js'
import Header from '../components/Header'
import api,{Paper} from "../api"

// marked use highlight Plugin
marked.setOptions({ highlight: (code, lang) => hljs.highlight(lang, code).value });

function createMarkup(data: string) {
    return { __html: marked(data) };
}

export default class PaperPage extends React.Component<{}, {paper: Paper}> {
    constructor(props: {}) {
        super(props)
        this.state = {
            paper: {
                title: "Load...",
                type: "",
                cover: "",
                create_at: "Load...",
                update_at: "Load...",
                summary: "",
                content: ""
            }
        }
    }
    componentDidMount() {
        const pathSplit = window.location.pathname.split("/");
        const key = pathSplit[pathSplit.length-1];
        api.fetchPaper(key).then(res => {
            this.setState({ paper: res })
        });
    }
    render() {
        return (
            <div id="paper">
                <Header />

                <div className="box-shadow-large m-4 py-3 border rounded-2 m-conter" style={{ maxWidth: "980px",height: '60px' }}>
                    <div className="text-align-center h5"> { this.state.paper.title } </div>

                    <div className="f6 d-flex flex-justify-center text-gray-light m-1">
                        <div className="px-1" >Create on: </div> <time>{this.state.paper.create_at.split(" ")[0]}</time>
                        <div className="px-1" >|</div> 
                        <div className="px-1" >Update on: </div> <time>{this.state.paper.update_at.split(" ")[0]}</time>
                    </div>
                </div>

                <article className="box-shadow-large m-4 border rounded-2 m-conter" style={{ maxWidth: "980px" }}>
                    <section className="py-2 px-4" style={{ minHeight: "500px" }}>
                        <div className="markdown-body" dangerouslySetInnerHTML={createMarkup(this.state.paper.content)} ></div>
                    </section>
                </article>
            </div>
        )
    }
}