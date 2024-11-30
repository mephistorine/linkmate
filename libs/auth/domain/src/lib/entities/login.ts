export type AuthLoginReqDto = {
    readonly email: string
    readonly password: string
}

export type AuthLoginResDto = {
    readonly accessToken: string
}
