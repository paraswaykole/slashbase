import React, { FunctionComponent, useEffect } from 'react'
import Header from './header'
import Footer from './footer'
import Sidebar from './sidebar'
import { Outlet } from "react-router-dom";
import { useAppSelector } from '../../redux/hooks';
import { selectIsShowingSidebar } from '../../redux/configSlice';

type PageLayoutPropType = {

}

const AppLayout: FunctionComponent<PageLayoutPropType> = () => {

    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)

    return (
        <React.Fragment>
            <div >
                <Header />
                <div className="appcontent">
                    {isShowingSidebar && <Sidebar />}
                    <main id="mainContainer" className={`maincontainer`}>
                        <Outlet />
                    </main>
                </div>
                <Footer />
            </div>
        </React.Fragment>
    )
}



export default AppLayout