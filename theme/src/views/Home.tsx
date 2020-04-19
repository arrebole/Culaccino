import React from 'react';
import api from "../api";
import * as types from '../types';
import Header from '../components/Header';
import SubNav from "../components/SubNav";
import Stage from '../components/Stage';

export default class Home extends React.Component<{}, { articles: types.Articles | null }> {
    constructor(props: {}) {
        super(props)
        this.state = {
            articles: null
        }
    }

    componentDidMount() {
        api.listArticles().then(res => this.setState({ articles: res }));
    }

    render() {
        return (
            <div style={{ "backgroundColor": "#fcfdfd" }}>
                <Header />
                <SubNav />
                <main >
                    {/* articles 列表 */}
                    <div className="flex flex-col max-w-screen-md mx-auto px-2">
                        {this.state.articles?.articles.map(article => <Stage {...article}/>)}
                    </div>
                </main>
            </div>
        )
    }
}