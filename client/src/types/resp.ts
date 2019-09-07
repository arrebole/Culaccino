export default interface IResp {
    code: number
    message: string
    data: {
        session: Session,
        Storage: Storage
        repo: Repo
        repos: Repo[]
        file: File
        count: Count
        issues:Issues
    }
}

export class Storage {
    constructor() {
        this.name = ""
        this.avatar = ""
    }
    name: string
    avatar: string
}

export class File {
    constructor() {
        this.hash = ""
        this.url = ""
    }
    hash: string
    url: string
}


export class Count {
    constructor() {
        this.remain = 0
        this.total = 0
    }
    remain: number
    total: number
}

export class Session {
    constructor() {
        this.cookie = ""
        this.secret = ""
        this.permission = 0
    }
    cookie: string
    secret: string
    permission: number
}

export class Repo {
    constructor() {
        this.area = ""
        this.cover = ""
        this.hate = 0
        this.parents = ""
        this.star = 0
        this.summary = ""
        this.tag = ""
    }
    star: number
    hate: number
    tag: string
    parents: string
    area: string
    cover: string
    summary: string
}


export class Chapter {
    constructor() {
        this.index = 0
        this.name = ""
        this.parents = ""
        this.contents = ""
    }
    index: number
    name: string
    parents: string
    contents: string
}


export class Issues {
    constructor() {
        this.parents = ""
        this.seat = 0
        this.contents = ""
    }
    parents: string
    seat: number
    contents: string
}



