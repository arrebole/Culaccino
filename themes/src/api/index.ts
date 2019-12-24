
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
    return fetch("http://www.mocky.io/v2/5e01dc192f00003788dcd365").then(response => response.json())
}

function getDetail(symbol?: string):Promise<Detail>{
    return fetch("http://www.mocky.io/v2/5e01e97d2f00001d88dcd3fc").then(response => response.json())
}


export default {
    getDashboard,
    getDetail,
}