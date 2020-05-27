// 传输接口，实现http方法
export default interface Http {
    GET<T>(url: string, headers?: object): Promise<T>
    POST<T>(url: string, body: string, headers?: object): Promise<T>
    PATCH<T>(url: string, body: string, headers?: object): Promise<T>
    DELETE<T>(url: string, headers?: object): Promise<T>
}
