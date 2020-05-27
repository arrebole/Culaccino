import React from "react";
import api from "../api";
import * as types from "../types";
import Header from "../components/Header";
import SubNav from "../components/SubNav";
import Stage from "../components/Stage";
import Pager from "../components/Pager";

interface State {
    articles: types.Articles
    isLoading: boolean
}

export default class Home extends React.Component<{}, State> {
    constructor(props: {}) {
        super(props)
        this.state = {
            articles: types.defatultArticles(),
            isLoading: true
        }
    }

    // 判断是否有更多内容
    get hasMore() {
        return this.state.articles.pagination.remain_size > 0;
    }

    // 计算下一页数据的请求
    get nextOffset() {
        const pagination = this.state.articles.pagination;
        return pagination.total_size - pagination.remain_size;
    }

    // 出发请求下一个数据
    async loadMore() {
        if (this.state.isLoading && !this.hasMore) return;
        const nextArticles = await api.listArticles(this.nextOffset)
        
        // 将获取到的下一页的内容，追加到存储区
        const articles = this.state.articles;
        articles.pagination = nextArticles.pagination;
        articles.articles.push(...nextArticles.articles);

        // 推送到组件状态
        this.setState({ articles })
    }

    componentDidMount() {
        api.listArticles().then(res => this.setState({ articles: res, isLoading: false }));
    }

    render() {
        return (
            <div style={{ "backgroundColor": "#fcfdfd" }}>
                <Header />
                <SubNav />
                <main >
                    <div className="flex flex-col max-w-screen-md mx-auto px-2">
                        {this.state.articles.articles.map(article => <Stage {...article} key={article.name} />)}

                        {this.hasMore && <Pager onClick={ () => this.loadMore() } />}
                    </div>

                </main>
            </div>
        )
    }
}