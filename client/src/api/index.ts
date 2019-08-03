import IResp, { IArticle, IArticleBase } from '@/types/resp';


export default {
    githubRenderAPI,
    getAllTable,
    delArticle,
    getArticle,
    updateArticle,
    createArticle,
    getTable,
    login,
    token,
}

function createFetch(url: string, method: string, data?: string) {
    return fetch(url, {
        method,
        mode: 'cors',
        credentials: "same-origin",
        headers: {
            'Content-Type': 'application/json',
        },
        body: data
    }).then(response => response.json())
}

function githubRenderAPI(aText: { text: string }): Promise<string> {
    return fetch('https://api.github.com/markdown', {
        method: "POST",
        mode: 'cors',
        credentials: "same-origin",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(aText)
    }).then(res => res.text())
}

function token(aToken: string): Promise<IResp> {
    return createFetch(`/api/token?token=${aToken}`, "GET");
}


function login(user: { userName: string, password: string }): Promise<IResp> {
    return createFetch(`/api/login?userName=${user.userName}&password=${user.password}`, "GET");
}

function getTable(page: number): Promise<IResp> {
    return createFetch(`/api/table/${page}`, "GET");
}

function getAllTable(): Promise<IResp> {
    return createFetch("/api/admin/table/all", "GET");
}

function delArticle(id: number): Promise<IResp> {
    return createFetch(`/api/admin/delete/${id}`, "DELETE");
}
function getArticle(id: number): Promise<IResp> {
    return createFetch(`/api/contents/${id}`, "GET");
}

function createArticle(aArticle: IArticleBase): Promise<IResp> {
    return createFetch(`/api/admin/add/text`, "POST", JSON.stringify(aArticle));
}

function updateArticle(id: number, aArticle: IArticleBase): Promise<IResp> {
    return createFetch(`/api/admin/update/${id}`, "PUT", JSON.stringify(aArticle));
}