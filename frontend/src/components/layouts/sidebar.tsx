import styles from './sidebar.module.scss'
import { useLocation } from 'react-router-dom'
import { DBConnection } from '../../data/models'
import { useAppSelector } from '../../redux/hooks'
import { selectDBConnection } from '../../redux/dbConnectionSlice'
import HomeSidebar from './sidebars/homesidebar'
import DatabaseSidebar from './sidebars/dbsidebar'
import SettingSidebar from './sidebars/settingsidebar'

enum SidebarViewType {
    HOME = "HOME", // home sidebar
    DATABASE = "DATABASE", // Used to show elements of database screen
    SETTINGS = "SETTINGS" // Used to show elements of settings screen
}

const Sidebar = () => {

    const location = useLocation()

    const sidebarView: SidebarViewType =
        (location.pathname.startsWith("/db")) ?
            SidebarViewType.DATABASE : (location.pathname.startsWith("/settings")) ? SidebarViewType.SETTINGS : SidebarViewType.HOME

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    return (
        <aside className={"menu " + styles.sidebar}>
            <div className={styles.spacebox}>
                {sidebarView === SidebarViewType.HOME &&
                    <HomeSidebar />
                }
                {sidebarView === SidebarViewType.DATABASE && dbConnection &&
                    <DatabaseSidebar />
                }
                {sidebarView === SidebarViewType.SETTINGS &&
                    <SettingSidebar />
                }
            </div>
        </aside>
    )
}


export default Sidebar