

interface ConstantsType {
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
    const CONFIG: {
        API_HOST: string;
    }
}

const Constants: ConstantsType = {
    APP_PATHS: {
        // APP
        HOME: {
            path: '/',
        },
        PROJECT: {
            path: '/project/[id]'
        },
        NEW_DB: {
            path: '/project/[id]/newdb'
        },
        DB: {
            path: '/db/[id]'
        },
        // SETTINGS
        SETTINGS: {
            path: '/settings'
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
    },
}

export default Constants

export const GetAPIConfig = function () {
    let API_HOST = String("http://localhost:22022")
    return {
        API_HOST: API_HOST,
        API_URL: API_HOST + "/api/v1"
    }
}