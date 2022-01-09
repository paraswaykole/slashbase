interface ConstantsType {
    IS_LIVE: boolean
    APP_PATHS: {
        [key: string]: {
            path: string
            isAuth: boolean
        }
    }
}

declare global {
    var CONFIG: {
        API_HOST: string;
    }
}

const LOCAL='local'
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
        ACCOUNT: {
            path: '/account',
            isAuth: true
        },
        SETTINGS_USER: {
            path: '/users',
            isAuth: true
        },
        SETTINGS_ADD_USER: {
            path: '/users/add',
            isAuth: true
        },
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
    }

}

export default Constants

export const GetAPIConfig = function () {
    let API_HOST = String(process.env.NEXT_PUBLIC_API_HOST)
    if (process.env.NEXT_PUBLIC_ENV_NAME === PRODUCTION && global.CONFIG?.API_HOST !== '#API_HOST#') {
        API_HOST = global.CONFIG?.API_HOST
    }
    return {
        API_HOST: API_HOST,
        API_URL: API_HOST + "/api/v1"
    }
}