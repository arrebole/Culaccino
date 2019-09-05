export default interface IResp {
    code: number
    message: string
    data: {
        archive: IArchive
        dir: IArchive[]
        user: {
            uid: number
            domain: string
            token: string
        },
        file: {
            hash: string
            url: string
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
        this.area = ""
        this.summary = ""
        this.target = ""
    }
    CreatedAt: string
    UpdatedAt: string
    ID: number
    views: number
    title: string
    author: string
    area: string
    target: string
    cover: string
    summary: string
    contents: string
}
export interface IArchiveBase {
    title: string
    target: string
    area: string
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
        this.target = ""
        this.cover = ""
        this.summary = ""
        this.contents = ""
        this.area = ""
        this.hide = ""
    }
    title: string
    target: string
    area: string
    cover: string
    summary: string
    contents: string
    hide: string
}






