import React from "react";
import Header from "../components/Header";
import * as types from "../types";
import * as utils from "../utils";
import api from "../api";

export default class Edit extends React.Component<{}, { article: types.Article }> {
    constructor(props: {}) {
        super(props)
        this.state = {
            article: types.defatultArticle(),
        }
    }

    handleSubmit() {
        api.updateArticle(this.state.article).then(res => window.location.reload())
    }

    handleChange(
        label: "name" | "cover" | "tag" | "summary" | "contents",
        e: React.ChangeEvent<HTMLInputElement> | React.ChangeEvent<HTMLTextAreaElement>
    ) {
        // 拷贝原有数据
        const newArticle = Object.assign({}, this.state.article);
        // 更新数据
        newArticle[label] = e.target.value;
        // 保存到状态
        this.setState({ article: newArticle });
    }

    componentDidMount() {
        api.getArticle(utils.mathArticleName()).then(res => this.setState({ article: res }))
    }

    render() {
        return (
            <div style={{ "backgroundColor": "#fcfdfd" }}>
                <Header></Header>
                <main >
                    <form className="max-w-screen-md bg-white shadow-md mx-auto pr-5 py-10 mt-12">

                        <div className="md:flex md:items-center mb-6">
                            <div className="md:w-1/6">
                                <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4" htmlFor="article-name">
                                    文章名称
                            </label>
                            </div>
                            <div className="md:w-5/6">
                                <input
                                    className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                                    id="article-name"
                                    type="text"
                                    placeholder="Name"
                                    value={this.state.article.name}
                                    onChange={e => this.handleChange("name", e)}
                                />
                            </div>
                        </div>

                        <div className="md:flex md:items-center mb-6">
                            <div className="md:w-1/6">
                                <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4" htmlFor="article-cover">
                                    封面地址
                            </label>
                            </div>
                            <div className="md:w-5/6">
                                <input className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                                    id="article-cover"
                                    type="text"
                                    placeholder="https://xxx.com/xxx.jpg"
                                    value={this.state.article.cover}
                                    onChange={e => this.handleChange("cover", e)}
                                />
                            </div>
                        </div>

                        <div className="md:flex md:items-center mb-6">
                            <div className="md:w-1/6">
                                <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4" htmlFor="article-tag">
                                    相关标签
                            </label>
                            </div>
                            <div className="md:w-5/6">
                                <input
                                    className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                                    id="article-tag"
                                    placeholder="算法; 数据库; 网络; ..."
                                    type="text"
                                    value={this.state.article.tag}
                                    onChange={e => this.handleChange("tag", e)}
                                />
                            </div>
                        </div>

                        <div className="md:flex md:items-center mb-6">
                            <div className="md:w-1/6"></div>
                            <textarea
                                className="bg-gray-200 w-full py-2 px-4 md:w-5/6"
                                placeholder="summary"
                                value={this.state.article.summary}
                                onChange={e => this.handleChange("summary", e)}
                            >

                            </textarea>
                        </div>

                        <div className="md:flex md:items-center mb-6">
                            <div className="md:w-1/6"></div>
                            <textarea
                                className="bg-gray-200 w-full py-2 px-4 md:w-5/6 h-64 min-h-full"
                                placeholder="contents"
                                value={this.state.article.contents}
                                onChange={e => this.handleChange("contents", e)}
                            ></textarea>
                        </div>

                        <div className="md:flex md:items-center mb-6">
                            <div className="md:w-1/6"></div>
                            <div className="md:w-5/6">
                                <button
                                    className="shadow bg-purple-500 hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded"
                                    type="button"
                                    onClick={e => this.handleSubmit()}
                                >
                                    Create
                            </button>
                            </div>
                        </div>
                    </form>
                </main>
            </div>
        )
    }
}