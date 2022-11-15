interface ConstantsType {
    IS_LIVE: boolean
    APP_PATHS: {
        [key: string]: {
            path: string
            isAuth: boolean
        }
    }
    EXTERNAL_PATHS: {
        [key: string]: string
    }
    SETTING_KEYS: {
        [key: string]: string
    }
    ROLES: {
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
        LOGIN: {
            path: '/login',
            isAuth: false
        },
        LOGOUT: {
            path: '/logout',
            isAuth: true
        },
        // APP
        HOME: {
            path: '/',
            isAuth: true
        },
        PROJECT: {
            path: '/project/[id]',
            isAuth: true
        },
        NEW_DB: {
            path: '/project/[id]/newdb',
            isAuth: true
        },
        PROJECT_MEMBERS: {
            path: '/project/[id]/members',
            isAuth: true
        },
        DB: {
            path: '/db/[id]',
            isAuth: true
        },
        DB_PATH: {
            path: '/db/[id]/[path]',
            isAuth: true
        },
        DB_QUERY: {
            path: '/db/[id]/query/[queryId]',
            isAuth: true
        },
        DB_HISTORY: {
            path: '/db/[id]/history',
            isAuth: true
        },
        // SETTINGS
        SETTINGS: {
            path: '/settings',
            isAuth: true
        },
        SETTINGS_ABOUT: {
            path: '/settings/about',
            isAuth: true
        },
        SETTINGS_SUPPORT: {
            path: '/settings/support',
            isAuth: true
        },
        SETTINGS_ACCOUNT: {
            path: '/settings/account',
            isAuth: true
        },
        SETTINGS_ACCOUNT_CHANGE_PASSWORD: {
            path: '/settings/account/password',
            isAuth: true
        },
        SETTINGS_USERS: {
            path: '/settings/users',
            isAuth: true
        },
        SETTINGS_ADD_USER: {
            path: '/settings/users/add',
            isAuth: true
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
        TELEMETRY_ENABLED: "TELEMETRY_ENABLED"
    },
    ROLES: {
        ADMIN: "Admin"
    }
}

export default Constants

export const GetAPIConfig = function () {
    let API_HOST = String(process.env.NEXT_PUBLIC_API_HOST)
    return {
        API_HOST: API_HOST,
        API_URL: API_HOST + "/api/v1"
    }
}