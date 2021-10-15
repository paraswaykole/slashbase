import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import Link from 'next/link'
import React, { useEffect, useState } from 'react'
import AppLayout from '../../../components/layouts/applayout'

import DefaultErrorPage from 'next/error'
import { getDBConnection, getDBDataModels, getDBQueries, selectDBConnection } from '../../../redux/dbConnectionSlice'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import DBShowDataFragment from '../../../components/dbfragments/showdata'
import { DBConnection } from '../../../data/models'
import Constants from '../../../constants'
import DBShowModelFragment from '../../../components/dbfragments/showmodel'

const DBPage: NextPage = () => {

    const router = useRouter()
    const { id, path, mschema, mname } = router.query

    const [error404, setError404] = useState(false)
    const dispatch = useAppDispatch()

    const dbConnection: DBConnection | undefined = useAppSelector(selectDBConnection)

    useEffect(()=>{
        (async () => {
            if (id){
                try {
                    await dispatch((getDBConnection({dbConnId: String(id)}))).unwrap() 
                } catch (e){
                    setError404(true)
                    return
                }
                dispatch((getDBDataModels({dbConnId: String(id)}))) 
                dispatch((getDBQueries({dbConnId: String(id)}))) 
            }
        })()
    }, [dispatch, router])

    if(error404 || !['data', 'model'].includes(String(path))) {
        return (<DefaultErrorPage statusCode={404} />)
    }

    return (
        <AppLayout title={dbConnection ? dbConnection.name + " | Slashbase" : "Slashbase"}>
            { mschema && mname && 
                <React.Fragment>
                    <div className="tabs is-toggle is-toggle-rounded">
                        <ul>
                            <Link 
                                href={{pathname: Constants.APP_PATHS.DB_PATH.path, query: {mschema, mname}}} 
                                as={Constants.APP_PATHS.DB_PATH.path.replace('[id]', String(id)).replace('[path]', String('data'))+"?mschema="+mschema+"&mname="+mname}
                                >
                                <li className={path === 'data' ? 'is-active': ''}>
                                    <a>
                                        <span className="icon is-small"><i className="fas fa-table" aria-hidden="true"/></span>
                                        <span>Data</span>
                                    </a>
                                </li>
                            </Link>
                            <Link 
                                href={{pathname: Constants.APP_PATHS.DB_PATH.path, query: {mschema, mname}}} 
                                as={Constants.APP_PATHS.DB_PATH.path.replace('[id]', String(id)).replace('[path]', String('model'))+"?mschema="+mschema+"&mname="+mname}
                                >
                                <li className={path === 'model' ? 'is-active': ''}>
                                    <a>
                                        <span className="icon is-small"><i className="fas fa-list-alt" aria-hidden="true"/></span>
                                        <span>Model</span>
                                    </a>
                                </li>
                            </Link>
                        </ul>
                    </div> 
                    { path === 'data' && <DBShowDataFragment key={String(mschema)+String(mname)}/> }
                    { path === 'model' && <DBShowModelFragment /> }
                </React.Fragment>
            }
        </AppLayout>
    )
}

export default DBPage
