import React from "react";
import Header from "../components/Header";
import * as types from "../types";
import api from "../api";

const Table: React.FC<{ data: types.Article[] }> = (props) => {
    const tableItem = (item: types.Article) => <tr key={item.name} className="">
        <td className="border border-gray-400 px-4 py-2 text-gray-800" >
            {item.name}
        </td>
        <td className="border border-gray-400 px-4 py-2 text-gray-800">
            {item.tag}
        </td>
        <td className="border border-gray-400 px-4 py-2 text-gray-800">
            {new Date(item.created_at).toLocaleString()}
        </td>
        <td className="border border-gray-400 px-4 py-2 text-gray-800">
            {new Date(item.updated_at).toLocaleString()}
        </td>
        <td className="border border-gray-400 px-4 py-2 text-gray-800">
            <a className="text-blue-700" href={`articles/${item.name}/edit`}>修改</a>
        </td>
    </tr>

    return (
        <table className="border-collapse border-2 border-gray-500 mx-auto">
            <thead>
                <tr>
                    <th className="border border-gray-400 px-4 py-2 text-gray-800">Title</th>
                    <th className="border border-gray-400 px-4 py-2 text-gray-800">Tag</th>
                    <th className="border border-gray-400 px-4 py-2 text-gray-800">Created</th>
                    <th className="border border-gray-400 px-4 py-2 text-gray-800">Updated</th>
                    <th className="border border-gray-400 px-4 py-2 text-gray-800">Done</th>
                </tr>
            </thead>
            <tbody>
                {
                    props.data?.map(item => tableItem(item))
                }
            </tbody>
        </table>
    )
}

export default class Admin extends React.Component<{}, { articles: types.Articles | null }> {
    constructor(props: {}) {
        super(props)
        this.state = {
            articles: null,
        }
    }

    componentDidMount() {
        api.listArticles().then(res => this.setState({ articles: res }));
    }

    render() {
        return (
            <div>
                <Header />
                <main>
                    <div className="max-w-screen-md mx-auto p-2">
                        <a href="/new">
                            <button className="border-2 border-gray-600 hover:border-blue-900 p-2">New Article</button>
                        </a>
                    </div>
                    {
                        this.state.articles && <Table data={this.state.articles.articles} ></Table>
                    }
                </main>
            </div>
        )
    }
}