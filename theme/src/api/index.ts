export interface Paper {
    title: string,
    type: string,
    cover: string,
    create_at: string,
    update_at: string,
    summary: string,
    content: string,
}

class API {
    // fetchPaper 通过title获取文章的详细内容
    fetchPaper(key: string): Promise<Paper> {
        return fetch(`/api/papers/${key}`).then(res => res.json());
    }

    // fetchPapers 获取所有文章列表
    fetchPapers(limit: number = 10): Promise<Paper[]> {
        return fetch(`/api/papers`).then(res => res.json());
    }

    // newPaper 创建新的文章
    newPaper(data: Paper): Promise<Paper> {
        return fetch(`/api/papers`, { method: "POST", body: JSON.stringify(data) }).then(res => res.json());
    }

    // updatePaper 更新的文章
    updatePaper(data: Paper): Promise<Paper> {
        return fetch(`/api/papers/${data.title}`, { method: "PATCH", body: JSON.stringify(data) })
            .then(res => res.json());
    }
}


export default new API()