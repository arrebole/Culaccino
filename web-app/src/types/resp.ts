export default interface IResp {
    code: number
    message: string
    data:{
        islogin: boolean
        token: string
        archive: IArchive
        dir: IArchive[]
        file:{
            hash:string
            url:string
        }
        count: {
            remain: number
            total: number
        }
    }
}

export class Archive {
    constructor() {
        this.ID = -1
        this.title = ""
        this.CreatedAt = ""
        this.UpdatedAt = ""
        this.views = 0
        this.author = ""
        this.contents = ""
        this.cover = ""
        this.summary = ""
        this.target = ""
    }
    CreatedAt: string
    UpdatedAt: string
    ID: number
    views: number
    title: string
    author: string
    target: string
    cover: string
    summary: string
    contents: string
}

export interface IArchive extends IArchiveBase {
    CreatedAt: string
    UpdatedAt: string
    ID: number
    views: number
}

export class ArchiveBase implements IArchiveBase {
    constructor() {
        this.title = ""
        this.author = ""
        this.target = ""
        this.cover = ""
        this.summary = ""
        this.contents = ""
    }
    title: string
    author: string
    target: string
    cover: string
    summary: string
    contents: string
}

export interface IArchiveBase {
    title: string
    author: string
    target: string
    cover: string
    summary: string
    contents: string
}




