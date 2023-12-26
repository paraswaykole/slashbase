

interface ConstantsType {
    IsLive: boolean
    Build: 'desktop' | 'server'
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
    ROLES: {
        [key: string]: string
    }
    ROLES_PERMISSIONS: {
        [key: string]: string
    }
}

declare global {
    const CONFIG: {
        API_HOST: string;
    }
}

const Constants: ConstantsType = {
    IsLive: import.meta.env.MODE === 'prodserver' || import.meta.env.MODE === 'proddesktop',
    Build: import.meta.env.MODE.endsWith("desktop") ? 'desktop' : 'server',
    APP_PATHS: {
        // APP
        HOME: {
            path: '/'
        },
        PROJECT: {
            path: '/project/[id]'
        },
        PROJECT_MEMBERS: {
            path: '/project/[id]/members'
        },
        NEW_DB: {
            path: '/project/[id]/newdb'
        },
        DB: {
            path: '/db/[id]'
        },
        LOGOUT: {
            path: '/logout'
        },
        // SETTINGS
        SETTINGS: {
            path: '/settings'
        },
        SETTINGS_ACCOUNT: {
            path: '/settings/account'
        },
        SETTINGS_ABOUT: {
            path: '/settings/about'
        },
        SETTINGS_SUPPORT: {
            path: '/settings/support'
        },
        SETTINGS_GENERAL: {
            path: '/settings/general'
        },
        SETTINGS_ADVANCED: {
            path: '/settings/advanced'
        },
        SETTINGS_USERS: {
            path: '/settings/users'
        },
        SETTINGS_ADD_USER: {
            path: '/settings/users/add'
        },
        SETTINGS_ROLES: {
            path: '/settings/roles'
        },
    },
    EXTERNAL_PATHS: {
        OFFICIAL_WEBSITE: "https://slashbase.com",
        DISCORD_COMMUNITY: "https://discord.gg/U6fXgm3FAX",
        REPORT_BUGS: "https://github.com/slashbaseide/slashbase/issues",
        CHANGELOG: "https://slashbase.com/updates",
        SUPPORT_MAIL: "mailto:slashbaseide@gmail.com",
    },
    SETTING_KEYS: {
        APP_ID: "APP_ID",
        TELEMETRY_ENABLED: "TELEMETRY_ENABLED",
        LOGS_EXPIRE: "LOGS_EXPIRE",
        OPENAI_KEY: "OPENAI_KEY",
        OPENAI_MODEL: "OPENAI_MODEL"
    },
    ROLES: {
        ADMIN: "Admin"
    },
    ROLES_PERMISSIONS: {
        READ_ONLY: "READ_ONLY"
    }
}

export default Constants

export const GetAPIConfig = function () {
    let API_HOST = Constants.IsLive ? "" : "http://localhost:3000"
    return {
        API_HOST: API_HOST,
        API_URL: API_HOST + "/api/v1"
    }
}
