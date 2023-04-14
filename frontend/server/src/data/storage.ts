import localforage from 'localforage'
import { User } from './models'

const slashbaseStore = localforage.createInstance({
    name: "SlashbaseStore"
})

const CURRENT_USER_KEY = 'currentUser'

const CONFIG_IS_SHOWING_SIDEBAR = 'configIsShowingSidebar'

const loginCurrentUser = async function (currentUser: User): Promise<User> {
    return await slashbaseStore.setItem(CURRENT_USER_KEY, currentUser)
}

const updateCurrentUser = async function (currentUser: User): Promise<User> {
    return await slashbaseStore.setItem(CURRENT_USER_KEY, currentUser)
}

const getCurrentUser = async function (): Promise<User | null> {
    return await slashbaseStore.getItem(CURRENT_USER_KEY)
}

const logoutUser = async function (): Promise<void> {
    return slashbaseStore.clear()
}

const isShowingSidebar = async function (): Promise<boolean> {
    const isShowing: boolean | null = await slashbaseStore.getItem(CONFIG_IS_SHOWING_SIDEBAR)
    if (isShowing == null) {
        return true
    }
    return isShowing
}

const setIsShowingSidebar = async function (isShowingSidebar: boolean): Promise<boolean> {
    return await slashbaseStore.setItem(CONFIG_IS_SHOWING_SIDEBAR, isShowingSidebar)
}

export default {
    loginCurrentUser,
    updateCurrentUser,
    getCurrentUser,
    logoutUser,
    isShowingSidebar,
    setIsShowingSidebar
}