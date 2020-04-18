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

    listArticles(offset = 0, limit = 10) {
        return this.httpStrategy
            .GET<types.Articles>(`/api/articles?offset=${offset}&limit=${limit}`)
    }
}

export default new Api(new Fetcher());