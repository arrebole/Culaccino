export interface Pagination {
    total_size: number;
    remain_size: number;
}

export interface Article {
    id: number;
    name: string;
    cover: string;
    summary: string;
    contents: string;
    tag: string;
    url: string;
    created_at: string;
    updated_at: string;
}

export interface Articles {
    articles: Article[];
    pagination: Pagination;
}

export interface Token {
    access_token: string
    token_type: string
    expires_at: string
}

export function defatultArticle(): Article {
    return {
        id: 0,
        name: "",
        cover: "",
        summary: "",
        contents: "",
        tag: "",
        url: "",
        created_at: "",
        updated_at: "",
    }
}

export function defatultArticles(): Articles {
    return {
        articles: new Array<Article>(),
        pagination: {
            total_size: 0,
            remain_size: 0,
        }
    }
}


