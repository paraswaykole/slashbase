import localforage from 'localforage'
import { User } from './models'
import Constants from '../constants'

const slashbaseStore = (() => {
    if (Constants.Build === 'server') {
        return localforage.createInstance({
            name: "SlashbaseStore"
        })
    }
    return undefined
})()


const CURRENT_USER_KEY = 'currentUser'

const CONFIG_IS_SHOWING_SIDEBAR = 'configIsShowingSidebar'

const loginCurrentUser = async function (currentUser: User): Promise<User> {
    return await slashbaseStore!.setItem(CURRENT_USER_KEY, currentUser)
}

const updateCurrentUser = async function (currentUser: User): Promise<User> {
    return await slashbaseStore!.setItem(CURRENT_USER_KEY, currentUser)
}

const getCurrentUser = async function (): Promise<User | null> {
    return await slashbaseStore!.getItem(CURRENT_USER_KEY)
}

const logoutUser = async function (): Promise<void> {
    return slashbaseStore!.clear()
}

const isShowingSidebar = async function (): Promise<boolean> {
    if (Constants.Build === 'desktop') {
        return true
    }
    const isShowing: boolean | null = await slashbaseStore!.getItem(CONFIG_IS_SHOWING_SIDEBAR)
    if (isShowing == null) {
        return true
    }
    return isShowing
}

const setIsShowingSidebar = async function (isShowingSidebar: boolean): Promise<boolean> {
    if (Constants.Build === 'desktop') {
        return true
    }
    return await slashbaseStore!.setItem(CONFIG_IS_SHOWING_SIDEBAR, isShowingSidebar)
}

export default {
    loginCurrentUser,
    updateCurrentUser,
    getCurrentUser,
    logoutUser,
    isShowingSidebar,
    setIsShowingSidebar
}