
export interface DashboardData {
    title: string,
    type: string,
    summary: string,
}

export interface DetailData{
    title: string,
    type: string,
    create_at: string,
    update_at: string,
    summary: string,
    content: string,
}

export interface Dashboard {
    code: number,
    message: string,
    data: DashboardData[]
}

export interface Detail {
    code: number,
    message: string,
    data: DetailData,
}


function getDashboard(): Promise<Dashboard> {
    return fetch("/api/dashboard").then(response => response.json())
}

function getDetail(title?: string): Promise<Detail> {
    return fetch(`/api/paper?title=${title}`).then(response => response.json())
}


export default {
    getDashboard,
    getDetail,
}