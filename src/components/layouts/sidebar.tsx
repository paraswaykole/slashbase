import styles from './sidebar.module.scss'
import React from 'react'
import { useRouter } from 'next/router'

type SidebarPropType = { }

const Sidebar = (_: SidebarPropType) => {

    const router = useRouter()

    return (
        <aside className={styles.sidebar}> 
            <div className={styles.spacebox}>
                <strong>Databases</strong>
            </div>
        </aside>
    )
}


export default Sidebar