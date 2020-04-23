import Http from "./http";
import Fetcher from "./fetch";
import * as types from "../types";

class Api {

    constructor(strategy: Http) {
        this.httpStrategy = strategy;
    }

    // 策略模式，可替换的实现http方法接口的对象
    private httpStrategy: Http;

    getArticle(name: string) {
        return this.httpStrategy
            .GET<types.Article>(`/api/articles/${name}`)
    }

    deleteArticle(name: string) {
        return this.httpStrategy
            .DELETE<types.Article>(`/api/articles/${name}`)
    }

    listArticles(offset = 0, page_size = 5) {
        return this.httpStrategy
            .GET<types.Articles>(`/api/articles?offset=${offset}&page_size=${page_size}`)
    }

    createArticle(article: types.Article) {
        return this.httpStrategy
            .POST<types.Articles>(`/api/articles`, JSON.stringify({
                name: article.name,
                tag: article.tag,
                cover: article.cover,
                summary: article.summary,
                contents: article.contents
            }))
    }

    updateArticle(article: types.Article) {
        return this.httpStrategy
            .PATCH<types.Articles>(`/api/articles/${article.name}`, JSON.stringify({
                tag: article.tag,
                cover: article.cover,
                summary: article.summary,
                contents: article.contents
            }))
    }
}

export default new Api(new Fetcher());