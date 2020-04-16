import Transfer from "./http"

export default class FetchApi implements Transfer {

    private fetchJSON<T>(url: string, options?: RequestInit | undefined): Promise<T> {
        return fetch(url, options).then(res => res.json()).then(res => res as T);
    }

    GET<T>(url: string): Promise<T> {
        return this.fetchJSON<T>(url, { method: "GET" });
    }

    POST<T>(url: string, body: string): Promise<T> {
        return this.fetchJSON<T>(url, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json'
            },
            body: body
        })
    }

    PATCH<T>(url: string, body: string): Promise<T> {
        return this.fetchJSON<T>(url, {
            method: "PATCH",
            headers: {
                'Content-Type': 'application/json'
            },
            body: body
        })
    }

    DELETE<T>(url: string): Promise<T> {
        return this.fetchJSON<T>(url, { method: "DELETE" })
    }

}