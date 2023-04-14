import styles from './tabsbar.module.scss'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { createTab, closeTab, selectTabs, setActiveTab } from '../../redux/tabsSlice'
import { TabType } from '../../data/defaults'
import { Tab } from '../../data/models'
import { selectDBConnection } from '../../redux/dbConnectionSlice'


type TabsBarPropType = {}

const TabsBar = (_: TabsBarPropType) => {
    const dispatch = useAppDispatch()

    const dbConnection = useAppSelector(selectDBConnection)
    const tabs: Tab[] = useAppSelector(selectTabs)

    const createNewTab = async () => {
        await dispatch(createTab({ dbConnId: dbConnection!.id, tabType: TabType.BLANK }))
    }

    const switchToTab = async (tabId: string) => {
        dispatch(setActiveTab(tabId))
    }

    const handleCloseTab = async (e: React.MouseEvent<HTMLElement>, tabId: string) => {
        e.stopPropagation()
        await dispatch(closeTab({ dbConnId: dbConnection!.id, tabId }))
    }

    if (!dbConnection) {
        return <></>
    }

    return (
        <div className={"tabs is-boxed " + styles.tabs}>
            <ul>
                {tabs.map(t => <li key={t.id} className={t.isActive ? "is-active" : ""} onClick={() => { switchToTab(t.id) }}>
                    <a>
                        <span>
                            {t.type === TabType.BLANK && "New Tab"}
                            {t.type === TabType.HISTORY && "History"}
                            {t.type === TabType.CONSOLE && "Console"}
                            {t.type === TabType.DATA && `${t.metadata.schema === '' ? t.metadata.name : `${t.metadata.schema}.${t.metadata.name}`}`}
                            {t.type === TabType.MODEL && `${t.metadata.schema === '' ? t.metadata.name : `${t.metadata.schema}.${t.metadata.name}`}`}
                            {t.type === TabType.QUERY && `${t.metadata.queryName ? t.metadata.queryName : "New Query"}`}
                        </span>
                        <span className={"icon " + (t.isActive ? styles.tabsCloseBtn : styles.tabsCloseBtnInActive)} onClick={(e) => { handleCloseTab(e, t.id) }}><i className="fas fa-times" aria-hidden="true"></i></span>
                    </a>
                </li>)}
                <li>
                    <a onClick={createNewTab}>
                        <span className="icon"><i className="fas fa-plus" aria-hidden="true"></i></span>
                    </a>
                </li>
            </ul>
        </div>
    )
}

export default TabsBar