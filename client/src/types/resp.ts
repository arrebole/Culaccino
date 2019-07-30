export default interface IResp {
    code: string
    token: string
    main: IArticle
    dir: IArticle[]
}

export interface IArticle {
    ID: number
    CreatedAt: string
    UpdatedAt: string

    title: string
    author: string
    kind: string
    cover: string
    summary: string

    views: number
    contents: string
}
