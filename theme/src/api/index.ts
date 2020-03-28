export interface Paper {
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

class API {
    // fetchPapers 获取所有文章列表
    fetchPapers(limit: number = 10): Promise<PapersResponse> {
        return fetch(`/api/papers`).then(response => response.json())
    }

    // newPaper 创建新的文章
    newPaper(data: Paper): Promise<PaperResponse> {
        return fetch(`/api/papers`, { method: "POST", body: JSON.stringify(data) }).then(response => response.json())
    }

    // fetchPaper 通过title获取文章的详细内容
    fetchPaper(key: string): Promise<PaperResponse> {
        return fetch(`/api/papers/${key}`).then(response => response.json())
    }

    // updatePaper 更新的文章
    updatePaper(data: Paper): Promise<PaperResponse> {
        return fetch(`/api/papers/${data.title}`, { method: "PATCH", body: JSON.stringify(data) }).then(response => response.json())
    }
}


export default new API()