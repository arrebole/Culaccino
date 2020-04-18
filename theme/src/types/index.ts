
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
    articles:  Article[];
    pagination: Pagination;
}

export function defatultArticle():Article{
    return {
        id: 0,
        name: "loading",
        cover: "loading",
        summary: "loading",
        contents: "loading",
        tag: "loading",
        url: "loading",
        created_at: "",
        updated_at: "",
    }
}




