import posthog from 'posthog-js'
import { AppID } from '../../wailsjs/go/app/App'

const posthogInit = async () => {
    let appID: string = await AppID()
    posthog.init(
        String(import.meta.env.VITE_POSTHOG_KEY),
        {
            api_host: import.meta.env.VITE_POSTHOG_API_HOST,
            capture_pageview: false,
            bootstrap: {
                distinctID: appID,
            }
        }
    )
}

export default {
    posthogInit
}