import React, { FunctionComponent, useEffect } from 'react'
import Header from './header'
import Footer from './footer'
import Sidebar from './sidebar'
import { Outlet, useLocation } from "react-router-dom";
import { useAppSelector } from '../../redux/hooks';
import { selectIsShowingSidebar } from '../../redux/configSlice';
import TabsBar from './tabsbar';

type PageLayoutPropType = {

}

const AppLayout: FunctionComponent<PageLayoutPropType> = () => {

    const location = useLocation()

    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)

    let showTabsBar = (location.pathname.startsWith("/db")) ? true : false

    return (
        <React.Fragment>
            <div >
                <Header />
                <div className="appcontent">
                    {isShowingSidebar && <Sidebar />}
                    <main className={"maincontainer" + (isShowingSidebar ? ' withsidebar' : '')}>
                        {showTabsBar && <TabsBar />}
                        <div id="maincontent" className="maincontent">
                            <Outlet />
                        </div>
                    </main>
                </div>
                <Footer />
            </div>
        </React.Fragment>
    )
}



export default AppLayout