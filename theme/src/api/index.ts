export interface Paper{
    title: string,
    type: string,
    cover: string,
    create_at: string,
    update_at: string,
    summary: string,
    content: string,
}

export interface PapersResponse {
    code: number,
    message: string,
    data: Paper[]
}

export interface PaperResponse {
    code: number,
    message: string,
    data: Paper,
}

class API{
    // fetchPapers 获取所有文章列表
    fetchPapers(limit:number = 10): Promise<PapersResponse> {
        return fetch(`/api/papers`).then(response => response.json())
    }
    // fetchPaper 通过title获取文章的详细内容
    fetchPaper(key?: string): Promise<PaperResponse> {
        return fetch(`/api/paper?key=${key}`).then(response => response.json())
    }
}


export default new API()