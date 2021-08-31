import React, { FunctionComponent, useEffect } from 'react'
import Head from 'next/head'
import Header from './header'
import Sidebar from './sidebar'
import { useAppSelector } from '../../redux/hooks'
import { selectIsShowingSidebar } from '../../redux/configSlice'

type PageLayoutPropType = {
    title: string
}    

const AppLayout: FunctionComponent<PageLayoutPropType> = ({title, children}) => {  

    const isShowingSidebar: boolean = useAppSelector(selectIsShowingSidebar)

    return (
        <React.Fragment>
            <Head>
                <title>{title || 'Slashbase'}</title>
            </Head>
            <div >
                <Header />    
                <div className="appcontent">
                    { isShowingSidebar && <Sidebar /> }
                    <main className={`maincontainer${isShowingSidebar?' withsidebar':''}`}>
                        { children }
                    </main>
                </div>
            </div>
        </React.Fragment>
    )
}


export default AppLayout