import {createInjectionToken} from "ngxtension/create-injection-token"

export const [injectLogger] = createInjectionToken(() => console, {isRoot: true})
