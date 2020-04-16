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
            .GET<types.FullArticle>(`/api/articles/${name}`)
    }

    listArticles() {
        return this.httpStrategy
            .GET<types.FullArticles>(`/api/articles`)
    }

    getSection(articleName: string, sectionName: string) {
        return this.httpStrategy
            .GET<types.FullSection>(`/api/articles/${articleName}/${sectionName}`)
    }
}

export default new Api(new Fetcher());