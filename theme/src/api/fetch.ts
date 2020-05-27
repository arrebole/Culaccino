import Transfer from "./http"

export default class FetchApi implements Transfer {

    private fetchJSON<T>(url: string, options?: RequestInit | undefined): Promise<T> {
        return fetch(url, options)
            .then(res => res.status >= 400 ? Promise.reject() : res.json())
            .then(res => res as T);
    }

    GET<T>(url: string, headers: object = {}): Promise<T> {
        return this.fetchJSON<T>(url, {
            method: "GET",
            headers: { ...headers }
        });
    }

    POST<T>(url: string, body: string, headers: object = {}): Promise<T> {
        return this.fetchJSON<T>(url, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
                ...headers
            },
            body: body
        })
    }

    PATCH<T>(url: string, body: string, headers: object = {}): Promise<T> {
        return this.fetchJSON<T>(url, {
            method: "PATCH",
            headers: {
                'Content-Type': 'application/json',
                ...headers
            },
            body: body
        })
    }

    DELETE<T>(url: string, headers: object = {}): Promise<T> {
        return this.fetchJSON<T>(url, {
            method: "DELETE",
            headers: {
                ...headers
            }
        })
    }
}
