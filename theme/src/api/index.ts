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

    deleteArticle(name: string, accessToken: string) {
        return this.httpStrategy
            .DELETE<types.Article>(`/api/articles/${name}`, { Authorization: accessToken })
    }

    listArticles(offset = 0, page_size = 5) {
        return this.httpStrategy
            .GET<types.Articles>(`/api/articles?offset=${offset}&page_size=${page_size}`)
    }

    createArticle(article: types.Article, accessToken: string) {
        return this.httpStrategy
            .POST<types.Articles>(`/api/articles`,
                JSON.stringify({
                    name: article.name,
                    tag: article.tag,
                    cover: article.cover,
                    summary: article.summary,
                    contents: article.contents
                }),
                { Authorization: accessToken }
            )
    }

    updateArticle(article: types.Article, accessToken: string) {
        return this.httpStrategy
            .PATCH<types.Articles>(`/api/articles/${article.name}`,
                JSON.stringify({
                    tag: article.tag,
                    cover: article.cover,
                    summary: article.summary,
                    contents: article.contents
                }),
                { Authorization: accessToken }
            )
    }

    login(user: { username: string, password: string }) {
        return this.httpStrategy
            .POST<types.Token>("/api/auth/token", JSON.stringify(user))
    }

    // checkToken 验证token的正确性
    checkToken(accessToken: string) {
        return this.httpStrategy.GET<any>("/api/auth/token/check", { Authorization: accessToken })
    }

}

export default new Api(new Fetcher());