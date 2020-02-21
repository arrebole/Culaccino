
export interface DashboardData{
    symbol: string,
    description: string,
}

export interface Dashboard{
    code: number,
    message: string,
    data: DashboardData[]
}

export interface Detail{
    code: number,
    message: string,
    data: string,
}


function getDashboard():Promise<Dashboard> {
    return fetch("/api/directory").then(response => response.json())
}

function getDetail(symbol?: string):Promise<Detail>{
    return fetch(`/api/source?symbol=${symbol}`).then(response => response.json())
}


export default {
    getDashboard,
    getDetail,
}