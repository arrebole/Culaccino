
export interface Pagination {
    total_size: number;
    remain_size: number;
}

export interface SectionsLink {
    name: string;
    url: string;
}

export interface Article {
    id: number;
    name: string;
    cover: string;
    summary: string;
    tag: string;
    url: string;
    created_at: string;
    updated_at: string;
}

export interface FullArticle extends Article {
    sections_link: SectionsLink[];
}

export interface FullArticles {
    articles: FullArticle[];
    pagination: Pagination;
}

export interface FullSection {
    id: number;
    article: Article;
    name: string;
    content: string;
    url: string;
    created_at: string;
    updated_at: string;
}






