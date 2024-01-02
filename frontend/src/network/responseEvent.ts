import { EventsOnce } from '../../wailsjs/runtime/runtime'

function responseEvent<T>(eventName: string) {
    return new Promise<T>((resolve) => {
        EventsOnce(eventName, (data) => {
            resolve(data)
        })
    })
}

export default responseEvent