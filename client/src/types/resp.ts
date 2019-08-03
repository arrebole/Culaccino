export default interface IResp {
    code: string
    power: string
    token: string
    articles: IArticle
    dir: IArticle[]
}

export interface githubApi{
  text: string
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




