import localforage from 'localforage'

const slashbaseStore = localforage.createInstance({
    name: "SlashbaseStore"
})

const CONFIG_IS_SHOWING_SIDEBAR = 'configIsShowingSidebar'

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
    isShowingSidebar,
    setIsShowingSidebar
}