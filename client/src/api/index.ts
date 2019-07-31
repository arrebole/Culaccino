import IResp from '@/types/resp';


export default {
    getAllTable,
    delArticle,
    login,
    token
}

function createFetch(url: string, method: string) {
    return fetch(url, {
        method,
        credentials: "include",
        headers: {
            'content-type': 'application/json'
        },
    }).then(response => response.json())
}


function token(aToken: string): Promise<IResp> {
    return createFetch(`/api/token?token=${aToken}`, "GET");
}


function login(user: { userName: string, password: string }): Promise<IResp> {
    return createFetch(`/api/login?userName=${user.userName}&password=${user.password}`, "GET");
}

function getAllTable(): Promise<IResp> {
    return createFetch("/api/admin/table/all", "GET");
}

function delArticle(id: number): Promise<IResp> {
    return createFetch(`/api/admin/delete/${id}`, "DELETE");
}