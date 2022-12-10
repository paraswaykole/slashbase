interface ConstantsType {
    IS_LIVE: boolean
    APP_PATHS: {
        [key: string]: {
            path: string
        }
    }
    EXTERNAL_PATHS: {
        [key: string]: string
    }
    SETTING_KEYS: {
        [key: string]: string
    }
}

declare global {
    var CONFIG: {
        API_HOST: string;
    }
}

const LOCAL = 'local'
const PRODUCTION = 'production'

const Constants: ConstantsType = {
    IS_LIVE: Boolean(process.env.NEXT_PUBLIC_ENV_NAME === PRODUCTION),
    APP_PATHS: {
        // APP
        HOME: {
            path: '/',
        },
        PROJECT: {
            path: '/project/[id]',
        },
        NEW_DB: {
            path: '/project/[id]/newdb',
        },
        DB: {
            path: '/db/[id]',
        },
        DB_PATH: {
            path: '/db/[id]/[path]',
        },
        DB_QUERY: {
            path: '/db/[id]/query/[queryId]',
        },
        DB_HISTORY: {
            path: '/db/[id]/history',
        },
        // SETTINGS
        SETTINGS: {
            path: '/settings',
        },
        SETTINGS_ABOUT: {
            path: '/settings/about',
        },
        SETTINGS_SUPPORT: {
            path: '/settings/support',
        },
        SETTINGS_ADVANCED: {
            path: '/settings/advanced',
        },
    },
    EXTERNAL_PATHS: {
        OFFICIAL_WEBSITE: "https://slashbase.com",
        DISCORD_COMMUNITY: "https://discord.gg/U6fXgm3FAX",
        REPORT_BUGS: "https://github.com/slashbaseide/slashbase/issues",
        CHANGELOG: "https://slashbase.com/updates",
    },
    SETTING_KEYS: {
        APP_ID: "APP_ID",
        TELEMETRY_ENABLED: "TELEMETRY_ENABLED",
        LOGS_EXPIRE: "LOGS_EXPIRE"
    },
}

export default Constants

export const GetAPIConfig = function () {
    let API_HOST = String(process.env.NEXT_PUBLIC_API_HOST)
    return {
        API_HOST: API_HOST,
        API_URL: API_HOST + "/api/v1"
    }
}