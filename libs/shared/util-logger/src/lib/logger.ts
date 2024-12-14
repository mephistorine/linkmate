import {createInjectionToken} from "ngxtension/create-injection-token"

export const [, , LOGGER] = createInjectionToken(() => console, {isRoot: true})
