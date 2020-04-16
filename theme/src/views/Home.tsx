import React from 'react';
import api from "../api";
import * as types from "../types";
import Header from '../components/Header';


interface State {
    fullArticle: types.FullArticle[] | null
    pagination: types.Pagination | null
}

export default class Home extends React.Component<{}, State> {
    constructor(props: {}) {
        super(props)
        this.state = {
            fullArticle: null,
            pagination: null,
        }
    }

    componentDidMount() {
        api.listArticles().then(res => {
            const data = {
                fullArticle: res.articles,
                pagination: res.pagination
            }
            this.setState(data)
        })
    }

    render() {
        return (
            <div>
                <Header />
            </div>
        )
    }
}