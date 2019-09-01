import IResp, { IArchiveBase, IArchive } from '@/types/resp';


export default {
    getArchiveAll,
    delArchive,
    getArchive,
    updateArchive,
    createArchive,
    getDashboard,
    sessionLogin,
    sessionExist,
    staticUpload
}

function createFetch(method: string, url: string, options?: { headers?: Headers, data?: any }) {
    return fetch(url, {
        method,
        mode: 'cors',
        credentials: "same-origin",
        headers: options && options.headers ? options.headers : {
            "Content-type":"application/json"
        },
        body: options && options.data ? options.data : null
    }).then(response => response.json())
}

function staticUpload(data: File): Promise<IResp> {
    const formData = new FormData()
    formData.append('file', data)

    return createFetch("POST", "/api/static/upload", {
        data: formData, headers: new Headers()
    })
}


function sessionExist(aToken: string): Promise<IResp> {
    return createFetch("GET", `/api/session/exist?token=${aToken}`);
}

function sessionLogin(user: { userName: string, password: string }): Promise<IResp> {
    return createFetch("GET", `/api/session/login?userName=${user.userName}&password=${user.password}`);
}

function getDashboard(page: number, per_page: number): Promise<IResp> {
    return createFetch("GET", `/api/archive/dashboard/${page}`);
}

function getArchiveAll(): Promise<IResp> {
    return createFetch("GET", "/api/archive/all");
}

function delArchive(id: number): Promise<IResp> {
    return createFetch("DELETE", `/api/archive/delete/${id}`);
}
function getArchive(id: number): Promise<IResp> {
    return createFetch("GET", `/api/archive/details/${id}`);
}

function createArchive(item: IArchiveBase): Promise<IResp> {
    return createFetch("POST", `/api/archive/create`, { data: JSON.stringify(item) });
}

function updateArchive(id: number | string, item: IArchiveBase): Promise<IResp> {
    return createFetch("PUT", `/api/archive/update/${id}`, { data: JSON.stringify(item) });
}