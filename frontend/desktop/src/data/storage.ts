

// const slashbaseStore = localforage.createInstance({
//     name: "SlashbaseStore"
// })

const CURRENT_USER_KEY = 'currentUser'

const CONFIG_IS_SHOWING_SIDEBAR = 'configIsShowingSidebar'

const isShowingSidebar = async function (): Promise<boolean> {
    return true
}

const setIsShowingSidebar = async function (isShowingSidebar: boolean): Promise<boolean> {
    return true
}

export default {
    isShowingSidebar,
    setIsShowingSidebar
}