export default interface IResp {
    code: string
    power: string
    token: string
    article: IArticle
    dir: IArticle[]
    count: number
}

export class Article{
    constructor(){
        this.ID = -1
        this.title = ""
        this.CreatedAt = ""
        this.UpdatedAt = ""
        this.views = 0
        this.author = ""
        this.contents = ""
        this.cover = ""
        this.summary = ""
        this.target = ""
    }
    CreatedAt: string
    UpdatedAt: string
    ID: number
    views: number
    title: string
    author: string
    target: string
    cover: string
    summary: string
    contents: string
}

export interface IArticle extends IArticleBase {
    CreatedAt: string
    UpdatedAt: string
    ID: number
    views: number
}

export interface IArticleBase {
    title: string
    author: string
    target: string
    cover: string
    summary: string
    contents: string
}




