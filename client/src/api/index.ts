import IResp, { IArchiveBase, IArchive } from '@/types/resp';


export default {
    getAllTable,
    delArchive,
    getArchive,
    updateArchive,
    createArchive,
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
    return createFetch("/api/admin/table", "GET");
}

function delArchive(id: number): Promise<IResp> {
    return createFetch(`/api/admin/delete/${id}`, "DELETE");
}
function getArchive(id: number): Promise<IResp> {
    return createFetch(`/api/archives/${id}`, "GET");
}

function createArchive(item: IArchiveBase): Promise<IResp> {
    return createFetch(`/api/admin/add`, "POST", JSON.stringify(item));
}

function updateArchive(id: number| string, item: IArchiveBase): Promise<IResp> {
    return createFetch(`/api/admin/update/${id}`, "PUT", JSON.stringify(item));
}