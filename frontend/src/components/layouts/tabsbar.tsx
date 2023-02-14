import styles from './tabsbar.module.scss'
import { useAppDispatch, useAppSelector } from '../../redux/hooks'
import { createTab, closeTab, selectTabs } from '../../redux/tabsSlice'
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

    const handleCloseTab = async (tabId: string) => {
        await dispatch(closeTab({ dbConnId: dbConnection!.id, tabId }))
    }

    if (!dbConnection) {
        return <></>
    }

    return (
        <div className={"tabs is-boxed " + styles.tabs}>
            <ul>
                {tabs.map(t => <li key={t.id} className={t.isActive ? "is-active" : ""}>
                    <a>
                        <span>
                            {t.type === TabType.BLANK && "New Tab"}
                        </span>
                        <span className="icon" onClick={() => { handleCloseTab(t.id) }}><i className="fas fa-times" aria-hidden="true"></i></span>
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