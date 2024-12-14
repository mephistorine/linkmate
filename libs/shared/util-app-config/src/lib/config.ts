import {createInjectionToken} from "ngxtension/create-injection-token"

export type AppConfig = {
    readonly apiUrl: string
    createShortUrl(key: string): string
}

function configFactory(): AppConfig {
    return {
        apiUrl: `${envs.LNKM_API_URL}/api`,
        createShortUrl(key: string) {
            return `${envs.LNKM_API_URL}/${key}`
        },
    }
}

export const [, , APP_CONFIG] = createInjectionToken(configFactory, {
    isRoot: true,
})
