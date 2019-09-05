import IResp, { IArchiveBase, IArchive } from '@/types/resp';


export default {
    getDashboard,
    
    getReposOwner,
    getRepo,

    newRepo,
    delRepo,
    commitRepo,
    
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
    return createFetch("GET", `/api/search/dashboard?page=${page}`);
}

function getReposOwner(domain:string): Promise<IResp> {
    return createFetch("GET", `/api/repo/${domain}`);
}

function newRepo(item: IArchiveBase): Promise<IResp> {
    return createFetch("POST", `/api/new`, { data: JSON.stringify(item) });
}

function commitRepo(item:{domain:string,repoName:string}, data: IArchiveBase): Promise<IResp> {
    return createFetch("PUT", `/api/repo/${item.domain}/${item.repoName}`, { data: JSON.stringify(data) });
}

function delRepo(item:{domain:string,repoName:string}): Promise<IResp> {
    return createFetch("DELETE", `/api/repo/${item.domain}/${item.repoName}`);
}
function getRepo(item:{domain:string,repoName:string}): Promise<IResp> {
    return createFetch("GET", `/api/repo/${item.domain}/${item.repoName}`);
}



