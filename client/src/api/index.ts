import IResp from '@/types/resp';


export default {
    getAllTable,
    delArticle,
}


function getAllTable(): Promise<IResp> {
    return fetch("/api/admin/table/all", {
        //mode: "cors",
        method: "GET",
        credentials: "include",
        headers: {
            'content-type': 'application/json'
        },
    }).then(response => response.json())
}

function delArticle(id:number): Promise<IResp> {
    return fetch(`/api/admin/delete/${id}`, {
        //mode: "cors",
        method: "DELETE",
        credentials: "include",
        headers: {
            'content-type': 'application/json'
        },
    }).then(response => response.json())
}