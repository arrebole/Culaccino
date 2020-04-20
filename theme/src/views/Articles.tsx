import React from 'react';
import * as types from "../types";
import * as utils from "../utils";
import Header from '../components/Header';
import Banner from "../components/banner";
import Markdown from "../components/Markdown";
import api from "../api";


export default class Articles extends React.Component<{}, { article: types.Article }> {
    constructor(props: {}) {
        super(props)
        this.state = {
            article: types.defatultArticle()
        }
    }

    componentDidMount() {
        api.getArticle(utils.mathArticleName()).then(res => this.setState({ article: res }))
    }

    render() {
        return (
            <div>
                <Header />
                <main>
                    <article className="max-w-screen-lg mx-auto p-2">
                        <Banner {...this.state.article} />
                        <Markdown contents={this.state.article.contents} />
                    </article>
                </main>
            </div>
        )
    }
}