interface ConstantsType {
    API_HOST: string
    API_URL: string
    IS_LIVE: boolean
    APP_PATHS: {
        [key: string]: {
            href: string
            as: string
            isAuth: boolean
        }
    }
}

const Constants: ConstantsType = {
    API_HOST: String(process.env.API_HOST),
    API_URL: process.env.API_HOST+"/api/v1",
    IS_LIVE: Boolean(process.env.IS_LIVE),


    APP_PATHS: {
        LOGIN: {
            href: '/login',
            as: '/login',
            isAuth: false
        },
        LOGOUT: {
            href: '/logout',
            as: '/logout',
            isAuth: true
        },
        PROJECT: {
            href: '/project/[id]',
            as: '/project/',
            isAuth: true
        },
        DB: {
            href: '/db/[id]',
            as: '/db/',
            isAuth: true
        },
        HOME: {
            href: '/',
            as: '/',
            isAuth: true
        },
    }

}

export default Constants