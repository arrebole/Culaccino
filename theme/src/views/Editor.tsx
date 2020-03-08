import React from "react"
import Header from "../components/Header"
import Editable from "../components/Editable"
import api from "../api"

interface State {
    title: string,
    type: string,
    cover: string,
    summary: string,
    content: string,
}

function queryString(name: string) {
    var reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)', 'i');
    var r = window.location.search.substr(1).match(reg);
    if (r != null) {
        return unescape(r[2]);
    }
    return "";
}

function createPaper(data: State) {
    return {
        title: data.title,
        cover: data.cover,
        summary: data.summary,
        content: data.content,
        type: "未分类",
        create_at: "",
        update_at: "",
    }
}

export default class EditorPage extends React.Component<{}, State> {
    constructor(props: {}) {
        super(props)
        this.state = {
            title: "",
            type: "",
            cover: "",
            summary: "",
            content: "",
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleCommit = this.handleCommit.bind(this);
    }

    handleChange(e: any, target?: string) {
        switch (target) {
            case "content":
                this.setState({ content: e.currentTarget.innerText })
                break;
            case "summary":
                this.setState({ "summary": e.target.value })
                break;
            case "cover":
                this.setState({ "cover": e.target.value })
                break;
            case "title":
                this.setState({ "title": e.target.value })
                break;
        }
    }

    handeCancel() {
        window.location.reload()
    }

    handleCommit() {
        switch (queryString("tab")) {
            case "":
            case "new":
                api.newPaper(createPaper(this.state)).then(res => alert(res.message))
                break;

            case "update":
                api.updatePaper(createPaper(this.state)).then(res => alert(res.message))
                break;
        }
    }

    componentDidMount() {
        if (queryString("title") !== "") {
            api.fetchPaper(queryString("title")).then(res => {
                this.setState({
                    title: res.data.title,
                    type: res.data.type,
                    cover: res.data.cover,
                    summary: res.data.summary,
                    content: res.data.content,
                })
            })
        }
    }

    render() {
        return (
            <div id="editor" className="bg-gray-light">
                <Header />
                <main className="d-flex flex-column flex-items-center">
                    <div className="py-2" style={{ width: '100%', maxWidth: '980px' }}>
                        <input onChange={e => this.handleChange(e, "title")} value={this.state.title} className="border p-2" placeholder="Title" type="text" style={{ height: '16px' }} />
                    </div>
                    <Editable value={this.state.content} handleChange={this.handleChange} />
                    <div className="bg-white border m-2 d-flex flex-column p-3" style={{ width: '100%', maxWidth: '980px', boxSizing: 'border-box' }}>
                        <h3 className="f3 m-2">Commit changes</h3>
                        <input value={this.state.cover} onChange ={e => this.handleChange(e, "cover")} className="m-2 border p-2 bg-gray-light box-shadow rounded-2 f5" type="text" style={{ outline: 'none' }} placeholder="Cover URL" />
                        <textarea value={this.state.summary} onChange={e => this.handleChange(e, "summary")} className="m-2 border p-2 bg-gray-light box-shadow rounded-2 f5" style={{ minHeight: '100px', outline: 'none' }} placeholder="Add an optional extended description…" />
                    </div>
                    <div>
                        <button className="btn-primary f5 p-2 m-1 rounded-2 border" onClick={this.handleCommit}>Commit changes</button>
                        <button className="btn-danger f5 p-2 m-1 rounded-2 border" onClick={this.handeCancel}>Cancel</button>
                    </div>
                </main>
                <footer className="p-3"></footer>
            </div>
        )
    }
}