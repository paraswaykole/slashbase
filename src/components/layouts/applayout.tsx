import React, { FunctionComponent, useEffect } from 'react'
import Head from 'next/head'
import Header from './header'
import Sidebar from './sidebar'

type PageLayoutPropType = {
    title: string
}    

const AppLayout: FunctionComponent<PageLayoutPropType> = ({title, children}) => {  

    return (
        <React.Fragment>
            <Head>
                <title>{title || 'Slashbase'}</title>
            </Head>
            <div >
                <Header />    
                <div className="appcontent">
                    <Sidebar />
                    { children }
                </div>
            </div>
        </React.Fragment>
    )
}


export default AppLayout