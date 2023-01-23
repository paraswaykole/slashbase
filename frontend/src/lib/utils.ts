declare var window: any;

const openInBrowser = (url: string) => { window.runtime.BrowserOpenURL(url) }

export default {
    openInBrowser
}