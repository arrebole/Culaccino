import React from 'react';
import api from "../api";
import * as types from '../types';
import Header from '../components/Header';
import SubNav from "../components/SubNav";
import Stage from '../components/Stage';

// State Home组件状态类型，文章列表和分页
interface State {
    articles: types.Articles | null
}

export default class Home extends React.Component<{}, State> {
    constructor(props: {}) {
        super(props)
        this.state = {
            articles: null
        }
    }

    // fetchData 获取首页数据
    fetchData() {
        return api.listArticles()
    }

    componentDidMount() {
        this.fetchData().then(res => this.setState({ articles: res }));
    }

    render() {
        return (
            <div style={{ "backgroundColor": "#fcfdfd" }}>
                <Header />
                <SubNav />
                <main >
                    {/* articles 列表 */}
                    <div className="flex flex-col max-w-screen-md mx-auto px-2">
                        {this.state.articles?.articles.map(article => Stage(article))}
                    </div>
                </main>
            </div>
        )
    }
}